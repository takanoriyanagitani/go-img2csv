package writer

import (
	"context"
	"encoding/csv"
	"io"

	ic "github.com/takanoriyanagitani/go-img2csv"
	. "github.com/takanoriyanagitani/go-img2csv/util"
)

type ImageWriter func(ic.Image) IO[Void]

type ImageWriterConfig struct {
	io.Writer
	ic.GrayToString8
}

func (c ImageWriterConfig) ToWriterGray8() ImageWriter {
	return func(img ic.Image) IO[Void] {
		return func(ctx context.Context) (Void, error) {
			var wtr *csv.Writer = csv.NewWriter(c.Writer)
			defer wtr.Flush()

			var height int = img.Height()
			var width int = img.Width()

			var buf []uint8 = make([]uint8, 0, width)
			var cbf []string = make([]string, 0, width)

			for y := range height {
				var row []uint8 = img.ToRowGray8(buf, y)
				var encoded []string = c.GrayToString8.RowToStringsGray8(row, cbf)
				e := wtr.Write(encoded)
				if nil != e {
					return Empty, e
				}
			}

			return Empty, nil
		}
	}
}
