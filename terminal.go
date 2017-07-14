package terminal

import (
	termbox "github.com/nsf/termbox-go"
	// "strconv"
	"fmt"
)

var (
	x, y    int = 0, 0
	out     string
	TabSize int = 16
)

func init() {
	// if err := termbox.Init(); err != nil {
	// 	panic(err)
	// }
}

func Close() {
	termbox.Close()
}

func Clear(fg, bg termbox.Attribute) error {
	x, y = 0, 0
	return termbox.Clear(fg, bg)
}

func printc(c rune) {
	w, _ := termbox.Size()
	if '\t' == c {
		x = (1 + x/TabSize) * TabSize
	}

	if x >= w || c == '\n' {
		y++
		x = 0
	}

	if '\t' == c || '\n' == c || '\r' == c {
		return
	}

	termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
	x++
}

// func PrintAt(s string, x, y int, attr ...termbox.Attribute) {
// 	w, h := termbox.Size()

// }

func Print(s string) {
	for _, c := range s {
		printc(c)
	}
	termbox.Flush()
}

func Printf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	Print(s)
}
