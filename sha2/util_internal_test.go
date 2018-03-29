package sha2

import (
	"fmt"
	"math"
	"testing"
)

/*
SHA-224 and SHA-256 use the same sequence of sixty-four constant 32-bit words.
These words represent the first thirty-two bits of the fractional parts of
the cube roots of the first sixty-four prime numbers.

SHA-384, SHA-512, SHA-512/224 and SHA-512/256 use the same sequence of eighty constant
64-bit words.  These words represent the first sixty-four bits of the
fractional parts of the cube roots of the first eighty prime numbers. In hex, these constant words
are (from left to right)
*/

func TestFirst64primes(t *testing.T) {
	a := first64primes()
	for i, p := range a {
		if p == 0 {
			t.Errorf("prime %d is zero", i)
		}
		////fmt.Printf("TestFirst64primes: Prime[%d]==%d\n", i, p)
	}
}

func first64primes() [64]int {
	// Return the first 64 primes, not counting One.  Do it the hard way.
	result := [64]int{}
	p := 0

	// bootstrap the first three primes
	result[p] = 2
	p++
	result[p] = 3
	p++

	// loop over all odd integers
	for i := 5; ; i += 2 {
		////fmt.Printf("first64primes: testing whether i = %d is prime\n", i)
		// loop over existing prime list
		// check for even division
		var j int
		for j = 1; j < len(result) && result[j] > 0; j++ {
			////fmt.Printf("first64primes: Is (i = %d) divisible by (j = %d), yes_if_zero (%d%%%d==%d)\n", i, j, i, result[j], i%result[j])
			if i%result[j] == 0 {
				// not a prime, stop checking
				////fmt.Printf("first64primes: Yes_divisible_by_existing_prime i = %d, j = %d, not a prime\n", i, j)
				break
			}
		}
		////fmt.Printf("first64primes: i = %d, j loop done, loop_result result[j] == %d\n", i, result[j])
		if result[j] == 0 {
			// record new prime
			////fmt.Printf("first64primes: p = %d, prime = %d\n", p, i)
			result[p] = i
			p++
		}
		if p >= 64 {
			////fmt.Printf("first64primes: p = %d, stopping\n", p)
			break
		}
	}

	return result
}

/*
func TestCubeRoot(t *testing.T) {
	a := first64primes()
	for _, p := range a {
		c := math.Cbrt(float64(p))
		fmt.Printf("Cube root of %d is %0.11f\n", p, c)
		f := int((c - math.Floor(c)) * math.Exp2(32))
		fmt.Printf("Fractional of Cube root of %d is %010d\n", p, f)
		fmt.Printf("Constant equality %b (f=%x vs k00=%x)\n", f == sha256k00, f, sha256k00)
		h := fmt.Sprintf("%8.8x", f)
		fmt.Printf("Hex of fractional of Cube root of %d is %s\n", p, h)
	}

}
*/

func first32BitsOfCubeRootsOfFirst64Primes() [64]uint32 {
	result := [64]uint32{}
	a := first64primes()
	for i, p := range a {
		cr := math.Cbrt(float64(p))
		result[i] = uint32((cr - math.Floor(cr)) * math.Exp2(32))
	}
	return result
}

func first32BitsOfSquareRootsOfFirst8Primes() [8]uint32 {
	result := [8]uint32{}
	a := first64primes()
	for i, p := range a[:8] {
		sr := math.Sqrt(float64(p))
		result[i] = uint32((sr - math.Floor(sr)) * math.Exp2(32))
	}
	return result
}

