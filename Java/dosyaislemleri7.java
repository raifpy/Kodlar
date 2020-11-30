import java.io.BufferedWriter;
import java.io.File;

import java.io.IOException;
import java.io.FileWriter;
import java.util.Scanner;

public class dosyaislemleri{
    public static void main(String[] args) throws IOException{
        File file = new File("Java_ile_olusturulan_eleman");
        
        if (!file.exists()){
            file.createNewFile();
            System.out.println("Dosya oluşturuldu !");
            BufferedWriter write = new BufferedWriter(new FileWriter(file));
            write.write("Merhaab Dünya !");
            System.out.println("Yazıldı !");
            write.close();
            //sc.close();
            // Bilmiyorum ki :)
        }
        else{
            Scanner sc = new Scanner(file);
            while (sc.hasNextLine()){
                System.out.println(sc.nextLine());
            }
            sc.close();

        }
        
    }
}