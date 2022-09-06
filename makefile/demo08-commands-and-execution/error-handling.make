# Add -k when running make to continue running even in the face of errors. Helpful if you want to see all the errors of Make at once.
# Add a - before a command to suppress the error
# Add -i to make to have this happen for every command.

blah:
	# This error will be printed but ignored, and make will continue to run
	-false
	touch blah


