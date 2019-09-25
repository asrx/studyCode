package AbstractFactory.Demo1.services;

import AbstractFactory.Demo1.interfaces.Color;

public class Red implements Color {
    @Override
    public void fill() {
        System.out.println("fill ... 红 ...");
    }
}
