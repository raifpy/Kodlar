public class static_client{
    public static void main(String[] args) {
        static_server.gui(); // static_serverçjava 'yı bulabilsin diye statşic_server.java'yı javac ile .class haline getirdim !
        // sınfı neseneye atamadan gui'i çağırabildim
        // çünkü gui() static method
        // static_server sc = new static_server(); isc.gui(); ile de çağırabildim

        // static_server.dondur(); //Hata verecek . Çünkü static değil && nesneye atamadık

        static_server sc = new static_server();
        String uwu = sc.dondur();
        System.out.println(uwu);

        sc.gui(); // Uyarıyı verdi ama sıkıntı yok nesneye atayıp da kullanabiliriz ..
        
    }
}