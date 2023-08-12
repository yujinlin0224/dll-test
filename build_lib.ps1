$Env:CGO_ENABLED = 1
go build -o lib.dll -buildmode=c-shared lib.go
$Env:CGO_ENABLED = 0
