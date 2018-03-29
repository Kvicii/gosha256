package sha2

import (
	"encoding/binary"
	"encoding/hex"
	"io"
	"log"
	"strconv"
)

var (
	LogTrace *log.Logger
	LogInfo  *log.Logger
	LogError *log.Logger
)

func InitLog(traceHandle io.Writer, infoHandle io.Writer, errorHandle io.Writer) {
	LogTrace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	LogInfo = log.New(infoHandle, "INFO: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	LogError = log.New(errorHandle, "ERROR: ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

// https://en.wikipedia.org/wiki/SHA-2#Pseudocode

/* you don't pad a single block, you pad the end of the entire message.  The padding may append an additional block.
func pad(n int16, buf []byte) {
	// pad to blocksizebits
	// n is the number of valid bits at the beginning
	// add a "1" bit
	// then zeroes
	// then n as a binary int in the last wordsizebits
	if n < blocksizebits {
		panic("pad function called with too many bits (%d vs %d)", n, blocksizebits)
	}
	if n < blocksizebytes {
		for i := n; i < blocksizebytes; i++ {
			buf[i] = 0
			fmt.Printf("Padding buf[%d]", i)
		}
	}
}
*/

// TODO stream 512-byte chunks (via channel?) and retain state between chunks
// any chunk len >= 440? bits will need an extra block
/*
   while not EOF
     read 512 bits or to EOF
     if EOF and len<440
       append stop bit, buffer, and message length to 512
     if EOF and len>=440 and len<512
       append stop bit, buffer to 512
       set flag to create next block with message length only (no stop bit)
     process block
     if flag
       create another 512 block with message length only (no stop bit)
       process final block

func Sha256Streamed(r Reader??, chan channel??) [32]byte {
}

func Sha256Bitwise(m []byte, lenBits uint64) [32]byte {
}
*/

//
func Sha256(m []byte) [32]byte {
	result := [32]byte{}

	LogTrace.Printf("Sha256: ========= input (%d bytes): 0x%s %s\n", len(m), hex.EncodeToString(m), strconv.QuoteToASCII(string(m)))
	LogInfo.Printf("Sha256: Input Message: %s", strconv.QuoteToASCII(string(m)))

	// assumes message lengths are a multiple of 8 bits (byte-aligned)
	mL := uint64(len(m) * 8) // length in bits
	//fmt.Printf("Sha256: message length => %d bits\n", mL)

	// calculate length of required padding
	mPadL := uint64(512 - ((mL + 8 + 64) % 512))
	LogTrace.Printf("Sha256: need to pad %d bits (%d bytes)\n", mPadL, mPadL/8)

	// create a buffer for padding and length additions
	// note: don't affect the source message and don't make a copy of it
	mBuf := make([]byte, (8+64+mPadL)/8)
	LogTrace.Printf("Sha256: mBuf size %d bits (%d bytes)\n", len(mBuf)*8, len(mBuf))

	// append "1" bit separator followed by "0" padding
	mBuf[0] = 0x80
	// append length
	binary.BigEndian.PutUint64(mBuf[len(mBuf)-8:], mL)

	LogTrace.Printf("Sha256: buffer status: 0x%s|%s\n", hex.EncodeToString(m), hex.EncodeToString(mBuf))

	// calculate the total hashed length
	hL := mL + uint64(len(mBuf)*8)
	LogTrace.Printf("Sha256: total size of hashed data including padding and length => %d bits (%d bytes)\n", hL, hL/8)

	// sanity check
	if hL%512 != 0 {
		panic("Sha256: total hashed length is not a multiple of 512")
	}

	// message schedule array
	w := [64]uint32{}

	h0 := sha256h00
	h1 := sha256h01
	h2 := sha256h02
	h3 := sha256h03
	h4 := sha256h04
	h5 := sha256h05
	h6 := sha256h06
	h7 := sha256h07

	// loop by 512-bit (64 bytes) chunks
	for i := uint64(0); i < (hL / 8); i += 64 {
		LogTrace.Printf("Sha256: new chunk, i==%d\n", i)
		var chunk []byte
		// see if this is the last part of original message slice or padding
		if i+64 >= (mL / 8) {
			// copy just the last part of the original message slice and append the padding slice
			chunk = make([]byte, 64)
			if i >= uint64(len(m)) {
				// final chunk
				copy(chunk[:], mBuf[len(mBuf)-64:]) // copy the tail of the buf to chunk
			} else {
				tail := len(m[i:])
				LogTrace.Printf("Sha256: chunk copying last original message slice from %d to %d\n", i, tail)
				copy(chunk[:], m[i:]) // copy the tail of message to the front of chunk

				final := 64 - tail
				LogTrace.Printf("Sha256: chunk copying buf len %d to chunk at %d\n", final, tail)
				copy(chunk[tail:], mBuf[:final]) // copy the buf to the tail of chunk

			}
		} else {
			// use the original message slice
			LogTrace.Printf("Sha256: chunk using original message slice from %d to %d\n", i, i+64)
			chunk = m[i : i+64]
		}

		LogTrace.Printf("Sha256: chunk status: 0x%s\n", hex.EncodeToString(chunk))

		// scattered, smothered, covered, padded, chunked

		// copy chunk into first 16 words of w
		for j := 0; j < 16; j++ {
			w[j] = uint32(chunk[j*4])<<24 | uint32(chunk[(j*4)+1])<<16 | uint32(chunk[(j*4)+2])<<8 | uint32(chunk[(j*4)+3])
			LogTrace.Printf("Sha256: copy chunk to w[%2.2d] 0x%8.8X\n", j, w[j])
		}

		{
			LogInfo.Printf("Block Contents:")
			for i := 0; i < 16; i++ {
				LogInfo.Printf("  W[%d] = %8.8X", i, w[i])
			}
		}

		// Extend the first 16 words into the remaining 48 words w[16..63] of the message schedule array:
		for t := 16; t < 64; t++ {
			w[t] = lowerSigma1(w[t-2]) + w[t-7] + lowerSigma0(w[t-15]) + w[t-16]
			LogTrace.Printf("Sha256: extend w[%2.2d] 0x%8.8X\n", t, w[t])
		}

		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		f := h5
		g := h6
		h := h7

		LogInfo.Printf("          A        B        C        D        E        F        G        H    \n")
		for t := 0; t < 64; t++ {
			uT1 := h + upperSigma1(e) + ch(e, f, g) + sha256kByIndex(t) + w[t]
			uT2 := upperSigma0(a) + maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + uT1
			d = c
			c = b
			b = a
			a = uT1 + uT2
			LogInfo.Printf("t=%2d: %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X\n", t, a, b, c, d, e, f, g, h)
		}

		LogInfo.Printf("H[0] = %8.8X + %8.8X = %8.8X\n", h0, a, a+h0)
		h0 = a + h0
		LogInfo.Printf("H[1] = %8.8X + %8.8X = %8.8X\n", h1, b, b+h1)
		h1 = b + h1
		LogInfo.Printf("H[2] = %8.8X + %8.8X = %8.8X\n", h2, c, c+h2)
		h2 = c + h2
		LogInfo.Printf("H[3] = %8.8X + %8.8X = %8.8X\n", h3, d, d+h3)
		h3 = d + h3
		LogInfo.Printf("H[4] = %8.8X + %8.8X = %8.8X\n", h4, e, e+h4)
		h4 = e + h4
		LogInfo.Printf("H[5] = %8.8X + %8.8X = %8.8X\n", h5, f, f+h5)
		h5 = f + h5
		LogInfo.Printf("H[6] = %8.8X + %8.8X = %8.8X\n", h6, g, g+h6)
		h6 = g + h6
		LogInfo.Printf("H[7] = %8.8X + %8.8X = %8.8X\n", h7, h, h+h7)
		h7 = h + h7

		LogTrace.Printf("Sha256: hash state: 0x%8.8X %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X\n", a, b, c, d, e, f, g, h)
	}

	LogInfo.Printf("Message Digest is  %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X %8.8X\n", h0, h1, h2, h3, h4, h5, h6, h7)

	// copy the final hash output
	binary.BigEndian.PutUint32(result[0:4], h0)
	binary.BigEndian.PutUint32(result[4:8], h1)
	binary.BigEndian.PutUint32(result[8:12], h2)
	binary.BigEndian.PutUint32(result[12:16], h3)
	binary.BigEndian.PutUint32(result[16:20], h4)
	binary.BigEndian.PutUint32(result[20:24], h5)
	binary.BigEndian.PutUint32(result[24:28], h6)
	binary.BigEndian.PutUint32(result[28:32], h7)
	return result
}
