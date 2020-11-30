import java.util.Arrays;

public class array_sort {
    public static void main(String[] args) {
        String[] array = {"4","2","1","3","alti","beş","yedi__"};

        for(String i:array){
            System.out.println(i);
        }

        Arrays.sort(array);
        System.out.println("-----------");

        for(String i:array){
            System.out.println(i);
        }
    }
}