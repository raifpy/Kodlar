void main(){
  String verBanaTextSatirIci() => "AlSanaText";
  String verBanaTextSatirIci2(){
    return "MerhabaDünya";
  }


  var buDaVar = (String text) => "Merhaba, $text!";

  print("Hatırlayalım");
  String abc = verBanaText();
  print(abc);
  print(verBanaText2());
  print(verBanaTextSatirIci());
  print(buDaVar("oye"));

  print(fonksiyon()());

  
}

String verBanaText(){
  return "AlSanaText";
}

String verBanaText2() => "AlSanaText";


Function fonksiyon(){
  // Bir fonksiyon döndürür!
  
  //return () => "MerhabaDünya";

  bool abc(){
    // Çeşitli işlemler
    return true;
  }
  return abc;

  /*return (){
    // çeşitli işlemler
    return true;
  };*/ //Aynı eleman
}