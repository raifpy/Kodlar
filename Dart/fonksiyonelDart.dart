class fonksiyon{
  int a = 10; // fonksiyonel değil
  int b() => 10; // fonksiyonel oye gobr
}

void main(){
  fonksiyon fn = new fonksiyon();
  fn.a; // fonksiyonel değil
  fn.b(); // fonksiyonel
}

/*Ne fark var?
a tanımı bir değişken. Sınıf oluşturulurken bellekte belli bir yer kaplıyor.
b tanım ise bir fonksiyon. Sınıf oluşturulurken bellekte belli bir yer kaplamıyor.
Senaryomuzda a adlı tanımı asla kullanmayacağımızı varsayalım. Bu durumda a = 10; belletek gereksiz yer kaplıyor olacaktır.
Fakat b kullanılsa da kullanılmasa da bellekte yer kaplamayacaktır.
*/