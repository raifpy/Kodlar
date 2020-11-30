class Merhabalar{
  String mesaj = "Merhaba Dart OOP";
  
  void printMesaj(){print(mesaj);}
  void setMesaj(String mesaj){this.mesaj = mesaj;}
  String getMesaj(){return mesaj;}
}

main(){
  Merhabalar abc = new Merhabalar();
  abc.printMesaj();
  abc.setMesaj("Merhaba Mesaj");
  abc.printMesaj();
  print(abc.getMesaj());

}


// Javadaki ile benzer bir sınıf düzeni var. Tabiki Java tamamen OOP bir dil olduğu için program sınıf içerisinde başlıyordu .
// Dart ise C gibi main fonksiyonundan başlıyor  ..

// fonksiyonları void yerine String , int deseydik birşeyler return etmek zorunda olacaktık . return edip edilmeyeceği belli değil ise dynamic yapısını kullanabiliriz :
/*

dynamic fonkisyon(){
  if (true){
    return "Hey"
  }
}

*/

// 