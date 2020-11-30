public class str2int {
    public static void main(String[] args) {
        // python int() methodu ile aynı eleman :)
        String stringDeger = "100"; // int'e dönüşemeyecek bir değer girildiğinde çökecektir .

        System.out.println(stringDeger);
        // Bu eleman String . int yapmaya çalışırsak :
        //int integral = stringDeger;
        // Tabiki hata verecek . AMA
        int eleman = Integer.valueOf(stringDeger);
        //int eleman = Integer.parseInt(stringDeger);

        // Python'daki int("100") ile aynı işi gördük .

        System.out.println(eleman);
        //System.out.println(stringDeger*3); # // Çarpamaya çalışırken hata verdi :)

        // type verisini nasıl gösteririm bilmediğim için :D

        

    }

}