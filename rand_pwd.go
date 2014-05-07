package pwd_gen 

import (
	"io"
    	"crypto/rand"
    	)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

func rand(length int) string {
    chars := StdChars
    new_pword := make([]byte, length)
    random_data := make([]byte, length+(length/4)) // storage for random bytes.
    clen := byte(len(chars))
    maxrb := byte(256 - (256 % len(chars)))
    i := 0
    for {
        if _, err := io.ReadFull(rand.Reader, random_data); err != nil {
            panic(err)
        }
        for _, c := range random_data {
            if c >= maxrb {
                continue
            }
            new_pword[i] = chars[c%clen]
            i++
            if i == length {
                return string(new_pword)
            }
        }
    }
    panic("unreachable")
}
