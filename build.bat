echo off
go generate
go build -tags vfs -ldflags "-s -w"