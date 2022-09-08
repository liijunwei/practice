# https://www.bilibili.com/video/BV1XZ4y1S7e1?p=112&share_source=copy_web

HEADER_PATH = ./ # #include "max.h" 会先在本目录找, 如果找不到才会去这里找

main:
	gcc -I $(HEADER_PATH) main.c max.c
	./a.out

main_intermediate:
	gcc --save-temps -c main.c # main.c 里的 "#include "max.h"" 仅仅是把 "max.h" 里的内容插入到main.c里
	ls -lh
	tail -n20 main.i | grep "max.h" -C 5

clean:
	rm main.i main.s main.o a.out

# man gcc
# /\s-I\s
#   -c: compile but not link, output the .o file
#   -g: Produce debugging information in the operating system's native format
#   -O: optimize code
