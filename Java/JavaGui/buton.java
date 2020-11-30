import javax.swing.*;

public class buton{
    public static void main(String[] args) {
        JFrame frame = new JFrame(); //Frame oluştu
        JButton buton = new JButton("Buton");
        JButton buton2 = new JButton("Buton2");
        buton.setBounds(20, 20, 300, 400); // buton boyutu
        buton2.setBounds(100, 20, 300, 400); // buton2 boyutu
        
        frame.add(buton);
        frame.add(buton2);

        
        
        frame.setSize(1000,1000); // en - boy // Etkileyici faktör
        
        //frame.setLayout(null); // 
        
        frame.setVisible(true); // Olmazsa ekran açılmaz
        
        System.out.println("Nedne olmadın lna");

        
    }
}