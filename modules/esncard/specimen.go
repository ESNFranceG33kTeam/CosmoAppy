package esncard

import (
	"math/rand"
	"strconv"
	"time"
)

func Specimen() {
	rand.Seed(time.Now().UnixNano())

	rand_numb := strconv.Itoa(rand.Intn(9999999-1000000) + 1000000)
	NewESNcard(&ESNcard{Id_adherent: 1, Esncard: rand_numb})

	rand_numb = strconv.Itoa(rand.Intn(9999999-1000000) + 1000000)
	NewESNcard(&ESNcard{Id_adherent: 2, Esncard: rand_numb})

	AutoNewMonthlyStat(time.Now().AddDate(0, 0, 0).Format("2006-01"))
}
