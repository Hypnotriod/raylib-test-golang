# raylib template for VS Code on Windows using golang
This is the template for `VS Code` on `Windows` of `raylib` C graphic library using golang.  

## Setup environment
* Download and install [Go](https://go.dev/doc/install)
* Download and install [raylib_installer_vX.X.mingw.64bit.exe](https://github.com/raysan5/raylib/releases)  
* Add `C:\raylib\w64devkit\bin` value to `PATH` in *system environment variables* to provide gcc compiler to Go  
* Add system environment variable `CGO_ENABLED` with value `1`

## Build and run
* Open `main.go` file and press `F5` to run debug. For the first time VSCode may prompt you to install `go tools`, so go for it:)
* To build release version press `Ctrl` + `~` to open terminal window in VSCode and type `make release` or `make release PROGRAM_NAME=your_game_name.exe`
* You may need to modify `winres/winres.json` file to provide application info. Checkout [winres](https://github.com/tc-hib/go-winres) repo for more info.
