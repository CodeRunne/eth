package utility

import (
	"math"
	"math/rand"
	"time"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("../.env")
    if err!= nil {
        log.Fatal(err)
    }
}

func GenerateToken(length int) string {
	var charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

    b := make([]byte, length)
    for i := range b {
      b[i] = charset[seededRand.Intn(len(charset))]
    }
    return string(b)
}

func Timer(timer time.Time) float64 {
	return math.Floor(time.Since(timer).Hours())
}