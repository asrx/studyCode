package ChainResponsibility;

public class PriceHandlerFactory {

    public static PriceHandler createPriceHandler(){
        PriceHandler sales = new Sales();
        PriceHandler man = new Man();
        PriceHandler vp = new Vp();
        PriceHandler ceo = new Ceo();

        sales.setSuccessor(man);
        man.setSuccessor(vp);
        vp.setSuccessor(ceo);

        return sales;
    }
}
