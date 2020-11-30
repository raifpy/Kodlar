class client{
    public static void main(String[] args) {
        print(server.getSaat()); // server.java 'yı javac ile .class hale getiriyoruz !
        // 03.45 String döndü en sonunda !

        // Şimdi saati değiştirelim > Için nesneye atayalım

        server sr = new server(); // 

        sr.setSaat("03.59");

        print(server.getSaat()); // Görüldüğü üzere static tarih_saat setSaat ile değişmiş (public void setSaat(String veri)

        print(sr.getSaat()); // static elemanı non-static ile çağırmaya çalıştığım için uyarı veriyor .. || Değer aynı

        for (String i :server.getArray()){
            System.out.println(String.format("Değer %s <--",i));
        }


    }
    public static void print(String text){
        System.out.println(text);
    }
}