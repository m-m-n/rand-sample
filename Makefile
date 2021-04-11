GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
EXECFILE=rand

all: build
build:
	$(GOBUILD) -o $(EXECFILE) -v $(EXECFILE).go
clean:
	$(GOCLEAN)
	rm -f $(EXECFILE)
install:
	ln -sfv $(HOME)/go/src/$(EXECFILE)/$(EXECFILE) $(HOME)/bin/$(EXECFILE)
