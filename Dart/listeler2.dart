void main(){
  List liste = ["Merhaba",true,"Dünya",false,10,10.0,["Merhaba","Bu bambaşka bir liste .."]];

  // List liste = new Liste(6) yapısını kullanmak zorunda değiliz . List liste = [, , , ,]; Yapısı ile çok daha basit devam edebiliyoruz ..
  print(liste);

  // Ama bu yapıyı for - each ile (in yapısı ile ) kullanmaya çalışırsak hata verecek şayet içerisinde aynı tür olmayan değerler var ..

  /*for (String i in liste){
    print(i);
  }*/

  //String olmayan veriyi String değere atamaya çalışıyoruz çünkü ..

  List liste2 = ["Merhaba","Dünya","Nasılsın ?"];

  for (String i in liste2){
    print(i);
  }
  // Burada herhangi bir hata alamayız çünkü tüm veirleri zaten String türünden 

  print(liste2);

}