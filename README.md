# plugin-interface

All plugins must follow this interface:
```go
type IPlugin interface {
	Setup() error
	Name() string
	Version() string
	Run() (error, string)
	ParseOutput() string
}
```
