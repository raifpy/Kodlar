import java.text.DecimalFormat;

public class okunaklidouble{
    public static void main(String[] args) {
        double sayi = 1.2340982348920384092384;

        System.out.println("İlk hali : "+sayi);
        DecimalFormat dc = new DecimalFormat("0.00");
        System.out.println("öbr hali : "+dc.format(sayi));

        DecimalFormat dc2 = new DecimalFormat("00.00");
        System.out.println("öbr hali : "+dc2.format(sayi));

        DecimalFormat dc3 = new DecimalFormat("0.000");
        System.out.println("öbr hali : "+dc3.format(sayi));


        DecimalFormat ds4 = new DecimalFormat("0.00 TL");
        System.out.println("str hali : "+ds4.format(sayi));
    }

}