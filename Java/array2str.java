import java.util.ArrayList; // 2. örnek | Bu elemanı kullanmak daha mantıklı ...
import java.util.Arrays; // 1. örnek . String[] için 

public class array2str {
    public static void main(String[] args) {
        String[] array = {"Selamlar","Ben","Liste"};

        System.out.println(array);

        String array_str = Arrays.toString(array);
        System.out.println(array_str);


        // YADA || java.util.ArrayList

        ArrayList<String> ARRAY = new ArrayList<String>();
        for (String i : "Selamlar Ben Liste".split(" ")){
            ARRAY.add(i);
        }
        System.out.println("-- java.util.ArrayList --");
        System.out.println(ARRAY);

        
    }
    
}


