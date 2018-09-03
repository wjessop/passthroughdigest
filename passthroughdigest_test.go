package passthroughdigest_test

import (
	"encoding/hex"
	"testing"

	"github.com/stvp/assert"
	"github.com/wjessop/passthroughdigest"
)

func isPassthroughDigest(t interface{}) bool {
	switch t.(type) {
	case *passthroughdigest.PassthroughDigest:
		return true
	default:
		return false
	}
}

// ShortWriter writes only the first 10 bytes in the first two Write calls.
// Subsequent write calls will write everything
type ShortWriter struct {
	Dst  []byte
	pass int
}

func (sw *ShortWriter) Write(p []byte) (n int, err error) {
	if sw.pass < 2 {
		sw.pass++
		sw.Dst = append(sw.Dst, p[:10]...)
		return 10, nil
	}

	sw.Dst = append(sw.Dst, p...)
	return len(p), nil
}

func TestNewPassthroughDigest(t *testing.T) {
	v := passthroughdigest.NewPassthroughDigest(&ShortWriter{})
	assert.True(t, isPassthroughDigest(v), "should be a PassthroughDigest")
}

func TestWrite(t *testing.T) {
	shortWriter := &ShortWriter{}

	v := passthroughdigest.NewPassthroughDigest(shortWriter)
	v.Write([]byte("This is a test string"))
	assert.Equal(t, "This is a test string", string(shortWriter.Dst))
}

func TestDigest(t *testing.T) {
	shortWriter := &ShortWriter{}

	v := passthroughdigest.NewPassthroughDigest(shortWriter)
	v.Write([]byte("This is a test string"))

	assert.Equal(t, "c639efc1e98762233743a75e7798dd9c", hex.EncodeToString(v.Digest()))
}
