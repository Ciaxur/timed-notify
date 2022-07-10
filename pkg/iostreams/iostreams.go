package iostreams

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
)

type IOStreams struct {
	In  io.ReadCloser
	Out io.Writer
	Err io.Writer
}

func System() *IOStreams {
	io := &IOStreams{
		In:  os.Stdin,
		Out: colorable.NewColorable(os.Stdout),
		Err: colorable.NewColorable(os.Stderr),
	}

	return io
}
