// page 125

// practice && cd csapp/ch03.程序的机器级表示/ && gcc -Og -S exchange.c && cat exchange.s

long exchange(long *xp, long y) {
  long x = *xp;
  *xp = y;

  return x;
}
