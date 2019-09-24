package Singleton;

import Singleton.Demo1.SingleObject;

public class TestClass {
    public static void main(String[] args) {
        SingleObject obj = SingleObject.getInstance();

        obj.showMessage();
    }
}
