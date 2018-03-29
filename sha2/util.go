package sha2

const (
	Sha256BlocksizeBits  = 512
	Sha256BlocksizeBytes = Sha256BlocksizeBits / 8
	Sha256WordsizeBits   = 32
	Sha256WordsizeBytes  = Sha256WordsizeBits / 8

	// Too bad there are no const static arrays
	sha256k00 uint32 = 0x428a2f98
	sha256k01 uint32 = 0x71374491
	sha256k02 uint32 = 0xb5c0fbcf
	sha256k03 uint32 = 0xe9b5dba5
	sha256k04 uint32 = 0x3956c25b
	sha256k05 uint32 = 0x59f111f1
	sha256k06 uint32 = 0x923f82a4
	sha256k07 uint32 = 0xab1c5ed5
	sha256k08 uint32 = 0xd807aa98
	sha256k09 uint32 = 0x12835b01
	sha256k10 uint32 = 0x243185be
	sha256k11 uint32 = 0x550c7dc3
	sha256k12 uint32 = 0x72be5d74
	sha256k13 uint32 = 0x80deb1fe
	sha256k14 uint32 = 0x9bdc06a7
	sha256k15 uint32 = 0xc19bf174
	sha256k16 uint32 = 0xe49b69c1
	sha256k17 uint32 = 0xefbe4786
	sha256k18 uint32 = 0x0fc19dc6
	sha256k19 uint32 = 0x240ca1cc
	sha256k20 uint32 = 0x2de92c6f
	sha256k21 uint32 = 0x4a7484aa
	sha256k22 uint32 = 0x5cb0a9dc
	sha256k23 uint32 = 0x76f988da
	sha256k24 uint32 = 0x983e5152
	sha256k25 uint32 = 0xa831c66d
	sha256k26 uint32 = 0xb00327c8
	sha256k27 uint32 = 0xbf597fc7
	sha256k28 uint32 = 0xc6e00bf3
	sha256k29 uint32 = 0xd5a79147
	sha256k30 uint32 = 0x06ca6351
	sha256k31 uint32 = 0x14292967
	sha256k32 uint32 = 0x27b70a85
	sha256k33 uint32 = 0x2e1b2138
	sha256k34 uint32 = 0x4d2c6dfc
	sha256k35 uint32 = 0x53380d13
	sha256k36 uint32 = 0x650a7354
	sha256k37 uint32 = 0x766a0abb
	sha256k38 uint32 = 0x81c2c92e
	sha256k39 uint32 = 0x92722c85
	sha256k40 uint32 = 0xa2bfe8a1
	sha256k41 uint32 = 0xa81a664b
	sha256k42 uint32 = 0xc24b8b70
	sha256k43 uint32 = 0xc76c51a3
	sha256k44 uint32 = 0xd192e819
	sha256k45 uint32 = 0xd6990624
	sha256k46 uint32 = 0xf40e3585
	sha256k47 uint32 = 0x106aa070
	sha256k48 uint32 = 0x19a4c116
	sha256k49 uint32 = 0x1e376c08
	sha256k50 uint32 = 0x2748774c
	sha256k51 uint32 = 0x34b0bcb5
	sha256k52 uint32 = 0x391c0cb3
	sha256k53 uint32 = 0x4ed8aa4a
	sha256k54 uint32 = 0x5b9cca4f
	sha256k55 uint32 = 0x682e6ff3
	sha256k56 uint32 = 0x748f82ee
	sha256k57 uint32 = 0x78a5636f
	sha256k58 uint32 = 0x84c87814
	sha256k59 uint32 = 0x8cc70208
	sha256k60 uint32 = 0x90befffa
	sha256k61 uint32 = 0xa4506ceb
	sha256k62 uint32 = 0xbef9a3f7
	sha256k63 uint32 = 0xc67178f2

	// For SHA-256, the initial hash value, H(0), obtained by taking the first 32 bits
	// of the fractional parts of the square roots of the first eight prime numbers.
	sha256h00 uint32 = 0x6a09e667
	sha256h01 uint32 = 0xbb67ae85
	sha256h02 uint32 = 0x3c6ef372
	sha256h03 uint32 = 0xa54ff53a
	sha256h04 uint32 = 0x510e527f
	sha256h05 uint32 = 0x9b05688c
	sha256h06 uint32 = 0x1f83d9ab
	sha256h07 uint32 = 0x5be0cd19
)

