package ChainResponsibility;

public class Man extends PriceHandler {

    @Override
    public void processDiscount(float discount) {
        if (discount <= 0.2){
            System.out.format("%s批准了折扣:%2%n",this.getClass().getName(),discount);
        }else{
            successor.processDiscount(discount);
        }
    }
}
