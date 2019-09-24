package ChainResponsibility;

public class Sales extends PriceHandler {
    @Override
    public void processDiscount(float discount) {
        if (discount <= 0.1){
            System.out.format("%s批准了折扣:%2%n",this.getClass().getName(),discount);
        }else{
            successor.processDiscount(discount);
        }
    }
}
