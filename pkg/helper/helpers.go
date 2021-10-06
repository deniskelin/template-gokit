package helper

import (
	"io"
	"reflect"
	"unsafe"

	"github.com/valyala/bytebufferpool"
)

// B2S converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func B2S(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

// S2B converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func S2B(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func StreamToString(stream io.Reader) (string, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	_, err := io.Copy(buf, stream)
	if err != nil {
		return "", err
	}
	out := make([]byte, 0, buf.Len())
	copy(out, buf.Bytes())
	return B2S(out), nil
}
