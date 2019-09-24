package Decorator.Demo2;

public class Milk extends CondimentDecorator {
    public Milk(Beverage beverage) {
        this.beverage = beverage;
    }

    @Override
    public float cost() {
        return 0.2f + beverage.cost();
    }

    @Override
    public String getDescription() {
        return beverage.getDescription() + "，Milk";
    }
}
