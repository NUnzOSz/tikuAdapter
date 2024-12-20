REM call buildweb.bat
go build -o tikuAdapter.exe -ldflags "-s -w -extldflags '-static' -H windowsgui" ./cmd/adapter-service
taskkill /F /IM tikuAdapter.exe
move /Y tikuAdapter.exe ..\ 
cd ..
.\tikuAdapter.exe
