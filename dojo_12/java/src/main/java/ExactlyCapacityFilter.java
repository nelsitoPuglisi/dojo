public class ExactlyCapacityFilter extends ExactMagnitudFilter {
    public ExactlyCapacityFilter(Magnitud needed) {
        super(needed, "capacity");
    }

    public boolean fulfilledBy(Magnitud value) {
        return this.needed.eq(value);
    }
}
