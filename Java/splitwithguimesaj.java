import javax.swing.JOptionPane;

public class splitwithguimesaj{
    public static void main(String[] args) {
        String bol = JOptionPane.showInputDialog("Bir cümle gir ","Ben örnek bir cümleyim .");
        
        if (bol == null || "".equals(bol)){
            JOptionPane.showMessageDialog(null, "Boş değer verdiniz .");
            System.exit(100);}
        

        String[] bollist = bol.split(" "); // " " 'dan böldük . yani her kelime array'in parçası olacak
        int sayi=0;
        for (String u:bollist){
            JOptionPane.showMessageDialog(null,String.format("KELIME : %s\nINDEX  : %s",u,sayi));
            sayi++;
        }

        // YADA
        JOptionPane.showMessageDialog(null,"Aynısı sayılı for döngüsü ile yapılıyor \n\n'for(int a=0;a<bollist.length;a++)'");


        for (int say=0;say<bollist.length;say++){
            JOptionPane.showMessageDialog(null,String.format("KELIME : %s\nINDEX  : %s",bollist[say],say));
        }
        
    }
    
}