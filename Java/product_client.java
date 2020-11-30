public class product_client{
    public static void main(String[] cuk){
        System.out.println("Merhaba");
        
        product pro = new product();
        System.out.println(pro.getName());

        pro.veri_yaz("Veri yazdan edinildi !");
        
        System.out.println("-------------");
        System.out.println(pro.getName());
        
    }
}