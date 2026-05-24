via github.com/MoeKoeMusic/MoeKoeMusic bd929528ecf4e2c1b45490a25c29b457745f414c
via github.com/MakcRe/KuGouMusicApi 566de02034493af31c5a44f124429cbc1cb11b13

build
CGO_ENABLED=0 go build -ldflags="-s -w" -o moekoe-go .

run
./moekoe-go --port=8080 --platform=lite
