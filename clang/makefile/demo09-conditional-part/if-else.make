# https://makefiletutorial.com/#conditional-part-of-makefiles

# make -f if-else.make foo=ok
# make -f if-else.make foo=notok
# foo = ok

all:
ifeq ($(foo), ok)
	@echo "foo equals ok"
else
	@echo "nope"
endif
