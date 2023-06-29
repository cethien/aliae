// doskey np=notepad++.exe $*

package shell

const (
	CMD = "cmd"
)

func (a *Alias) Cmd() *Alias {
	if a.Type == Command {
		a.template = `local p = assert(io.popen("doskey {{ .Alias }}={{ .Value }}"))
p:close()`
	}

	return a
}