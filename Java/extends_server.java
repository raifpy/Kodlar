public class extends_server{
    public static void main(String[] args) {
        System.out.println("Server doğrudan çalıştırılamaz !");
        javax.swing.JOptionPane.showMessageDialog(null,"Server doğrudan çalıştırılamaz !");
        System.exit(0);
    }
    private String mesaj;
    public extends_server(){
        this.mesaj = "Default mesajı görüntülüyorsunuz .";
    }
    public extends_server(String mesaj){
        this.mesaj = mesaj;
    }
    public void modul(){
        javax.swing.JOptionPane.showMessageDialog(null, mesaj);
    }
    public String getMesaj(){
        return mesaj;
    }
}