package main

import "testing"

func TestGraph(t *testing.T) {
	data := make([]pair, 0)
	data = append(data, pair{1, 2})
	data = append(data, pair{3.3, 4.66666})
	jsonString := PairsJSON(data)
	OutputHTML("test", "test.html", jsonString)
}
