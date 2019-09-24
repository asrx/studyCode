package AbstractFactory.services;

import AbstractFactory.interfaces.Shape;

public class Circle implements Shape {
    @Override
    public void Draw() {
        System.out.println("Draw ... 圆方形 ...");
    }
}
