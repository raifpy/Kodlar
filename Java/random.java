import java.util.Random;

public class random{
    public static void main(String[] args) {
        Random random = new Random();

        int a = random.nextInt(200); // Random int | MAX VALUE (200)
        p(a);

        double b = random.nextDouble(); // Random double | Aynı işler geçerli | 0-1 arası değerler ..
        p(b);

        p(random.nextGaussian()); // Ne olduğunu anlamadım ama double ve -3-1 arası döndü şimdilik | -0 'da dönüyor :D

        //random.wait(); // Hata veriyor anlamadım

        //p(random.hashCode()); // int birşeyler | random ile alakası yok , aynı veriler

        p(random.toString()); // elemanı str yapıyor O_O (java.util.Random@587c290d)

        p(random.nextLong()); // long dönüyor
        
        
        
    }
    public static void p(String args) {System.out.println(args);}
    public static void p(int args) {System.out.println(args);}
    public static void p(double args) {System.out.println(args);}
    public static void p(float args) {System.out.println(args);}
}