GO = go

RUN = run
BUILD = build

SOURCE = main.go

EXEC = main

all: run

run:
	$(GO) $(RUN) $(SOURCE)

build:
	$(GO) $(BUILD) -o $(EXEC) $(SOURCE)

run_build:
	./$(EXEC)

clean:
	rm -rf $(EXEC)