public class degismeyen_degisken{
    public static void main(String[] args) {

        String degisebilir = "Değişebilirim .";
        System.out.println(degisebilir="Değiştim");

        final String degisemez = "Değişemem ."; //Beni böyle kabul et 
        System.out.println(degisemez="Değiştim"); // HATA :         The final local variable degisemez cannot be assigned. It must be blank and not using a compound assignment

        //Görüldüğü üzere final ile atanmış elemanlar değişemez .


    }
}