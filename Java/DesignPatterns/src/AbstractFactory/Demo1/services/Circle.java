package AbstractFactory.Demo1.services;

import AbstractFactory.Demo1.interfaces.Shape;

public class Circle implements Shape {
    @Override
    public void Draw() {
        System.out.println("Draw ... 圆方形 ...");
    }
}
