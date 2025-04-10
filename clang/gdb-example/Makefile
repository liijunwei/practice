install-debug-info:
	sudo debuginfo-install glibc-2.17-326.el7_9.x86_64

build-demo01:
	gcc demo01.c

run-demo01: build-demo01
	echo 1 | ./a.out # Segmentation fault

debug-demo01:
	gcc demo01.c -g
	gdb a.out

run-gdb-script:
	gcc demo01.c -g
	gdb -x test.gdb a.out
