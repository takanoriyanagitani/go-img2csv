package main

import (
	"context"
	"image"
	"io"
	"log"
	"os"

	ic "github.com/takanoriyanagitani/go-img2csv"
	ws "github.com/takanoriyanagitani/go-img2csv/csv/writer/std"
	. "github.com/takanoriyanagitani/go-img2csv/util"
)

var output io.Writer = os.Stdout

var cfg ws.ImageWriterConfig = ws.ImageWriterConfig{
	Writer:        output,
	GrayToString8: ic.GrayToString8Default,
}

var writer ws.ImageWriter = cfg.ToWriterGray8()

var imgRdr io.Reader = os.Stdin

var decoded IO[image.Image] = Bind(
	Of(imgRdr),
	Lift(ic.ReaderToImage),
)

var img IO[ic.Image] = Bind(
	decoded,
	Lift(func(i image.Image) (ic.Image, error) {
		return ic.Image{Image: i}, nil
	}),
)

var img2wtr IO[Void] = Bind(img, writer)

func main() {
	_, e := img2wtr(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}
