public class x{
    String i = "Merhaba .";

    public static void main(String[] args){
        x iks = new x();
        System.out.println(iks.i);
    }

    public String getI(){
        return i;
    }

    public void setI(String i){
        this.i = i;
    }
}