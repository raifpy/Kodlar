import java.util.ArrayList;

public class arraylist_2{
    public static void main(String[] args) {
        ArrayList<String> yeni = new ArrayList<String>();
        yeni.add("Merhaba");
        yeni.add("Hey");
        System.out.println(yeni);
        System.out.println("\n");
        for (String hey:yeni){
            System.out.println(hey);
        }
        System.out.println(yeni.get(0));
        System.out.println(yeni.size());

        String[] yeni2 = {"Böyle","Yaptık"};
        System.out.println(yeni2);
        for (String hey:yeni2){
            System.out.println(hey);
        }

        System.out.println(yeni2.length);

    }
}