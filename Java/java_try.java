import java.util.InputMismatchException;
import java.util.Scanner;

public class java_try{
    public static void main(String[] args) {
        // try catch kullanımı .

    try{
        Scanner sc = new Scanner(System.in);
        System.out.print("INT Gir : ");
        int veri = sc.nextInt();
        System.out.println(veri);
    } 
    catch (InputMismatchException e){
        System.out.println("Hata , Rakam girmek durumundasınız , karakter değil : "+e.toString());
    }
    finally{
        System.out.println("ne olursa olsun finalley çalışacak ..");
    }
}    
}