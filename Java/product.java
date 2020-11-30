public class product{
    public static void main(String[] args) {
        System.out.println("Maınden gelen eleman");
    }
    private String veri="Veri";

    public void veri_yaz(String mesaj){
        System.out.println(mesaj);
        this.veri = mesaj;
    }

    public String getName(){
        return this.veri;
    }

}