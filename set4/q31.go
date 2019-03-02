package set4

import (
	"errors"
	"fmt"
	"matasano/hash"
	"matasano/types"
	"math/rand"
	"net/http"
	"time"
)

var h func(msg []byte) string

func HMACSign() func(msg []byte) string {
	//rand.Seed(int64(time.Now().Second()))
	l := rand.Intn(20) + 1
	key := make([]byte, l)
	rand.Read(key)

	return func(msg []byte) string {
		return hash.HMAC(key, msg)
	}
}

func insecureCompare(sha1, sha2 string) bool {
	hex1 := types.Hex{B: []byte(sha1)}
	hex2 := types.Hex{B: []byte(sha2)}

	b1, err := hex1.Decode()
	if err != nil {
		return false
	}
	b2, err := hex2.Decode()
	if err != nil {
		return false
	}

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
		time.Sleep(time.Millisecond * 50)
	}
	return true
}

func hello(w http.ResponseWriter, r *http.Request) {
	file, ok := r.URL.Query()["file"]
	if !ok || len(file) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("file missing"))
		return
	}

	signature, ok := r.URL.Query()["signature"]
	if !ok || len(signature) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("signature missing"))
		return
	}

	expectedSha := h([]byte(file[0]))

	if insecureCompare(signature[0], expectedSha) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Matching signature"))
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong signature"))
		return
	}
}

func server() {
	h = HMACSign()
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

func Solve31() ([]byte, error) {
	file := "foo"
	url := "http://localhost:8000/?file=%s&signature=%s"

	mac := make([]byte, 20)

	for i := 0; i < 20; i++ {
		for j := 0; j < 256; j++ {
			mac[i] = byte(j)
			h := types.Hex{}
			h.Encode(mac)
			fUrl := fmt.Sprintf(url, file, h.Get())

			startTime := time.Now()
			resp, err := http.Get(fUrl)
			duration := time.Now().Sub(startTime)

			ms := duration.Seconds() * 1000

			if err != nil {
				return []byte{}, err
			}
			if resp.StatusCode == 200 {
				return mac, nil
			}

			if ms > float64(i+1)*50 {
				break
			}
			resp.Body.Close()
		}
	}

	return []byte{}, errors.New("didn't get the mac")
}
