public class GreatestCommonDivisor {
    public static void main(String[] args) {
      int a = 15;
      int b = 25;
      System.out.println(gcd(a, b));
    }

    public static int gcd(int p, int q) {
      if (q == 0) return p;
      int r = p % q;
      return gcd(q, r);
    }
}
