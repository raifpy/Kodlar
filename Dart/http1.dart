import "dart:io";
import 'dart:convert';

void main() {
  mainSync();
}

void mainNormal() {
  // iç içe fonksiyon yapısı sayesinde program http isteğinin atılmasını beklemeden devam edecek
  HttpClient()
      .openUrl("GET", Uri.parse("http://ip-api.com/json"))
      .then((request) => request.close()) // isteği yolladı
      .then((response) => response.transform(Utf8Decoder()).listen(print));
  //.then((r) => r.cancel());
}

// async version;

void mainSync() async {
  var request =
      await HttpClient().openUrl("GET", Uri.parse("http://ip-api.com/json"));
  var response = await request.close();
  await for (var abc in response.transform(Utf8Decoder())) {
    print("*  : " + abc);
  }
}
