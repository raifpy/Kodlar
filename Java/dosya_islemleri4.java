import java.io.File;
public class dosya_islemleri4{
    public static void main(String[] args) {
        File dosya = new File("dosya_islemleri4");
        System.out.println(dosya.toPath()); // dosya_islemleri4 adını dönüştürüyor
        System.out.println(dosya.getAbsolutePath()); // Dosyanın tam konumunu dönüyor
        System.out.println(dosya.lastModified()); // Dosyanınson değiştirilem tarhini döner
        //tring[] liste = dosya.list(); // NullPoint hatası veriyor        
    }
}