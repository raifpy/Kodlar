import java.io.File;
public class dosya_islemleri2{
    public static void main(String[] args) {
        File dosya = new File("dosya_islemleri2");
        dosya.exists(); // dosya_islemleri2 var ise true yok ise false dönecek
    }
}