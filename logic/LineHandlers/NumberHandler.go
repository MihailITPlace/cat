package LineHandlers

import (
	"fmt"
	"strconv"
)

type NumberHandler struct {
	numberAll  bool
	lineNumber int
}

func (self *NumberHandler) Handle(line string, printLine bool) (string, bool) {
	if !printLine {
		return line, printLine
	}

	if self.numberAll || line != "" {
		line = fmt.Sprintf("%6v  %s", strconv.Itoa(self.lineNumber), line)
		self.lineNumber++
	}

	return line, printLine
}

func NewNumberHandler(numberAll bool) *NumberHandler {
	h := new(NumberHandler)
	h.lineNumber = 1
	h.numberAll = numberAll
	return h
}
