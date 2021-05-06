default: compile clean run

.PHONY: compile run clean

compile:
	ghc main.hs

run: main
	./main

main:
	ghc main.hs

clean:
	rm -f main.o main.hi
