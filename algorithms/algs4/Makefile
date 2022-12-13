# code
# https://algs4.cs.princeton.edu/code/

TMP_FILE_FOLDER := tmp
JAVA_JARS_FOLDER := $(HOME)/java-jars
DOWNLOAD_CODE_TO := wget --timestamping --directory-prefix

# tutorial: https://www.guru99.com/download-installation-junit.html
# download from: https://github.com/junit-team/junit4/wiki/Download-and-Install
JAVA_CLASSPATH = -classpath .:$(JAVA_JARS_FOLDER)/algs4.jar:$(JAVA_JARS_FOLDER)/hamcrest-core-1.3.jar:$(JAVA_JARS_FOLDER)/junit-4.13.2.jar

setup_folders:
	mkdir -p $(JAVA_JARS_FOLDER)
	mkdir -p $(TMP_FILE_FOLDER)

download_jars: setup_folders
	$(DOWNLOAD_CODE_TO) $(JAVA_JARS_FOLDER) https://algs4.cs.princeton.edu/code/algs4.jar
	wget --timestamping -O $(JAVA_JARS_FOLDER)/hamcrest-core-1.3.jar https://search.maven.org/remotecontent?filepath=org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar
	wget --timestamping -O $(JAVA_JARS_FOLDER)/junit-4.13.2.jar https://search.maven.org/remotecontent?filepath=junit/junit/4.13.2/junit-4.13.2.jar

gcd:
	javac GreatestCommonDivisor.java && java GreatestCommonDivisor

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

test_runner:
	@javac $(JAVA_CLASSPATH) TestRunner.java
	@java $(JAVA_CLASSPATH) TestRunner

test_flips:
	@javac $(JAVA_CLASSPATH) Flips.java
	@java $(JAVA_CLASSPATH) Flips 10

test_flips_max:
	@javac $(JAVA_CLASSPATH) FlipsMax.java
	@java $(JAVA_CLASSPATH) FlipsMax 10000000

test_accumulator:
	@javac $(JAVA_CLASSPATH) TestAccumulator.java
	@java $(JAVA_CLASSPATH) TestAccumulator 10000000

test_visual_accumulator:
	@javac $(JAVA_CLASSPATH) TestVisualAccumulator.java
	@java $(JAVA_CLASSPATH) TestVisualAccumulator 1000

# https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/UF.java.html
download_tinyUF.txt: setup_folders
	@$(DOWNLOAD_CODE_TO) $(TMP_FILE_FOLDER) https://algs4.cs.princeton.edu/15uf/tinyUF.txt
	@$(DOWNLOAD_CODE_TO) $(TMP_FILE_FOLDER) https://algs4.cs.princeton.edu/15uf/mediumUF.txt
	@$(DOWNLOAD_CODE_TO) $(TMP_FILE_FOLDER) https://algs4.cs.princeton.edu/15uf/largeUF.txt

# download_tinyUF.txt -> test_quick_union
build_quick_union:
	@javac $(JAVA_CLASSPATH) union_find/UF.java

test_quick_union_with_tinyUF: build_quick_union
	@java $(JAVA_CLASSPATH) union_find.UF < $(TMP_FILE_FOLDER)/tinyUF.txt

test_quick_union_with_mediumUF: build_quick_union
	@java $(JAVA_CLASSPATH) union_find.UF < $(TMP_FILE_FOLDER)/mediumUF.txt

test_quick_union_with_largeUF: build_quick_union
	@java $(JAVA_CLASSPATH) union_find.UF < $(TMP_FILE_FOLDER)/largeUF.txt




