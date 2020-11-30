import javax.swing.*;

public class guinput{
    public static void main(String[] args) {
        String str = JOptionPane.showInputDialog(null,"Değer gir : ");
        if (str==null){System.out.println("NULL Döndü");}
        else if ("".equals(str)){System.out.println("BOŞ Döndü");}
        else{System.err.println(str);}

    }
}
