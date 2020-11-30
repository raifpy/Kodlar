import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class dosya_islemleri5 {
    public static void main(String[] args) throws FileNotFoundException{
        String dosya = "dosya_islemleri5.java"; // Bu dosyayı okuyacağız ..
        File file = new File(dosya);
        Scanner sc = new Scanner(file);
        //int satir_sayisi = sc.nextInt();
        //System.out.println(satir_sayisi+"\n");

        while (sc.hasNextLine()){ // Başka satır var ise döngüye devam
            String yazi = sc.nextLine();
            System.out.println(yazi);
        }

        
    }
}


// Scanner kullanarak dosyayı okumak ..

// Scanner(System.in) yerine file nesnesini veriyoruz