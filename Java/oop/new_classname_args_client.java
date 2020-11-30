public class new_classname_args_client{
    public static void main(String[] args) {
        System.out.println(new_classname_args_server.dogrudan()); // ..server.java'yı javac ile derledim !

        new_classname_args_server new_c = new new_classname_args_server();
        new_c.getVeri();

        System.out.println("-----------");
        new_classname_args_server new_c2 = new new_classname_args_server("Bazı STRINGler");
        new_c2.getVeri();

    }
}




// ilk olarak static methodunu kullanarak nesneye atamadan sınıfı doğrudan çağırdık | server.java
// ikinicisinde nesneyi herhangi bir değer vermeden atadık . verimiz ona göre şekillendi
// üçüncüsünde nesnee String değer vererek atadık . Verimiz ona göre şekillendi .