package AbstractFactory.Demo1.services;

import AbstractFactory.Demo1.interfaces.Shape;

public class Square implements Shape {
    @Override
    public void Draw() {
        System.out.println("Draw ... 正方形 ...");
    }
}
