package ChainResponsibility;

public abstract class PriceHandler {
    protected PriceHandler successor;

    public void setSuccessor(PriceHandler successor){
        this.successor = successor;
    }

    public abstract void processDiscount(float discount);
}
