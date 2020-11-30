public class cok_boyutlu_dizeler{
    public static void main(String[] args) {
        // Liste içi liste .
        // Syntax : 
        //            String[][] strArrayArray = {{"Hey","sdfsdf"},{"Veri","Veri2"},{"Veri bilmem ne ","taam"}..     }
        //              for (String[] i : strArrayArray) sysout(i.length) // Devam
        //     YADA :
        //            String[][] strArrayArray = new String[3][] ; // ilk 3 kaç liste elemanından oluşacağını belirtiyoruz
        //            strArrayArray[0] = new String[][]{{"Merhaba","Evet"}}


        //String[][][][] str = new String[10][][][]; // Ne kadar eklenirse o kadar oluyormuş :D liste içindeki liste içindeki liste içindeki list olabiliyor yani .

        //System.out.println(str.length); // Tabiki bokunu çıkartmadan örnek verelim .

        String[][] StrArrayArray = {{"Liste1","Liste1.1"},{"Liste2","Liste2.2"}};
        for (String[] i : StrArrayArray) for (String ii : i) System.out.println(ii); // Bu syntax müthiş .

        // Sınırları zorlayalım

        Integer[][][] intArrayList = {        {         {1,2,3,4,5},{6,7,8,9,10}    },{     {11,12,13,14,15},{16,17,18,19,20}          }           };
        // int[] içinde int[] olacak ve int[] 'in içinde int'ler olacak //


        System.out.println(intArrayList);

        for (Integer[][] i:intArrayList){
            for (Integer[] ii:i){
                // son liste
                for (Integer iii:ii){
                    System.out.println(iii);
                }
            }
        }


        // Bu da böyle çok işimize yaracak önemli birşey .. 
    }
}