void main(){
  List liste = ["Merhaba","Dünya",true]; // farklı türlerde veri eklerken hata almayız şayet List nesnesi tek başına dnnymic * halde
  // yani bu liste'ye her türden veri ekleyebilirsiniz

  List<String> liste_str = ["Merhaba","Dünya"]; // Bu eleman ise sadece String alabilir
  liste_str.add("string değer eklerim paşa");
  //liste_str.add(1); // Hata verir

  List<int> liste_int = [1,2,3,4,5,6,7];
  liste_int.add(10);
  //iste_int.add("ONBEŞ"); // Hata verir

  // Yani List 'e tüm veri türlerini ekleyebilirken List<tür> 'e sadece tür eklemesi yapılabilir

  print(liste_int);

  



}