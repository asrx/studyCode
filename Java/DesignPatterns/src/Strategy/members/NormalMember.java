package Strategy.members;

/**
 * 普通会员
 */
public class NormalMember implements MemberInterface{
    @Override
    public Double getPrice(Double price) {
        // 原价
        return price;
    }

    @Override
    public boolean freeShipping() {
        // 不包邮
        return false;
    }
}
