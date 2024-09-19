# plugin-interface

All plugins must follow this interface:
```go
type IPlugin interface {
	Name() string
	Version() string
	Run() string
	ParseOutput() string
}
```
