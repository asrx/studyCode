package AbstractFactory;

import AbstractFactory.Demo1.AbstractFactory;
import AbstractFactory.Demo1.FactoryProducer;
import AbstractFactory.Demo1.interfaces.Color;
import AbstractFactory.Demo1.interfaces.Shape;

public class TestClass {
    public static void main(String[] args) {
        AbstractFactory shapeFactory = FactoryProducer.getFactory("SHAPE");

        Shape shape1 = shapeFactory.getShape("CIRCLE");

        shape1.Draw();

        AbstractFactory colorFactory = FactoryProducer.getFactory("COLOR");

        Color color1 = colorFactory.getColor("GREEN");

        color1.fill();
    }
}
