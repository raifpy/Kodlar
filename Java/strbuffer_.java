class strbuffer{
    public static void main(String[] args) {
        String text="Merhaba Dünya";
        StringBuffer strbuffer = new StringBuffer(text); // text
        //StringBuffer strbuffer = new StringBuffer(20); // 20 haneli
        //StringBuffer strbuffer = new StringBuffer(); // 16 haneli

        print(strbuffer.toString()); // Stringi veriyor
        
        print( strbuffer.substring(2, 5)  ); // Aralığı veriyor
        print( strbuffer.substring(6)); //  Python'da : [6:] | Sadeçe start'ı verdik devam
        
        strbuffer.append(" Ben append ile eklendim !");
        print(strbuffer.toString());
        
        strbuffer.delete(2, 5); // Aralığı sildik
        print(strbuffer.toString());

        // Doğrudan strbuffer'ı yazalım ekrana
        System.out.println(strbuffer); // Böylede oluyor , ama str yapmak daha mantıklı tabiki ..

        strbuffer.replace(1, 3, "Birden 3'e kadar olan metni bu yazı ile değiştirmiş olmam lazım"); // Evet öyleymiş :)
        print(strbuffer.toString());


        print(strbuffer.reverse().toString()); // .reverse ile buffer nesnesini tersine çevirmiş olduk sonra .toString()

        
    }
    public static void print(String text){
        System.out.println(text);
    }
}