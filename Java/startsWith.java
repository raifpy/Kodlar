public class startsWith{
    public static void main(String[] args) {
        String veri = "Merhabalar ";
        System.out.println(veri.startsWith("M"));
        System.out.println(veri.startsWith("m"));

        System.out.println(veri.endsWith("r")); // Sonunda " " var :))
        veri = veri.trim();
        System.out.println(veri.endsWith("r"));

    }
}