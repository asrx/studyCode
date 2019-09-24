package Singleton.Demo2;

public class Singleton5 {
    private static class SingletonHolder{
        private static final Singleton5 INStANCE = new Singleton5();
    }
    private Singleton5(){}

    public static final Singleton5 getInstance(){
        return SingletonHolder.INStANCE;
    }

}
