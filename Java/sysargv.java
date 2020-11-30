public class sysargv {
    public static void main(String[] args) {
        for (String text:args){
            print(text);
        }
        
    }
    
    public static void print(String text) {
        System.out.println(text);
    }
}