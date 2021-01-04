package data

import (
	"database/sql"
	"log"
	"math/rand"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=radix10 password=12345 dbname=gwp")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// createUUID create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID(uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7f
	// Set the four most signficant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
}
