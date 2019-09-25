package SimpleFactory.Demo2.CJG;

import SimpleFactory.Demo2.Pizza;

public class CJGStyleCheesePizza extends Pizza {
    public CJGStyleCheesePizza() {
        name = "Chicago Style Sauce and cheese Pizza";
        dough = "Extra Thin Crust Dough";
        sauce = "Plum Tomato Sauce";

        toppings.add("Shredded Mozzarella Cheese");
    }


}
