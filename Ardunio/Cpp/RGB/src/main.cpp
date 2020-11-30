#include <Arduino.h>
int red_light_pin= 11;
int green_light_pin = 10;
int blue_light_pin = 9;
void setup() {
  pinMode(13,1);
  pinMode(red_light_pin, OUTPUT);
  pinMode(green_light_pin, OUTPUT);
  pinMode(blue_light_pin, OUTPUT);
}

void RGB_color(int red_light_value, int green_light_value, int blue_light_value)
 {
  analogWrite(red_light_pin, red_light_value);
  analogWrite(green_light_pin, green_light_value);
  analogWrite(blue_light_pin, blue_light_value);
}

void aloop() {
  digitalWrite(13,1);
  RGB_color(255, 0, 0); // Red
  delay(500);
  RGB_color(0, 255, 0); // Green
  delay(500);
  RGB_color(0, 0, 255); // Blue
  delay(500);
  RGB_color(255, 255, 125); // Raspberry
  delay(500);
  RGB_color(0, 255, 255); // Cyan
  delay(500);
  RGB_color(255, 0, 255); // Magenta
  delay(500);
  RGB_color(255, 255, 0); // Yellow
  delay(500);
  RGB_color(255, 255, 255); // White
  delay(500);
  //RGB_color(random(0,255),random(0,255),random(0,255));
  RGB_color(255,0,0);
  delay(500);
  digitalWrite(13,0);delay(15);
  //delay(1000);
  
}

void loop() {
  RGB_color(255,0,0);
  delay(1000);
  RGB_color(0,255,0);
  delay(1000);
  RGB_color(0,0,255);
  delay(1000);
  RGB_color(0,255,255);
  delay(1000);
  
}
