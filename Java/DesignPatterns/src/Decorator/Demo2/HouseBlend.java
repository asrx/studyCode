package Decorator.Demo2;

public class HouseBlend extends Beverage{

    public HouseBlend() {
        description = "House Blend";
    }

    @Override
    public float cost() {
        return 10.0f;
    }
}
