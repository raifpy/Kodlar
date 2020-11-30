public class splitstr{
    public static void main(String[] args) {
        String mesaj = "Merhaba dostum ben bir mesajım .";
        String[] mesajsplit = mesaj.split("");
        for (String i:mesajsplit){
            System.out.println(i);
        }
    }
}