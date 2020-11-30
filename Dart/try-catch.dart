main(){
  print("Heyt");
  // hataver(); // Merhaba Dünya Hatası böylece gelmiş oluyor 

  try{
    hataver();
  }
  catch (Exception){ //  throw "Merhaba Dünya Hatası"; şeklindeki hatalar Exception olarak geçiyor 
    print("Merhaba Dünya Hatası Bro "); // Normalde programın çökmesi gerekiyor ama catch ile çökmeyi engelledik
  }
  
}

void hataver(){
  throw "Merhaba Dünya Hatası";
}
