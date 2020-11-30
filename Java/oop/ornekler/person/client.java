import javax.swing.*;

class client{
    public static void main(String[] args) {
        String dongu = (JOptionPane.showInputDialog("Kaç adet person oluşsun ","3"));
        devam(dongu);

        int while_ = Integer.parseInt(dongu);
        server[] sv_array = new server[while_];
        
        String age;
        String name;
        for (int b=0;b<while_;b++){
            age = JOptionPane.showInputDialog(null,"Yaş :");
            devam(age);
            name = JOptionPane.showInputDialog(null,"Isim :");
            devam(name);
            sv_array[b] = new server(age,name);
        }
        String son="\n";
        for(server i : sv_array){
            son += server.listele(i)+"\n";
            //System.out.println(i.toStr()); // Aynısı .
        }

        JTextArea text = new JTextArea(son);
        text.setEditable(false);
        
        JOptionPane.showMessageDialog(null,text);

    }
    public static void devam(String veri){
        if (veri == null || "".equals(veri)){
            System.out.println("Çıkılıyor .");
            System.exit(0);
        }

    }
    
}