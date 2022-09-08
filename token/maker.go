package token

import (
	"fmt"
	"log"
	"time"
)

// Maker is an interface to manage the token life cycle.
type Maker interface {
	// CreateToken create a new token for a specific username and duration.
	CreateToken(username string, ttl time.Duration) (string, error)

	// VerifyToken Checks if the provided token is valid.
	VerifyToken(token string) (*Payload, error)
}

// after running this function on both approaches (JWT, PASETO)
// it turns out that Paseto is more faster then the jwt method.
func MeasureApproach(maker Maker) {
	t1 := time.Now()
	t, err := maker.CreateToken("islam", time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now().Sub(t1))
	t1 = time.Now()
	p, err := maker.VerifyToken(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now().Sub(t1))
	fmt.Println(p)
}
