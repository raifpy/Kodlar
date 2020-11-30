import java.util.Scanner;

public class issame2 {
    public static void main(String[] args) {
        boolean same = isim();
        if (same){
            print("Yum There are Same");

        }

        else if (same==false){
            print("Ymu There are not Same");
        }
    }
    
    public static void print(String veri){
        System.out.println(veri);
    }

    public static String input(String veri){
        Scanner sc = new Scanner(System.in);
        System.out.print(veri);
        veri = sc.nextLine();
        sc.close(); // kapat , hata verecek
        return veri;
    }

    public static boolean isim(){
        print("Merhaba");
        String a =input("ISIM : ");
        String b =input("ISIM : ");
        
        if (a==b){
            return true;
        }
        else{
            print(a);
            print(b);
            return false;
        }
    }
}