package kilogo

const KILO_VERSION string = "0.0.1"

/* Syntax highlight types */
const (
	HL_NORMAL = iota
	HL_NONPRINT
	HL_COMMENT   /* Single line comment. */
	HL_MLCOMMENT /* Multi-line comment. */
	HL_KEYWORD1
	HL_KEYWORD2
	HL_STRING
	HL_NUMBER
	HL_MATCH /* Search match. */
)

const (
	HL_HIGHLIGHT_STRINGS = iota + 1
	HL_HIGHLIGHT_NUMBERS
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
	syntax  *EditorSyntax /* Current syntax highlight, or NULL. */
}

type KEY_ACTION = uint16

const (
	KEY_NULL  = 0   /* NULL */
	CTRL_C    = 3   /* Ctrl-c */
	CTRL_D    = 4   /* Ctrl-d */

	CTRL_F    = 6   /* Ctrl-f */
	CTRL_H    = 8   /* Ctrl-h */
	TAB       = 9   /* Tab */
	CTRL_L    = 12  /* Ctrl+l */
	ENTER     = 13  /* Enter */
	CTRL_Q    = 17  /* Ctrl-q */
	CTRL_S    = 19  /* Ctrl-s */
	CTRL_U    = 21  /* Ctrl-u */
	ESC       = 27  /* Escape */
	BACKSPACE = 127 /* Backspace */
	/* The following are just soft codes not really reported by the
	 * terminal directly. */
	ARROW_LEFT = 1000
	ARROW_RIGHT
	ARROW_UP
	ARROW_DOWN
	DEL_KEY
	HOME_KEY
	END_KEY
	PAGE_UP
	PAGE_DOWN
)

type Kilo interface {
	editorSetStatusMessage(fmt string)
}