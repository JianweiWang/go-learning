package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"io"
	"log"
)

func main() {
	//dul(os.Args[1:])
	//urlGet(os.Args[1])
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe("localhost:8080", nil ))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to here," + r.RemoteAddr)
}

func echo() {
	for i := 0; i < len(os.Args); i++ {

		fmt.Printf("%d %s", i, os.Args[i] + "\r\n")
	}

	s := strings.Join(os.Args, " ");

	fmt.Printf(s)
}

func urlGet(url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("fetch error: %v\n" , err)
	}

	_, err1 := io.Copy(os.Stdout, resp.Body)//avoid allocating memory

	if err1 != nil {
		fmt.Printf("fetch error: %v\n" , err1)
	}

	fmt.Println(resp.StatusCode)

}

func dul(fileNames []string)  {

	counts := make(map[string]int)
	line2File := make(map[string]string)
	for _, fileName:= range(fileNames) {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("error while opening file: " + fileName + err.Error())
			continue
		}

		strs := strings.Split(string(data), "\n")

		for _ , s := range(strs) {
			counts[s]++
			line2File[s] += (fileName + ",")
		}

	}

	for line, count := range(counts) {
		if count >= 2 {
			fmt.Printf("%s: %s %d", line2File[line], line, count)
		}
	}
}
