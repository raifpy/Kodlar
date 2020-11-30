public class equelsdeneme{
    public static void main(String[] args) {
        
        String bir = "Merhaba";
        String iki = new String("Merhaba");
        
        System.out.println(bir == iki);
        System.out.println(bir.equals(iki));

        String bir_bir = "Merhaba";

        System.out.println(bir == bir_bir);

        System.out.println(iki == "Merhaba");
        

    }
}
