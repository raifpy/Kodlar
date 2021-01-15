import "dart:html";

void main() {
  HttpRequest.getString("http://ip-api.com/json").then(print);
}
