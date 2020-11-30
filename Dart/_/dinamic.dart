// Dinamic veri tipi 
void main(){
  var veri = funk("abc");
  print(veri);


  // golang interface gibi düşünülebilir...
}

dynamic funk(String key){
  if (key == "abc"){
    return 10;
  }else if (key == 10){
    return "abc";
  }else{
    return false;
  }

}

/*
func funk(key string) interface{}{
  // devam
}
*/