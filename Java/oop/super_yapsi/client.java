public class client {
    public static void main(String[] args) {
        server2.main(new String[]{}); // main metodu static olduğu için nesne olarak tanımlamamıza gerek yok
        // Yada çağıralım fark etmez ;
        server2 sv = new server2("Bir takım örnek");
        sv.yolla();

    }
}