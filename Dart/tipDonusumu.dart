

void main(){
  var abc = "10";
  int abcInt ;

  // abcInt = abc; // Tabiki hata verecek abey
  abcInt = int.parse(abc);
  

  //String abcdfgh = "$abcInt"; // Çok zekice :D
  String abcdfgh = abcInt.toString();
  
  print(abcInt); // Doğrulama için assert
  
  print(abcdfgh+" : Evet bu bir string"); // o ye

  // double to str yapablım abey

  double duble = 10.1010;

  print(duble.toString());
  print(duble.toStringAsFixed(1)); // güzel düşünülmüş  | .'dan sonra kaç eleman gözüksün diye soruyor

  // tipleri edinelim

  print("abc : "+abc.runtimeType.toString());
  print("abcInt : "+abcInt.runtimeType.toString());
  print("abcdfgh : "+abcdfgh.runtimeType.toString());
  

}