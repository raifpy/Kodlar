import java.util.Scanner;
import java.util.Random;

public class cliOyun{
    static Random random = new Random();
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String nick;
        String isim;

        
        anadongu:
        while (true){

            isimdongu:
            while (true){
                System.out.print("AD : ");
                isim = sc.nextLine();
                if (!isim.equals("")){
                    
                    break isimdongu;
                }
                else System.out.println("\n\033[31mLütfen\033[0m , bir isim gir ...\n");
            }
            nickdongu:
            while (true){
                System.out.print("Nick : ");
                nick = sc.nextLine();
                if (!nick.equals("")){
                    break nickdongu;
                }
                else{
                    nick = getNick(getRandom(8, 10));
                    System.out.println(String.format("Senin için '%s' nick'ini hazırladım .",nick));
                    System.out.print("Başka bir nick almak için [Enter] , devam etmek için herhangi bir tuşa bas ");
                    String dogrulama = sc.nextLine();
                    if (!dogrulama.equals("")){
                        break nickdongu;
                    }
                    
                    }
                }
            
            System.out.println(String.format("Isim :  %s \nNick : %s\n\n[Devam edil]", isim,nick));
            String dogrulama = sc.nextLine();
            

            
            }

        }

    public static int getRandom(int ilk,int son){
        while (true){
            int ham = random.nextInt(son);
            if (ham >= ilk) return ham;
        }
    }
    public static String getNick(int hane){
        String nick = "";
        for (int i=0;i<hane;i++){
            nick += String.valueOf((char) getRandom(100, 150));
            //System.out.println(nick);
        }
        
        return nick;
    }
}