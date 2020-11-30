public class issame{
    public static void main(String[] args){
        
      String str1 = "Hi";
      String str2=new String("Hi");
      String str3 = "Hi";
     

        System.out.println(str1); //Hi
        System.out.println(str2); //Hi
        System.out.println(str3); //Hi


        System.out.println(str1==str2); //false
        System.out.println(str1==str3); //true
        
      	// | openjdk 11                  
    }
}

// Yeni havuzda takıldığı için elemanlar aynı olsa bile aynı olarak görülmüyorlar ... CcC