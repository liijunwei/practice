gcd:
	javac GreatestCommonDivisor.java && java GreatestCommonDivisor

download_jars:
	wget https://algs4.cs.princeton.edu/code/algs4.jar --directory-prefix algs4-jars

LIB_FOLDER := edu/princeton/cs/algs4
# download and put lib into `edu/princeton/cs/algs4/`
# https://algs4.cs.princeton.edu/code/
StdRandom:
	wget --directory-prefix $(LIB_FOLDER) https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/StdRandom.java
	wget --directory-prefix $(LIB_FOLDER) https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/StdDraw.java

binary_search_demo:
# 	javac -Xlint:deprecation BinarySearch.java
	javac BinarySearch.java
	java BinarySearch tinyW.txt < tinyT.txt

random_seq_demo:
	javac RandomSeq.java
	java RandomSeq 3 1 100

average_demo:
	javac Average.java
	java Average < tinyW.txt

draw_function_graph_demo:
	javac DrawFunctionGraph.java
	java DrawFunctionGraph
