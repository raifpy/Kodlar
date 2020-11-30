public class ters_f{
    public static void main(String[] args) {
        String[] array = "Merhaba Merhaba2 Merhaba3 Merhaba4".split(" ");

        System.out.println(String.format("%s\f%s\f%s\f%s",array[0],array[1],array[2],array[3]));

        // \f alt satır , tab şeklinde . Form Feed ()
    }
}