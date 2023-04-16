clean:
	rm -r ./release

release: clean
	go build -ldflags "-H=windowsgui" -o ./release/program.exe .
