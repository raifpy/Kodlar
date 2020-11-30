class operatorler{
    public static void main(String[] args) {
        String bir = "Mer";
        byte bir_len = (byte) bir.length();

        if (bir != "Merhaba"){
            println("Merhaba değil !\n\n");
        }

        if (bir == "Merhaba" && bir_len > 5){
            println("Tüm gereklilikleri karşılıyor");
        }

        else if (bir == "Merhaba" || bir_len > 5){
            println("Bazı gereklilikler çalışmıyor , else if çalıştı");
        }
        else{
            println("Hiçbiri olmuyor , else çalıştı.");
        }
    }

    public static void println(String mesaj){
        System.out.println(mesaj);
    }
}