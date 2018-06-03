public class Magnitud {
    private Integer value;

    public Magnitud(int value) {
        this.value = value;
    }

    public Magnitud() {
        this.value = 0;
    }

    @Override
    public boolean equals(Object other) {
        Magnitud m = (Magnitud)other;

        if ( m != null) {
            return this.value.equals(m.value);
        }

        return false;
    }

    public boolean le(Magnitud needed) {
        return this.value.compareTo(needed.value) == 0 || this.value.compareTo(needed.value) == -1;
    }

    public boolean eq(Magnitud needed) {
        return this.value.compareTo(needed.value) == 0;
    }
}
