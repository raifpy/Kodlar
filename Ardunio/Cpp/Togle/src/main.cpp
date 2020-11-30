#include <Arduino.h>

void setup() {
  pinMode(13,1); //OUTPUT
  
}

void loop() {
  digitalWrite(13,1);
  delay(5000); // 5 sec
  //delay(100);
  digitalWrite(13,0);
  delay(100); // 0.1 sec
  //delay(5000);
}