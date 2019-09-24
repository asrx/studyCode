package AbstractFactory.services;

import AbstractFactory.interfaces.Shape;

public class Square implements Shape {
    @Override
    public void Draw() {
        System.out.println("Draw ... 正方形 ...");
    }
}
