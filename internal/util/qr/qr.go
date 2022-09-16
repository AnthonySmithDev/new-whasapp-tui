package qr

import (
	"bytes"
	"strings"

	"rsc.io/qr"
)

const QUIET_ZONE = 2

const BLACK_WHITE = "▄"
const BLACK_BLACK = " "
const WHITE_BLACK = "▀"
const WHITE_WHITE = "█"

func Generate(text string) string {
	code, _ := qr.Encode(text, qr.L)
	return generateHalfBlock(code)
}

func generateHalfBlock(code *qr.Code) string {
	// Frame the barcode in a 2 pixel border
	qz := QUIET_ZONE

	ww := WHITE_WHITE
	bb := BLACK_BLACK
	wb := WHITE_BLACK
	bw := BLACK_WHITE

	var b bytes.Buffer

	// top border
	if qz%2 != 0 {
		b.WriteString((stringRepeat(bw, code.Size+qz*2) + "\n"))
		b.WriteString((stringRepeat(stringRepeat(ww,
			code.Size+qz*2)+"\n", qz/2)))
	} else {
		b.WriteString((stringRepeat(stringRepeat(ww,
			code.Size+qz*2)+"\n", qz/2)))
	}
	for i := 0; i <= code.Size; i += 2 {
		b.WriteString((stringRepeat(ww, qz))) // left border
		for j := 0; j <= code.Size; j++ {
			next_black := false
			if i+1 < code.Size {
				next_black = code.Black(j, i+1)
			}
			curr_black := code.Black(j, i)
			if curr_black && next_black {
				b.WriteString((bb))
			} else if curr_black && !next_black {
				b.WriteString((bw))
			} else if !curr_black && !next_black {
				b.WriteString((ww))
			} else {
				b.WriteString((wb))
			}
		}
		b.WriteString((stringRepeat(ww, qz-1) + "\n")) // right border
	}
	// bottom border
	if qz%2 == 0 {
		b.WriteString((stringRepeat(stringRepeat(ww,
			code.Size+qz*2)+"\n", qz/2-1)))
		b.WriteString((stringRepeat(wb, code.Size+qz*2) + "\n"))
	} else {
		b.WriteString((stringRepeat(stringRepeat(ww,
			code.Size+qz*2)+"\n", qz/2)))
	}
	return b.String()
}

func stringRepeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return strings.Repeat(s, count)
}
