package lispinterpreterlogic

const (
	DEFUN = iota

	COND
	LET

	ADD
	SUB
	MUL
	DIV

	SETQ
	SET
	QUOTE
	EVAL
	AND
	OR
	NOT
)

var tokens = [...]string{
	DEFUN: "DEFUN",
	COND:  "COND",
	LET:   "LET",
	ADD:   "ADD",
	SUB:   "SUB",
	MUL:   "MUL",
	DIV:   "DIV",

	SETQ:  "SETQ",
	SET:   "SET",
	QUOTE: "QUOTE",
	EVAL:  "EVAL",
	AND:   "AND",
	OR:    "OR",
	NOT:   "NOT",
}
