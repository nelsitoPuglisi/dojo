import java.util.Arrays;
import java.util.List;

public class Request {
    private final List<ClassRoom> classRooms;

    public Request(ClassRoom... classRooms) {
        this.classRooms = Arrays.asList(classRooms);
    }

    public RequestResult havingAtLeast( Capacity c) {
        CapacityAtLeastFilter f = new CapacityAtLeastFilter(c);
        return new RequestResult(this.classRooms, f);
    }

    public RequestResult havingAtLeast(SquareMeter squareMeter) {
        SquareMeterAtLeastFilter f = new SquareMeterAtLeastFilter(squareMeter);
        return new RequestResult(this.classRooms, f);
    }

    public RequestResult havingExactly(Capacity c) {
        ExactlyCapacityFilter f = new ExactlyCapacityFilter(c);
        return new RequestResult(this.classRooms, f);
    }
}
