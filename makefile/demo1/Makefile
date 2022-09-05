# https://www.youtube.com/watch?v=zeEMISsjO38

CC = gcc

blah: blah.o
	$(CC) blah.o -o blah # runs third

blah.o: blah.c
	$(CC) -c blah.c -o blah.o # runs second

blah.c:
	echo "int main() {return 0;}" > blah.c # runs first

clean:
	rm blah*