func TestConst(t *testing.T) {
	r := first32BitsOfCubeRootsOfFirst64Primes()
	// copy the consts into an array
	/*
		c := [64]uint32{}
		c[0] = sha256k00
		c[1] = sha256k01
		c[2] = sha256k02
		c[3] = sha256k03
		c[4] = sha256k04
		c[5] = sha256k05
		c[6] = sha256k06
		c[7] = sha256k07
		c[8] = sha256k08
		c[9] = sha256k09
		c[10] = sha256k10
		c[11] = sha256k11
		c[12] = sha256k12
		c[13] = sha256k13
		c[14] = sha256k14
		c[15] = sha256k15
		c[16] = sha256k16
		c[17] = sha256k17
		c[18] = sha256k18
		c[19] = sha256k19
		c[20] = sha256k20
		c[21] = sha256k21
		c[22] = sha256k22
		c[23] = sha256k23
		c[24] = sha256k24
		c[25] = sha256k25
		c[26] = sha256k26
		c[27] = sha256k27
		c[28] = sha256k28
		c[29] = sha256k29
		c[30] = sha256k30
		c[31] = sha256k31
		c[32] = sha256k32
		c[33] = sha256k33
		c[34] = sha256k34
		c[35] = sha256k35
		c[36] = sha256k36
		c[37] = sha256k37
		c[38] = sha256k38
		c[39] = sha256k39
		c[40] = sha256k40
		c[41] = sha256k41
		c[42] = sha256k42
		c[43] = sha256k43
		c[44] = sha256k44
		c[45] = sha256k45
		c[46] = sha256k46
		c[47] = sha256k47
		c[48] = sha256k48
		c[49] = sha256k49
		c[50] = sha256k50
		c[51] = sha256k51
		c[52] = sha256k52
		c[53] = sha256k53
		c[54] = sha256k54
		c[55] = sha256k55
		c[56] = sha256k56
		c[57] = sha256k57
		c[58] = sha256k58
		c[59] = sha256k59
		c[60] = sha256k60
		c[61] = sha256k61
		c[62] = sha256k62
		c[63] = sha256k63
	*/

	// compare the arrays
	for i, _ := range r {
		if r[i] != sha256kByIndex(i) {
			t.Errorf("Const sha256k%02d incorrect => code gave 0x%X, test wants 0x%X", i, sha256kByIndex(i), r[i])
		}
	}

	s := first32BitsOfSquareRootsOfFirst8Primes()
	// copy the consts into an array
	h := [8]uint32{}
	h[0] = sha256h00
	h[1] = sha256h01
	h[2] = sha256h02
	h[3] = sha256h03
	h[4] = sha256h04
	h[5] = sha256h05
	h[6] = sha256h06
	h[7] = sha256h07

	// compare the arrays
	for j, _ := range s {
		if s[j] != h[j] {
			t.Errorf("Const sha256h%02d incorrect => code gave 0x%X, test wants 0x%X", j, h[j], s[j])
		}
	}

}

// Ch(x, y, z)=(x and y) xor ( complement x and z)
// Ch(x, y, z)=(x & y) ^ ( ^x & z)
// "Choose" the bit from y or z based on the bit in x
// if the bit in x is 1, choose the bit in y
// if the bit in x is 0, choose the bit in z
func TestCh(t *testing.T) {
	v := []struct {
		x   uint32
		y   uint32
		z   uint32
		out uint32
	}{
		{0x00000000, 0x00000000, 0x00000000, 0x00000000},
		{0x11111111, 0x11111111, 0x11111111, 0x11111111},
		{0x11111111, 0x11111111, 0x00000000, 0x11111111},
		{0x00000000, 0x11111111, 0x00000000, 0x00000000},
		{0x00000001, 0x00000001, 0x00000000, 0x00000001},
		{0x00000000, 0x00000000, 0x00000001, 0x00000001},
		{0x00000000, 0x00000001, 0x00000000, 0x00000000},
		{0x00000001, 0x00000000, 0x00000001, 0x00000000},
		{0x01010101, 0x11111111, 0x00000000, 0x01010101},
		{0x01010101, 0x00000000, 0x11111111, 0x10101010},
	}
	for i, a := range v {
		o := ch(a.x, a.y, a.z)
		if o != a.out {
			t.Errorf("Ch failure #%d: ch(0x%X, 0x%X, 0x%X) => code gave 0x%X, test wants 0x%X", i, a.x, a.y, a.z, o, a.out)
		}
	}
}

