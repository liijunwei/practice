passwd = File.open('/etc/passwd')
puts passwd.fileno

hosts = File.open('/etc/hosts')
puts hosts.fileno

# Close the open passwd file. The frees up its file descriptor
# number to be used by the next opened resource.
passwd.close

null = File.open('/dev/null')
puts null.fileno

# There are two key takeaways from this example.

# File descriptor numbers are assigned the lowest unused value. The first file we opened, passwd, got file descriptor #3, the next open file got #4 because #3 was already in use.

# Once a resource is closed its file descriptor number becomes available again. Once we closed the passwd file its file descriptor number became available again. So when we opened the file at dev/null it was assigned the lowest unused value, which was then #3.
