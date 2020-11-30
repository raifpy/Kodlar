main(){ // @raifpy
  // var ve tür olarak tanımlamayı biliyorduk zaten

  var isim = "Merhaba";
  String isim2 = "Dünya";

  print(isim+" | "+isim2);

  // Peki isim 'i (var-String) başka bir tür yapmak istersek ne olur ?

  isim = "Hey"; // Bunda elbette problem yok .

  // isim = true; // Hata : Error: A value of type 'bool' can't be assigned to a variable of type 'String'
  
  // Aynı şekilde isim2^de hata verecekti .


  dynamic nesne = "Merhaba Dost .";
  print(nesne); 

  nesne = true;

  print(nesne); //

  // ama var yerine dynamic yapısını kullanarak tür değişikliği yaparsak herhangi bir problem olmuyor :)

  // 

  // dynamic 
  // var
  // tür
   // Python değişkenleri dynamic yapısını kullanıyormuş .




}