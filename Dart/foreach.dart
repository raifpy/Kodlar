main(){
  List liste = "Merhaba Dünya , Bu da .split() yapısı | Java'dan geçme :)(".split(" "); // " " ile boşlukları silmiş olduk
  
  // liste 'nin bütün elemanları String olduğu için aşağıda problem olmaz 

  // ama doğru kullanım List<String>
  for (String i in liste){ // Java'da in yerine : koyuyorduk :)
    print(i);

  }
}



// Daha karışık olan for (int i) 'li default for döngüsü yerine daha pratik olan for (String i in liste) yapısını kullandık.abstract
/// Bu yapı Java'da for (String i:liste) şeklinde kullanılıyor ..