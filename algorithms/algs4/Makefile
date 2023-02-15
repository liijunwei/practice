# code
# https://algs4.cs.princeton.edu/code/

TMP_FILE_FOLDER := tmp
ALGS4_DATA_FOLDER := data
JAVA_JARS_FOLDER := $(HOME)/java-jars
DOWNLOAD_CODE_TO := wget --timestamping --directory-prefix

# tutorial: https://www.guru99.com/download-installation-junit.html
# download from: https://github.com/junit-team/junit4/wiki/Download-and-Install
JAVA_CLASSPATH = -classpath .:$(JAVA_JARS_FOLDER)/algs4.jar:$(JAVA_JARS_FOLDER)/hamcrest-core-1.3.jar:$(JAVA_JARS_FOLDER)/junit-4.13.2.jar

setup_folders:
	mkdir -p $(JAVA_JARS_FOLDER)
	mkdir -p $(TMP_FILE_FOLDER)
	mkdir -p $(ALGS4_DATA_FOLDER)

download_jars: setup_folders
	$(DOWNLOAD_CODE_TO) $(JAVA_JARS_FOLDER) https://algs4.cs.princeton.edu/code/algs4.jar
	wget --timestamping -O $(JAVA_JARS_FOLDER)/hamcrest-core-1.3.jar https://search.maven.org/remotecontent?filepath=org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar
	wget --timestamping -O $(JAVA_JARS_FOLDER)/junit-4.13.2.jar https://search.maven.org/remotecontent?filepath=junit/junit/4.13.2/junit-4.13.2.jar

# see README.md
download_data: setup_folders
	ruby parse_data_urls.rb | parallel "$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) {}"

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

testBagQueueStack:
	javac test/testBagQueueStack.java
	java org.junit.runner.JUnitCore test.testBagQueueStack

download_ch1.4:
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/14analysis/1Kints.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/14analysis/2Kints.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/14analysis/4Kints.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/14analysis/8Kints.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/14analysis/16Kints.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/14analysis/1Mints.txt

1.4.2.1:
	java edu.princeton.cs.algs4.ThreeSum data/1Kints.txt
	java edu.princeton.cs.algs4.ThreeSum data/2Kints.txt
	java edu.princeton.cs.algs4.ThreeSum data/4Kints.txt
	java edu.princeton.cs.algs4.ThreeSum data/8Kints.txt
DoublingTest:
	java edu.princeton.cs.algs4.DoublingTest


TwoSumFast:
	javac TwoSumFast.java
	java TwoSumFast data/8Kints.txt
	java TwoSumFast data/16Kints.txt
	java TwoSumFast data/1Mints.txt


ThreeSumFast:
	javac ThreeSumFast.java
	java ThreeSumFast data/8Kints.txt
	java ThreeSumFast data/16Kints.txt
# 	java ThreeSumFast data/1Mints.txt


# https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/UF.java.html
download_ch1.5: setup_folders
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/15uf/tinyUF.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/15uf/mediumUF.txt
	@$(DOWNLOAD_CODE_TO) $(ALGS4_DATA_FOLDER) https://algs4.cs.princeton.edu/15uf/largeUF.txt

build_quick_find:
	@javac $(JAVA_CLASSPATH) union_find/QuickFind.java

test_quick_find: build_quick_find
	@java $(JAVA_CLASSPATH) union_find.QuickFind < $(ALGS4_DATA_FOLDER)/tinyUF.txt
# 	@java $(JAVA_CLASSPATH) union_find.QuickFind < $(ALGS4_DATA_FOLDER)/mediumUF.txt
# 	@java $(JAVA_CLASSPATH) union_find.QuickFind < $(ALGS4_DATA_FOLDER)/largeUF.txt

build_quick_union:
	@javac $(JAVA_CLASSPATH) union_find/QuickUnion.java

test_quick_union: build_quick_union
	@java $(JAVA_CLASSPATH) union_find.QuickUnion < $(ALGS4_DATA_FOLDER)/tinyUF.txt
	@java $(JAVA_CLASSPATH) union_find.QuickUnion < $(ALGS4_DATA_FOLDER)/mediumUF.txt
	@java $(JAVA_CLASSPATH) union_find.QuickUnion < $(ALGS4_DATA_FOLDER)/largeUF.txt

test_junit:
	javac bag_queue_stack/MyQueue.java
	java org.junit.runner.JUnitCore bag_queue_stack.MyQueue

DoublingRatio:
	java edu.princeton.cs.algs4.DoublingRatio
