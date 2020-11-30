public class new_classname_args_server{
    String veri;
    public new_classname_args_server(){
        System.out.println("Doğrudan çağırıldı . ARGS yok .");
        veri = "Arg verilmedi ";
    }
    public new_classname_args_server(String veri){
        System.out.println(String.format("ARGS Var : %s",veri));
        this.veri = veri;
    }

    public void getVeri(){
        System.out.println(veri);
    }

    public static String dogrudan(){
        return "Doğrudan dönen veri efenim";
    }
}