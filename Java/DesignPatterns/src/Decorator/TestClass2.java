package Decorator;

import Decorator.Demo1.Circle;
import Decorator.Demo1.RedShapeDecorator;
import Decorator.Demo1.Shape;
import Decorator.Demo1.ShapeDecorator;
import Decorator.Demo2.*;

public class TestClass2 {
    public static void main(String[] args) {
        Beverage beverage = new HouseBlend();
        System.out.println(beverage.getDescription() + " $ " +beverage.cost());

        Beverage beverage1 = new DarkRoast();
        beverage1 = new Mocha(beverage1);
        beverage1 = new Mocha(beverage1);
        beverage1 = new Milk(beverage1);
        System.out.println(beverage1.getDescription() + " $ " +beverage1.cost());

    }
}
