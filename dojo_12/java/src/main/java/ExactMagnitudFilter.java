public class ExactMagnitudFilter implements Filter {
    protected final Magnitud needed;
    private final String key;

    public ExactMagnitudFilter(Magnitud needed, String key) {
        this.needed = needed;
        this.key = key;
    }

    @Override
    public boolean fulfilledBy(DataBag data) {

        return this.fulfilledBy((Magnitud) data.value(this.key));
    }

    public boolean fulfilledBy(Magnitud value) {
        return this.needed.eq(value);
    }
}
