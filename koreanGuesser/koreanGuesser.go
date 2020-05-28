// Tuomo Miettinen
// http://terokarvinen.com/2020/go-programming-course-2020-w22/
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// https://www.geeksforgeeks.org/fmt-scanln-function-in-golang-with-examples/
// https://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
// https://gobyexample.com/time
// https://gobyexample.com/time-formatting-parsing

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// slicet, joista verrataan toisiinsa round() funktiossa
var korean []string

func roundKorean() {
	korean = append(korean, "ㄱ")
	korean = append(korean, "ㄴ")
	korean = append(korean, "ㄷ")
	korean = append(korean, "ㄹ")
	korean = append(korean, "ㅁ")
	korean = append(korean, "ㅂ")
	korean = append(korean, "ㅅ")
	korean = append(korean, "ㅇ, last character")
	korean = append(korean, "ㅇ, first character")
	korean = append(korean, "ㅈ")
	korean = append(korean, "ㅊ")
	korean = append(korean, "ㅋ")
	korean = append(korean, "ㅌ")
	korean = append(korean, "ㅍ")
	korean = append(korean, "ㅎ")
	korean = append(korean, "ㅏ")
	korean = append(korean, "ㅑ")
	korean = append(korean, "ㅓ")
	korean = append(korean, "ㅕ")
	korean = append(korean, "ㅗ")
	korean = append(korean, "ㅛ")
	korean = append(korean, "ㅜ")
	korean = append(korean, "ㅠ")
	korean = append(korean, "ㅡ")
	korean = append(korean, "ㅣ")
	korean = append(korean, "ㄲ")
	korean = append(korean, "ㄸ")
	korean = append(korean, "ㅃ")
	korean = append(korean, "ㅆ")
	korean = append(korean, "ㅉ")
	korean = append(korean, "ㅐ")
	korean = append(korean, "ㅒ")
	korean = append(korean, "ㅔ")
	korean = append(korean, "ㅖ")
	korean = append(korean, "ㅚ")
	korean = append(korean, "ㅙ")
	korean = append(korean, "ㅟ")
	korean = append(korean, "ㅞ")
	korean = append(korean, "ㅢ")
	korean = append(korean, "ㅝ")
}

var vastaukset []string

func roundVastaukset() {

	vastaukset = append(vastaukset, "g")
	vastaukset = append(vastaukset, "n")
	vastaukset = append(vastaukset, "d")
	vastaukset = append(vastaukset, "l")
	vastaukset = append(vastaukset, "m")
	vastaukset = append(vastaukset, "b")
	vastaukset = append(vastaukset, "s")
	vastaukset = append(vastaukset, "ng")
	vastaukset = append(vastaukset, "")
	vastaukset = append(vastaukset, "j")
	vastaukset = append(vastaukset, "ch")
	vastaukset = append(vastaukset, "k")
	vastaukset = append(vastaukset, "t")
	vastaukset = append(vastaukset, "p")
	vastaukset = append(vastaukset, "h")
	vastaukset = append(vastaukset, "a")
	vastaukset = append(vastaukset, "ya")
	vastaukset = append(vastaukset, "eo")
	vastaukset = append(vastaukset, "yeo")
	vastaukset = append(vastaukset, "o")
	vastaukset = append(vastaukset, "yo")
	vastaukset = append(vastaukset, "u")
	vastaukset = append(vastaukset, "yu")
	vastaukset = append(vastaukset, "eu")
	vastaukset = append(vastaukset, "i")
	vastaukset = append(vastaukset, "gg")
	vastaukset = append(vastaukset, "dd")
	vastaukset = append(vastaukset, "bb")
	vastaukset = append(vastaukset, "ss")
	vastaukset = append(vastaukset, "jj")
	vastaukset = append(vastaukset, "ae")
	vastaukset = append(vastaukset, "yae")
	vastaukset = append(vastaukset, "e")
	vastaukset = append(vastaukset, "ye")
	vastaukset = append(vastaukset, "oe")
	vastaukset = append(vastaukset, "wae")
	vastaukset = append(vastaukset, "wi")
	vastaukset = append(vastaukset, "we")
	vastaukset = append(vastaukset, "ui")
	vastaukset = append(vastaukset, "wo")
}

// funktio, joka tulostaa merkit, jotta voi nähdä oikeat vastaukset tai muistella ja oppia
func tulostaMerkit() {
	for i := 0; i < 40; i++ {
		fmt.Println(korean[i], " = ", vastaukset[i])
	}
}

