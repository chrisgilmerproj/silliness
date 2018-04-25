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
#define LATCH_PIN 0

ShiftBrite sb = ShiftBrite(NUM_PIXELS, LATCH_PIN);

void setup() {
  sb.begin();
  sb.show();
}

void loop() {
  for (int i = 0; i < NUM_PIXELS; ++i) {
    sb.setPixelRGB(i, 1023, 0, 0);
  }
  sb.show();
  delay(1000);
}
