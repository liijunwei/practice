# 生产者-消费者 模型
# https://jyywiki.cn/OS/2023/build/lect9.ipynb

go-demo:
	gofmt -w demo.go
	go build -o a.out demo.go && ./a.out 3 4 | python3 pc-check.py 3 100000

ruby-demo:
	ruby demo.rb 1 1 | python3 pc-check.py 1 10000 # 单线程速度还可以...
# 	ruby demo.rb 3 4 | python3 pc-check.py 3 10000

java-demo:
	javac demo.java && java demo 3 4 | python3 pc-check.py 3 100000

fish-demo:
	ruby fish.rb 1 1
