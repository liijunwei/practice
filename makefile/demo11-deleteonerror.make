# https://makefiletutorial.com/#delete_on_error

# The make tool will stop running a rule (and will propogate back to prerequisites) if a command returns a nonzero exit status.
# It's a good idea to always use this, even though make does not for historical reasons.

.DELETE_ON_ERROR:
all: one two

one:
	touch one
	false

two:
	touch two
	false
