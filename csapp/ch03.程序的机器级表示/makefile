clean:
	rm -rf a.out prog

main_demo: clean
	gcc -Og -o prog main.c mstore.c
	./prog

exchange_demo: clean
	gcc -Og -S exchange.c && cat exchange.s
	./a.out
