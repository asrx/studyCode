package AbstractFactory.Demo1.services;

import AbstractFactory.Demo1.interfaces.Shape;

public class Rectangle implements Shape {
    @Override
    public void Draw() {
        System.out.println("Draw ... 长方形 ...");
    }
}
