package main

import (
	"errors"
	"fmt"
	"github.com/tetsuzawa/dxb"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "usage: DDB_to_DDA in.DDB out.DDA")
		os.Exit(1)
	}
}

func run() error {
	args := os.Args
	if len(args) != 3 {
		return errors.New("argument is insufficient")
	}
	in := args[1]
	out := args[2]

	f, err := os.Open(in)
	if err != nil {
		return err
	}
	samples, err := dxb.ReadDDB(f)
	err = f.Close()
	if err != nil {
		return err
	}
	f, err = os.Create(out)
	if err != nil {
		return err
	}
	err = dxb.WriteDDA(f, samples)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

