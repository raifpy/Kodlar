class sinif {
  static int staticRakam = 0;
  final String kelime;
  sinif(this.kelime);
  String get repeatKelime => kelime * staticRakam;
  void set rakam(int rakam) {
    sinif.staticRakam = rakam;
  }

  Map<String, Object> get map {
    return {
      "staticRakam": sinif.staticRakam,
      "kelime": this.kelime,
      "repeatKelime": this.repeatKelime,
    };
  }
}

void main() {
  sinif.staticRakam = 2;
  var eleman = new sinif("Merhaba Dünya! | ");
  print(eleman.kelime);
  print(eleman.repeatKelime + sinif.staticRakam.toString());

  eleman.rakam = 5;

  print(eleman.repeatKelime + sinif.staticRakam.toString());

  var json = eleman?.map;

  print(json);

  print(json.cast());

  print(json.toString());

  json.forEach((key, value) => print(key + " : " + value.toString()));
}
