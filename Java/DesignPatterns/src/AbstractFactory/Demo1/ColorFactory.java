package AbstractFactory.Demo1;

import AbstractFactory.Demo1.interfaces.Color;
import AbstractFactory.Demo1.interfaces.Shape;
import AbstractFactory.Demo1.services.Green;
import AbstractFactory.Demo1.services.Red;

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
