package AbstractFactory;

import AbstractFactory.interfaces.Color;
import AbstractFactory.interfaces.Shape;

public abstract class AbstractFactory {
    public abstract Shape getShape(String shapeType);
    public abstract Color getColor(String colorType);

}
