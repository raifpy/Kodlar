public class pwd{
    public static void main(String[] args) {
        String pwd = System.getProperty("user.dir");
        System.out.println(String.format("My Java code run in : ",pwd));
    }
}