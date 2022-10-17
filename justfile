default: build run
build:
    if [ ! -d build ]; then mkdir ./build; fi
    go build -o ./build/objdeliv ./cmd/main.go
build-cross: build
    GOOS=windows GOARCH=amd64 go build -o ./build/objdeliv-windows-amd64 ./cmd/main.go
    GOOS=windows GOARCH=386 go build -o ./build/objdeliv-windows-i386 ./cmd/main.go
    GOOS=windows GOARCH=arm go build -o ./build/objdeliv-windows-arm ./cmd/main.go
    GOOS=windows GOARCH=arm64 go build -o ./build/objdeliv-windows-arm64 ./cmd/main.go
    GOOS=linux GOARCH=amd64 go build -o ./build/objdeliv-linux-amd64 ./cmd/main.go
    GOOS=linux GOARCH=386 go build -o ./build/objdeliv-linux-i386 ./cmd/main.go
    GOOS=linux GOARCH=arm64 go build -o ./build/objdeliv-linux-arm64 ./cmd/main.go
    GOOS=linux GOARCH=mips go build -o ./build/objdeliv-linux-mips ./cmd/main.go
    GOOS=linux GOARCH=mipsle go build -o ./build/objdeliv-linux-mipsle ./cmd/main.go
    GOOS=linux GOARCH=mips64 go build -o ./build/objdeliv-linux-mips64 ./cmd/main.go
    GOOS=linux GOARCH=mips64le go build -o ./build/objdeliv-linux-mips64le ./cmd/main.go
    GOOS=linux GOARCH=arm go build -o ./build/objdeliv-linux-arm ./cmd/main.go
    GOOS=linux GOARCH=ppc64 go build -o ./build/objdeliv-linux-ppc64 ./cmd/main.go
    GOOS=linux GOARCH=ppc64le go build -o ./build/objdeliv-linux-ppc64le ./cmd/main.go
    GOOS=linux GOARCH=s390x go build -o ./build/objdeliv-linux-s390x ./cmd/main.go
    GOOS=linux GOARCH=riscv64 go build -o ./build/objdeliv-linux-riscv64 ./cmd/main.go
    GOOS=darwin GOARCH=amd64 go build -o ./build/objdeliv-darwin-amd64 ./cmd/main.go
    GOOS=darwin GOARCH=arm64 go build -o ./build/objdeliv-darwin-arm64 ./cmd/main.go
    GOOS=solaris GOARCH=amd64 go build -o ./build/objdeliv-solaris-amd64 ./cmd/main.go
    GOOS=plan9 GOARCH=386 go build -o ./build/objdeliv-plan9-386 ./cmd/main.go
    GOOS=plan9 GOARCH=amd64 go build -o ./build/objdeliv-plan9-amd64 ./cmd/main.go
    GOOS=plan9 GOARCH=arm go build -o ./build/objdeliv-plan9-arm ./cmd/main.go
    # GOOS=android GOARCH=arm go build -o ./build/objdeliv-android-arm ./cmd/main.go
    # GOOS=android GOARCH=arm64 go build -o ./build/objdeliv-android-arm64 ./cmd/main.go
    # GOOS=android GOARCH=386 go build -o ./build/objdeliv-android-386 ./cmd/main.go
    # GOOS=android GOARCH=amd64 go build -o ./build/objdeliv-android-amd64 ./cmd/main.go
    GOOS=openbsd GOARCH=arm go build -o ./build/objdeliv-openbsd-arm ./cmd/main.go
    GOOS=openbsd GOARCH=arm64 go build -o ./build/objdeliv-openbsd-arm64 ./cmd/main.go
    GOOS=openbsd GOARCH=386 go build -o ./build/objdeliv-openbsd-386 ./cmd/main.go
    GOOS=openbsd GOARCH=amd64 go build -o ./build/objdeliv-openbsd-amd64 ./cmd/main.go
    GOOS=freebsd GOARCH=386 go build -o ./build/objdeliv-freebsd-386 ./cmd/main.go
    GOOS=freebsd GOARCH=amd64 go build -o ./build/objdeliv-freebsd-amd64 ./cmd/main.go
    GOOS=freebsd GOARCH=arm go build -o ./build/objdeliv-freebsd-arm ./cmd/main.go
    GOOS=freebsd GOARCH=arm64 go build -o ./build/objdeliv-freebsd-arm64 ./cmd/main.go
    GOOS=openbsd GOARCH=mips64 go build -o ./build/objdeliv-openbsd-mips64 ./cmd/main.go
    GOOS=js GOARCH=wasm go build -o ./build/objdeliv-js-wasm ./cmd/main.go
run:
    ./build/objdeliv
clean:
    rm ./build -Rf