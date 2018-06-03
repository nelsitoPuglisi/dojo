import java.util.List;

public class ClassRoom implements DataBag {
    private final Capacity capacity;
    private final SquareMeter squareMeter;
    private final Computers computers;

    public ClassRoom(Label lab_a, Capacity capacity) {
        this.capacity = capacity;
        this.squareMeter = new SquareMeter();
        this.computers = new Computers();
    }

    public ClassRoom(Label label, Capacity capacity, SquareMeter squareMeter) {
        this.capacity = capacity;
        this.squareMeter = squareMeter;
        this.computers = new Computers();
    }

    public ClassRoom(Label lab_b, Capacity capacity, SquareMeter squareMeter, Computers computers) {
        this.capacity = capacity;
        this.squareMeter = squareMeter;
        this.computers = computers;
    }

    public boolean equals(Object other) {
        return other.getClass().equals(this.getClass());
    }

    public boolean fitsAll(List<Filter> filters) {
        boolean result = true;

        for( Filter f : filters) {
            if ( !f.fulfilledBy(this)) {
                result = false;
                break;
            }
        }

        return result;
    }

    @Override
    public Object value(String key) {
        switch (key) {
            case "squareMeter": return this.squareMeter;
            case "capacity": return this.capacity;
            case "computers": return this.computers;
        }
        return new None();
    }

    public LabelForBlind labelForBlind() {
        return new LabelForBlind();
    }
}
