import java.util.Scanner;

class input_str{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int eleman = sc.nextInt();
        yaz(eleman);
        sc.close();
    }
    public static void yaz(int mesaj){
        
        System.out.println(mesaj);
    }
}