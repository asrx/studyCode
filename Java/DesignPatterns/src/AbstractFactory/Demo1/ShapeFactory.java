package AbstractFactory.Demo1;

import AbstractFactory.Demo1.interfaces.Color;
import AbstractFactory.Demo1.interfaces.Shape;
import AbstractFactory.Demo1.services.Circle;
import AbstractFactory.Demo1.services.Rectangle;
import AbstractFactory.Demo1.services.Square;

public class ShapeFactory extends AbstractFactory {

    public Shape getShape(String shapeType){
        if (shapeType == null) return null;

        if (shapeType.equalsIgnoreCase("CIRCLE")){
            return new Circle();
        }else if(shapeType.equalsIgnoreCase("RECTANGLE")){
            return new Rectangle();
        }else if(shapeType.equalsIgnoreCase("SQUARE")){
            return new Square();
        }
        return null;
    }

    @Override
    public Color getColor(String colorType) {
        return null;
    }
}
