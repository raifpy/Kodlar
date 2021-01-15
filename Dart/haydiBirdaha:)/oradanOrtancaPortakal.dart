import 'dart:io';

class biftek {
  static int rusvet = 0;
  String _cesit;
  bool _mevcut;
  biftek(this._cesit) {
    if (this._cesit == "" && biftek.rusvet == 0) {
      throw "Ne istiyorsun kardeşim?";
    }
    if (biftek.rusvet != 0) {
      print("Oooo ${Platform.localHostname} abim gelmiş. Hoş geldin");
    }
  }

  bool varMi() {
    if (_cesit == "çitos") {
      this._mevcut = true;
      return true;
    }
    this._mevcut = false;
    return false;
  }

  String verFiyat() {
    if (!this._mevcut) {
      return "";
    }
    return "10 tıl";
  }
}

void main() {
  print("Biftek:");
  String veri = stdin.readLineSync();

  //biftek et = ....;//
  biftek.rusvet = 100;
  var et = biftek(veri);
  if (!et.varMi()) {
    print("üzgünüm elimizde kalmamış!");
    exit(1);
  }
  print("${et._cesit} fiyatı: " + et.verFiyat());
}
