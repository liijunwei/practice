import java.util.Objects;

public class Date {
    private final int month;
    private final int day;
    private final int year;

    public Date(int m, int d, int y) {
        month = m;
        day = d;
        year = y;
    }

    public int month() {
        return month;
    }

    public int day() {
        return day;
    }

    public int year() {
        return year;
    }

    @Override
    public String toString() {
        return month() + "/" + day() + "/" + year();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Date that = (Date) o;
        return month == that.month && day == that.day && year == that.year;
    }

    @Override
    public int hashCode() {
        return Objects.hash(month, day, year);
    }
}
