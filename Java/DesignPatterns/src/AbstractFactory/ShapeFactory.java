package AbstractFactory;

import AbstractFactory.interfaces.Color;
import AbstractFactory.interfaces.Shape;
import AbstractFactory.services.Circle;
import AbstractFactory.services.Rectangle;
import AbstractFactory.services.Square;

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
