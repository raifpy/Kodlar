import java.awt.*;
import javax.swing.*;
public class GrafiklerGiris extends JLabel {
    
    public GrafiklerGiris(String s)
    {
      super();
     }
    
    @Override
    protected void paintComponent(Graphics g)
    {
       super.paintComponent(g);
       g.setColor(Color.red);
       g.setFont(new Font("Verdana",Font.BOLD,40));
       
       g.drawLine(0, 0,200, 150);
    }
}

import javax.swing.*;
public class GrafiklerGirisPencere extends JFrame {
    
    public GrafiklerGirisPencere()
    {
     add(new GrafiklerGiris("GrafiklerGirisEtiket"));
    
     }
}

 public static void main(String[] args) {
         GrafiklerGirisPencere pencere = new GrafiklerGirisPencere();
         pencere.setSize(500,400);
         
         pencere.setTitle("pencere");
         pencere.setLocationRelativeTo(null);
         pencere.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
         pencere.setVisible(true);
        
          }


---------------------------------------------------------------------------------------

import javax.swing.*;
import java.awt.*;

public class PanelGrafikCizimi extends JPanel {
    
    @Override protected void paintComponent(Graphics g)
    {
       super.paintComponent(g);
       g.drawLine(0, 0, 200,200);
       g.drawString("PanelGrafikÇizimi",200,150);
       g.fillRect(205, 155, 150, 60);
       g.fill3DRect(105, 80, 75, 30, true);
      }
}

import javax.swing.*;
public class GrafiklerGirisPencere extends JFrame {
    
    public GrafiklerGirisPencere()
    {
     add(new PanelGrafikCizimi());
    
     }
}

public static void main(String[] args) {
         GrafiklerGirisPencere pencere = new GrafiklerGirisPencere();
         pencere.setSize(500,400);
         
         pencere.setTitle("pencere");
         pencere.setLocationRelativeTo(null);
         pencere.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
         pencere.setVisible(true);
        
          }

--------------------------------------------------------------------------------

import java.awt.*;
import javax.swing.*;

public class Poligoncizimi extends JPanel {
    @Override protected void paintComponent(Graphics g)
    {
       super.paintComponent(g);
       
       int[] x  ={25,35,5,55,65};
       int[] y = {100,60,150,80,190};
       
       Polygon poligon = new Polygon(x,y,5);
       g.drawPolygon(poligon);
       
       int[] x1 = {125,135,85,155,165};
       int[] y1 = {50,30,150,40,90};
       
       Polygon poligon2 = new Polygon(x1,y1,5);
       g.fillPolygon(poligon2);
       
       int[] x2={75,80,105,155,90};
       int[] y2={150,250,200,180,190};
      
       g.drawPolyline(x2,y2,5);
     }
}

import javax.swing.*;
public class GrafiklerGirisPencere extends JFrame {
    
    public GrafiklerGirisPencere()
    {
     add(new Poligoncizimi());
    
     }
}

public static void main(String[] args) {
         GrafiklerGirisPencere pencere = new GrafiklerGirisPencere();
         pencere.setSize(400,300);
         
         pencere.setTitle("pencere");
         pencere.setLocationRelativeTo(null);
         pencere.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
         pencere.setVisible(true);
        
          }

---------------------------------------------------------------
import java.awt.*;
import javax.swing.*;
public class PanelSekiller extends JPanel {
    
    
   public static final int cizgi =1;
   public static final int Dikdortgen =2;
   public static final int yuvarlak_koseli_dikdortgen =3;
   public static final int oval =4;
   
   private int sekilTip;
   private boolean dolu;
   
   public PanelSekiller()
   { 
       
     }
   
   public PanelSekiller(int sekilTip,boolean dolu)
   {
     this.sekilTip=sekilTip;
     this.dolu=dolu;
    }
   
   public PanelSekiller(int sekilTip)
   {
     this.sekilTip=sekilTip;
     }

    public int getSekilTip() {
        return sekilTip;
    }

    public void setSekilTip(int sekilTip) {
        this.sekilTip = sekilTip;
    }

    public boolean isDolu() {
        return dolu;
    }

    public void setDolu(boolean dolu) {
        this.dolu = dolu;
    }
   
   @Override protected void paintComponent(Graphics g)
   {
       super.paintComponent(g);
       int en = getSize().width;
       int boy = getSize().height;
       
       switch(sekilTip)
   {
       case Dikdortgen:  
            g.setColor(Color.BLUE);
          if(dolu)
            {
              g.fillRect((int)(0.1*en),(int)(0.1*boy),
       (int)(en*0.8),(int)(0.8*boy));
         }
       else
       {
       g.drawRect((int)(0.1*en),(int)(0.1*boy),
       (int)(en*0.8),(int)(0.8*boy));
   }
   break;
       
       
       case yuvarlak_koseli_dikdortgen:  
          g.setColor(Color.ORANGE);
          if(dolu)
          {
            g.fillRoundRect((int)(0.1*en),(int)(0.1*boy),
       (int)(en*0.8),(int)(0.8*boy),30,30);
            }
          else
          {
            g.drawRoundRect((int)(0.1*en),(int)(0.1*boy),
       (int)(en*0.8),(int)(0.8*boy),30,30);
           }
   break;
   
       case oval:
           g.setColor(Color.GREEN);
           
           if(dolu)
           {
             g.fillOval((int)(0.1*en),(int)(0.1*boy),
       (int)(en*0.8),(int)(0.8*boy));
            
            }
           else{
             g.drawOval((int)(0.1*en),(int)(0.1*boy),
       (int)(en*0.8),(int)(0.8*boy));
           
           }
   break;
         }
   
     }
   
   
}



import java.awt.*;
import javax.swing.*;
public class GrafiklerGirisPencere extends JFrame {
    
    public GrafiklerGirisPencere()
    {
     setLayout(new GridLayout(2,3,10,10));
     add(new PanelSekiller(PanelSekiller.Dikdortgen));
     add(new PanelSekiller(PanelSekiller.yuvarlak_koseli_dikdortgen));
     add(new PanelSekiller(PanelSekiller.oval));
     
     add(new PanelSekiller(PanelSekiller.Dikdortgen,true));
     add(new PanelSekiller(PanelSekiller.yuvarlak_koseli_dikdortgen,true));
     add(new PanelSekiller(PanelSekiller.oval,true));
    
     }
}



public static void main(String[] args) {
         GrafiklerGirisPencere pencere = new GrafiklerGirisPencere();
         pencere.setSize(400,300);
         
         pencere.setTitle("pencere");
         pencere.setLocationRelativeTo(null);
         pencere.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
         pencere.setVisible(true);
        
          }
    
    