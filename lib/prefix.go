package lib

import (
	"io"
	"strings"
)

func PrefixLines(w io.Writer, prefix string) *replaceWriter {
	w.Write([]byte(prefix))
	return &replaceWriter{
		replacer: strings.NewReplacer("\n", "\n"+prefix),
		next:     w,
	}
}

type replaceWriter struct {
	replacer *strings.Replacer
	next     io.Writer
}

func (r *replaceWriter) Write(p []byte) (int, error) {
	replaced := r.replacer.Replace(string(p))
	return r.next.Write([]byte(replaced))
}
