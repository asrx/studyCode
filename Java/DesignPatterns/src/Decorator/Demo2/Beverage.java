package Decorator.Demo2;

/**
 * 饮料
 */
public abstract class Beverage {
    String description = "Unkown Beverage";

    public String getDescription() {
        return description;
    }

    public abstract float cost();
}
