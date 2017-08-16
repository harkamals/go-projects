package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// User defined type
type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {

	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}

	defer resp.Body.Close()

	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}

}

func (w *webPage) isOK() bool {
	return w.err == nil
}

// User defined type (or as an alias to built-in type)
type SummableSlice []int

// Why no *pointer here ?!
func (s SummableSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {

	// Get webPage
	//------------------------------------------------------
	// Method -1
	// w := &webPage{url: "http://google.com"}

	// Method -2
	// w := &webPage{}
	//w.url = "http://google.com"

	// Method -3
	w := new(webPage)
	w.url = "http://google.com"

	w.get()
	fmt.Printf("Url:%s Err:%s Len:%d\n", w.url, w.err, len(w.body))
	if w.isOK() {
		fmt.Println("WebPage downloaded")
	} else {
		fmt.Println("Error while downloading")
	}

	// Summable slice - method call
	s := &SummableSlice{1, 2, 3}
	fmt.Println(s.sum())

}
