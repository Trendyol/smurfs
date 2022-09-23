package verifier

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"hash"
	"io"
	"os"
)

// Verifier can check a reader against its correctness.
type Verifier interface {
	io.Writer
	Verify() error
	VerifyFile(ctx context.Context, filePath string) error
}

var _ Verifier = &sha256Verifier{}

type sha256Verifier struct {
	hash.Hash
	wantedHash []byte
}

// NewSha256Verifier creates a Verifier that tests against the given hash.
func NewSha256Verifier(hashed string) Verifier {
	raw, _ := hex.DecodeString(hashed)
	return &sha256Verifier{
		Hash:       sha256.New(),
		wantedHash: raw,
	}
}

func (v *sha256Verifier) Verify() error {
	if bytes.Equal(v.wantedHash, v.Sum(nil)) {
		return nil
	}
	return errors.Errorf("checksum does not match, want: %x, got %x", v.wantedHash, v.Sum(nil))
}

func (v *sha256Verifier) VerifyFile(ctx context.Context, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrapf(err, "could not open file %s", filePath)
	}

	_, err = io.Copy(v, file)
	return v.Verify()
}
