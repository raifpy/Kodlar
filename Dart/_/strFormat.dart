void main(){
  int sayi = -10;
  print("Merhaba, Sayın : $sayi");
  print("abs & 2'ye bölümü : ${sayi.abs() / 2}"); // float dönüyor, bu şekilde int yapabiliyoruz (float).toInt()
  print("2'ye bölümü (int) : ${(sayi / 2).toInt()}"); // Bu şekilde de int olabilir ama!
  print("2'ye bölümü (int) : ${sayi ~/ 2}"); // editörün bana verdiği bilgiyi kullanarak / .toInt() yapmak yerine ~/ yapmanın daha iyi olduğunu öğrendim . 
  // ~/ float yerine int dönüyor.
  
}