package Observer;

import Observer.Demo1.BinaryObserver;
import Observer.Demo1.OctalObserver;
import Observer.Demo1.Subject;

public class TestClass {
    public static void main(String[] args) {
        Subject subject = new Subject();

        new OctalObserver(subject);
        new BinaryObserver(subject);

        System.out.println("First state change: 15");
        subject.setState(15);
        System.out.println("");

        System.out.println("Second state change: 10");
        subject.setState(10);
    }
}
