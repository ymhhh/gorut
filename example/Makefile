all: ddd ddd_sample ddd_isolator

build:
	./build.sh

ddd: build
	gorut create -f -t=ddd
	@echo "ddd done"

ddd_sample: build
	gorut create -f -t=ddd_sample
	./run.sh
	@echo "ddd_sample done"


ddd_isolator: build
	gorut deps
	gorut create -f -t=ddd_isolator
	./run.sh
	@echo "ddd_isolator done"

