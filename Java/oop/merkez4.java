class sabunum{
    String dujtu = "Sabunum dujtu";

    
    
    public String str(){
        return dujtu;
    }
}

public class merkez4{
    public static void main(String[] args){
        sabunum eleman = new sabunum();
        eleman.dujtu = "DENEME";
        System.out.println(eleman.dujtu);
        //System.out.println(eleman.str());

    }
}
