rm -rf ./gosuv_upx
go generate
go build -tags vfs -ldflags "-s -w"
upx -9 -o ./gosuv_upx ./gosuv