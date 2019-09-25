package AbstractFactory.Demo1.services;

import AbstractFactory.Demo1.interfaces.Color;

public class Green implements Color {
    @Override
    public void fill() {
        System.out.println("fill ... 绿 ...");
    }
}
