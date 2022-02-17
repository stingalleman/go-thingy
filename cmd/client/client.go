package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

var (
	posx int
	posy int
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	textStyle := defStyle
	// boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	width, height := s.Size()

	// draw borders
	for x := 0; x < width; x++ {
		s.SetContent(x, height-1, '+', nil, textStyle)
		s.SetContent(x, 0, '+', nil, textStyle)
	}
	for y := 0; y < height; y++ {
		s.SetContent(width-1, y, '+', nil, textStyle)
		s.SetContent(0, y, '+', nil, textStyle)
	}

	// Event loop
	// ox, oy := -1, -1
	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	char := 'S'
	posx = (width - 2) / 2
	posy = (height - 2) / 2

	for {
		// Update screen
		s.Show()

		s.SetContent(posx, posy, char, nil, textStyle)

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			key := ev.Key()

			{
				s.SetContent(posx, posy, ' ', nil, textStyle)

				if key == tcell.KeyUp {
					posy--
				} else if key == tcell.KeyDown {
					posy++
				} else if key == tcell.KeyLeft {
					posx--
				} else if key == tcell.KeyRight {
					posx++
				}

				s.SetContent(posx, posy, char, nil, textStyle)
			}

			if key == tcell.KeyEscape || key == tcell.KeyCtrlC {
				quit()
			} else if key == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		}
	}
}

func handleMovement(ev *tcell.EventKey) {
	key := ev.Key()

	if key == tcell.KeyUp {
		posy--
	} else if key == tcell.KeyDown {
		posy++
	} else if key == tcell.KeyLeft {
		posx--
	} else if key == tcell.KeyRight {
		posx++
	}
}

// Poll

// case *tcell.EventMouse:
// 	x, y := ev.Position()
// 	button := ev.Buttons()
// 	// Only process button events, not wheel events
// 	button &= tcell.ButtonMask(0xff)

// 	if button != tcell.ButtonNone && ox < 0 {
// 		ox, oy = x, y
// 	}
// 	switch ev.Buttons() {
// 	case tcell.ButtonNone:
// 		if ox >= 0 {
// 			label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
// 			drawBox(s, ox, oy, x, y, boxStyle, label)
// 			ox, oy = -1, -1
// 		}
// 	}
// }
