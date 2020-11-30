main(){
// throw ile yapıyoruz :')
  int a = 110;

  if (a > 10){
    throw FormatException("Bu ne bilmiyorum ama hata hatadır .");
  }
  else if (a<10) throw "Merhaba Dünya bu 10'dan küçük"; // bu eleman ile Exception() elemanı aynı !
  else throw Exception("2 farklı şekilde kullanılabiliyor ..");
}