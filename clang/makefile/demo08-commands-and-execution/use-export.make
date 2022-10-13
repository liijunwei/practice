# https://makefiletutorial.com/#use-export-for-recursive-make

# The export directive takes a variable and makes it accessible to sub-make commands.

new_contents = "hello:\n\techo \$$(cooly)"

all:
	mkdir -p subdir
	printf $(new_contents) | sed -e 's/^ //' > subdir/makefile
	@echo "---MAKEFILE CONTENTS---"
	@cd subdir && cat makefile
	@echo "---END MAKEFILE CONTENTS---"
	cd subdir && $(MAKE)

# Note that variables and exports. They are set/affected globally.
cooly = "The subdirectory can see me!"
export cooly
# This would nullify the line above: unexport cooly

clean:
	rm -rf subdir

one=this will only work locally
export two=we can run subcommands with this

another:
	@echo $(one)
	@echo $$one
	@echo $(two)
	@echo $$two

# .EXPORT_ALL_VARIABLES exports all variables for you.
