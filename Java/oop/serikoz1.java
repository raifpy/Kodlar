//package oop;

public class serikoz1{
    public static void main(String[] args){
        System.exit(130); 
    }
    String str = null;

    public String getName(){
        return this.str;
    }

    public boolean setName(String name){
        this.str = name;
        return true;
    }
}