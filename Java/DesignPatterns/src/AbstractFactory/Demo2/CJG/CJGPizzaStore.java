package AbstractFactory.Demo2.CJG;

import AbstractFactory.Demo2.Factory.CJGPizzaIngredientFactory;
import AbstractFactory.Demo2.Factory.PizzaIngredientFactory;
import AbstractFactory.Demo2.Pizza;
import AbstractFactory.Demo2.PizzaStore;
import AbstractFactory.Demo2.Pizzas.*;

public class CJGPizzaStore extends PizzaStore {
    @Override
    protected Pizza createPizza(String type) {
        Pizza pizza = null;
        PizzaIngredientFactory ingredientFactory =
                new CJGPizzaIngredientFactory();


        if (type.equals("cheese")){
            pizza = new CheesePizza(ingredientFactory);
            pizza.setName("芝加哥 Style Cheese Pizza");
        }else if (type.equals("pepperoni")){
            pizza = new VeggiePizza(ingredientFactory);
            pizza.setName("芝加哥 Style Veggie Pizza");
        }else if (type.equals("clam")){
            pizza = new ClamPizza(ingredientFactory);
            pizza.setName("芝加哥 Style Clam Pizza");
        }else if (type.equals("veggie")){
            pizza = new Pepperoni(ingredientFactory);
            pizza.setName("芝加哥 Style Pepperoni Pizza");
        }

        return pizza;
    }
}
