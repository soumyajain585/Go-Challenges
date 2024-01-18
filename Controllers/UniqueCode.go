package Controllers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUniqueCode() {
	rand.Seed(time.Now().UnixNano())
	uniquePart := rand.Intn(1000000)
	uniquePartStr := fmt.Sprintf("%06d", uniquePart)
	prefix := "ET"
	code := fmt.Sprintf("%s%s", prefix, uniquePartStr)
	fmt.Println(code)
}
