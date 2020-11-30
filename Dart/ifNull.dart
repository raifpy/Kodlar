// obje yok ise hata verecek doğal olarak , ama ?. yapısı ile bunu aşabiliriz

class codeksiyon{
  var abc = "Merhaba Codeksiyon";

}
void main(){

  var n;

  //print(n.runtimeType);
  
  //print(n.deneme); // çökecek şayet eleman yok
  print(n?.deneme); // aynı eleman yine yok ama ? yapısı , eleman yok ise null dönmesine yarıyor
  // dönen elemanın null değilde başka bir şey olmasını istiyorsak ;

  var iim = codeksiyon();
  print(iim.abc);
  print(iim?.abc ?? false); // eğer iim nesnesinde abc yoksa false döndür , abc varsa abc'yi döndür
  print(n?.abc ?? false); // eleman yok ise null dönmek yerine false dönecek

  // Güzel yapı ayrı mesele :)
}