# Learning GO!

All my GO tutorials

## Setup

### asdf package manager install

- https://github.com/asdf-vm/asdf
- https://github.com/kennyp/asdf-golang

```bash
asdf plugin-add golang https://github.com/kennyp/asdf-golang.git
asdf install golang latest
asdf list
asdf global golang 1.16.3
asdf current
go help
```

### VSCode

- https://marketplace.visualstudio.com/items?itemName=golang.go
- https://github.com/golang/vscode-go/blob/master/README.md

### Changing default`$GOPATH`

```bash
echo "export GOPATH=$HOME/code/go" >> ~/.bashrc
source ~/.bashrc

go env -w GOPATH=$HOME/code/go
```

- https://github.com/golang/go/wiki/SettingGOPATH

### Initializing go.mod

```bash
go mod init
```

## Notes

If it starts with an upper case it's public, else it's private. Genius!

Methods are implemented with a reciever declaration:

```go
// Method copies reciever
func (receiver struct_type) Method(arg int) {
	receiver.value *= arg
}

// Method manipulates a pointer receiver
func (receiver *struct_type) Method(arg int) {
	receiver.value *= arg
}
```

## Docs

- https://golang.org/doc/articles/wiki/

## Resources

- https://blog.golang.org/maps
- https://go101.org/article/type-system-overview.html
- http://www.golangpatterns.info/object-oriented/constructors