// Too bad there are no const static arrays
func sha256kByIndex(i int) uint32 {
	switch i {
	case 0:
		return sha256k00
	case 1:
		return sha256k01
	case 2:
		return sha256k02
	case 3:
		return sha256k03
	case 4:
		return sha256k04
	case 5:
		return sha256k05
	case 6:
		return sha256k06
	case 7:
		return sha256k07
	case 8:
		return sha256k08
	case 9:
		return sha256k09
	case 10:
		return sha256k10
	case 11:
		return sha256k11
	case 12:
		return sha256k12
	case 13:
		return sha256k13
	case 14:
		return sha256k14
	case 15:
		return sha256k15
	case 16:
		return sha256k16
	case 17:
		return sha256k17
	case 18:
		return sha256k18
	case 19:
		return sha256k19
	case 20:
		return sha256k20
	case 21:
		return sha256k21
	case 22:
		return sha256k22
	case 23:
		return sha256k23
	case 24:
		return sha256k24
	case 25:
		return sha256k25
	case 26:
		return sha256k26
	case 27:
		return sha256k27
	case 28:
		return sha256k28
	case 29:
		return sha256k29
	case 30:
		return sha256k30
	case 31:
		return sha256k31
	case 32:
		return sha256k32
	case 33:
		return sha256k33
	case 34:
		return sha256k34
	case 35:
		return sha256k35
	case 36:
		return sha256k36
	case 37:
		return sha256k37
	case 38:
		return sha256k38
	case 39:
		return sha256k39
	case 40:
		return sha256k40
	case 41:
		return sha256k41
	case 42:
		return sha256k42
	case 43:
		return sha256k43
	case 44:
		return sha256k44
	case 45:
		return sha256k45
	case 46:
		return sha256k46
	case 47:
		return sha256k47
	case 48:
		return sha256k48
	case 49:
		return sha256k49
	case 50:
		return sha256k50
	case 51:
		return sha256k51
	case 52:
		return sha256k52
	case 53:
		return sha256k53
	case 54:
		return sha256k54
	case 55:
		return sha256k55
	case 56:
		return sha256k56
	case 57:
		return sha256k57
	case 58:
		return sha256k58
	case 59:
		return sha256k59
	case 60:
		return sha256k60
	case 61:
		return sha256k61
	case 62:
		return sha256k62
	case 63:
		return sha256k63
	default:
		return 0
	}
}

// Ch(x, y, z)=(x and y) xor ( complement x and z)
// Ch(x, y, z)=(x & y) ^ ( ^x & z)
// "Choose" the bit from y or z based on the bit in x
// if the bit in x is 1, choose the bit in y
// if the bit in x is 0, choose the bit in z
func ch(x, y, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

// Maj(x, y, z)=(x and y) xor (x and z) xor (y and z)
// Maj(x, y, z)=(x & y) ^ (x & z) ^ (y & z)
// "Majority" between x,y,z of each bit
// the value of each bit that is most common in the 3 inputs
func maj(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func upperSigma0(x uint32) uint32 {
	return rotr(2, x) ^ rotr(13, x) ^ rotr(22, x)
}

func upperSigma1(x uint32) uint32 {
	return rotr(6, x) ^ rotr(11, x) ^ rotr(25, x)
}

func lowerSigma0(x uint32) uint32 {
	return rotr(7, x) ^ rotr(18, x) ^ shr(3, x)
}

func lowerSigma1(x uint32) uint32 {
	return rotr(17, x) ^ rotr(19, x) ^ shr(10, x)
}

// rotate right
// (x >> n) | (x << sizeof(x) - n).
func rotr(n uint8, x uint32) uint32 {
	n = n % 32
	return (x >> n) | (x << (32 - n))
}

// rotate left
// (x << n) | (x >> sizeof(x) - n)
func rotl(n uint8, x uint32) uint32 {
	n = n % 32
	return (x << n) | (x >> (32 - n))
}

// shift right
func shr(n uint8, x uint32) uint32 {
	if n >= 32 {
		return 0
	} else {
		return x >> n
	}
}
