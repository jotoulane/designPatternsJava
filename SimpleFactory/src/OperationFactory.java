public class OperationFactory {
    public static Operation createOperate(String operate) {
        Operation open = null;
        switch (operate) {
            case "+":
                open = new OperationAdd();
                break;
            case "-":
                open = new OperationSub();
                break;
            case "*":
                open = new OperationMul();
                break;
            case "/":
                open = new OperationDiv();
                break;
        }
        return open;
    }
}
