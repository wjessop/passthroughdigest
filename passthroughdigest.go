package passthroughdigest

import (
	"crypto/md5"
	"hash"
	"io"
)

// DigestWriter is an interface that describes PassthroughDigest
type DigestWriter interface {
	Write(p []byte) (n int, err error)
	Digest() []byte
}

// PassthroughDigest is a construct that updates an internal digest
// with any data written to the output io
type PassthroughDigest struct {
	digest hash.Hash
	dst    io.Writer
}

// NewPassthroughDigest creates a new PassthroughDigest object
func NewPassthroughDigest(out io.Writer) *PassthroughDigest {
	return &PassthroughDigest{md5.New(), out}
}

// Write writes bytes to the output and digest calculator
func (pd *PassthroughDigest) Write(p []byte) (written int, err error) {
	for {
		nw, ew := pd.dst.Write(p[written:])
		if nw > 0 {
			pd.digest.Write(p[written : written+nw])
			written += nw
		}
		if ew != nil {
			err = ew
			break
		}

		if len(p) == written {
			break
		}
	}
	return written, err
}

// Digest returns the current calculated hash of all bytes written to dst
func (pd *PassthroughDigest) Digest() []byte {
	return pd.digest.Sum(nil)
}
