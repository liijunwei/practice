import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class TestDate {
    @Test
    public void testDate() {
        Date date = new Date(12, 2, 2022);

        assertEquals(date.month(), 12);
        assertEquals(date.day(), 2);
        assertEquals(date.year(), 2022);

        assertEquals("12/2/2022", date.toString());

        Date anotherDate = new Date(12, 2, 2022);
        assertEquals(date, anotherDate);
    }
}
