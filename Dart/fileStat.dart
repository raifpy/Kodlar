import 'dart:io';

void main() {
  var dosya = File("/etc/os-release");
  print(dosya);
  print("Path: " + dosya.path);
  print("Absolute: " + dosya.isAbsolute.toString()); // absolute = mutlak

  dosya.exists().then((bool ok) {
    print("normal: " + ok.toString());
  });

  bool dosyaVarmi = dosya.existsSync();
  print("async: " + dosyaVarmi.toString());

  var stat = dosya.statSync();
  print(stat.accessed.toString());
  print(stat.changed.toString());
  print(stat.mode);
  print(stat.modified.toString());
  print((stat.size / 1024).toString() + " kb");
  print(stat.modeString());
  print(stat.toString()); // Güzel çıktı
}
