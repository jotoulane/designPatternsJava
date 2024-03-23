public class Operation {
    private double numberA = 0;
    private double numberB = 0;

    public double getNumberA() {
        return numberA;
    }

    public double getNumberB() {
        return numberB;
    }

    public void setNumberA(double numberA) {
        this.numberA = numberA;
    }

    public void setNumberB(double numberB) {
        this.numberB = numberB;
    }

    public double GetResult() {
        double res = 0;
        return res;
    }
}

class OperationAdd extends Operation {
    @Override
    public double GetResult() {
        double res;
        res = getNumberA() + getNumberB();
        return res;
    }
}

class OperationSub extends Operation {
    @Override
    public double GetResult() {
        double res;
        res = getNumberA() - getNumberB();
        return res;
    }
}

class OperationMul extends Operation {
    @Override
    public double GetResult() {
        double res;
        res = getNumberA() * getNumberB();
        return res;
    }
}

class OperationDiv extends Operation {
    @Override
    public double GetResult() {
        double res;
        res = getNumberA() / getNumberB();
        return res;
    }
}