// ohjelmanlaajuiset muuttujat, jotta ne voi nollata helposti ja käyttää round() funktiossa
var oikeatVastaukset int
var kaikkiVastaukset int

// itse peli funktio, jossa vastataan kysymyksiin
func round() {

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(40)

	fmt.Println("\nEnter 'q' to stop and return")
	fmt.Println(" | How is this")
	fmt.Println(" ↓")
	fmt.Println(" ", korean[r])
	fmt.Println(" ↑")
	fmt.Println(" | pronounced? Write now...")

	var vastaus string
	fmt.Scanln(&vastaus)
	fmt.Println("\n\n\n############################################")
	fmt.Println("############################################\n")

	fmt.Println("\nYou answered: ", vastaus)

	switch vastaus {
	case vastaukset[r]:
		fmt.Println("\n----- You are right! -----\n")
		oikeatVastaukset++
	case "q":
		ending(oikeatVastaukset, kaikkiVastaukset)
		return
	default:
		fmt.Println("\n----- Wrong! -----")
		fmt.Println("The right answer for : '", korean[r], "' was: '", vastaukset[r], "'")
	}
	kaikkiVastaukset++
	round()
}

// pääfunktio, tästä käyttäjä liikkuu muualle, tarkastelee listoja tai poistuu hienostuneesti
func start() {

	fmt.Println("\n\n\n############################################")
	fmt.Println("############################################\n")
	fmt.Println("What would you like to do?")
	fmt.Println("Enter 'start' to start the guessing game")
	fmt.Println("Enter 'list' to open pronunciation list")
	fmt.Println("Enter 'hs' to open current highscore")
	fmt.Println("Enter 'exit' to exit")

	var vastaus string
	fmt.Scanln(&vastaus)
	fmt.Println("You chose: ", vastaus)
	fmt.Println("____________________________________________")

	switch vastaus {
	case "start":
		oikeatVastaukset = 0
		kaikkiVastaukset = 0
		round()
	case "list":
		tulostaMerkit()
	case "hs":
		lueHighscore()
	case "exit":
		return
	default:
		fmt.Println("Not a valid argument.")
	}

	start()
}

// kun käyttäjä lopettaa round() tätä kutsutaan, jolloin käyttäjä näkee oikeiden arvausten määrän
func ending(oikeat int, kaikki int) {

	highscore(oikeat, kaikki)

	fmt.Println("\n\n\n############################################")
	fmt.Println("############################################\n")
	fmt.Println("Your score was", oikeat, " out of ", kaikki)
	fmt.Println("")
	fmt.Println("\nPress Enter to continue...")

	var vastaus string
	fmt.Scanln(&vastaus)

}

// ennätystä käsittelevä funktio, joka kutsuu tarvittaessa lueHighscore() ja kirjoitaHighscore()
func highscore(oikeat int, kaikki int) {

	file, err := os.Open("./highscore.txt")
	if err != nil {
		kirjoitaHighscore(oikeat, kaikki)
		return
	}
	defer file.Close()

	var hscore string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hscore = string(scanner.Text()[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	i, err := strconv.Atoi(hscore)
	if err != nil {
		fmt.Println(err)
		return
	}

	if oikeat > i {
		kirjoitaHighscore(oikeat, kaikki)
	}
}

// tätä kutsutaan, kun oikeiden vastauksien määrä ylittää highscore.txt tiedostossa olevan numeron tai kun highscore.txt ei löydy
func kirjoitaHighscore(oikeat int, kaikki int) {
	nyt := time.Now()

	text := fmt.Sprintf("%d%s%d%s %d %s %d", oikeat, " out of ", kaikki, " ---- The date of the record: ", nyt.Year(), nyt.Month(), nyt.Day())
	t := []byte(text)
	err := ioutil.WriteFile("./highscore.txt", t, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// jos haluaa tarkastella ennätystä start() funktiosta voi
func lueHighscore() {
	file, err := os.Open("./highscore.txt")
	if err != nil {
		fmt.Println("\n\n highscore.txt not found.")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("\n\n\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// alustaa sliste korean[] ja vastaukset[], joita käytetään round() määrittelemään oikeat vastaukset
// tästä lähetään start() :illa ja palataan, kun ollaan valittu exit start() :ista
func main() {

	roundKorean()
	roundVastaukset()
	start()

	fmt.Println("Thanks for running")
}
