package plugin_interface

type IPlugin interface {
	Name() string
	Version() string
	Run() string
	ParseOutput() string
}
