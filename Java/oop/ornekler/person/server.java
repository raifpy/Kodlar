public class server{
    public static void main(String[] args) {
        System.out.println("'client' dan çalıştırın . ");
        System.exit(0);
    }
    private String yas;
    private String isim;

    public server(){
        yas = "Tanımlanmadı";
        isim = "Tanımlanmadı";
    }

    public server(String yas,String isim){
        this.yas = yas;
        this.isim = isim;
    }

    String[] PersonArray = {yas,isim};

    public void setName(String name){
        this.isim = name;
    }
    public void setAge(String age){
        this.yas = age;
    }


    public String getName(){
        return this.isim;
    }
    public String getAge(){
        return this.yas;
    }

    public String toStr(){
        return String.format("NAME : %s | AGE  : %s", isim,yas);
    }

    public static String listele(server nesne){
        return String.format("YAŞ  : %s | Isım : %s", nesne.getAge(),nesne.getName());
    }
}