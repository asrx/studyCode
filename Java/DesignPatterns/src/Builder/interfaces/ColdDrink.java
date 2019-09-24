package Builder.interfaces;

import Builder.Services.Bottle;
import Builder.Services.Wrapper;

public abstract class ColdDrink implements Item{
    @Override
    public Packing packing() {
        return new Bottle();
    }

    @Override
    public abstract float price();
}
