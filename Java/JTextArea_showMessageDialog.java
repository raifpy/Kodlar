import javax.swing.*;

class JTextArea_{
    public static void main(String[] args) {
        JTextArea text = new JTextArea(10,10); //  boy en
        text.setText("Merhaba JTextArea");
        text.setEditable(false);
        //text.setEnabled(false);

        JOptionPane.showMessageDialog(null, text);
        System.err.println(text);
    }
}