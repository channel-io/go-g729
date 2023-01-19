package main

import (
	"fmt"
	"io"
	"os"

	"github.com/channel-io/go-g729/pkg"
)

func main() {
	// encoding example
	inputWAV, err := os.Open("./assets/input.wav")
	if err != nil {
		panic(err)
	}

	wavHeader := make([]byte, 44)
	if _, err = inputWAV.Read(wavHeader); err != nil {
		panic(err)
	}

	outG729, _ := os.Create("./enc.g729")
	if err = encode(inputWAV, outG729); err != nil {
		panic(fmt.Sprintf("failed to encode: %v", err))
	}

	if err = inputWAV.Close(); err != nil {
		panic(err)
	}
	if err = outG729.Close(); err != nil {
		panic(err)
	}

	// decoding example
	inputG729, err := os.Open("./enc.g729")
	if err != nil {
		panic(err)
	}

	outWAV, err := os.Create("./dec.wav")
	if err != nil {
		panic(err)
	}

	if _, err = outWAV.Write(wavHeader); err != nil {
		panic(err)
	}

	if err = decode(inputG729, outWAV); err != nil {
		panic(fmt.Sprintf("failed to decode: %v", err))
	}
}

func encode(r io.Reader, w io.Writer) error {
	enc := pkg.NewEncoder(false)

	for {
		buf := make([]byte, 160)
		if n, err := r.Read(buf); err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if n != 160 {
			// ignore last frame if frame size is invalid
			break
		}

		if encoded, err := enc.Encode(buf); err != nil {
			return err
		} else if _, err = w.Write(encoded); err != nil {
			return err
		}
	}

	return nil
}

func decode(r io.Reader, w io.Writer) error {
	dec := pkg.NewDecoder()

	for {
		buf := make([]byte, 10)
		if n, err := r.Read(buf); err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if n != 10 {
			// ignore last frame if frame size is invalid
			break
		}

		if decoded, err := dec.Decode(buf); err != nil {
			return err
		} else if _, err = w.Write(decoded); err != nil {
			return err
		}
	}

	return nil
}
