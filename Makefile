windows:
	go build -o service-man.exe ./cmd/main.go

linux: export GOOS=linux
linux:
	go build -o service-man.bin ./cmd/main.go
