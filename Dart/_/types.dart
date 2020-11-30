
import 'dart:io';


void main(){

  final int sayi = 10;
  final kelime = "Merhaba";
  //final String kelime = "Merhaba";

  const sayiC = 10;
  const String kelimeC  ="Merhaba";

  // final veya const ile tanımlanan veri türleri bir kez değer alabilir.

  //Aralarındaki fark kısaca final daha oop'de kullanılabilirken const daha düşük yerlerde kullanılır**




  
  
  
  
  var sayi2 = 10;
  var kelime2 = "Merhaba";

  // Golang,JS,Java benzeri

  dynamic dinamik;
  dinamik = 10;
  dinamik = "Merhaba Dünya -raifpy";
  dinamik = true;
  
  // Tüm veri türlerini içerebilen eleman. (Golang; interface{} Python; normal-veri-türü)

  print(dinamik); // true

  var foo = const["Merhaba Dünya",10,false,10.05]; // dinamik ama const liste (liste içeriği değiştirelemez)
  var foo2 = []; // foo 'dan farkı liste içeriğinin değiştirilebilir olması.

  // foo.add("gobr"); // çökmek 
  
  // Python'a benziyor.
  
  //foo = 10; //Yapılamaz: List<Dynamic> veri türü Dynamic veri türü değildir!

  for (var item in foo) { // Python; for item in foo:  Golang; for _,item := range foo{ |Java for-each yapısına oldukça benzer (":" yerine "in")
    print(item);
  }

  List<String> liste1 = ["Merhaba","Dünya"];
  List<String> liste2 = ["Merhaba","Dart",...liste1];
  if (liste2.length != 4){
    print("Olmaması gereken bir şey oldu!");
    exit(1);
  }
  print(liste2);

  /*while (true){
    print("h");
  }*/

  //List<String> liste3;
  var liste3;

  
  // liste2 = [...liste2,...liste3]; // null eleman olduğu için hata veriyor
  liste2 = [...liste2,...?liste3]; // null elemanı göz ardı ediyor! | ignore not list type (null)
  
  print("u");
  print(liste2);

  liste1.add("Liste 1'e sonradan eklenen veri");
  liste1.remove("Merhaba");

  print(liste1);

  // liste içerisinde if;
  var ok = true;
  List<String> raifpy = [
    "Merhaba",
    "Codeksiyon",
    if (ok) "raifpy",
    if (!ok) "t.me/raifblog",
  ];

  print("list with if : "+raifpy.join(" "));

  List<String> ehehehe = [
    "Merhaba",
    for (var i in liste2) "$i-ü",
    "Dünya",
  ];

  print(ehehehe);

  Map<String,String> json = {"Merhaba":"Dünya"};
  print(json["Merhaba"]);

  // Map = Map<Dynamic,Dynamic>

  var json2 = {
    "a":"b",
    "c":"d",
  };
  // String,String

  var json3 = {
    10:"Merhaba",
    "Merhaba":10,
  };
  // Dynamic,Dynamic
  print(json3["olmayanveri"]); // null
  print(json3.containsKey("olmayanveri")); // false
  
  json3["olmayanveri"] = "Artık Var!";
  
  print(json3["olmayanveri"]); // Artık Var!
  print(json3.containsKey("olmayanveri")); // true

  var lamba = "Merhaba Dünya";
  String son = lamba.substring(lamba.length-1);
  print(son);
  
  






  



  
  
}