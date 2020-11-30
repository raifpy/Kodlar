public class server{
    private String veri;
    public server(){
        veri = "Merhaba , herhangi bir değer verilmedi .";
    }
    public server(String deger){
        veri = String.format("Merhaba , %s değeri verildi .", deger);
    }
    public server(String[] array){
        veri = String.format("Merhaba , %s karakterli String array verildi : \n",array.length);
        for(String i:array){
            veri += i+"\n";
        }
    }
    
    public void print_veri(){
        System.out.println(veri);
    }
}