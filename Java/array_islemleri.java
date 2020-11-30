
public class array_islemleri {
    public static void main(String[] args) {
        boolean bool = false;
        int[] init = new int[]{1,2,3,4,5,6,7,8,9,10,11,12,13,14};
        
        for(int veri:init){
            System.out.println(veri);
            if (veri==9){
                System.out.println("Sayı bulundu ! "+ veri);
                bool = true;
                break;
            };
        };
        if (bool){
            System.out.println("Viyuuuuuuuuuu");
            for (int a=0;a<=10.0;a++){
                System.out.println("Fiyüüüüü "+a);
            }
        }
        init = new int[10];
        //int[] init2 = new int[]{1,2,3,4,5,6,7};
        //init[10] = init2;
        System.out.println("---------------");
        for (int eleman:init){
            System.out.println(eleman);
        }
        System.out.println(init[3]);


    }
}