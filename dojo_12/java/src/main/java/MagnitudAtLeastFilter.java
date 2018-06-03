public class MagnitudAtLeastFilter implements Filter {
    protected final Magnitud needed;
    private final String key;

    public MagnitudAtLeastFilter(Magnitud needed, String key) {
        this.needed = needed;
        this.key = key;
    }

    @Override
    public boolean fulfilledBy(DataBag data) {

        return this.fulfilledBy((Magnitud) data.value(this.key));
    }

    public boolean fulfilledBy(Magnitud value) {
        return this.needed.le(value);
    }
}
