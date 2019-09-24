package Strategy.members;

/**
 * 钻石会员
 */
public class DiamondMember implements MemberInterface{

    /**
     * 会员价格，打八折
     * @param price
     * @return
     */
    @Override
    public Double getPrice(Double price) {
        return price * 0.8;
    }

    /**
     * 包邮
     * @return
     */
    @Override
    public boolean freeShipping() {
        return true;
    }
}
