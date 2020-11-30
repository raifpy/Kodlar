import java.util.Scanner;
import javax.swing.JOptionPane;


public class giumesaj_sysarrgv{
    public static void main(String[] args) {
        byte len = (byte) args.length;
        if (len==0){
            JOptionPane.showMessageDialog(null,String.format("Toplam args : %s\nDoğru çalıştırmak için argüman eklemek zorundasın\n\n\"java guimesaj_sysarrgv.java argm1 agv2 agv3\"",len));
            System.exit(0);
        }
        System.out.print(String.format("0-%s aralığında sayı giriniz : ",len-1));
        Scanner sc = new Scanner(System.in);
        byte eleman = (byte) sc.nextInt();
        if (eleman > len-1 || eleman < 0){
            System.out.println(String.format("%s sayısını girdiniz . Lakin 0-%s aralığında değer bekleniyor !", eleman,len-1));
            System.exit(10);
        }
        else{
            JOptionPane.showMessageDialog(null,String.format("%s Sayısını verdiniz .", eleman));
        }



    }
}