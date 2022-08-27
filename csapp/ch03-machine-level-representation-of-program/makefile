clean:
	rm -rf a.out prog

main_demo: clean
	gcc -o prog main.c mstore.c
	./prog
	objdump -d prog

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

recursive_factorial_demo:
	gcc -Og -S rfact.c
	cat rfact.s

gets_demo:
	gcc -Og -S  gets-demo.c
	cat gets-demo.s


