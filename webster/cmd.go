package jen

import (
	"fmt"
)

// Run the command
func Run(args []string) (string, error) {

    if len(args) < 2 {
    	return "", fmt.Errorf("expect two args (src & dst) but only got one")
	}

	// acquire our source and destinations
	src, dst := args[0], args[1]
	if src == "" || dst == "" {
		return "", fmt.Errorf("expected a src or dst but one is missing")
	}

	wt, err := NewTransplantor(src)
	if err != nil {
		return "", fmt.Errorf("new trans src: %s err %v", src, err)
	}

	// Scan sources and generate destination
	err = wt.GenerateSite(dst)
	if err != nil {
		return "", fmt.Errorf("failed on %s err %v", dst, err)
	}
	return "This is the end of the gen.Run() command", err
}
