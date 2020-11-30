public class server2 extends server{
    public server2(){
        super();
    }
    public server2(String veri){
        super(veri);
    }
    public server2(String[] array){
        super(array);
    }
    

    public void yolla(){
        super.print_veri();
    }
    public static void main(String[] args) {
        server2 sv2 = new server2("Ey Türk gençliği! Birinci vazifen, Türk istiklâlini, Türk Cumhuriyet'ini, ilelebet, muhafaza ve müdafaa etmektir\nMevcudiyetinin ve istikbalinin yegâne temeli budur\nBu temel, senin, en kıymetli hazinendir\nİstikbalde dahi, seni bu hazineden mahrum etmek isteyecek, dahilî ve haricî bedhahların olacaktır\n".split("\n"));    
        sv2.yolla();
    }
    
    
}