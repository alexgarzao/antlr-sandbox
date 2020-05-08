ANTLR := java -Xmx500M -cp "$(HOME)/apps/antlr-4.8-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool

gen-grammar:
	$(ANTLR) -Dlanguage=Go -o parser GoLexer.g4 GoParser.g4

run:
	go run main.go
