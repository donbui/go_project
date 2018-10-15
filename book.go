// 1. Create a go program that has a book structure. the structure should contain relevant book properties
// 2. Create a slice of books and print them to stdout
// 3. (Advanced) Sort the books by title
// 4. (Advanced) Output the books to json or xml

package main

import "fmt"
import "sort"
import "encoding/json"
// import "os"

type book struct {
  Title string `json:"title"`
  Author string `json:"author"`
	Edition int `json:"edition"`
	Illustrator string `json:"illustrator"`
	Abridged bool `json:"abridged"`
	NumberOfPages int `json:"numberOfPages"`
	ISBN string `json:"isbn"`
}

type Library []book

func (s Library) Len() int {
	return len(s)
}
func (s Library) Less(i, j int) bool {
	return s[i].Title < s[j].Title
}

func (s Library) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// loop and read these from a file?
	var s Library = Library{
		{Title: "Blankets", Author: "Craig Thompson", Edition: 1, Illustrator: "Craig Thompson", Abridged: false, NumberOfPages: 592, ISBN: "978-1770462182"},
		{Title: "Collected Poems of Philip Larkin", Author: "Philip Larkin", Edition: 2, Abridged: false, NumberOfPages: 192, ISBN: "978-0571216543"},
		{Title: "Lucky Jim", Author: "Kingsley Amis", Edition: 3, Abridged: false, NumberOfPages: 320, ISBN: "978-0141046716"},
		{Title: "A Canticle for Leibowitz", Author: "Walter Miller", Edition: 4, Abridged: false, NumberOfPages: 368, ISBN: "978-0553273816"},
		{Title: "The Drowned World", Author: "J.G Ballard", Edition: 1, Abridged: false, NumberOfPages: 176, ISBN: "978-0007221837"},
	}

	sort.Sort(s)
	// fmt.Println(s)

	j, _ := json.Marshal(s)
  fmt.Println(string(j))
}
