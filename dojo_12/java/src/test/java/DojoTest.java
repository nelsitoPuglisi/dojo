import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

/**
 * Tests for the dojo.
 */
public class DojoTest {
	
	@Before
	public void setup() {

	}

    @Test
    public void test_usuario_pide_aula_para_albergar_10_estudiantes_como_minimo_label_no_vidente() {
        ClassRoom labA = new ClassRoom(new Label("Lab A"), new Capacity(5));
        ClassRoom labB = new ClassRoom(new Label("Lab B"), new Capacity(10));

        RequestResult result = new Request(labA, labB).havingAtLeast( new Capacity(10));
        ClassRoom actual = result.chooseOne();
        Assert.assertEquals(labB.labelForBlind(), actual.labelForBlind());
    }

    @Test
    public void test_usuario_pide_aula_para_albergar_10_estudiantes_como_minimo() {
        ClassRoom labA = new ClassRoom(new Label("Lab A"), new Capacity(5));
        ClassRoom labB = new ClassRoom(new Label("Lab B"), new Capacity(10));

        RequestResult result = new Request(labA, labB).havingAtLeast( new Capacity(10));

        Assert.assertEquals(labB, result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_para_albergar_30_estudiantes_como_minimo_y_obtiene_ninguna() {
        ClassRoom labA = new ClassRoom(new Label("Lab A"), new Capacity(5));
        ClassRoom labB = new ClassRoom(new Label("Lab B"), new Capacity(10));

        RequestResult result = new Request(labA, labB).havingAtLeast( new Capacity(30));

        Assert.assertEquals(new NoRoom(), result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_con_30_m2_como_minimo_y_obtiene_una() {
        ClassRoom labA = new ClassRoom(
                new Label("Lab A"),
                new Capacity(5),
                new SquareMeter(45)
        );
        ClassRoom labB = new ClassRoom(new Label("Lab B"), new Capacity(10));

        RequestResult result = new Request(labA, labB).havingAtLeast( new SquareMeter(30));

        Assert.assertEquals( labA, result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_con_30_m2_como_minimo_y_obtiene_ninguna() {
        ClassRoom labA = new ClassRoom(
                new Label("Lab A"),
                new Capacity(5),
                new SquareMeter(45)
        );
        ClassRoom labB = new ClassRoom(new Label("Lab B"), new Capacity(10));

        RequestResult result = new Request(labA, labB).havingAtLeast( new SquareMeter(55));

        Assert.assertEquals( new NoRoom(), result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_con_30_m2_y_capacidad_20_alumnos_como_minimo_y_obtiene_una() {
        ClassRoom labA = new ClassRoom(
                new Label("Lab A"),
                new Capacity(5),
                new SquareMeter(45)
        );
        ClassRoom labB = new ClassRoom(
                new Label("Lab B"),
                new Capacity(10),
                new SquareMeter(32)
        );

        RequestResult result = new Request(labA, labB)
                .havingAtLeast(new Capacity(10))
                .havingAtLeast( new SquareMeter(30));

        Assert.assertEquals( labB, result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_con_30_m2_y_capacidad_20_con_20_computadoras_como_minimo_y_obtiene_una() {
        ClassRoom labA = new ClassRoom(
                new Label("Lab A"),
                new Capacity(5),
                new SquareMeter(45)
        );
        ClassRoom labB = new ClassRoom(
                new Label("Lab B"),
                new Capacity(10),
                new SquareMeter(32),
                new Computers(20)
        );

        RequestResult result = new Request(labA, labB)
                .havingAtLeast(new Capacity(10))
                .havingAtLeast( new SquareMeter(30))
                .havingAtLeast( new Computers(20));

        Assert.assertEquals( labB, result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_con_30_m2_y_capacidad_20_con_20_computadoras_exactos_obtiene_una() {
        ClassRoom labA = new ClassRoom(
                new Label("Lab A"),
                new Capacity(5),
                new SquareMeter(45)
        );
        ClassRoom labB = new ClassRoom(
                new Label("Lab B"),
                new Capacity(10),
                new SquareMeter(30),
                new Computers(20)
        );

        RequestResult result = new Request(labA, labB)
                .havingExactly(new Capacity(10))
                .havingExactly( new SquareMeter(30))
                .havingExactly( new Computers(20));

        Assert.assertEquals( labB, result.chooseOne());
    }

    @Test
    public void test_usuario_pide_aula_con_30_m2_y_capacidad_20_con_20_computadoras_exactos_obtiene_ninguna() {
        ClassRoom labA = new ClassRoom(
                new Label("Lab A"),
                new Capacity(5),
                new SquareMeter(45)
        );
        ClassRoom labB = new ClassRoom(
                new Label("Lab B"),
                new Capacity(10),
                new SquareMeter(30),
                new Computers(20)
        );

        RequestResult result = new Request(labA, labB)
                .havingExactly(new Capacity(10))
                .havingExactly( new SquareMeter(30))
                .havingExactly( new Computers(21));

        Assert.assertEquals( new NoRoom(), result.chooseOne());
    }

}
