# raylib template for VS Code on Windows using golang
This is the template for `VS Code` on `Windows` of `raylib` C graphic library using golang.  

## Setup environment
* Download and install [Go](https://go.dev/doc/install)
* Download latest [w64devkit-1.XX.X.zip](https://github.com/skeeto/w64devkit/releases/latest) and unzip it to `C:\` (`C:\w64devkit` folder should appear).  
* Add `C:\w64devkit\bin` value to `PATH` in *system environment variables* to provide gcc compiler to Go.  
* Add system environment variable `CGO_ENABLED` with value `1`

## Run and Debug
* Open `main.go` file in VSCode and press `F5` to run debug. For the first time VSCode may prompt you to install `go tools`, so go for it:)

## Build Release
* `Ctrl` + `~` to open terminal window in VSCode.
* Install `go-winres` by running `go install github.com/tc-hib/go-winres@latest` command in terminal.  
* You may need to modify `winres/winres.json` file to provide application info. Checkout [winres](https://github.com/tc-hib/go-winres) repo for guidance.
* Type `make release` or `make release FILE_NAME=your_file_name.exe` command in terminal.

