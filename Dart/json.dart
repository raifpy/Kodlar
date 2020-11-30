main(){
  var json = {"Merhaba":"Dünya"};
  print(json["Merhaba"]);

  // Map olarak geçiyor :

  Map json2 = {1:2};
  print(json2[1]); // Python ile aynı neredeyse

  json2["Hey"] = "Umarım olur";
  print(json2);

  // Aynı Python'daki gibi ..

  // birde hata alalım :

  print(json2[10]); // Olmayan bir keyi çağırmaya çalışırsak : > null dönüyor

  // Olmayan neseneyi çağırmak hata vermiyor , null dönüyor

  print(json["heyt"]);
  // json'un value'si null ise sorun olacak gibi :)

  print(json2.containsKey("olmayan key"));
  print(json2.containsKey("Hey"));
  print(" key | value |");
  print(json2.containsValue("Olmayan bir value"));
  print(json2.containsValue("Umarım olur"));

  // containsKey yapısı ile anahtarın var olup olmadığını kontrol edebiliyormuşuz

  print(json2.keys); // Tüm keyleri dönüyor (1,Hey)
  print(json2.values); // Tüm valueleri dönüyor
  print(json2.length); // len'i dönüyor 

  print(json2.toString());

  json2.remove(1); // json2 'nun içinden 1 keyini ve valuesini sildik , 

  print(json2.toString()); // Kontrol edelim

  print("----");
  print(json);
  json.clear(); // json'un tamamını sıfırladım
  print(json);

  print(json2.isNotEmpty); // liste yapısında da var .| string'de de
  print(json.isNotEmpty); // isEmpty 'de var aynı şekilde

  print(json2.runtimeType); // Güzel yapı :)
  
  
}