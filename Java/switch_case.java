import java.util.Scanner;

public class switch_case{
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
            System.out.print("CASE'e mesaj : ");
            String veri = sc.nextLine();
            sc.close();

            switch (veri) {
                case "merhaba":
                    System.out.println("CASE1 çalışıyor dostum ");
                    break;
                case "2":
                    System.out.println("2 dost .");
                    break;
                
                case "":
                    System.out.println("Boş değer dostum °°");
                    break;
            

                default:
                    System.out.println("DEFAULT > ELSE çalışıyor dostum");
                    break;
            }

        
    }
}

/* Kullanışlı baya  *///