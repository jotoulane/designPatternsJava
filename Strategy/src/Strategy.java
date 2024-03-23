//抽象算法类
abstract class Strategy {
    public abstract void AlgorithmInterface();
}


//算法A
class ConcreteStrategyA extends Strategy {

    @Override
    public void AlgorithmInterface() {
        System.out.println("算法A实现");
    }
}

//算法B
class ConcreteStrategyB extends Strategy {

    @Override
    public void AlgorithmInterface() {
        System.out.println("算法B实现");
    }
}

//算法C
class ConcreteStrategyC extends Strategy {

    @Override
    public void AlgorithmInterface() {
        System.out.println("算法C实现");
    }
}
