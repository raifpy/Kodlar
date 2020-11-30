
public class int2str {
    public static void main(String[] args) {
        int eleman = 1000; // 1000 int'i oluşturduk .
        
        System.out.println(eleman + 2); //2 ile toplayıp kontrol ediyoruz int mi değil mi :D

        String stringEleman = String.valueOf(eleman);

        System.out.print(stringEleman + 2);

        // int olan ile 1000 + 2'yi yaptık ve 1002 oldu .
        // int2str yaptk ve 1000+2 yaptık 10002 oldu . Dönüşme işi başarılı .


    }

}