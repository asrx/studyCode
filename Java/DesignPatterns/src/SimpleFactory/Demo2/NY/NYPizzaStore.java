package SimpleFactory.Demo2.NY;

import SimpleFactory.Demo1.CheesePizza;
import SimpleFactory.Demo1.ClamPizza;
import SimpleFactory.Demo1.PepperoniPizza;
import SimpleFactory.Demo1.VeggiePizza;
import SimpleFactory.Demo2.Pizza;
import SimpleFactory.Demo2.PizzaStore;

public class NYPizzaStore extends PizzaStore {
    @Override
    protected Pizza createPizza(String type) {

        if (type.equals("cheese")){
            return new NYStyleCheesePizza();
        }else if (type.equals("pepperoni")){
            return new NYStylePepperoniPizza();
        }else if (type.equals("clam")){
            return new NYStyleClamPizza();
        }else if (type.equals("veggie")){
            return new NYStyleVeggiePizza();
        }

        return null;
    }
}
