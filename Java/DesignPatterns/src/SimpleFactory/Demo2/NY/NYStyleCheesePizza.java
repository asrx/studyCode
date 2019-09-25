package SimpleFactory.Demo2.NY;

import SimpleFactory.Demo2.Pizza;

public class NYStyleCheesePizza extends Pizza {
    public NYStyleCheesePizza() {
        name = "NY Style Sauce and cheese Pizza";
        dough = "Thin Crust Dough";
        sauce = "Marinara Sauce";

        toppings.add("Grated Reggiano Cheese");
    }

    @Override
    public void cut() {
        System.out.println("Cutting the pizza into seuare slices");
    }
}
