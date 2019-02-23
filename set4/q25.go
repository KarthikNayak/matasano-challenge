package set4

import (
	"crypto/aes"
	"io/ioutil"
	"log"
	"matasano/cipher"
	"matasano/helpers"
	"matasano/types"
)

func EncryptCTR() ([]byte, error) {
	b, err := ioutil.ReadFile("q25_data.txt") // just pass the file name
	if err != nil {
		return []byte{}, err
	}

	b64 := types.Base64{B: b}
	data, err := b64.Decode()
	if err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data); i += 16 {
		block.Decrypt([]byte(data[i:i+16]), []byte(data[i:i+16]))
	}

	var c cipher.CTR
	err = c.Init([]byte("YELLOW SUBMARINE"), 128, uint64(0))
	if err != nil {
		return []byte{}, err
	}

	return c.Encode(&types.Text{T: data})
}

func edit(ciphertext, key, newtext []byte, offset int) ([]byte, error) {
	var c cipher.CTR
	err := c.Init(key, 128, uint64(0))
	if err != nil {
		return []byte{}, err
	}

	decoded, err := c.Decode(&types.Text{T: ciphertext})
	if err != nil {
		return []byte{}, err
	}

	finalText := append(decoded[:offset], newtext...)
	if (offset + len(newtext)) < len(decoded) {
		finalText = append(finalText, decoded[offset+len(newtext):]...)
	}
	return c.Encode(&types.Text{T: finalText})
}

func SolveQ25() error {
	data, err := EncryptCTR()
	if err != nil {
		return err
	}
	key := []byte("YELLOW SUBMARINE")

	keystream := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		newCT, err := edit(data, key, []byte("A"), i)
		if err != nil {
			return err
		}
		keystream[i] = byte('A') ^ newCT[i]
	}

	val, _ := helpers.Xor(&types.Text{T: keystream}, &types.Text{T: data})
	_ = val
	return nil
}
