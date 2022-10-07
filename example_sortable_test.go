// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package fuzzy

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"bufio"
	"os"
	"io"
	"bytes"
)

// Player is a very simple data model.
type Player struct {
	Firstname string
	Lastname  string
}

// Name returns the full name of the Player.
func (p *Player) Name() string {
	return strings.TrimSpace(p.Firstname + " " + p.Lastname)
}

// Team is a collection of Player items. This is where fuzzy.Sortable
// must be implemented to enable fuzzy sorting.
type Team []Player

// Default sort.Interface methods
func (t Team) Len() int      { return len(t) }
func (t Team) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// Less is used as a tie-breaker when fuzzy match score is the same.
func (t Team) Less(i, j int) bool { return t[i].Name() < t[j].Name() }

// Keywords implements Sortable.
// Comparisons are based on the the full name of the player.
func (t Team) Keywords(i int) string { return t[i].Name() }


func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

// Fuzzy sort players by name.
func TestExampleSort(u *testing.T) {
	fmt.Println("===================")
	readFile, err := os.Open("dictionary2.txt")

	count, _ := lineCounter(readFile)
	t := make(Team, count + 1)

	readFile, err = os.Open("dictionary2.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
	
	i := 0
    for fileScanner.Scan() {
        fmt.Println(i)
		t[i] = Player{Firstname: string(fileScanner.Text())}
        fmt.Println(fileScanner.Text())
		i++
    }
  
    readFile.Close()
	// var t = Team{
	// 	&Player{Firstname: "nguyễn thị tú"},
	// 	&Player{Firstname: "hồ_chí_minh"},
	// 	&Player{Firstname: "nguyễn thanh tú"},
	// 	&Player{Firstname: "tú mập ú"},
	// 	&Player{Firstname: "tú xương"},
	// 	&Player{Firstname: "túng"},
	// 	&Player{Firstname: "nghèo túng"},
	// 	&Player{Firstname: "súng"},
	// 	&Player{Firstname: "mùng"},
	// 	&Player{Firstname: "mụn"},
	// 	&Player{Firstname: "cụng"},
	// 	&Player{Firstname: "hàng tuấn thiên"},
	// 	&Player{Firstname: "lê thị trang"},
	// 	&Player{Firstname: "trương văn lanh"},
	// 	&Player{Firstname: "đoàn minh vương"},
	// 	&Player{Firstname: "hà nội"},
	// 	&Player{Firstname: "đà nẵng"},
	// 	&Player{Firstname: "hải thượng lãn ông"},
	// 	&Player{Firstname: "đồng nai"},
	// 	&Player{Firstname: "hò xuân hương"},
	// 	&Player{Firstname: "hồ tùng mậu"},
	// }
	// Unsorted
	// fmt.Println(t[0].Name())

	// Initials
	// Sort(t, "taa")
	// fmt.Println(t[0].Name())

	// // Initials beat start of string
	// Sort(t, "al")
	// fmt.Println(t[0].Name())

	// // Start of word
	// Sort(t, "ox")
	// fmt.Println(t[0].Name())

	// // Earlier in string = better match
	// Sort(t, "x")
	// fmt.Println(t[0].Name())

	// Diacritics ignored if query is ASCII

	fmt.Println("==========================")

	input := "ho chis minh"
	start := time.Now()
	Sort(t, input)
	fmt.Println("time elapse: ", time.Since(start))
	// fmt.Println(time.Since(start))
	fmt.Println("input: " + input)
	fmt.Println(t[0].Name())
	fmt.Println(t[1].Name())
	fmt.Println(t[2].Name())
	fmt.Println(t[3].Name())
	fmt.Println(t[4].Name())
	fmt.Println(t[5].Name())
	fmt.Println(t[6].Name())
	fmt.Println(t[7].Name())
	fmt.Println(t[8].Name())
	fmt.Println(t[9].Name())
	fmt.Println(t[10].Name())

	fmt.Println("==========================")

	

	// for i, _ := range t {
	// 	fmt.Println(t[i].Name())
	// }

	// But not if query isn't
	// Sort(t, "né")
	// fmt.Println(t[0].Name())
	// Sort(t, "ne")
	// fmt.Println(t[0].Name())
}
