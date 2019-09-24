package Decorator;

import Decorator.Demo1.Circle;
import Decorator.Demo1.RedShapeDecorator;
import Decorator.Demo1.Shape;
import Decorator.Demo1.ShapeDecorator;

public class TestClass {
    public static void main(String[] args) {
        Shape circle = new Circle();

        ShapeDecorator redCircle = new RedShapeDecorator(circle);

        System.out.println("Circle with normal border:");
        circle.draw();

        System.out.println("\nCircle of red border:");
        redCircle.draw();
    }
}
