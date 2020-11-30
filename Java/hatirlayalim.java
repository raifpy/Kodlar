public class hatirlayalim{ //@raifpy 05.04.2020
    public static void main(String[] args){
        System.out.println("\nÖzetleniyor .\n\n");

        int rakam = 12;
        double duble=12.78;
        
        boolean True=false;
        boolean False=false;


        System.out.println("INT  : "+rakam);
        System.out.println("DUBL : "+duble);
        System.out.println("\033[31m---\t---\t---\033[0m");
        if (True){
            System.out.println("True'nın true değeri var !");
        }
        else{
            System.out.println("True'nın false değeri var !");
        }

        //

        if (False){
            System.out.println("False'ın true değeri var !");
        }
        else{
            System.out.println("False'ın false değeri var !");
        }

        System.out.println("String değer atayalım !");

        String abc="Merhaba ben string";

        System.out.println("Char değer atayalım !");

        char cba = 'a';

        System.out.println("\n\n- - - - - - -\n\n");

        System.out.println("STRING : "+abc);
        System.out.println("CHAR   : "+cba);
        System.out.println("\n\n\n");
        System.out.println("DOUBLE TO INT VERI DONUSUMLERI YAPALIUM :)");

        double _duble = 12.70;
        int _int = 10;

        System.out.println("DOUBLE DEGER : "+_duble);
        System.out.println("INTEGR DEGER : "+_int);

        _int = (int) duble; //Double'yi doğrudan int yapamayız . ama (int) yparametresi kullanılırsa bu yapılabilir

        System.out.println("INT DEGER EVRILDI : "+_int);

        System.out.println("\n\n\n");

        String str_int = "340";

        System.out.println(str_int+" değerinin int'a çevirilip 2'ye bölünmüş hali : "+Integer.valueOf(str_int)/2);
        System.out.println(str_int+" değerinin double'a çevirilip 2'ye bölünmüş hali : "+Double.valueOf(str_int)/2);
        System.out.println("\n");
        System.out.println(str_int+" değerinin int'a çevirilip 169'ye bölünmüş hali : "+Integer.valueOf(str_int)/169);
        System.out.println(str_int+" değerinin double'a çevirilip 169'ye bölünmüş hali : "+Double.valueOf(str_int)/169);

        //str_init dediğimiz eleman String ama Integer.valueOf() | Double.valueOf() 'ları kullanarak strin'i rakam olarak değiştirdik .
        // int("150") 'den farkı yok .


    }
}