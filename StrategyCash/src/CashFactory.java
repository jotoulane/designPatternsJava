class CashFactory {
    public static CashSuper createCashAccept(String type) {
        CashSuper cs = null;
        switch (type) {
            case "正常收费":
                cs = new CashNormal();
                break;
            case "满300返100":
                CashReturn cs1 = new CashReturn(300, 100);
                cs = cs1;
                break;
            case "打八折":
                CashRebate cs2 = new CashRebate(0.8);
                cs = cs2;
                break;
        }
        return cs;
    }
}
