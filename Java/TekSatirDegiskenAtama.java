import java.util.Scanner;

public class TekSatirDegiskenAtama{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        
        System.out.print("Mesaj gir : ");
        String ifade = sc.nextLine(); // inputdan veri alıp bunun ile karşılaştırma yapacağız

        sc.close();



        String veri = (ifade.toLowerCase().equals("acaba")) ? "Acaba yok . Olacak , oluyor bile . ":String.format("%s Verisini girdiniz .",ifade);
        System.out.println(veri);

        // degiskentürü degiskenadı = (boolean_doner) ? "true ise dönecek veri":"false ise dönecek veri";
        // Güzel , pratik . Şans vermeye değer 

        String veri2 = (true) ? "True ise dönecek veri": "False ise dönecek veri"; // İşin sonunda veri2 ya true ise yada false ise olacak
        System.out.println(veri2);

        // true değerini verdiğimiz için sonuç her türlü true ise dönecek bu yüzden ide bu kod ölü , bu blok asla çalışamz diyor
        
        
    }
}