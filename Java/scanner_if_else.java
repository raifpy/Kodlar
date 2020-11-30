import java.util.Scanner;
public class scanner_if_else {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.print("Bir sayı gir : ");
        int sayi = sc.nextInt();
        if (sayi==1024){
            System.out.println("1024 1024");
        }
        else if (sayi==2048){
            System.out.println("2048 2048");
        }

        else if (sayi==0){
            System.out.println("Sıfır exit'ı tetikler ..");
            System.exit(1);
        }
            

        else{
            System.out.println(sayi+" SAYI "+sayi);
        }

        sc.close();
    }
    
}
