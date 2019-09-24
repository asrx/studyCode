package Decorator.Demo2;

public class DarkRoast extends Beverage {
    public DarkRoast() {
        description = "Dark Roast";
    }

    @Override
    public float cost() {
        return 0.9f;
    }
}
