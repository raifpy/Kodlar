import java.io.File;
import java.io.IOException;

public class dosyaislemleri{
    public static void main(String[] args) throws IOException{
        File file = new File("silinecek");
        if (file.exists()){file.delete();System.out.println(String.format("Dosya ( %s ) zaten mevcut ! Yok edildi ...",file.toString()));}
        else{

            file.createNewFile();
            System.out.println(file.canExecute());
            System.out.println(file.canRead());
            System.out.println(file.canWrite());
        }
    }
}