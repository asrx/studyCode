package AbstractFactory.services;

import AbstractFactory.interfaces.Shape;

public class Rectangle implements Shape {
    @Override
    public void Draw() {
        System.out.println("Draw ... 长方形 ...");
    }
}
