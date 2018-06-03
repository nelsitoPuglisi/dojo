import java.util.ArrayList;
import java.util.List;

public class RequestResult {
    private final List<ClassRoom> classRooms;
    private final List<Filter> filters;

    public RequestResult(List<ClassRoom> classRooms, CapacityAtLeastFilter capacityFilter) {
        this.classRooms = classRooms;
        this.filters = new ArrayList<>();
        this.filters.add(capacityFilter);

    }

    public RequestResult(List<ClassRoom> classRooms, SquareMeterAtLeastFilter smFilter) {
        this.classRooms = classRooms;
        this.filters = new ArrayList<>();
        this.filters.add(smFilter);
    }

    public RequestResult(List<ClassRoom> classRooms, ExactlyCapacityFilter f) {
        this.classRooms = classRooms;
        this.filters = new ArrayList<>();
        this.filters.add(f);
    }

    public RequestResult havingAtLeast( Capacity c) {
        this.filters.add( new CapacityAtLeastFilter(c));
        return this;
    }

    public RequestResult havingAtLeast(SquareMeter sm) {
        this.filters.add( new SquareMeterAtLeastFilter(sm));
        return this;
    }

    public RequestResult havingAtLeast(Computers computers) {
        this.filters.add( new ComputersAtLeastFilter(computers));
        return this;
    }

    public RequestResult havingExactly(SquareMeter squareMeter) {
        this.filters.add( new ExactlySquareFilter(squareMeter));
        return this;
    }

    public RequestResult havingExactly(Computers computers) {
        this.filters.add( new ExactlyComputerFilter(computers));
        return this;
    }

    public ClassRoom chooseOne() {
        for( ClassRoom clazz : this.classRooms) {
            if ( clazz.fitsAll( filters )) {
                return clazz;
            }
        }
        return new NoRoom();
    }

}
