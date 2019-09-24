package AbstractFactory.services;

import AbstractFactory.interfaces.Color;

public class Green implements Color {
    @Override
    public void fill() {
        System.out.println("fill ... 绿 ...");
    }
}
