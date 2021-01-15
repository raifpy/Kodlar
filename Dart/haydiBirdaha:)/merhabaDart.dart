import 'dart:math';

void main() {
  print("Merhaba Dünya!");
  print(dinamikFonksiyon());
  print(r"'\' devre dışı bırakıldı"); // Tıpkı Python'daki gibi
  print(statikFonksiyon());
}

dynamic dinamikFonksiyon() {
  var rand = new Random().nextBool();
  switch (rand) {
    case true:
      {
        print("true");
        return "Berber";
      }
      break;

    case false:
      {
        print("false");
        return 99;
      }
      break;

    /*case "aslaOlmayacakBirDeğer": // 
      print("nasıl olur yav :=)");
      return false;
      break;
      */
    default:
      {
        print("asla olmayacak değer 2");
        return 9.1;
        // break; // Gereksiz
      }
  }
}

int statikFonksiyon() {
  //return "J";//
  return Random().nextInt(99999);
}
