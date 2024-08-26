package morse

import (
	"strings"

	"github.com/paramsiddharth/esp32-morse/blink"
)

type MorseCodeType uint8

const DURATION = 1000

const (
	MORSE_CODE_DOT MorseCodeType = iota
	MORSE_CODE_DASH
	MORSE_CODE_WAIT
	MORSE_CODE_SEP
)

var MorseCode = struct {
	Dot  MorseCodeType
	Dash MorseCodeType
	Wait MorseCodeType
	Sep  MorseCodeType
}{
	Dot:  MORSE_CODE_DOT,
	Dash: MORSE_CODE_DASH,
	Wait: MORSE_CODE_WAIT,
	Sep:  MORSE_CODE_SEP,
}

var charMorseCodeMap = map[rune][]MorseCodeType{
	'A':  {MorseCode.Dot, MorseCode.Dash},
	'N':  {MorseCode.Dash, MorseCode.Dot},
	'B':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'O':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dash},
	'C':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'P':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot},
	'D':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot},
	'Q':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash},
	'E':  {MorseCode.Dot},
	'R':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'F':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'S':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'G':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dot},
	'T':  {MorseCode.Dash},
	'H':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'U':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'I':  {MorseCode.Dot, MorseCode.Dot},
	'V':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'J':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash},
	'W':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash},
	'K':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dash},
	'X':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'L':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot},
	'Y':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash},
	'M':  {MorseCode.Dash, MorseCode.Dash},
	'Z':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot},
	'1':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash},
	'6':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'2':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash},
	'7':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'3':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash},
	'8':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot},
	'4':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'9':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot},
	'5':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'0':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash},
	' ':  {},
	'?':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot},
	';':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	':':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot},
	'/':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'-':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'\'': {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot},
	'"':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'(':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot},
	')':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash},
	'=':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'+':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'*':  {MorseCode.Dash, MorseCode.Dot, MorseCode.Dot, MorseCode.Dash},
	'@':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot},
	'Á':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash},
	'Ä':  {MorseCode.Dot, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash},
	'É':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dot, MorseCode.Dot},
	'Ñ':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash},
	'Ö':  {MorseCode.Dash, MorseCode.Dash, MorseCode.Dash, MorseCode.Dot},
	'Ü':  {MorseCode.Dot, MorseCode.Dot, MorseCode.Dash, MorseCode.Dash},
	'.':  {MorseCode.Wait},
}

func RenderMorseCodeChar(ch MorseCodeType) *rune {
	dash := '_'
	dot := '.'
	wait := ' '
	switch ch {
	case MorseCode.Dash:
		return &dash
	case MorseCode.Dot:
		return &dot
	case MorseCode.Wait, MorseCode.Sep:
		return &wait
	default:
		return nil
	}
}

func RenderMorseCode(ch_s *[]MorseCodeType) *string {
	chs := *ch_s
	str := ""

	for i := range len(chs) {
		ch := chs[i]
		chr := RenderMorseCodeChar(ch)
		str += string(*chr)
	}

	return &str
}

func StrToMorseCode(str string) *[]MorseCodeType {
	strCaps := strings.ToUpper(str)
	code := []MorseCodeType{}

	for i := range len(strCaps) {
		ch := rune(strCaps[i])
		if val, ok := charMorseCodeMap[ch]; ok {
			code = append(code, val...)
			if len(val) > 0 {
				code = append(code, MorseCode.Sep)
			}
		}
	}

	return &code
}

func BlinkMorseCodeChar(c MorseCodeType) {
	switch c {
	case MorseCode.Dash:
		blink.Blink(DURATION, 0.8, 0)
		break
	case MorseCode.Dot:
		blink.Blink(DURATION, 0.2, 0)
		break
	case MorseCode.Wait:
		blink.Blink(DURATION, 0, 0)
		break
	default:
	}
}

func BlinkMorseCode(c_s *[]MorseCodeType) {
	cs := *c_s

	for i := range len(cs) {
		BlinkMorseCodeChar(cs[i])
	}
}
