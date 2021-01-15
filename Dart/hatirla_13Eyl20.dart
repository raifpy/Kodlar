//Bu dil en çok Java'ya benziyor . Değişkenler olarak Python'a , fonksiyonlar olarak da C'ye.

void main(List args) {
  print("merhaba");
  if (args.length == 0) {
    print("Olamaz dostum , bana arguman vermen gerek !");
  } else {
    print(args);
  }

  String degiskenstring = "Merhaba Ben String değişken! Sadece String'im ...";
  var degiskenvar =
      "Merhaba Ben de String değişken! Ama var ile tanımlandığım için veri tipim, ilk atama ile sabit kalmak şartı ile değişebilir ";
  dynamic degiskendynamic =
      "Merhaba Ben Dynamic! Şuan da String'im ama veri türüm istenilen zaman  istenildiği gibi değiştirilebilir [Python gibi]";

  var listeEleman = [
    degiskenstring,
    degiskenvar,
    degiskendynamic
  ]; // manuel list
  List listeAmaBos = new List(10); // 10 null 'lu bir liste yaptık .
  List<String> listeAmaSadeceString =
      new List(10); // 10 elemanlı , sadece String alabilen liste yaptık
  print(listeAmaSadeceString); // null
  //listeAmaSadeceString.add("Merhaba Dünya"); // New List() ile olşuturulan listeler fixed size'dır . Ekleme yapılamaz ! ( Go - unslices (lists))

  print(listeEleman.join("\n")); // for

  for (int i = 0; i < 5; i++) {
    // for
    print(i);
  }

  for (String i in listeEleman) {
    // for-each
    print("for-each : " + i);
  }
  print("map oluşturuluyor şef");
  Map benMapim = {"a": "in ali bin ali"};
  benMapim["c"] = "Codeksiyon :)";
  print(benMapim);
  print("c çağırıldı : " + benMapim["c"]);
  print("Olmayan çağırılıypr : ");
  print(benMapim["yokki"]);

  // olmayan eleman null dönüyor

  var abc = "raifpy";
  print("\nabc'ye 'raifpy' değeri verildi , switch 'e giriliyor");
  switch (abc) {
    case "raif":
      print("Tım bu raif");
      break;
    case "raifp":
      print("Tım bu çakma raif");
      break;
    default:
      print("Olamaz hiç biri değil");
      break;
  }

  try {
    throw "Hüüü bu bir hata";
  } catch (Exception) {
    print(
        "Olamaz , program çöktü :( Neyseki catch ile çökmesi engellendi o ye");
  }

  print("Hatırla raif bitti .");

// GoLang 'den sonra "gofmt" nin ne kadar kullanışlı olduğunu bir kez daha anladım.
// Genel olarak zor bir dil değil Dart. Import edilen modüllerin, neyi import ettiğini bir bakışta söyleyemiyoruz; bana göre bu bir eksi
}
