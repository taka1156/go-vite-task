package auth

import "crypto/md5"

func HashMd5(password string) [16]byte {
	md5Password := md5.Sum([]byte(password))

	return md5Password
}
