package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
)

var format = flag.String("format", "", "format to output to")

func main() {
	flag.Parse()
	format := flag.Arg(0)
	var myFunc func(io.Writer, image.Image) error
	switch format {
	case "jpeg":
		myFunc = func(out io.Writer, img image.Image) error { return jpeg.Encode(out, img, &jpeg.Options{Quality: 95}) }
	case "png":
		myFunc = png.Encode
	case "gif":
		myFunc = func(out io.Writer, img image.Image) error { return gif.Encode(out, img, nil) }
	default:
		fmt.Errorf("Invalid format %s", format)
		os.Exit(1)
	}

	if err := convert(os.Stdin, os.Stdout, myFunc); err != nil {
		fmt.Fprintf(os.Stderr, "convert: %format\n", err)
		os.Exit(2)
	}
}

func convert(in io.Reader, out io.Writer, f func(io.Writer, image.Image) error) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "input format =", kind)
	return f(out, img)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
