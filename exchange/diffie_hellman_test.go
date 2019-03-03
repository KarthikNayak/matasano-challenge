package exchange

import (
	"fmt"
	"testing"
)

func TestDiffieHellman(t *testing.T) {
	Alice := DF{}
	Bob := DF{}

	Alice.Setup()
	Bob.Setup()

	Alice.SendOnPublic(&Bob)
	Bob.SendOnPublic(&Alice)

	Alice.GenerateSecret()
	Bob.GenerateSecret()

	fmt.Println(Alice.Hash)
	fmt.Println(Bob.Hash)
}
