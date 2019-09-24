package Strategy.demo1;

public class StrategyDemo {

    public static void main(String[] args) {
        Context context1 = new Context(new OperationAdd());
        System.out.println("10 + 5 = " + context1.executeStrategy(10,5));

        Context context2 = new Context(new OperationSubstract());
        System.out.println("10 - 5 = " + context2.executeStrategy(10,5));

        Context context3 = new Context(new OperationMultiply());
        System.out.println("10 * 5 = " + context3.executeStrategy(10,5));
    }
}
