package AbstractFactory.services;

import AbstractFactory.interfaces.Color;

public class Red implements Color {
    @Override
    public void fill() {
        System.out.println("fill ... 红 ...");
    }
}
