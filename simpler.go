package pwd_gen

import "crypto/rand"

func SimpleRand(str_size int) string {
    alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, str_size)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b%byte(len(alphanum))]
    }
    return string(bytes)
}
