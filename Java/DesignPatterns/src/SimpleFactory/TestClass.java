package SimpleFactory;

import SimpleFactory.Demo2.CJG.CJGPizzaStore;
import SimpleFactory.Demo2.NY.NYPizzaStore;
import SimpleFactory.Demo2.Pizza;
import SimpleFactory.Demo2.PizzaStore;

public class TestClass {

    public static void main(String[] args) {

        PizzaStore nyStore = new NYPizzaStore();
        PizzaStore cjgStore = new CJGPizzaStore();

        Pizza pizza = nyStore.orderPizza("cheese");
        System.out.println("Ethan ordered a " + pizza.getName() + "\n");

        pizza = cjgStore.orderPizza("cheese");
        System.out.println("Joel ordered a " + pizza.getName() + "\n");
    }
}
