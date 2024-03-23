public class Main {
    public static void main(String[] args) {
        Operation op;
        op = OperationFactory.createOperate("+");
        op.setNumberA(1);
        op.setNumberB(1);
        double res = op.GetResult();
        System.out.println(res);
    }
}
