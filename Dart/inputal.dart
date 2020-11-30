import 'dart:io';

void main(){
  print("Merhaba Şef ; Adın ne ? ");
  String sefAdi = stdin.readLineSync();
  if (sefAdi == ""){
    print("Ayıp oluyor efendi !");
  }else{
  //print("Merhaba "+sefAdi+" şef ! ");
  print("Merhaba $sefAdi şef !"); // $varName , güzel yapı
    }
}

// ulan gofmt neredesin ...