// "Majority" between x,y,z of each bit
// the value of each bit that is most common in the 3 inputs
func TestMaj(t *testing.T) {
	v := []struct {
		x   uint32
		y   uint32
		z   uint32
		out uint32
	}{
		{0x00000000, 0x00000000, 0x00000000, 0x00000000},
		{0x11111111, 0x00000000, 0x00000000, 0x00000000},
		{0x00000000, 0x11111111, 0x00000000, 0x00000000},
		{0x00000000, 0x00000000, 0x11111111, 0x00000000},
		{0x00000000, 0x11111111, 0x11111111, 0x11111111},
		{0x11111111, 0x00000000, 0x11111111, 0x11111111},
		{0x11111111, 0x11111111, 0x00000000, 0x11111111},
		{0x11111111, 0x11111111, 0x11111111, 0x11111111},
		{0x00000001, 0x00000001, 0x00000000, 0x00000001},
		{0x00000000, 0x00000000, 0x00000001, 0x00000000},
		{0x00000000, 0x00000001, 0x00000000, 0x00000000},
		{0x00000001, 0x00000000, 0x00000001, 0x00000001},
		{0x01010101, 0x11111111, 0x00000000, 0x01010101},
		{0x01010101, 0x00000000, 0x11111111, 0x01010101},
		{0x00001111, 0x00110011, 0x01010101, 0x00010111},
	}
	for i, a := range v {
		o := maj(a.x, a.y, a.z)
		if o != a.out {
			t.Errorf("Maj failure #%d: maj(0x%X, 0x%X, 0x%X) => code gave 0x%X, test wants 0x%X", i, a.x, a.y, a.z, o, a.out)
		}
	}
}

func TestUpperSigma0(t *testing.T) {
	v := []struct {
		in  uint32
		out uint32
	}{
		{0x00000000, 0x00000000},
		{0xF73108CE, 0x3F98C067},
		{0x00000003, 0xC0180C00},
		// TODO more test cases
	}
	for i, a := range v {
		o := upperSigma0(a.in)
		if o != a.out {
			t.Errorf("UpperSigma0 failure #%d: upperSigma0(0x%X) => code gave 0x%X, test wants 0x%X", i, a.in, o, a.out)
		}
	}
}

func TestUpperSigma1(t *testing.T) {
	v := []struct {
		in  uint32
		out uint32
	}{
		{0x00000000, 0x00000000},
		{0x00000001, 0x04200080},
		// TODO more test cases
	}
	for i, a := range v {
		o := upperSigma1(a.in)
		if o != a.out {
			t.Errorf("UpperSigma1 failure #%d: upperSigma1(0x%X) => code gave 0x%X, test wants 0x%X", i, a.in, o, a.out)
		}
	}
}

func TestLowerSigma0(t *testing.T) {
	v := []struct {
		in  uint32
		out uint32
	}{
		{0x00000000, 0x00000000},
		{0x00000001, 0x02004000},
		// TODO more test cases
	}
	for i, a := range v {
		o := lowerSigma0(a.in)
		if o != a.out {
			t.Errorf("LowerSigma0 failure #%d: lowerSigma0(0x%X) => code gave 0x%X, test wants 0x%X", i, a.in, o, a.out)
		}
	}
}

func TestLowerSigma1(t *testing.T) {
	v := []struct {
		in  uint32
		out uint32
	}{
		{0x00000000, 0x00000000},
		{0x00000001, 0x0000A000},
		{0x000003FF, 0x01806000},
		{0x00000078, 0x00330000},
		{0x00060000, 0xC0000183},
		{0x00018000, 0xF0000060},
		{0x2CCFFFED, 0x00000001},
		{0x9F0E663F, 0xFFFFFFFF},
	}
	for i, a := range v {
		o := lowerSigma1(a.in)
		if o != a.out {
			t.Errorf("LowerSigma1 failure #%d: lowerSigma1(0x%X) => code gave 0x%X, test wants 0x%X", i, a.in, o, a.out)
		}
	}
}

