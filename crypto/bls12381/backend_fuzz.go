// +build gofuzz,amd64,!generic

package bls12381

func FuzzBackendMultiplication(data []byte) int {
	if len(data) != 48*2 {
		return 0
	}
	a := new(fe).setBytes(data[:48])
	b := new(fe).setBytes(data[48:])
	for !a.isValid() {
		subAssign(a, &modulus)
	}
	for !b.isValid() {
		subAssign(b, &modulus)
	}
	c0, c1 := new(fe), new(fe)
	mulADX(c0, a, b)
	mulNoADX(c1, a, b)
	if !c0.equal(c1) {
		panic("bad multiplication")
	}
	return 1
}
