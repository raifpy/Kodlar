void main(){
  sayMerhaba();
  sayString("Tüh");

  codeksiyon modul = new codeksiyon();
  print(modul.getText());
  modul.setText("Osuruk kemiksiz kakadır yiğen");
  print(modul.getText());


}

void sayMerhaba(){
  print("Merhaba");
}

void sayString(String text){
  print(text);
}


class codeksiyon{
  String text = "Merhaba Codeksiyon";
  String getText(){return this.text;}
  void setText(String text){this.text = text;}

}