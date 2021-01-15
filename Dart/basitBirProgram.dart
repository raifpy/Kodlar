import 'dart:io';

void main(List<String> args){ // dart program.dart abc abc2
/*
argümanın'ın ilk elemanı program.dart olması gerekirken abc oluyor.
Bu durumda dart kendi argümanını listeye eklemiyor. Dikkat edilmesi gereken bir olay

Python:
sys.argv[0] == <betikAdi>.py

Go: 
os.Args[0] == <programAdi>

Dart:
args[0] == <argüman>sanaa 

*/
  if (args.length == 0){
    stderr.write("basitBirProgram: <argüman> ...");
    return;
  }
  args.forEach((element) {
    print("Argüman - "+element);
  });

  if (args.first.split(" ").length > 1) { // dart basitBirProgram.dart "merhaba dünya"
      print("İlk arguman birden fazla argumanın birleşimi. | split-length"); 
      // args listesinin ilk elemanını " " 'dan bölerek başka bir liste haline getirdi ve length'i kontrol etti
  }
  
  if(args.first.indexOf(" ") > -1){ // ' ' yok ise -1 dönecek
    print("İlk arguman birden fazla argumanın birleşimi. | indexOf ' ' ");
    // args listesinin ilk elemanınındaki bütün ' ' 'ları arayarak kaç tane olduğunu döndürdü ve dönen sayıyı kontrol etti { kaç tane olduğunu değil, ilk ' 'ın hangi indexte olduğunu döndürdü. Yine de işimize yarar }
  }
  var liste = [
    "Merhaba Dünya!",
    if (args.length == 1) "1 tane Argüman var" else "1'den fazla Argüman var",
    ...args,
    for (String veri in args) if (veri != "argüman") "Peki, geç $veri",
  ];

  print(liste);
  print("\n");
  print("Çalıştırılan: "+Platform.executable);
  print("env listeleniyor: ");
  var env = Platform.environment;
  env.forEach((key, value) { 
    print(key.toString()+" : "+value.toString());
    sleep(Duration(milliseconds: 500));
  });
}