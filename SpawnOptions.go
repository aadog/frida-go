package frida_go

const (
	FRIDA_STDIO_INHERIT = iota
	FRIDA_STDIO_PIPE
)
type FridaStdio int32
type SpawnOptions struct {
	Argv []string
	Envp []string
	Env []string
	Cwd string
	Stdio FridaStdio
}
