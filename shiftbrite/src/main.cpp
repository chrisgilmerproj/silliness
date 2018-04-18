/***************************************************
  This is an example for our Adafruit 16-channel PWM & Servo driver
  Servo test - this will drive 16 servos, one after the other

  Pick one up today in the adafruit shop!
  ------> http://www.adafruit.com/products/815

  These displays use I2C to communicate, 2 pins are required to
  interface. For Arduino UNOs, thats SCL -> Analog 5, SDA -> Analog 4

  Adafruit invests time and resources providing this open source code,
  please support Adafruit and open-source hardware by purchasing
  products from Adafruit!

  Written by Limor Fried/Ladyada for Adafruit Industries.
  BSD license, all text above must be included in any redistribution
 ****************************************************/

#include <Arduino.h>
#include <shiftbrite.h>

#define NUM_PIXELS 2
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
}
