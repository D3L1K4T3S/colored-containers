package main

import (
	"bufio"
	"colored_containers/src"
	"log/slog"
	"os"
)

func main() {

	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			slog.Error("error clear buffer: ", err)
		}
	}(out)

	err := src.App(in, out)
	if err != nil {
		slog.Error("error with solve: ", err.Error())
	}
}
