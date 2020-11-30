void funk(String text){
  print("Ç : "+text);
}

void main(){
  var liste = "String .split fonksiyonu aracılığı ile bu stringi list yapacağız.".split(" ");
  liste.forEach(funk);

  liste.forEach(
    (eleman) => print("Ş $eleman")
  );

  liste.forEach((String element) { // sadece element 'de yazılabilir
    print("Ğ! $element");
  });

  for(var i in liste){
    print("Ü : $i");
  }



}