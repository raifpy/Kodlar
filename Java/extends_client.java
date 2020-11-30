public class extends_client extends extends_server{
    public static void main(String[] args){
        extends_client ex = new extends_client("Merhaba Java");
        //ex.getMesaj();
        ex.getModul();
        System.out.println(ex.getMesaj());
        ex.mesajyaz();
    }

    public extends_client(){
        super();
    }
    public extends_client(String mesaj){
        super(mesaj);
    }
    public void mesajyaz(){
        System.out.println(super.getMesaj());
    }

    public void getModul(){
        super.modul();
    }


}