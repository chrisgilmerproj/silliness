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

#define NUM_PIXELS 5
#define LATCH_PIN 9

ShiftBrite sb = ShiftBrite(NUM_PIXELS, LATCH_PIN);

void setup() {
  sb.begin();
  sb.show();
  sb.allOff();
  delay(1000);
}

void loop() {
  for (int i = 0; i < NUM_PIXELS; ++i) {
    sb.setPixelRGB(i, 0, 0, 1023);
    sb.show();
    delay(1000);
  }
  for (int i = 0; i < NUM_PIXELS; ++i) {
    sb.unsetPixel(i);
    sb.show();
    delay(1000);
  }
}
