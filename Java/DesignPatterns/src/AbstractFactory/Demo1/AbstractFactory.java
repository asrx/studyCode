package AbstractFactory.Demo1;

import AbstractFactory.Demo1.interfaces.Color;
import AbstractFactory.Demo1.interfaces.Shape;

public abstract class AbstractFactory {
    public abstract Shape getShape(String shapeType);
    public abstract Color getColor(String colorType);

}
