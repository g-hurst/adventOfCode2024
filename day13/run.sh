

export CGO_CFLAGS="-I/usr/include/lpsolve"
export CGO_LDFLAGS="-llpsolve55 -lm -ldl -lcolamd"
go run main.go
