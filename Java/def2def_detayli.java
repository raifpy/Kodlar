
public class def2def_detayli {
    public static void main(final String[] args) {
        System.out.println("\033[31mDeğer atama işlemleri başlıyor !\033[0m\n");
        byte __byte__ = 10;
        int __int__ = 10;
        float __float__ = 10.9F;
        double __double__ = 10.0;
        //boolean bool=true;
        
        System.out.println("__byte__   : "+__byte__ + "\n__int__    : "+__int__+"\n__float__  : "+__float__+"\n__double__ : "+__double__);
        System.out.println("\n\033[32mDeğer atama işlemi bitti . \033[0m\nAtanan elemanların min,max ve size değerleri geliyor !\033[0m");
        System.out.println("\n\t__byte__   "+__byte__+" (Değer Aralığı) : "+Byte.MIN_VALUE+" - "+Byte.MAX_VALUE+" | Bit : "+Byte.SIZE);
        System.out.println("\n\t__int__    "+__int__+" (Değer Aralığı) : "+Integer.MIN_VALUE+" - "+Integer.MAX_VALUE+" | Bit : "+Integer.SIZE);
        System.out.println("\n\t__float__  "+__float__+" (Değer Aralığı) : "+Float.MIN_VALUE+" - "+Float.MAX_VALUE+" | Bit : "+Float.SIZE);
        System.out.println("\n\t__double__ " +__double__+" (Değer Aralığı) : "+Double.MIN_VALUE+" - "+Double.MAX_VALUE+" | Bit : "+Double.SIZE);
        System.out.println("\n\033[35mDönüştürülüyor !\033[0m");
        System.out.println("-_-_-_-_-_-_-_-_-_-_-\n");
        __int__ = (int)__float__;
        System.out.println("__init__ = __float__ | __float__'ı int yapmak için (int)'i kullandık \n\t\\__init__ : "+__int__+"\n");
        
        
        if (__int__ == __byte__){
            System.out.println("Görünüşe göre değişen birşey olmadı :D");
            //bool = false;
        }

        // dönüşüm işlemleri

        //int __int2__ = (int) __float__;
        
    }
}