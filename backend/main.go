package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// countResult count package
type countResult struct {
	Count int `json:"count"`
}

// columnName colune name package
type columnName struct {
	Name []string `json:"name"`
}

// result content package
type results struct {
	Result []result `json:"result"`
}

// result content
type result struct {
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	Maintainer string    `json:"maintainer"`
	Content    string    `json:"content"`
}

/*
Count return result count
ex:
{
	count:10
}
*/
func Count(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if ok := recover(); ok != nil {
			log.Println(ok)
		}
	}()
	cr := &countResult{Count: 10}
	enc := json.NewEncoder(w)
	err := enc.Encode(cr)
	if err != nil {
		log.Panic(err)
	}
}

/*
ColumnName return columns name
ex:
{
	"name":["title","date","maintainer","content"]
}
*/
func ColumnName(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if ok := recover(); ok != nil {
			log.Println(ok)
		}
	}()
	cn := &columnName{Name: []string{"title", "date", "maintainer", "content"}}
	enc := json.NewEncoder(w)
	err := enc.Encode(cn)
	if err != nil {
		log.Panic(err)
	}
}

/*
Result  return result content
ex:
{
	{
		"title": "Happy"
		"date": time.Now()
		"maintainer": "Penny"
		"content": "Keep on going never give up"
	},
	{
		"title": "Go Go Go"
		"date": time.Now()
		"maintainer": "Joe"
		"content": "Winners do what losers don't want to do"
	}
}
*/
func Result(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if ok := recover(); ok != nil {
			log.Println(ok)
		}
	}()
	rs := &results{}
	lenght := 100
	rs.Result = make([]result, lenght)
	i := 0
	for i < lenght {
		if i%2 == 0 {
			rs.Result[i] = result{
				Title:      "Go Go Go",
				Date:       time.Now(),
				Maintainer: "Joe",
				Content:    "Winners do what losers don't want to do",
			}
		} else {
			rs.Result[i] = result{
				Title:      "Happy",
				Date:       time.Now(),
				Maintainer: "Penny",
				Content:    "Keep on going never give up",
			}
		}

		i++
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(rs)
	if err != nil {
		log.Panic(err)
	}

}

func main() {
	// http://127.0.0.1:8080/count
	http.HandleFunc("/count", Count)

	// http://127.0.0.1:8080/ColumnName
	http.HandleFunc("/ColumnName", ColumnName)

	// http://127.0.0.1:8080/result
	http.HandleFunc("/result", Result)

	log.Println("Listening on ", ":8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
