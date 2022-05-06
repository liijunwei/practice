# did the user request help?
p ARGV.include?('--help') ? "usage" : nil

# get the value of the -c option
p ARGV.include?('-c') && (config = ARGV[ARGV.index('-c') + 1]) ? config : nil

