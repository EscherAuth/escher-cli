package runner

import (
	"fmt"
	"regexp"
	"strconv"
)

func (r *runner) setPortInCommandArgs() error {
	transformedArgs := make([]string, 0, len(r.command.Args))

	srcPort, isGiven := r.env.Port.Source()

	if !isGiven {
		return nil
	}

	expression := fmt.Sprintf(`(\b%v\b)`, strconv.Itoa(srcPort))
	rgx, err := regexp.Compile(expression)

	if err != nil {
		return err
	}

	for _, arg := range r.command.Args {
		transformedArgs = append(transformedArgs, rgx.ReplaceAllLiteralString(arg, r.env.EnvDifferencesForSubProcess()["PORT"]))
	}

	r.command.Args = transformedArgs

	return nil
}
