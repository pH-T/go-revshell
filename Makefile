osx:
	GOOS=darwin go build -ldflags "-s -w -X main.shellPath=/bin/bash" -o shell *.go

linux:
	GOOS=linux go build -ldflags "-s -w -X main.shellPath=/bin/bash" -o shell *.go	

win:
	GOOS=windows go build -ldflags "-s -w -X main.shellPath=C:\\Windows\\System32\\cmd.exe" -o shell *.go

winps:
	GOOS=windows go build -ldflags "-s -w -X main.shellPath=C:\\Windows\\system32\\WindowsPowerShell\\v1.0\\powershell.exe" -o shell *.go