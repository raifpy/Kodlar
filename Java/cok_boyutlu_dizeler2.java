public class cok_boyutlu_diziler2{
    public static void main(String[] args) {
        
        String[] strArray = "Hi I Am Simple String Array".split(" ");
        for (String i:strArray) System.out.println(i);
        
        System.out.println("\n\t\033[32m--------\033[0m\n");
        
        String[][] strArrayArray = {/**/ {"Hi I Am Not Hard ArrayArray"} , {"Hi I Don't Know Why I Writing This Codes Now"} /**/ };
        for(String[] i:strArrayArray) for (String ii:i) System.out.println(ii);

        System.out.println("\n\t\033[32m--------\033[0m\n");

        String[][][] strArrayArrayArray = {/**/ {"Hi I Am Hard ArrayArray".split(" ")} , {"Hi I Don't Know Why I Writing This Codes Now".split(" ")} /**/ };
        for (String[][] i : strArrayArrayArray) for (String[] ii : i) for (String iii : ii) System.out.println(iii);

        // Evet zorladı :)

        /*
        strArray = null;
        strArrayArray = null;
        strArrayArrayArray = null;*/


        // Bu değerleri tekrar atayaraj kullanacağım . Bu yüzden nullladım

        
    }
}