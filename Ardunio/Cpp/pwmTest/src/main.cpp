#include <Arduino.h>

//Initializing LED Pin
int led_pin = 11;

void setup() {
  //Declaring LED pin as output
  pinMode(led_pin, OUTPUT);
}

void loop() {
  for (int i=0;i<255;i++){
    analogWrite(led_pin,i); // analog olduğu önemli 
    delay(5);
  }

    for (int i=255;i>0;i--){
    analogWrite(led_pin,i);
    delay(5);
  }

  
}