bar =
foo = $(bar)

all:
ifdef foo
	@echo "foo is defined"
endif

ifndef bar
	@echo "but bar is not"
endif
