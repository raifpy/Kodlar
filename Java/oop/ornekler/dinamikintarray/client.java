class client{
    public static void main(String[] args) {
        server liste1 = new server(); // default ayarlar ile liste oluşacak
        //System.out.println(liste1.getSize());
        System.out.print(liste1.getArray()); // Default ayarlarda da oluşabiliyor , görseli çıksa yeter ..
        System.out.println(" Default ArrayList \n------\n"+liste1.toObjArray()+"\n\n");
    

        server liste3 = new server(new Integer[]{1,2,3,7,9,8,0,9}); // Integer[] vererek istediğimiz gibi işleyebiliyoruz
        liste3.add(900);
        System.out.print(liste3.getArray());
        System.out.println(" Int[] \n\n");
        
        server liste2 = new server(10);  // eleman sayısını biz belirttik | 0-10 arasında (10 dahil değil) 0,1,2,3,4,5,6,7,8,9 for döngüsü ile oluştu .add() ile eklendi ..
        System.out.println(liste2.getSize()); // arraylist 'in kaç elemandan oluştuğunu aldık
        System.out.println(liste2.getArray()); // arraylist 'i döner ( ArrayList<String> dönecek || java.util.ArrayList )
        System.out.println(liste2.getIndex(5)); // arraylist 'in 5. elemanı (4.index 'i alıyoruz)
        System.out.println("-------------"); 
        liste2.remove(5); // 5.elemanı sildik *
        System.out.println(liste2.getArray()); // arraylist 'i döner ( ArrayList<String> dönecek || java.util.ArrayList )
        System.out.println(liste2.getIndex(5)); // arraylist 'in 5. elemanı (4.index 'i alıyoruz)


    }
}