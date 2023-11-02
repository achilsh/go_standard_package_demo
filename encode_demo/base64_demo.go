package encode_demo

import (
	"encoding/base64"
	"fmt"
)

type EncodeDemo struct {
	enc *base64.Encoding
}

func RunEncodeDemo() {
	t := NewEncodeDemo()
	s1 := []byte("this is demo test for encoding.")
	d1, e := t.Encode(s1)
	if e != nil {
		panic(e)
	}
	fmt.Println("encoded data: ", string(d1))

	d2, e := t.Decode(d1)
	if e != nil {
		panic(e)
	}
	fmt.Println("decode data: ", string(d2))
	//

	d3 := []byte("this decode/encode string demo")
	d4, e := t.EncodeToString(d3)
	if e != nil {
		panic(e)
	}
	t.DecodeFromString(d4)
}

func NewEncodeDemo() *EncodeDemo {
	r := &EncodeDemo{
		enc: base64.StdEncoding, //URLEncoding, URLEncoding, RawStdEncoding, RawURLEncoding.
	}
	return r
}
func (p *EncodeDemo) Encode(src []byte) ([]byte, error) {
	fmt.Println("run encode......")
	if src == nil || len(src) <= 0 {
		return nil, nil
	}
	dst := make([]byte, p.enc.EncodedLen(len(src)))
	p.enc.Encode(dst, src)
	return dst, nil
}
func (p *EncodeDemo) Decode(dst []byte) ([]byte, error) {
	fmt.Println("run decode......")
	if dst == nil || len(dst) == 0 {
		panic("is nil")
	}
	src := make([]byte, p.enc.DecodedLen(len(dst)))
	n, e := p.enc.Decode(src, dst)
	if e != nil {
		fmt.Println("decode fail, e: ", e)
		return nil, e
	}
	return src[:n], nil
}

// /
func (p *EncodeDemo) EncodeToString(src []byte) (string, error) {
	fmt.Println("run encode to string......")
	if src == nil || len(src) <= 0 {
		return "", nil
	}
	dst := p.enc.EncodeToString(src)
	fmt.Println("encode to string val: ", string(dst))
	return dst, nil
}
func (p *EncodeDemo) DecodeFromString(src string) ([]byte, error) {
	fmt.Println("run decode from string......")
	if len(src) == 0 {
		panic("is nil")
	}
	dst, e := p.enc.DecodeString(src)
	if e != nil {
		fmt.Println("decode fail, e: ", e)
		return nil, e
	}
	fmt.Println("decode from string, val: ", string(dst))
	return dst[:], nil
}
