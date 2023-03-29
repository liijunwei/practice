# 生产者-消费者 模型
# https://jyywiki.cn/OS/2023/build/lect9.ipynb

go-demo:
	gofmt -w demo.go
	go build -o a.out demo.go && ./a.out 3 4 | python3 pc-check.py 3 100000

ruby-demo:
	ruby demo.rb 3 4 | python3 pc-check.py 3 10

java-demo:
	javac demo.java && java demo 3 4 | python3 pc-check.py 3 100000
