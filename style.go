package owy

type Style struct {
	Stages []string
	Interval int
}

var (
	LineStyle = Style{
		[]string{"-", "\\", "|", "/"},
		60,
	}
	TripleStyle = Style{
		[]string{"- - -", "\\ \\ \\", "| | |", "/ / /"},
		60,
	}
	DotStyle = Style{
		[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		60,
	}
)

