#include <Arduino.h>
#define sleepDef 200
int sleep = sleepDef;
bool arttir = false;

void setup() {
  pinMode(13,1);
  pinMode(12,1);
  pinMode(11,1);
  pinMode(10,1);
  pinMode(9,1);
}

void loop() {
  //digitalWrite(13,0)
  for (int i=9;i<12;i++){
    digitalWrite(i,1);
    delay(sleep);
    digitalWrite(i,0);
  }

  for (int i=12;i>9;i--){
    digitalWrite(i,1);
    delay(sleep);
    digitalWrite(i,0);
  }
  if (sleep >= sleepDef){
    arttir = false;
  }
  if (arttir){
    sleep++;
    return;
  }
  
 sleep--;
 if (sleep <= 0){
   arttir = true;
     for (int i=9;i<13;i++){
    digitalWrite(i,1);
  }
   delay(2000);
  
    for (int i=9;i<13;i++){
    digitalWrite(i,0);
  }
 }
  
}