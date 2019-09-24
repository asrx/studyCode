package Strategy.members;

/**
 * 会员接口
 */
interface MemberInterface {
    /**
     * 会员价格
     * @param price
     * @return
     */
    Double getPrice(Double price);

    /**
     * 是否包邮
     * @return
     */
    boolean freeShipping();
}
