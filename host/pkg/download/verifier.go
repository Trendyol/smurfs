package download

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"
)

// Verifier can check a reader against it's correctness.
type Verifier interface {
	io.Writer
	Verify() error
}

var _ Verifier = sha256Verifier{}

type sha256Verifier struct {
	hash.Hash
	wantedHash []byte
}

// NewSha256Verifier creates a Verifier that tests against the given hash.
func NewSha256Verifier(hashed string) Verifier {
	raw, _ := hex.DecodeString(hashed)
	return sha256Verifier{
		Hash:       sha256.New(),
		wantedHash: raw,
	}
}

func (v sha256Verifier) Verify() error {
	klog.V(1).Infof("Compare sha256 (%s) signed version", hex.EncodeToString(v.wantedHash))
	if bytes.Equal(v.wantedHash, v.Sum(nil)) {
		return nil
	}
	return errors.Errorf("checksum does not match, want: %x, got %x", v.wantedHash, v.Sum(nil))
}
