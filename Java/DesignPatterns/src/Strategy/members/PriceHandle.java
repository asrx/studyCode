package Strategy.members;

/**
 * 价格处理类
 */
public class PriceHandle {
    private MemberInterface memberInterface;

    public void setMemberInterface(MemberInterface memberInterface){
        this.memberInterface = memberInterface;
    }

    public Double getMemberPrice(Double price){
        Double postage = 5.0;

        if (memberInterface.freeShipping()){
            postage = 0.0;
        }

        return memberInterface.getPrice(price) + postage;
    }
}
