# plugin-interface

All plugins must follow this interface:
```go
type IPlugin interface {
	Init() error
	Name() string
	Version() string
	Run() error
	ParseOutput() string
}
```

You need to have installed protobuf before you run any commands in the Makefile.  
For example, if you're using a mac, please run `brew install protobuf` before running `make install`.
