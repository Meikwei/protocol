@echo off
setlocal

rem Define array elements
set "PROTO_NAMES=auth conversation errinfo friend group msg msggateway push rtc sdkws third user statistics wrapperspb"

rem Loop through each element in the array
for %%i in (%PROTO_NAMES%) do (
    rem Generate Go code for Protobuf
    protoc --go_out=./%%i --go_opt=module=github.com/Meikwei/protocol/%%i %%i/%%i.proto
    rem Generate Go gRPC code for Protobuf
    protoc --go-grpc_out=./%%i --go-grpc_opt=module=github.com/Meikwei/protocol/%%i %%i/%%i.proto
    if ERRORLEVEL 1 (
        echo error processing %%i.proto
        exit /b %ERRORLEVEL%
    )
)

rem Replace "omitempty" in *.pb.go files with UTF-8 encoding
for /r %%f in (*.pb.go) do (
    powershell -Command "(Get-Content -Path '%%f' -Encoding UTF8) -replace ',omitempty\"`"', '\"`"' | Set-Content -Path '%%f' -Encoding UTF8"
)


endlocal
