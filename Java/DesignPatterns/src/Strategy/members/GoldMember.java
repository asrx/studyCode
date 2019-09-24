package Strategy.members;

/**
 * 黄金会员
 */
public class GoldMember implements MemberInterface{

    /**
     * 会员价格，打九折
     * @param price
     * @return
     */
    @Override
    public Double getPrice(Double price) {
        return price * 0.9;
    }

    /**
     * 不包邮
     * @return
     */
    @Override
    public boolean freeShipping() {
        return false;
    }
}
