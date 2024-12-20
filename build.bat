call buildweb.bat
go build -o tikuAdapter.exe -ldflags "-s -w -extldflags '-static'" ./cmd/adapter-service
pause
