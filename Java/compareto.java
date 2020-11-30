// str kıyaslama

public class compareto{
    public static void main(String[] args) {
        System.out.println("mrb".compareTo("mrb")); // 0
        System.out.println("mrb".compareTo("mrb1")); // -1 | ilk str ikinci str'dan -1 kadar küçük (length mantığı)
        System.out.println("mrb1".compareTo("mrb")); //1 |  1. str 2. str'dan 1 kadar büyük

    // length gibi ..
    }
}