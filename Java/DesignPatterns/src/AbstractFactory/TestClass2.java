package AbstractFactory;

import AbstractFactory.Demo1.AbstractFactory;
import AbstractFactory.Demo1.FactoryProducer;
import AbstractFactory.Demo1.interfaces.Color;
import AbstractFactory.Demo1.interfaces.Shape;
import AbstractFactory.Demo2.NY.NYPizzaStore;
import AbstractFactory.Demo2.Pizza;

public class TestClass2 {
    public static void main(String[] args) {
        NYPizzaStore nyPizzaStore = new NYPizzaStore();
        Pizza pizza = nyPizzaStore.orderPizza("cheese");

        pizza.toString();
    }
}
