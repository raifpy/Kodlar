import "package:http/http.dart" as http;

/*
Paketi import edebilmek için 'pubspec.yaml' hazırlamanız gerekiyor;

name: <KlasörAdi>
dependencies:
  http:

Ardından terminale 'pub get' yazmanız yeterli olacaktır.
*/

void main() {
  print("sync çağırılıyor");
  mainSync();
  // sync fonksiyon çağırılığı zaman bekletmiyor
  print("normal çağırılıyor");
  mainNormal();
}

void mainSync() async {
  var response = await http.get("http://ip-api.com/json");
  print("StatusCode: " + response?.statusCode.toString() ?? "statuscode yok?");
  print(response.body);

  print("İşlem Bitti {async}");
}

void mainNormal() {
  http.get("http://ip-api.com/json", headers: {"User-Agent": "blabla"}).then(
      (response) => print("StatusCode: " +
              response.statusCode.toString() +
              "\n" +
              response?.body ??
          "body yok?"));
  print("İşlem Bitti {normal}");
  // Async'den farkı; async son satıra gelmek için istek atmasını beklerken .then ile yazılan fonksiyon isteği arkaplanda atıp fonksiyonunu çağırıyor
}
