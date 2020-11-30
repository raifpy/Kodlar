class private_client{
    public static void main(String[] args) {
        private_server server = new private_server();
        //System.out.println(server.eleman); server.eleman diye birşey yok ..
        System.out.println(server.getEleman());
        server.setEleman("\n\t\t---\n\nserver.setEleman(String); ile beni değiştirebilirsin .\n\nAma server.eleman = String; ile değiştiremezsin ..");
        System.out.println(server.getEleman());
    }
}