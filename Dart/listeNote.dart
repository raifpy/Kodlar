import 'dart:io';

void main(List<String> list){
  if (list.isEmpty){
    stderr.writeln("Argüman listesi boş!"); // Ne değişti? | Unix sistemlerde komuttan sonra | > gibi kavramlar ile yardımcı programlar çalıştırılır. 
    /* Örnek:
    dart merhabaDart.dart > dosya
    dart merhabaDart.dart 'ın çıktısı "Merhaba Dünya" olduğunu varsayarak
    bu çıktıyı dosya adlı dosyaya kayıt etmiş olduk.

    stderr ile yaptığımız çıktı ise bu kayıt işlemini geçersiz bırakır. *kendini kullandırtmaz
    */
  }
}