package ChainResponsibility;

public class Vp extends PriceHandler {

    @Override
    public void processDiscount(float discount) {
        if (discount <= 0.5){
            System.out.format("%s批准了折扣:%2%n",this.getClass().getName(),discount);
        }else{
            successor.processDiscount(discount);
        }
    }
}
