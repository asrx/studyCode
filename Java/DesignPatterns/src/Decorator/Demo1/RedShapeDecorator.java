package Decorator.Demo1;

public class RedShapeDecorator extends ShapeDecorator {
    public RedShapeDecorator(Shape decoatorShape) {
        super(decoatorShape);
    }

    @Override
    public void draw() {
        super.draw();
        this.setRedBorder(this.decoatorShape);
    }

    private void setRedBorder(Shape decoratedShap){
        System.out.println("Border color: Red");
    }
}
