import java.util.Scanner;

public class TekSatirIfElse{
    public static void main(String[] args) {
        Scanner sc= new Scanner(System.in);
        System.out.print("SAYI : ");
        int a = sc.nextInt();
        sc.close(); // Scanner'ı kapattık
        
        if(a == 9) System.out.println("Sayımız 9 'dur");
        else if (a==8) System.out.println("Sayımız 8 .");
        else System.out.println(String.format("%s Sayısı bertilen durumlar ile uyuşmuyor",a));
    }
}