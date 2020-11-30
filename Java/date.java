import java.util.Date;

class date{
    public static void main(String[] args) {
        Date date = new Date();
        String strdate = date.toString(); // String formatda tam tarihi verdi > Fri May 29 01:32:34 TRT 2020
        String[] strlist = strdate.split(" "); // Boşluklarıdan böldük ona göre işleyeceğiz
        String Day = strlist[0];
        String Mon = strlist[1];
        String MonDay = strlist[2];
        String Clock = strlist[3];
        String Location = strlist[4];
        String Year = strlist[5];

        System.out.println("GUN : "+Day);
        System.out.println("AY  : "+Mon);
        System.out.println("KAÇ : "+MonDay);
        System.out.println("Sat : "+Clock);
        System.out.println("Loc : "+Location);
        System.out.println("YIL : "+Year);
        p("\n");
        p(date.getHours()); // Saati veriyor 
        p(date.getMinutes());
        p(date.getSeconds()); // Saniyeleri dönüyor 
        
        p(date.getDate()); // Ayın gününü dönüyor (Bugün 29 Mayıs 2020) // 29 döndü
        
        p(date.getMonth());
        p(date.getDay()); // Haftanın gününü veriyor (1-2-3-4-5-6-7)



        //date.notifyAll();
    }
    public static void p(String text){System.out.println(text);}
    public static void p(int text){System.out.println(text);}
    public static void p(float text){System.out.println(text);}

}