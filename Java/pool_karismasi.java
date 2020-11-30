public class pool_karismasi{
    public static void main(String[] args){
        String str1 = "Merhaba";
        String str2 = new String("Merhaba");
        String str3 = "Merhaba";

        System.out.println(str1);
        System.out.println(str2);
        System.out.println(str3);

        System.out.println(str1==str2);
        System.out.println(str1==str3);

        
    }
}