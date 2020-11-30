import java.util.ArrayList;

public class server{
    // Bu sefer hepsini non-static yapacağım
    // non-static olursa aynı betikde aynı classı kullanarak birden fazla hesap oluşturabiliriz
    // 1 PC 4 user gibi
    // ama staticde 1 PC 1 User gibi gibi ..

    public static void main(String[] args) {
        System.out.println("client 'dan başlat !");
        System.exit(0);
    }


    private ArrayList<Integer> liste_int = new ArrayList<Integer>(); // Böylece liste_int'e başka sınıflar doğrudan erişemeyecek
    
    public server(){
        // Hiç değer girilmedi . liste_int 'ı default ayarlayalım !
        for (int ab=0;ab<5;ab++){
            liste_int.add(ab);
        }
    }

    public server(Integer sayi){ // Bize bir sayi verecek , o sayı kadar listeye ekleme yapacağız .. cCC
        for (int a=0;a<sayi;a++){ // For döngüsü hazır
            liste_int.add(a);
        }
    }

    public server(Integer[] array){ // Bize Int[] gelecek , for ile tek tek int'e atayıp ArrayList<String> 'e .add ile ekleyeceğiz ..
        for (int i:array){
            liste_int.add(i);
        }
    }

    public ArrayList<Integer> getArray(){
        return liste_int;
    }

    public Integer getSize(){
        return liste_int.size();
    }

    public Integer getIndex(Integer index){
        if (index <= getSize() && index > -1){
            return liste_int.get(index-1);}
        else{
            System.out.println("Olmayan index girildi !");
            return -1;
        }
    }

    public void remove(Integer index){
        liste_int.remove(index-1);
        
    }
    public void add(int veri){
        liste_int.add(veri);
    }

    public Object[] toObjArray(){
        return liste_int.toArray();
    }


}