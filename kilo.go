package main

import "fmt"

const KILO_VERSION string = "0.0.1"

/* Syntax highlight types */
const (
	HlNormal = iota
	HlNonprint
	HlComment   /* Single line comment. */
	HlMlcomment /* Multi-line comment. */
	HlKeyword1
	HlKeyword2
	HlString
	HlNumber
	HlMatch /* Search match. */
)

const (
	HlHighlightStrings = iota + 1
	HlHighlightNumbers
)

type EditorSyntax struct {
	fileMatch              *string
	keywords               *string
	singlelineCommentStart [2]rune
	multilineCommentStart  [3]rune
	multilineCommentEnd    [3]rune
	flags                  int8
}

// EditorRow This structure represents a single line of the file we are editing.
type EditorRow struct {
	idx          uint   /* Row index in the file, zero-based. */
	size         uint   /* Size of the row, excluding the null term. */
	renderedSize uint   /* Size of the rendered row. */
	chars        []rune /* Row content. */
	render       []rune /* Row content "rendered" for screen (for TABs). */
	hl           []rune /* Syntax highlight type for each character in render.*/
	hl_oc        int    /* Row had open comment at end in last syntax highlight
	   check. */
}

type HLColor struct {
	r, g, b uint8
}

type editorConfig struct {
	cx, cy         int        /* Cursor x and y position in characters */
	rowoff         int        /* Offset of row displayed. */
	coloff         int        /* Offset of column displayed. */
	screenrows     int        /* Number of rows that we can show */
	screencols     int        /* Number of cols that we can show */
	numrows        int        /* Number of rows */
	rawmode        int        /* Is terminal raw mode enabled? */
	row            int        /* Rows */
	dirty          *EditorRow /* File modified but not saved. */
	filename       int        /* Currently open filename */
	statusmsg      [80]string
	statusmsg_time string
	syntax         *EditorSyntax /* Current syntax highlight, or NULL. */
}

type KEY_ACTION = uint16

const (
	KeyNull = 0 /* NULL */
	CtrlC   = 3 /* Ctrl-c */
	CtrlD   = 4 /* Ctrl-d */

	CtrlF     = 6   /* Ctrl-f */
	CtrlH     = 8   /* Ctrl-h */
	Tab       = 9   /* Tab */
	CtrlL     = 12  /* Ctrl+l */
	Enter     = 13  /* Enter */
	CtrlQ     = 17  /* Ctrl-q */
	CtrlS     = 19  /* Ctrl-s */
	CtrlU     = 21  /* Ctrl-u */
	Esc       = 27  /* Escape */
	Backspace = 127 /* Backspace */
	/* The following are just soft codes not really reported by the
	 * terminal directly. */

	ArrowLeft = 1000
	ArrowRight
	ArrowUp
	ArrowDown
	DelKey
	HomeKey
	EndKey
	PageUp
	PageDown
)

type Kilo interface {
	editorSetStatusMessage(fmt string)
}

type HighlightDatabase struct {
	extensions []string
	keywords   []string
	comments   []string
	flags      int
}

var HLDB = HighlightDatabase{
	CLanguage{}.GetHLKeywords(),
	CLanguage{}.GetHLExtensions(),
	[]string{"//", "/*", "*/"},
	HlHighlightNumbers | HlHighlightNumbers,
}

func main() {
	fmt.Printf("%s", CLanguage{}.GetHLExtensions())
}
