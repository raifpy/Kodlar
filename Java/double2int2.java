public class double2int2{
    public static void main(String[] args) {
        double duble = 10.40;
        int inte;
        //inte = duble; HATA VERİR
        duble = Double.valueOf("10.20");
        inte = (int) duble;
        System.out.println(duble);
        System.out.println(inte);
    }
}