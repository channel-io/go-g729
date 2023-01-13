package pkg

/*
#cgo LDFLAGS: -lbcg729
#include <bcg729/encoder.h>
*/
import "C"
import (
	"encoding/binary"
	"fmt"
)

type Encoder struct {
	ctx *C.bcg729EncoderChannelContextStruct
}

func NewEncoder(vadEnabled bool) *Encoder {
	enableVAD := int8(0)
	if vadEnabled {
		enableVAD = 1
	}

	return &Encoder{
		ctx: C.initBcg729EncoderChannel(C.uint8_t(enableVAD)),
	}
}

func (e *Encoder) Close() {
	C.closeBcg729EncoderChannel(e.ctx)
}

// Encode encodes 160 bytes of PCM data into 10 bytes of G729 data.
// The input (PCM data) must be 160 bytes, otherwise it will panic.
// The output (G729 data) will be 10 bytes.
func (e *Encoder) Encode(input []byte) ([]uint8, error) {
	if len(input) != 160 {
		return nil, fmt.Errorf("input must be 160 bytes")
	}

	inputFrame := make([]int16, 80)
	for i := 0; i < 80; i++ {
		inputFrame[i] = int16(binary.LittleEndian.Uint16(input[i*2 : i*2+2]))
	}

	bitStream := make([]uint8, 10)
	bitStreamLength := uint8(0)
	C.bcg729Encoder(e.ctx, (*C.int16_t)(&inputFrame[0]), (*C.uint8_t)(&bitStream[0]), (*C.uint8_t)(&bitStreamLength))
	if bitStreamLength != 10 {
		return nil, fmt.Errorf("failed to encode: encoded bisStream length is invalid (len=%d)", bitStreamLength)
	}

	return bitStream, nil
}
