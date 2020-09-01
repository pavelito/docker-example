package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var counter int
var file *os.File

func main() {
	path, _ := os.Getwd()
	file, _ = os.OpenFile(fmt.Sprintf("%s/data/counter.txt", path), os.O_RDWR, os.ModePerm)
	defer file.Close()
	counter = getCounter()

	http.HandleFunc("/", CounterServer)
	http.ListenAndServe(":8080", nil)

}

func CounterServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %d!", counter)
	counter = counter + 1
	saveCounter(counter)
}

func getCounter() int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		log.Println(fmt.Sprintf("Original Counter - %d", i))
		return i
	}
	return 0
}

func saveCounter(value int) {
	log.Println(fmt.Sprintf("Saving Counter Value - %d", value))
	file.Seek(0, 0)
	_, err := file.WriteAt([]byte(fmt.Sprintf("%d\n", value)), 0)
	if err != nil {
		log.Fatalln(err)
	}
}
