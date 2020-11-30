void main(){
  List eleman = new List(2); // 2 null elemanlaı bir oluşturmuş olduk ..
  eleman.add("Merhaba 1"); // Hata
  eleman.add("Merhaba 2"); // Hata
  print(eleman);

  // Hata verecek . .add yapısı ile fixed listeleri eleman ekleyemiyoruz . [0] , [1] ile ekleyebilirsiz ..
}