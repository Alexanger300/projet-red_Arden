package css

const (
	Reset      = "\033[0m"
	Yellow     = "\033[33m"
	Green      = "\033[32m"
	SteelBlue  = "\033[38;5;27m"
	Red        = "\033[31m"
	Violet     = "\033[35m"
	Orange     = "\033[38;5;208m"
	LightBlue  = "\033[94m"
	Gold       = "\033[38;5;220m"
	LightGreen = "\033[38;5;120m"
	Gray       = "\033[90m"
	Beige      = "\033[38;5;229m"
	Bold       = "\033[1m"
	Blue       = "\033[34m"
	Underline  = "\033[4m"
)

func Clear() {
	print("\033[H\033[2J")
}
