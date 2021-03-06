package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/felixhudson/peakdection"
)

func main() {
	//printCSV(result)
	peaks, _ := peakdection.Findpeaks([]int{1, 5, 1, 5, 1, 5, 1, 9, 5, 1})
	fmt.Printf("peaks = %+v\n", peaks)
	files := GetAllFiles()
	// if we have any files to process check if the output directory exists
	if len(files) >= 1 {
		os.Mkdir("output", 0750)
	}
	for _, f := range files {
		processFile(f)
	}
}

func processFile(filename string) {
	fmt.Println("Processing file ", filename)
	data := readTcx(filepath.Join("tcx", filename))
	if len(data) == 0 {
		fmt.Println("Skipping file ", filename)
	} else {
		result := calculate(data)
		jsonData := PowerJSON(result)
		OutputHTML("output", "results "+filename+".html", jsonData)
	}
}

type power struct {
	Time  int     `json:"x"`
	Power float64 `json:"y"`
}

func readTcx(name string) []power {
	byteinput, _ := ioutil.ReadFile(name)
	//input := "<Trackpoint><Time>2017-08-06T07:22:21.000Z</Time><DistanceMeters>4.0</DistanceMeters><Cadence>21</Cadence><Extensions><ns2:TPX><ns2:Speed>1.337171</ns2:Speed><ns2:Watts>37</ns2:Watts></ns2:TPX></Extensions></Trackpoint>"
	return extractData(byteinput)
}

func extractData(byteinput []byte) []power {
	input := string(byteinput)
	var start, end int
	var result = make([]power, 0)
	//var val float64
	count := 0
	lines := strings.Split(input, "<Trackpoint>")
	if len(lines) <= 10 {
		fmt.Printf("Found %d lines\n", len(lines))
		fmt.Println(input)
		fmt.Println(byteinput)
	}
	for _, value := range lines {
		count++
		if len(value) == 0 {
			//fmt.Println("Read an empty line")
		} else {
			start = strings.Index(value, "<ns2:Watts>") + 11
			end = strings.Index(value, "</ns2:Watts>")
			if end-1 > 0 {
				val, err := strconv.ParseFloat(value[start:end], 64)
				if err == nil {
					result = append(result, power{count, val})
				}
			}
		}

	}
	return result

}

func xmlgrep(data string, searchElement string) []string {
	//output any element that starts and ends with the search element

	//input := "<Trackpoint><Time>2017-08-06T07:22:21.000Z</Time><DistanceMeters>4.0</DistanceMeters><Cadence>21</Cadence><Extensions><ns2:TPX><ns2:Speed>1.337171</ns2:Speed><ns2:Watts>37</ns2:Watts></ns2:TPX></Extensions></Trackpoint>"
	// we will use a basic io.Reader and read 100chars at a time.

	return []string{"temp"}
}

func sum(data []power) (float64, int) {
	total := 0.0
	count := 0
	for _, v := range data {
		total += v.Power
		count++
	}
	return total, count

}
func calculate(data []power) []power {
	result := make([]power, 0)
	if len(data) == 0 {
		panic("Error reading data")
	}
	start := 0
	end := len(data) - 1
	count := end
	sum, count := sum(data)
	var avg float64
	for start != end {
		avg = float64(sum) / float64(count)
		//fmt.Printf("Best over %d is %f\n",count, avg)
		//fmt.Printf("start %d and %d\n", start, end)
		//fmt.Println(count,sum)
		result = append(result, power{count, avg})
		if data[start].Power >= data[end].Power {
			sum = sum - data[end].Power
			// move the end
			end--
		} else {
			sum = sum - data[start].Power
			start++
		}
		count--
	}
	return result
}

func printCSV(data []power) {
	fmt.Println("Sec,Power(watts)")
	for _, v := range data {
		fmt.Println(v.Time, ",", v.Power)
	}
}
