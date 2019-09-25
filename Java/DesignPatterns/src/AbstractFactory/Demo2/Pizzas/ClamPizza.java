package AbstractFactory.Demo2.Pizzas;

import AbstractFactory.Demo2.Factory.PizzaIngredientFactory;
import AbstractFactory.Demo2.Pizza;

public class ClamPizza extends Pizza {
    PizzaIngredientFactory factory;

    public ClamPizza(PizzaIngredientFactory factory) {
        this.factory = factory;
    }

    @Override
    public void prepare() {
        System.out.println("Preparing " + name);
        dough = factory.createDough();
        sauce = factory.createSauce();
        cheese = factory.createCheese();

        clam = factory.createClams();
    }
}
