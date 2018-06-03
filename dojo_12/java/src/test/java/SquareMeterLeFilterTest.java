import org.junit.Test;

import static junit.framework.TestCase.assertTrue;
import static org.junit.Assert.assertFalse;

public class SquareMeterLeFilterTest {
    @Test
    public void testSquareMeterIsFulfilledByRoom() {
        assertTrue(new SquareMeterAtLeastFilter(5).fulfilledBy(new SquareMeter(10)) );
    }

    @Test
    public void testSquareMeterIsNotFulfilledByRoom() {
        assertFalse(new SquareMeterAtLeastFilter(20).fulfilledBy(new SquareMeter(10)) );
    }
}
