file0.txt:
	echo "foo" > file0.txt

file1.txt: file0.txt
	cat file0.txt > file1.txt
	@cat file0.txt >> file1.txt
	-@echo 'hello'

file2.txt file3.txt: file0.txt file1.txt
	touch file2.txt
	touch file3.txt

all: maker process

maker:
	touch ex0.txt ex1.txt

.PHONY: all maker process

ex0.txt ex1.txt: maker

process: file*.txt
	@echo ^...
	@echo $^
	@echo @...
	@echo $@
	@echo \<...
	@echo $<
	@echo ?...
	@echo $?
	@echo +...
	@echo $+

process: ex1.txt file0.txt
