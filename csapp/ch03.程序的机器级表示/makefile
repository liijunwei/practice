clean:
	rm -rf a.out prog

main_demo: clean
	gcc -o prog main.c mstore.c
	./prog

exchange_demo: clean
	gcc -o a.out exchange_demo.c exchange.c
	./a.out

scale_demo: clean
	gcc -Og -S scale.c
	cat scale.s
