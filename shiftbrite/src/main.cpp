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
uint8_t XP_PIN = 3;
uint8_t YP_PIN = 1;
uint8_t XM_PIN = 0;
uint8_t YM_PIN = 2;
uint16_t RX_VAL = 870;  // Avg Resistance across screen

int16_t lims[4] = {140, 870, 0, 1023};  // X lims, Y lims
int16_t scale[2] = {lims[1]-lims[0], lims[3]-lims[2]};
long rgb[3];

ShiftBrite sb = ShiftBrite(NUM_PIXELS, LATCH_PIN);
TouchScreen ts = TouchScreen(XM_PIN, YP_PIN, XM_PIN, YM_PIN, RX_VAL);


// ---- Converting HSV to RGB color
long HSV_to_RGB( float h, float s, float v ) {
  // Range of color values
  long range = 1023;

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

  if(coords.x < lims[1] && coords.y <lims[3]){ // Thresholding for touch activity

    // Detect out of range values and trim them.
    if(coords.x > lims[1]){coords.x = lims[1];}
    if(coords.y > lims[3]){coords.y = lims[3];}
    if(coords.x < lims[0]){coords.x = lims[0];}
    if(coords.y < lims[2]){coords.y = lims[2];}

    // Scale the coords to match 1023 resolution
    coords.x = int(1023.0 * (coords.x - lims[0]) / scale[0]);
    coords.y = int(1023.0 * (coords.y - lims[2]) / scale[1]);

    float floatcoords[2]={float(coords.x),float(coords.y)}; //cheap type conversion.

    long rgbval = HSV_to_RGB(6*floatcoords[0]/1023,1,floatcoords[1]/1023); // Get RGB from coords via HV
    //long rgbval = HSV_to_RGB(6*floatcoords[0]/1023,floatcoords[1]/1023,.1); // Get RGB from coords via HS

    // Shifting out the returns
    rgb[0] = (rgbval & 0x00FF0000) >> 16; // there must be better ways
    rgb[1] = (rgbval & 0x0000FF00) >> 8;
    rgb[2] = rgbval & 0x000000FF;

    Serial.print(coords.x);
    Serial.print("\t");
    Serial.print(coords.y);
    Serial.print("\t");
    Serial.print(rgbval);
    Serial.print("\t");
    Serial.print(rgb[0]);
    Serial.print("\t");
    Serial.print(rgb[1]);
    Serial.print("\t");
    Serial.println(rgb[2]);


    for (int i = 0; i < NUM_PIXELS; ++i) {
      sb.unsetPixel(i);
      sb.setPixelRGB(i, rgb[0], rgb[1], rgb[2]);
    }
    sb.show();

  }
  delay(1000);
}
