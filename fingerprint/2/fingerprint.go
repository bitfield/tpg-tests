package fingerprint

import "crypto/sha256"

func Hash(data []byte) [sha256.Size]byte {
	return sha256.Sum256(data)
}
