package plugin_interface

type IPlugin interface {
	Setup() error
	Name() string
	Version() string
	Run() (error, string)
	ParseOutput() string
}
