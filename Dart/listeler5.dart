void main() {
  var eleman = new List<String>(10);
  print(eleman);
  print(eleman.isEmpty); // false

  eleman[eleman.length - 3] = "Merhaba Dünya";

  //eleman.add("hüü");
  //eleman.clear();
  eleman.asMap().forEach((int index, String value) {
    //
    print(index.toString() +
        " : " +
        value
            .toString()); // Zaten String olanı neden String yaptık?: String value String ya da null veri türünü alabilir. null veri türünde de toString() olduğu için hata almadan String'e çevirmiş olduk
  });
}
