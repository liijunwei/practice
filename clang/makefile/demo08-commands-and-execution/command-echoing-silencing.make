# https://makefiletutorial.com/#command-echoingsilencing

# Add an @ before a command to stop it from being printed
# You can also run make with -s to add an @ before each line

all:
	echo "This make line will not be printed1"
	@echo "This make line will not be printed2"
	echo "But this will"
