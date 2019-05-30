package explore

import "io"

// since there is no state for commands (yet) we'll just use a function
type Command func(adventure *Adventure, cmd []string, out io.Writer) bool

var commandMap = make(map[string]Command)

const (
	welcome = `
Welcome {{ if .Turns }}back {{ end }}to {{.Name}}
{{.Description}}

Type help for additional instructions.
Ready...
{{$cl :=  index .Locations .CurrentLocation}}
You are in a {{$cl.Name}}

=>`
)

func init() {
	commandMap["bye"] = exit
	commandMap["exit"] = exit
	commandMap["end"] = exit

	commandMap["move"] = move
	commandMap["mv"] = move
	commandMap["go"] = move
	commandMap["walk"] = move
	commandMap["run"] = move

}

// ProcessCommand finds the proper command and executes it.
// The first element in cmd must be the command.  It is assumed that all of cmd are in lower case.
// Only single work commands are supported, sub commands can be used if supported by the particular command
func ProcessCommand(a *Adventure, cmd []string, out io.Writer) bool {
	if len(cmd) == 0 {

		return false
	}

	c, ok := commandMap[cmd[0]]

	if !ok {
		return false
	}
	return c(a, cmd, out)

}

func exit(_ *Adventure, _ []string, out io.Writer) bool {
	_, _ = out.Write([]byte("Thank you for exploring with us.\nGoodbye..."))
	return true
}

var directionMap = map[string]string{
	"n":     "n",
	"north": "n",
	"s":     "s",
	"south": "s",
}

func move(a *Adventure, args []string, out io.Writer) bool {

	if len(args) < 2 {
		_, _ = out.Write([]byte("A direction is required.\n=>"))
		return false
	}
	d := directionMap[args[1]]
	if len(d) == 0 {
		_, _ = out.Write([]byte("Not a valid direction.\n=>"))
		return false
	}
	cl := a.Locations[a.CurrentLocation]
	nl := cl.Paths[d]
	if nl == nil {
		_, _ = out.Write([]byte("You can't go that way.\n=>"))
		return false
	}
	a.CurrentLocation = nl.LocationID

	_, _ = out.Write([]byte("\nYou are in a " + a.Locations[a.CurrentLocation].Name() + "\n=>"))
	return false
}
