public class server{
    public static void main(String[] args) {
        System.out.println("client 'dan çalıştırın !");
        System.exit(0);
    }

    private static String tarih_yil="2020"; // private başka sınıfdan doğrudan erişilememsi için , static ise static tanımalamada kullanılabilsin , return edilebilsin diye ...
    private static String tarih_ay = "05"; //     //
    private static String tarih_gun = "23"; //    // 
    private static String tarih_saat = "03.47"; //    // 

    public static String getYil(){
        return tarih_yil;
    }
    public static String getAy(){
        return tarih_ay;
    }
    public static String getGun(){
        return tarih_gun;
    }
    public static String getSaat(){
        return tarih_saat;
    }
    public static String[] getArray(){
        String[] abc = {tarih_saat,tarih_gun,tarih_ay,tarih_yil};
        return abc;
    }

    // public static || static || yaparak nesneye atamadan doğrudan return edilebilmesini sağladık
    // ama tarihler değişecekse bir zahmet nesneye atansın :)

    public void setYil(String veri){
        this.tarih_yil = veri; // static veri olduğu için uyarı veriyor | Mühim değil
    }
    public void setAy(String veri){
        this.tarih_ay = veri; // static veri olduğu için uyarı veriyor | Mühim değil
    }
    public void setGun(String veri){
        this.tarih_gun = veri; // static veri olduğu için uyarı veriyor | Mühim değil
    }
    public void setSaat(String veri){
        this.tarih_saat = veri; // static veri olduğu için uyarı veriyor | Mühim değil
    }

    // Buradan öğrenmiş olduk ki static elemanı non-static elemanda kullanmaya başlayınca tartışma çıkıyor . En iyisi hepsini static yapmak yada hepsini non-static yapmak ..
    





}