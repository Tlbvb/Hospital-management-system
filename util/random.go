package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
var specialties = [10]string{
    "Emergency Medicine",
    "Critical Care",
    "Cardiology",
    "Neurology",
    "Pediatrics",
    "Obstetrics and Gynecology",
    "Orthopedics",
    "Radiology",
    "Oncology",
    "General Surgery",
}

var departments = [10]string{
	"Emergency",
	"Intensive Care Unit",
	"Cardiology",
	"Neurology",
	"Pediatrics",
	"Obstetrics and Gynecology",
	"Orthopedics",
	"Radiology",
	"Oncology",
	"Surgery",
}



func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomFullName() string {
	return RandomString(6)
}

func RandomAddress()string{
	return RandomString(10)
}

func RandomSpecialization() string{
	return specialties[RandomInt(0,9)]
}

func RandomDep()string{
	return departments[RandomInt(0,9)]
}

func RandomTime() time.Time{
	dur:=RandomInt(time.Hour.Nanoseconds(),time.Hour.Nanoseconds()*24*7)
	return time.Now().Add(time.Duration(dur))
}



// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}