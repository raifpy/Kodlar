import java.util.InputMismatchException;

public class herdefasindabirartaraktopla{
    public static void main(String[] args) {
        java.util.Scanner sc = new java.util.Scanner(System.in);
        int max_sayi=0;
        
        try{
            System.out.print("Bir Rakam Gir Dostum : ");
            max_sayi = sc.nextInt();
        } catch (InputMismatchException e){System.out.println("Sadece sayı ! Hata ( "+e+" )\n");System.exit(100);}
        
        System.out.println(String.format("RAKAM : %s \n", max_sayi));
        
        for (int i=0,toplam=0;i < max_sayi;i++){
            toplam += i;
            System.out.println(String.format("Toplam : %s ",toplam));
        }
    }
}


// Bence oldu :D