public class array_islemleri3 {
    public static void main(String[] args) {

        String[] str_list = new String[]{"uwuu","uuuuwu"};
        String[] str_list2 = new String[]{"elma","armut"};

        System.out.println(str_list[1]);

        System.out.println(str_list2[1]);

        str_list = str_list2;  //str_list str_liste2 ile aynı olmuş oldu !

        System.out.println(str_list[1]);

        System.out.println(str_list2[1]);

        System.out.println(str_list);

        System.out.println(str_list2); // görüldüğü üzere id * 'leri de aynı . Yani str_list 'de ne değişiklik yaparsan str_list2 'de de değişiklik olmuş olacak


        str_list[1] = "MERCİMEK"; // değişkiliği str_list'e yaptım

        System.out.println(str_list2[1]); // ama str_list2 'yi çağırdım'
        
        System.out.println("Ve görüldüğü üzere iki liste tek eleman oldu ...");


    }
    
}