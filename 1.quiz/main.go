package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	flagFliePath string
	flagTime     int
	flagRandom   bool
)

func init() {
	flag.StringVar(&flagFliePath, "f", "questions.csv", "questions file path")
	flag.IntVar(&flagTime, "t", 10, "time")
	flag.BoolVar(&flagRandom, "r", false, "quesetions index random or not")
	flag.Parse()
}

func main() {
	//fmt.Printf("time: %v\n", flagTime)
	//fmt.Printf("filePath: %v\n", flagFliePath)
    fmt.Printf("file: %q\t time: %ds\t random: %v\n", flagFliePath, flagTime, flagRandom)
	f, err := os.Open(flagFliePath)
	if err != nil {
		log.Fatal("can not open file: "+flagFliePath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+flagFliePath, err)
	}

	//fmt.Printf("csv : %v", records)

	var questions []string
	var answers []string
	responses := make([]string, 0, 20)

	for _, v := range records {
		questions = append(questions, v[0])
		answers = append(answers, v[1])
	}

	//fmt.Printf("questions : %v\n", questions)
	//fmt.Printf("answers : %v\n", answers)

	fmt.Println("PLEASE INPUT [ENTER] TO START")
	bufio.NewScanner(os.Stdout).Scan()

	timeup := time.After(time.Second * time.Duration(flagTime))

	if flagRandom {
		rand.Seed(time.Now().UnixNano())
	}
	randPool := rand.Perm(len(questions))

	response := make(chan string)
	go func() {
		for _, v := range randPool {

			fmt.Println(questions[v])
			reader := bufio.NewReader(os.Stdin)
			value, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal("read string err :", err)
			}

			response <- value
		}

        close(response)

	}()

	for {

		select {
		case <-timeup:
			fmt.Println("TIME UP")
			caculate(responses, answers, randPool)
			return
		case v, ok := <-response:
			if !ok {
				caculate(responses, answers, randPool)
				return
			}
			responses = append(responses, v)
		}
	}
}

func caculate(responses []string, answers []string, randPool []int) {
    //fmt.Printf("responses is %+q\n", responses)
	var counter int
	for i, v := range randPool {
        if len(responses) <= i{
            break
        }
        //fmt.Printf("answers is %+q\n", answers[v])
        //fmt.Printf("responses is %+q\n", responses[i])
		result := strings.EqualFold(answers[v], strings.TrimSuffix(responses[i], "\n"))
		if result {
			counter++
		}
	}

	fmt.Printf("correct number is %d", counter)
}
