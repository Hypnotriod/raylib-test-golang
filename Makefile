PROGRAM_NAME=program.exe

clean:
	rm -rf ./release && rm -rf *.syso

winres-init:
	go-winres init

winres-make:
	go-winres make

release: clean winres-make
	go build -ldflags "-s -w -H=windowsgui" -o ./release/${PROGRAM_NAME} .
