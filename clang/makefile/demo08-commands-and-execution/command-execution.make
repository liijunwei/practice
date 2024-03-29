# Each command is run in a new shell (or at least the effect is as such)

all:
	cd ..
	# The cd above does not affect this line, because each command is effectively run in a new shell
	echo `pwd`

	# This cd command affects the next because they are on the same line
	cd ..;echo `pwd`

	# Same as above
	cd ..; \
	echo `pwd`

# The default shell is /bin/sh. You can change this by changing the variable SHELL:
SHELL=/bin/bash


