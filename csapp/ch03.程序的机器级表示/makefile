clean:
	rm -rf a.out prog

main_demo: clean
	gcc -o prog main.c mstore.c
	./prog
	objdump -d prog
	rm -rf prog

exchange_demo: clean
	gcc -o a.out exchange_demo.c exchange.c
	./a.out

scale_demo: clean
	gcc -Og -S scale.c
	cat scale.s

store_uprod_demo: clean
	gcc -Og -S store_uprod.c

loop_while_demo: clean
	gcc -Og -S loop_while.c

