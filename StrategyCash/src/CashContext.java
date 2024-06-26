class CashContext {
    CashSuper cs;

    public CashContext(String type) {
        switch (type) {
            case "正常收费":
                cs = new CashNormal();
                break;
            case "满300返100":
                cs = new CashReturn(300, 100);
                break;
            case "打八折":
                cs = new CashRebate(0.8);
                break;
        }
    }

    public double GetResult(double money) {
        return cs.acceptCash(money);
    }
}