func xTestIterateLowerSigma1(t *testing.T) {
	var lowi, lowo, highi, higho uint32
	lowo = (^uint32(0))
	var i uint32
	for i = 1; i < (^uint32(0)); i++ {
		o := lowerSigma1(i)
		if i%10000 == 0 {
			fmt.Printf("BenchmarkIterateLowerSigma1 0x%8.8X 0x%8.8X\n", i, o)
		}
		if o < lowo {
			lowi = i
			lowo = o
		}
		if o > higho {
			highi = i
			higho = o
		}
	}
	fmt.Printf("Low  0x%8.8X 0x%8.8X\n", lowi, lowo)
	fmt.Printf("High 0x%8.8X 0x%8.8X\n", highi, higho)
}

// rotate right
func TestRotr(t *testing.T) {
	v := []struct {
		in  uint32
		sh  uint8
		out uint32
	}{
		{2, 1, 1},
		{4, 1, 2},
		{^uint32(0), 1, ^uint32(0)},
		{0x00000000, 0, 0x00000000},
		{0xA5A5A5A5, 0, 0xA5A5A5A5},
		{0xA5A5A5A5, 1, 0xD2D2D2D2},
		{0xA5A5A5A5, 2, 0x69696969},
		{0xA5A5A5A5, 3, 0xB4B4B4B4},
		{0xA5A5A5A5, 4, 0x5A5A5A5A},
		{0x00000000, 0, 0x00000000},
		{0x00000001, 1, 0x80000000},
		{0x80000000, 1, 0x40000000},
		{0x80000000, 2, 0x20000000},
		{0x80000000, 3, 0x10000000},
		{0x80000000, 4, 0x08000000},
		{0x80000000, 5, 0x04000000},
		{0x80000000, 6, 0x02000000},
		{0x80000000, 7, 0x01000000},
		{0x80000000, 8, 0x00800000},
		{0x80000000, 9, 0x00400000},
		{0x80000000, 31, 0x00000001},
		{0x80000000, 32, 0x80000000},
		{0xF73108CE, 0, 0xF73108CE},
		{0xF73108CE, 1, 0x7B988467},
		{0xF73108CE, 2, 0xBDCC4233},
		{0xF73108CE, 3, 0xDEE62119},
		{0xF73108CE, 4, 0xEF73108C},
		{0xF73108CE, 5, 0x77B98846},
		{0xF73108CE, 6, 0x3BDCC423},
		{0xF73108CE, 7, 0x9DEE6211},
		{0xF73108CE, 8, 0xCEF73108},
		{0xF73108CE, 9, 0x677B9884},
		{0xF73108CE, 10, 0x33BDCC42},
		{0xF73108CE, 11, 0x19DEE621},
		{0xF73108CE, 12, 0x8CEF7310},
		{0xF73108CE, 13, 0x4677B988},
		{0xF73108CE, 14, 0x233BDCC4},
		{0xF73108CE, 15, 0x119DEE62},
		{0xF73108CE, 16, 0x08CEF731},
		{0xF73108CE, 17, 0x84677B98},
		{0xF73108CE, 18, 0x4233BDCC},
		{0xF73108CE, 19, 0x2119DEE6},
		{0xF73108CE, 20, 0x108CEF73},
		{0xF73108CE, 21, 0x884677B9},
		{0xF73108CE, 22, 0xC4233BDC},
		{0xF73108CE, 23, 0x62119DEE},
		{0xF73108CE, 24, 0x3108CEF7},
		{0xF73108CE, 25, 0x9884677B},
		{0xF73108CE, 26, 0xCC4233BD},
		{0xF73108CE, 27, 0xE62119DE},
		{0xF73108CE, 28, 0x73108CEF},
		{0xF73108CE, 29, 0xB9884677},
		{0xF73108CE, 30, 0xDCC4233B},
		{0xF73108CE, 31, 0xEE62119D},
		{0xF73108CE, 32, 0xF73108CE},
		{0xF73108CE, 33, 0x7B988467},
	}
	for i, a := range v {
		o := rotr(a.sh, a.in)
		if o != a.out {
			t.Errorf("Rotr failure #%d: rotr(%d, 0x%X) => code gave 0x%X, test wants 0x%X", i, a.sh, a.in, o, a.out)
		}
	}
}

