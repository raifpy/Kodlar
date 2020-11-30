public class objectArray{
    public static void main(String[] args) {
        // int[] 'ye sadece int nesneler atayabilrisiniz
        // keza String[] 'e de öyle ..

        // ama Object[] 'e her türlü nesnesiyi ekleyebiliyoruz

        String str = "Merhaba Dünya";
        Integer int_ = 999;
        double double_ = 999.9;
        float float_ = 888.8f;
        char char_ = 'a';
        String[] strArray = "Merhaba Dünya".split(" ");

        Object[] obj = {str,int_,double_,float_,char_,strArray};

        //System.out.println(obj); // 
        for (Object i : obj){
            System.out.println(i);
        }
    }
}