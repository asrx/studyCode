package AbstractFactory;

import AbstractFactory.interfaces.Color;
import AbstractFactory.interfaces.Shape;
import AbstractFactory.services.Green;
import AbstractFactory.services.Red;

public class ColorFactory extends AbstractFactory {
    @Override
    public Shape getShape(String shapeType) {
        return null;
    }

    public Color getColor(String colorType){
        if (colorType == null) return null;

        if (colorType.equalsIgnoreCase("RED")){
            return new Red();
        }else if(colorType.equalsIgnoreCase("GREEN")) {
            return new Green();
        }
        return null;
    }
}
