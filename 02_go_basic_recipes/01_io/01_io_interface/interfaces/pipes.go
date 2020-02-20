package interfaces

import (
	"io"
	"os"
)

func PipeExample() error {
	// pipe reader and writer implement io.Reader
	// and io.Writer interface
	r, w := io.Pipe()

	// Run in go routine as it blocks waiting for close
	// from reader at cleanup
	go func() {
		// Use this for marshaling/unmarshaling, etc.
		w.Write([]byte("test\n"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
