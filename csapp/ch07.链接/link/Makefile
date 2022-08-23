clean:
	rm -rf ./prog

main: clean
	gcc -Og -o prog main.c sum.c
	./prog

verbose_main: clean
	gcc -Og -v -o prog main.c sum.c
	./prog
