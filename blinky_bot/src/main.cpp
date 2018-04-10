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
#include <Wire.h>
#include <Adafruit_PWMServoDriver.h>

// Use the default address for the first board
Adafruit_PWMServoDriver pwm = Adafruit_PWMServoDriver(0x40);

// Depending on your servo make, the pulse width min and max may vary, you
// want these to be as small/large as possible without hitting the hard stop
// for max range. You'll have to tweak them as necessary to match the servos you
// have!
#define SERVOMIN 200 // this is the 'minimum' pulse length count (out of 4096)
#define SERVOMAX 450 // this is the 'maximum' pulse length count (out of 4096)

// our servo # counter
uint8_t servo_00 = 4;
uint8_t servo_01 = 5;

// Continuous servo positions
uint16_t BACKWARD = map(0, 0, 180, SERVOMIN, SERVOMAX); // ms, "0"
uint16_t     STOP = map(90, 0, 180, SERVOMIN, SERVOMAX); // ms, "90"
uint16_t  FORWARD = map(180, 0, 180, SERVOMIN, SERVOMAX); // ms, "180"

// MyoWare Sensor
int ledPin = 13;      // select the pin for the LED

void setup() {
  Serial.begin(9600);
  Serial.println("16 channel Servo test!");

  // LED Indicator
  pinMode(ledPin, OUTPUT);

  pwm.begin();
  pwm.setPWMFreq(60);  // Analog servos run at ~60 Hz updates
  pwm.setPWM(servo_00, 0, STOP);
  pwm.setPWM(servo_01, 0, STOP);

  delay(1000);
  yield();
}

void loop() {

  digitalWrite(ledPin, LOW);

  pwm.setPWM(servo_00, 0, FORWARD);
  delay(1000);
  pwm.setPWM(servo_01, 0, FORWARD);
  delay(1000);

  digitalWrite(ledPin, HIGH);

  pwm.setPWM(servo_00, 0, BACKWARD);
  delay(1000);
  pwm.setPWM(servo_01, 0, BACKWARD);
  delay(1000);
}
