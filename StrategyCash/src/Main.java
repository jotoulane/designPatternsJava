public class Main {
    static double PRICE = 5d;
    static double NUM = 10d;

    public static void main(String[] args) {
        double total = 0.0d;
        CashContext csuper = new CashContext("正常收费");
        double totalPrices = 0d;
        totalPrices = csuper.GetResult(PRICE * NUM);
        total+=totalPrices;

    }
}
