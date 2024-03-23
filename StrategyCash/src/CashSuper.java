//现金收费抽象类
abstract class CashSuper {
    public abstract double acceptCash(double money);
}

//正常收费子类
class CashNormal extends CashSuper {

    @Override
    public double acceptCash(double money) {
        return money;
    }
}

//打折收费子类
class CashRebate extends CashSuper {

    private double moneyRebate = 1d;

    public CashRebate(double moneyRebate) {
        this.moneyRebate = moneyRebate;
    }

    @Override
    public double acceptCash(double money) {
        return money * moneyRebate;
    }
}

//返利收费子类
class CashReturn extends CashSuper {
    private double moneyCondition = 0d;
    private double moneyReturn = 0d;

    public CashReturn(double moneyCondition, double moneyReturn) {
        this.moneyCondition = moneyCondition;
        this.moneyReturn = moneyReturn;
    }

    @Override
    public double acceptCash(double money) {
        double result = money;
        if (money >= moneyCondition) {
            result = money - Math.floor(money / moneyCondition) * moneyReturn;
        }
        return result;
    }
}