// rotate left
func TestRotl(t *testing.T) {
	v := []struct {
		in  uint32
		sh  uint8
		out uint32
	}{
		{2, 1, 4},
		{4, 1, 8},
		{^uint32(0), 1, ^uint32(0)},
		{0x00000000, 0, 0x00000000},
		{0xA5A5A5A5, 0, 0xA5A5A5A5},
		{0xA5A5A5A5, 1, 0x4B4B4B4B},
		{0xA5A5A5A5, 2, 0x96969696},
		{0xA5A5A5A5, 3, 0x2D2D2D2D},
		{0xA5A5A5A5, 4, 0x5A5A5A5A},
		{0x00000000, 0, 0x00000000},
		{0x00000001, 1, 0x00000002},
		{0x80000000, 1, 0x00000001},
		{0x80000000, 2, 0x00000002},
		{0x80000000, 3, 0x00000004},
		{0x80000000, 4, 0x00000008},
		{0x80000000, 5, 0x00000010},
		{0x80000000, 6, 0x00000020},
		{0x80000000, 7, 0x00000040},
		{0x80000000, 8, 0x00000080},
		{0x80000000, 9, 0x00000100},
		{0x80000000, 31, 0x40000000},
		{0x80000000, 32, 0x80000000},
		{0xF73108CE, 0, 0xF73108CE},
		{0xF73108CE, 1, 0xEE62119D},
		{0xF73108CE, 2, 0xDCC4233B},
		{0xF73108CE, 3, 0xB9884677},
		{0xF73108CE, 4, 0x73108CEF},
		{0xF73108CE, 5, 0xE62119DE},
		{0xF73108CE, 6, 0xCC4233BD},
		{0xF73108CE, 7, 0x9884677B},
		{0xF73108CE, 8, 0x3108CEF7},
		{0xF73108CE, 9, 0x62119DEE},
		{0xF73108CE, 10, 0xC4233BDC},
		{0xF73108CE, 11, 0x884677B9},
		{0xF73108CE, 12, 0x108CEF73},
		{0xF73108CE, 13, 0x2119DEE6},
		{0xF73108CE, 14, 0x4233BDCC},
		{0xF73108CE, 15, 0x84677B98},
		{0xF73108CE, 16, 0x08CEF731},
		{0xF73108CE, 17, 0x119DEE62},
		{0xF73108CE, 18, 0x233BDCC4},
		{0xF73108CE, 19, 0x4677B988},
		{0xF73108CE, 20, 0x8CEF7310},
		{0xF73108CE, 21, 0x19DEE621},
		{0xF73108CE, 22, 0x33BDCC42},
		{0xF73108CE, 23, 0x677B9884},
		{0xF73108CE, 24, 0xCEF73108},
		{0xF73108CE, 25, 0x9DEE6211},
		{0xF73108CE, 26, 0x3BDCC423},
		{0xF73108CE, 27, 0x77B98846},
		{0xF73108CE, 28, 0xEF73108C},
		{0xF73108CE, 29, 0xDEE62119},
		{0xF73108CE, 30, 0xBDCC4233},
		{0xF73108CE, 31, 0x7B988467},
		{0xF73108CE, 32, 0xF73108CE},
		{0xF73108CE, 33, 0xEE62119D},
	}
	for i, a := range v {
		o := rotl(a.sh, a.in)
		if o != a.out {
			t.Errorf("Rotl failure #%d: rotl(%d, 0x%X) => code gave 0x%X, test wants 0x%X", i, a.sh, a.in, o, a.out)
		}
	}
}

