package lib

import (
	"fmt"
	"io"
	"text/tabwriter"
)

var (
	version = "dev"
	commit  = "n/a"
	date    = "n/a"
)

func ShowInfo(w io.Writer) {
	tw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.StripEscape)
	fmt.Fprintf(tw, "Date:\t%s\nVersion:\t%s\nCommit:\t%s\n", date, version, commit)
	tw.Flush()
}
