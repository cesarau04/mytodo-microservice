GOGET = GO GET
GORUN = GO RUN
GOBUILD = GO BUILD
GOMAIN = main.go
GOOTHERS = handlers.go
TARGET = main

all: $(TARGET)
	$(GOGET)
	$(GORUN) $(GOMAIN) $(GOOTHERS)

build: $(TARGET)
	$(GOGET)
	$(GOBUILD) $(GOMAIN) $(GOOTHERS)
	
clean:
	rm -f $(TARGET)