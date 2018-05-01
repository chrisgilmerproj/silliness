/*
 * Shiftbrite
 *
 * Arduino Uno:
 * DI -> D11
 * LI -> D9
 * EI -> GND
 * CI -> D13
 */

#include <Arduino.h>
#include <shiftbrite.h>
#include <TouchScreen.h>

#define NUM_PIXELS 5
#define LATCH_PIN 9
#define YP_PIN A1
#define XM_PIN A0
#define YM_PIN A2
#define XP_PIN A3

#define MINPRESSURE 10
#define MAXPRESSURE 20000

uint16_t RX_VAL = 870;  // Avg Resistance across screen

int16_t lims[4] = {140, 870, 0, 1023};  // X lims, Y lims
int16_t scale[2] = {lims[1]-lims[0], lims[3]-lims[2]};
int16_t rgbval;
int16_t rgb[3];
int16_t newcoords[2];
float floatcoords[2];

ShiftBrite sb = ShiftBrite(NUM_PIXELS, LATCH_PIN);
TouchScreen ts = TouchScreen(XM_PIN, YP_PIN, XM_PIN, YM_PIN, RX_VAL);


// ---- Converting HSV to RGB color
long HSV_to_RGB( float h, float s, float v ) {
  // Range of color values
  long range = 255;

  // Inits
  int i;
  float m, n, f;

  // Return black for out of range S or V
  if ((s<0.0) || (s>1.0) || (v<0.0) || (v>1.0)) {
    return 0L;
  }

  // Return values of white for out of range H
  if ((h < 0.0) || (h > 6.0)) {
    return long( v * range ) + long( v * range ) * (range + 1) + long( v * range ) * 65536;
  }

  // Calculate RBG from HSV
  i = floor(h);
  f = h - i;
  if ( !(i&1) ) {
    f = 1 - f; // if i is even
  }
  m = v * (1 - s);
  n = v * (1 - s * f);

  switch (i) {
    case 6:
    case 0: // RETURN_RGB(v, n, m)
      return long(v * range ) * 65536 + long( n * range ) * (range + 1) + long( m * range);
    case 1: // RETURN_RGB(n, v, m)
      return long(n * range ) * 65536 + long( v * range ) * (range + 1) + long( m * range);
    case 2:  // RETURN_RGB(m, v, n)
      return long(m * range ) * 65536 + long( v * range ) * (range + 1) + long( n * range);
    case 3:  // RETURN_RGB(m, n, v)
      return long(m * range ) * 65536 + long( n * range ) * (range + 1) + long( v * range);
    case 4:  // RETURN_RGB(n, m, v)
      return long(n * range ) * 65536 + long( m * range ) * (range + 1) + long( v * range);
    case 5:  // RETURN_RGB(v, m, n)
      return long(v * range ) * 65536 + long( m * range ) * (range + 1) + long( n * range);
  }
  return 0;
}

void setup() {
  Serial.begin(115200);

  sb.begin();
  sb.show();
  sb.allOff();
  delay(500);
}

void loop() {

  TSPoint coords = ts.getPoint();
  Serial.print(coords.x);
  Serial.print("\t");
  Serial.print(coords.y);
  Serial.print("\t");
  Serial.println(coords.z);

  // Thresholding for touch activity
  if (coords.z > MINPRESSURE && coords.z < MAXPRESSURE) {

    newcoords[0] = coords.x;
    newcoords[1] = coords.y;

    // Detect out of range values and trim them.
    if(newcoords[0] > lims[1]){newcoords[0] = lims[1];}
    if(newcoords[1] > lims[3]){newcoords[1] = lims[3];}
    if(newcoords[0] < lims[0]){newcoords[0] = lims[0];}
    if(newcoords[1] < lims[2]){newcoords[1] = lims[2];}

    // Scale the coords to match 1023 resolution
    newcoords[0] = int(1023.0 * (newcoords[0] - lims[0]) / scale[0]);
    newcoords[1] = int(1023.0 * (newcoords[1] - lims[2]) / scale[1]);

    // Cheap type conversion.
    floatcoords[0] = float(newcoords[0]);
    floatcoords[1] = float(newcoords[1]);

    rgbval = HSV_to_RGB(6 * floatcoords[0]/1023, 1, floatcoords[1]/1023); // Get RGB from coords via HV
    //rgbval = HSV_to_RGB(6 * floatcoords[0]/1023, floatcoords[1]/1023, .1); // Get RGB from coords via HS

    // Shifting out the returns and normalize back to 1023
    rgb[0] = 4 * (rgbval & 0x00FF0000) >> 16;
    rgb[1] = 4 * (rgbval & 0x0000FF00) >> 8;
    rgb[2] = 4 * (rgbval & 0x000000FF) >> 0;

    Serial.print(newcoords[0]);
    Serial.print("\t");
    Serial.print(newcoords[1]);
    Serial.print("\t");
    Serial.print(rgb[0]);
    Serial.print("\t");
    Serial.print(rgb[1]);
    Serial.print("\t");
    Serial.println(rgb[2]);

    sb.setPixelRGB(0, rgb[0], rgb[1], rgb[2]);
    sb.setPixelRGB(1, 0, 0, 0);
    sb.setPixelRGB(2, 0, 0, 0);
    sb.setPixelRGB(3, 0, 0, 0);
    sb.setPixelRGB(4, 0, 0, 0);
    sb.show();

  }
  delay(1000);
}
