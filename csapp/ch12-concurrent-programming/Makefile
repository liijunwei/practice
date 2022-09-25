LDLIBS = -lpthread

CC = gcc -std=gnu99
CSAPP_INC = ../include
CSAPP_SRC = ../src
CFLAGS = -I $(CSAPP_INC) $(LDLIBS) -Wall -g

csapp.o:
	@$(CC) $(CFLAGS) -c $(CSAPP_SRC)/csapp.c -o csapp.o

conc_badcnt: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_badcnt.c
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000

conc_goodcnt: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_goodcnt.c
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000
	./a.out 100000000

conc_sharing: csapp.o
	$(CC) $(CFLAGS) csapp.o conc_sharing.c
	./a.out

# OK 为什么有错误?
# time ./a.out 1 32
# Error: result=9223372034707292160
# ./a.out 1 32  70.94s user 0.05s system 99% cpu 1:11.04 total
#
# 原因: 计算的中间值产生了long类型的最大值溢出
# 处理方法: 先做除法, 再做乘法, 并且保证做除法的那个数是偶数
conc_psum-mutex: clean format csapp.o
	$(CC) $(CFLAGS) csapp.o conc_psum-mutex.c
	./a.out

conc_psum-array: clean format csapp.o
	$(CC) $(CFLAGS) csapp.o conc_psum-array.c
	./a.out

conc_psum-local: clean format csapp.o
	$(CC) $(CFLAGS) csapp.o conc_psum-local.c
	./a.out

conc_race: clean format csapp.o
	@$(CC) $(CFLAGS) csapp.o conc_race.c
	./a.out
	./a.out
	./a.out

.PHONY: clean
clean:
	@rm csapp.o

.PHONY: format
format:
	@clang-format -style=LLVM -i *.c
