package Decorator.Demo2;

public class Mocha extends CondimentDecorator {
    public Mocha(Beverage beverage) {
        this.beverage = beverage;
    }

    @Override
    public float cost() {
        return 0.7f + this.beverage.cost();
    }

    @Override
    public String getDescription() {
        return this.description + ", Mocha";
    }
}
