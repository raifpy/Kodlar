public class array_copy{
    public static void main(String[] args) {
        String[] array1 = {"Merhaba","Ben","Array","1"};
        String[] array2 = new String[array1.length]; // array1 'in len'inde yeni bir String array oluşturdk
        
        //array2 = array1;  // 2 arrayi tek array yapıyor bu şekilde .
        
        System.arraycopy(array1, 0, array2, 0, array1.length); // arrray1 'in 0 . indexini , array2 'nin 0. indexinden itibaren yaz , boyutu da array1.length olsun 
        
        System.out.println(array1); // id | pid ** leri aynı değil buradan da görebiliyoruz .
        System.out.println(array2); // id | pid ** leri aynı değil buradan da görebiliyoruz .
        System.out.println("");

        for (String i : array2){
            System.out.println(i);
        }
    }
}