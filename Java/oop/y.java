public class y{
    public static void main(String[] args) {
        x iks = new x();                           // x.java 'nın derlenmiş olmasına dikkat et ! (javac x.java)
        System.out.println(iks.i);
        iks.setI("Local değişiklik ");
        //iks.i = "Local Değişkilik"; // 'de olabilir
        System.out.println(iks.i);
        System.out.println("-----");
        String[] abc = {"@raifpy"}; // String[] 'i oluşturduk | bilmemne.java ^nın maininde String[] olmak zorunda !
        iks.main(abc); // Doğrudan x.java 'nın main'ini çağırabiliriz .
    }
}