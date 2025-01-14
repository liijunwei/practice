# ruby: 运行时间 49ms 占用内存 13068KB
# go:   运行时间 3ms  占用内存 844KB
input = STDIN.read
puts input.split("\n").map {|e| e.split(',').sort.join(',')}.join("\n")
