package AbstractFactory.Demo2.NY;

import AbstractFactory.Demo2.Factory.NYPizzaIngredientFactory;
import AbstractFactory.Demo2.Factory.PizzaIngredientFactory;
import AbstractFactory.Demo2.Pizza;
import AbstractFactory.Demo2.PizzaStore;
import AbstractFactory.Demo2.Pizzas.CheesePizza;
import AbstractFactory.Demo2.Pizzas.ClamPizza;
import AbstractFactory.Demo2.Pizzas.Pepperoni;
import AbstractFactory.Demo2.Pizzas.VeggiePizza;

public class NYPizzaStore extends PizzaStore {
    @Override
    protected Pizza createPizza(String type) {

        Pizza pizza = null;

        PizzaIngredientFactory ingredientFactory =
                new NYPizzaIngredientFactory();

        if (type.equals("cheese")){
            pizza = new CheesePizza(ingredientFactory);
            pizza.setName("New York Style Cheese Pizza");
        }else if (type.equals("pepperoni")){
            pizza = new VeggiePizza(ingredientFactory);
            pizza.setName("New York Style Veggie Pizza");
        }else if (type.equals("clam")){
            pizza = new ClamPizza(ingredientFactory);
            pizza.setName("New York Style Clam Pizza");
        }else if (type.equals("veggie")){
            pizza = new Pepperoni(ingredientFactory);
            pizza.setName("New York Style Pepperoni Pizza");
        }

        return pizza;
    }
}
