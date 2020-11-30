import javax.swing.JOptionPane;

public class Guilogin{
    public static void main(String[] args) {
        login();
    }

    public static void login(){
        boolean dongu = true;
        byte deneme_sayisi = 0;

        final String username = "raifpy";
        final String password = "123";
        
        while (dongu){
            if (deneme_sayisi == 3){
                System.out.println("Çok fazla giriş denemesi yapıldı !");
                JOptionPane.showMessageDialog(null,"Çok fazla giriş denemesi !\nLütfen bir süre sonra tekrar dene .");
                System.exit(0);
            }
            String username_attemp = JOptionPane.showInputDialog(null,"Kullanıcı adını giriniz");
            if (!bosmu(username_attemp)){
                String password_attemp = JOptionPane.showInputDialog(null,String.format("%s Şifresini giriniz",username_attemp)    );
                if (!bosmu(password_attemp)){
                    if (password_attemp.equals(password)== false || username_attemp.equals(username) == false){
                        JOptionPane.showMessageDialog(null,"Kullanıcı adı yada Şifre yanlış !");
                        ++deneme_sayisi;}
                    else{
                        System.out.println(String.format("Giriş başarılı !\n\nUsername : %s \nPassword : %s\n\njavax.swing",username_attemp,password_attemp));
                        JOptionPane.showMessageDialog(null,String.format("Giriş başarılı !\n\nUsername : %s \nPassword : %s\n\njavax.swing",username_attemp,password_attemp));
                        break;
                    }
        }}}
        // BURADAN DEVAM
    }
    public static boolean bosmu(String veri){
        if (veri == null){
            System.out.println("Işlem iptal ediliyor !");
            System.exit(0);
            return true; /* must be return boolean işte | Yoksa olayı yok */}
        
        else if (veri.equals("")){
            JOptionPane.showMessageDialog(null,"Değer Boş olamaz !");
            return true;}


              
        else{
            return false;}
    }
}