package main

import (
	"fmt"
	"math/rand"
	s "strings"
)

var p = fmt.Println
var t = "Tätä tekstiä testataan strings funktioilla"

func main() {
	p(t)
	p("onko tekstissä 'fun':", s.Contains(t, "fun"))
	p("kuinka monta s:ää tekstissä on:", s.Count(t, "s"))
	p("Monesko merkki eka 'e' on:", s.Index(t, "e"))
	p("Toistetaan kaksi kertaa:\n", s.Repeat(t, 2))
	p("vaihdetaan kaikki i:t ö:ksi\n", s.Replace(t, "i", "ö", -1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("kaikki kirjaimet isoiksi:\n", s.ToUpper(t))
	p("Kuinka pitkä teksi on:", len(t))
	p()
	p("lopuksi vielä tulostetaan teksti sattuman monta kertaa")
	r := (rand.Intn(9))
	for i := 0; i < r; i++ {
		p(t)
	}

}
