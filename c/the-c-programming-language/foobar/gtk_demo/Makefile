all: clean
	gcc 007_gtk.c -o a.out `pkg-config --cflags gtk+-3.0` `pkg-config --libs gtk+-3.0`
	GDK_SCALE=3 ./a.out

clean:
	rm -rf a.out

