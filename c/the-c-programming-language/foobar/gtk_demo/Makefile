# brew install gtk+3
# brew install gtk4

all: clean
	gcc 007_gtk.c -o a.out `pkg-config --cflags gtk+-3.0` `pkg-config --libs gtk+-3.0`
	GDK_SCALE=3 ./a.out

clean:
	rm -rf a.out

tuto01: clean
	gcc tuto01.c -o a.out `pkg-config --cflags gtk+-3.0` `pkg-config --libs gtk+-3.0`
	./a.out

tuto02: clean
	gcc tuto02.c -o a.out `pkg-config --cflags gtk4` `pkg-config --libs gtk4`
	./a.out

tuto03: clean
	gcc tuto03.c -o a.out `pkg-config --cflags gtk4` `pkg-config --libs gtk4`
	./a.out

