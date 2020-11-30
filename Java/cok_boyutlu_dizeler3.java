import javax.swing.JOptionPane;

public class cok_boyutlu_dizeler3{
    public static void main(String[] args) {

        byte dongu = Byte.parseByte(nullcheck(InputGUI("2","Kaç elemanlı bir ArrayArray oluşturulsun ? ")));
        String[][] ArrayArray = new String[dongu][];
        for (int i=0;i<dongu;i++){
            ArrayArray[i] = nullcheck(InputGUI("", String.format("%s. index'in liste elemanları (boşlukdan bölünerek array hale gelecek)",i))).split(" ");
        }
        String mesaj = "";
        mesaj+= String.format("Ana Dizi Sayısı = %s\nElemanları : \n\n",ArrayArray.length);
        for (String[] i : ArrayArray){
            
            for (String ii : i){
                mesaj += ii+" , ";
            }
           
        }
        ShowGUI(mesaj);



        
    }

    public static String nullcheck(String veri){
        if (veri == null || veri.trim().equals("")){System.out.println("Boş değer , çıkılıyor ");System.exit(0);return veri;} // .trim() boşlukları silmek için // sondaki return compiler'da hata vermemesi için . Hiçbirşey dönemyecek , dönemez ..
        else return veri;
    }
    public static String InputGUI(String textbox_mesaj , String mesaj){
        return JOptionPane.showInputDialog(mesaj,textbox_mesaj);
    }
    public static void ShowGUI(String mesaj){
        JOptionPane.showMessageDialog(null, mesaj);
    }
}