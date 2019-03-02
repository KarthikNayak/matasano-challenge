package set4

import (
	"errors"
	"fmt"
	"matasano/types"
	"net/http"
	"time"
)

const (
	waitTime = 5
)

func insecureCompare2(sha1, sha2 string) bool {
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
		time.Sleep(time.Millisecond * waitTime)
	}
	return true
}

func hello2(w http.ResponseWriter, r *http.Request) {
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

	if insecureCompare2(signature[0], expectedSha) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Matching signature"))
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("wrong signature"))
		return
	}
}

func server2() {
	h = HMACSign()
	http.HandleFunc("/", hello2)
	http.ListenAndServe(":8000", nil)
}

func maxIndex(f []float64) int {
	max := 0.0
	index := 0
	for i, val := range f {
		if val > max {
			index = i
			max = val
		}
	}
	fmt.Println("max:", max, index)
	return index
}

func Solve32() ([]byte, error) {
	file := "foo"
	url := "http://localhost:8000/?file=%s&signature=%s"

	mac := make([]byte, 20)

	for i := 0; i < 20; i++ {
		fmt.Println(i)
		var times []float64
		for j := 0; j < 256; j++ {
			norm := 2

			mac[i] = byte(j)
			h := types.Hex{}
			h.Encode(mac)
			fUrl := fmt.Sprintf(url, file, h.Get())

			td := 0.0

			for k := 0; k < norm; k++ {
				startTime := time.Now()
				resp, err := http.Get(fUrl)
				duration := time.Now().Sub(startTime)

				if err != nil {
					return []byte{}, err
				}
				if resp.StatusCode == 200 {
					return mac, nil
				}
				resp.Body.Close()

				td += duration.Seconds() * 1000
			}

			td /= float64(norm)
			times = append(times, td)
		}
		fmt.Println(times)
		mac[i] = byte(maxIndex(times))

	}

	return []byte{}, errors.New("didn't get the mac")
}
