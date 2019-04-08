# build for win, mac, linux
GOOS=windows go build -ldflags "-s -w";zip dots-win64.zip dots.exe;GOOS=darwin go build -ldflags "-s -w";zip dots-mac.zip dots;go build -ldflags "-s -w";tar czf dots-linux.tgz dots
