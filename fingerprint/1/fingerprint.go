package fingerprint

import "crypto/md5"

func Hash(data []byte) [md5.Size]byte {
	return md5.Sum(data)
}
