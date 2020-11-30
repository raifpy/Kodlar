main(){
  String eleman = "Kodlar benim lna çalmayın :D (raifpy) ";

  List liste = new List(eleman.length); // String türünün kaç harfdan oluştuğunu lenght ile alabiliyoruz . 
  // eleman 'ın harfleri kadar bir liste oluşturmuş oldum
  
  for (int i=0;i<eleman.length;i++){
    liste[i] = eleman[i];
  }

  print(liste);
  
}