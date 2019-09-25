package SimpleFactory.Demo2.CJG;

import SimpleFactory.Demo2.Pizza;
import SimpleFactory.Demo2.PizzaStore;

public class CJGPizzaStore extends PizzaStore {
    @Override
    protected Pizza createPizza(String type) {

        if (type.equals("cheese")){
            return new CJGStyleCheesePizza();
        }else if (type.equals("pepperoni")){
            return new CJGStylePepperoniPizza();
        }else if (type.equals("clam")){
            return new CJGStyleClamPizza();
        }else if (type.equals("veggie")){
            return new CJGStyleVeggiePizza();
        }

        return null;
    }
}
