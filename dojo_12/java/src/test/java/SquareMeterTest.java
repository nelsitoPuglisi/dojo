import org.junit.Test;

import static junit.framework.TestCase.assertTrue;
import static org.junit.Assert.assertFalse;

public class SquareMeterTest {
    @Test
    public void testLeTrue() {
        assertTrue( new SquareMeter(10).le( new SquareMeter(20)));
    }

    @Test
    public void testLeFalse() {
        assertFalse( new SquareMeter(30).le( new SquareMeter(20)));
    }
}
