public class compare{
    public static void main(String[] args) {
        int a;
        a = "Hey".compareTo("Hey1");
        System.out.println(a); // -1

        int ab = "hey".compareTo("Hey"); // lenler aynı ama baş Case farkı olduğu için şaşıyor :D
        System.out.println(ab); // 32

        int abc = "hey".compareToIgnoreCase("Hey"); // karşılaştır , case' büyük küçük kelime ayrımını görmezden gel 
        
        System.out.println(abc); // 0

        /*int b = "Hey".length() - "Hey1".length();
        System.out.println(b);*/

        // compareToIgnoreCase yapısını kullanmak çok daha mantıklı 
    }
}