package Decorator.Demo1;

/**
 * 抽象装饰器类
 */
public class ShapeDecorator {
    protected Shape decoatorShape;

    public ShapeDecorator(Shape decoatorShape) {
        this.decoatorShape = decoatorShape;
    }

    public void draw(){
        this.decoatorShape.draw();
    }
}