// shift right
func TestShr(t *testing.T) {
	v := []struct {
		in  uint32
		sh  uint8
		out uint32
	}{
		{2, 1, 1},
		{4, 1, 2},
		{^uint32(0), 1, 2147483647},
		{0x00000000, 0, 0x00000000},
		{0xA5A5A5A5, 0, 0xA5A5A5A5},
		{0xA5A5A5A5, 1, 0x52D2D2D2},
		{0xA5A5A5A5, 2, 0x29696969},
		{0xA5A5A5A5, 3, 0x14B4B4B4},
		{0xA5A5A5A5, 4, 0x0A5A5A5A},
		{0x00000000, 0, 0x00000000},
		{0x00000001, 1, 0x00000000},
		{0x80000000, 1, 0x40000000},
		{0x80000000, 2, 0x20000000},
		{0x80000000, 3, 0x10000000},
		{0x80000000, 4, 0x08000000},
		{0x80000000, 5, 0x04000000},
		{0x80000000, 6, 0x02000000},
		{0x80000000, 7, 0x01000000},
		{0x80000000, 8, 0x00800000},
		{0x80000000, 9, 0x00400000},
		{0x80000000, 31, 0x00000001},
		{0x80000000, 32, 0x00000000},
		{0xF73108CE, 0, 0xF73108CE},
		{0xF73108CE, 1, 0x7B988467},
		{0xF73108CE, 2, 0x3DCC4233},
		{0xF73108CE, 3, 0x1EE62119},
		{0xF73108CE, 4, 0x0F73108C},
		{0xF73108CE, 5, 0x07B98846},
		{0xF73108CE, 6, 0x03DCC423},
		{0xF73108CE, 7, 0x01EE6211},
		{0xF73108CE, 8, 0x00F73108},
		{0xF73108CE, 9, 0x007B9884},
		{0xF73108CE, 10, 0x003DCC42},
		{0xF73108CE, 11, 0x001EE621},
		{0xF73108CE, 12, 0x000F7310},
		{0xF73108CE, 13, 0x0007B988},
		{0xF73108CE, 14, 0x0003DCC4},
		{0xF73108CE, 15, 0x0001EE62},
		{0xF73108CE, 16, 0x0000F731},
		{0xF73108CE, 17, 0x00007B98},
		{0xF73108CE, 18, 0x00003DCC},
		{0xF73108CE, 19, 0x00001EE6},
		{0xF73108CE, 20, 0x00000F73},
		{0xF73108CE, 21, 0x000007B9},
		{0xF73108CE, 22, 0x000003DC},
		{0xF73108CE, 23, 0x000001EE},
		{0xF73108CE, 24, 0x000000F7},
		{0xF73108CE, 25, 0x0000007B},
		{0xF73108CE, 26, 0x0000003D},
		{0xF73108CE, 27, 0x0000001E},
		{0xF73108CE, 28, 0x0000000F},
		{0xF73108CE, 29, 0x00000007},
		{0xF73108CE, 30, 0x00000003},
		{0xF73108CE, 31, 0x00000001},
		{0xF73108CE, 32, 0x00000000},
		{0xF73108CE, 33, 0x00000000},
	}
	for i, a := range v {
		o := shr(a.sh, a.in)
		if o != a.out {
			t.Errorf("Shr failure #%d: shr(%d, 0x%X) => code gave 0x%X, test wants 0x%X", i, a.sh, a.in, o, a.out)
		}
	}
}
