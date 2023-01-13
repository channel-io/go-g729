package pkg

/*
#cgo LDFLAGS: -lbcg729
#include <bcg729/decoder.h>
*/
import "C"
import (
	"encoding/binary"
	"fmt"
)

type Decoder struct {
	ctx *C.bcg729DecoderChannelContextStruct
}

func NewDecoder() *Decoder {
	return &Decoder{
		ctx: C.initBcg729DecoderChannel(),
	}
}

// Decode decodes 10 bytes of G729 data into 160 bytes of PCM data.
// The input (G729 data) must be 10 bytes, otherwise it will panic.
// The output (PCM data) will be 160 bytes.
func (d *Decoder) Decode(input []byte) ([]byte, error) {
	if len(input) != 10 {
		return nil, fmt.Errorf("input must be 10 bytes")
	}
	outputFrame := make([]int16, 80)

	C.bcg729Decoder(
		d.ctx,
		(*C.uint8_t)(&input[0]),
		C.uint8_t(10),
		C.uint8_t(0),
		C.uint8_t(0),
		C.uint8_t(0),
		(*C.int16_t)(&outputFrame[0]),
	)

	out := make([]byte, 0, 160)
	for i := 0; i < 80; i++ {
		out = binary.LittleEndian.AppendUint16(out, uint16(outputFrame[i]))
	}
	return out, nil
}
