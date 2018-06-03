public class SquareMeterAtLeastFilter extends MagnitudAtLeastFilter {

    public SquareMeterAtLeastFilter(Integer needed) {
        super(new SquareMeter(needed), "squareMeter");
    }

    public SquareMeterAtLeastFilter(SquareMeter needed) {
        super(needed, "squareMeter");
    }
}
