void main(){
  List liste = new List(5);
  // 5 elemanlı bir liste oluşturduk
  //
  // Object[] liste = new Object[5]}; // Umarım yanlış değildir :D | Unuttum Java karşılığını :D

  liste[0] = "Merhaba";
  liste[1] = "Dünya";
  liste[2] = true;
  
  liste[3] = 10;
  

  print(liste); // liste4 eklenmediği için null olarak çıkıyor .
  // [Merhaba, Dünya, true, 10, null]

  // List ile tanımlanan elemanlar tüm veri tiplerini tutabiliyor | Java - Object[]

}