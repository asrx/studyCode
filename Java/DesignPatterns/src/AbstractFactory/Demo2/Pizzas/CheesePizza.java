package AbstractFactory.Demo2.Pizzas;

import AbstractFactory.Demo2.Factory.PizzaIngredientFactory;
import AbstractFactory.Demo2.Pizza;

public class CheesePizza extends Pizza {
    PizzaIngredientFactory factory;

    public CheesePizza(PizzaIngredientFactory factory) {
        this.factory = factory;
    }

    @Override
    public void prepare() {
        System.out.println("Preparing " + name);
        dough = factory.createDough();
        sauce = factory.createSauce();
        cheese = factory.createCheese();
    }
}
