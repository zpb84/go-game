package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var (
	version = "dev"
	commit  = "n/a"
	date    = "n/a"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.StripEscape)
	fmt.Fprintf(w, "Date:\t%s\nVersion:\t%s\nCommit:\t%s\n", date, version, commit)
	w.Flush()
}
