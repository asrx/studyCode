package Strategy.members;

public class TestClass {

    public static void main(String[] args) {
        PriceHandle priceHandle = new PriceHandle();
        int memberType = 0;
        switch (memberType){
            case 0:
                priceHandle.setMemberInterface(new NormalMember());
                break;
            case 1:
                priceHandle.setMemberInterface(new GoldMember());
                break;
            case 2:
                priceHandle.setMemberInterface(new DiamondMember());
                break;
        }

        System.out.println(priceHandle.getMemberPrice(10.9));
    }
}
