package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var template = `

<html>

<head>
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
      <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.6.0/Chart.bundle.min.js"></script>
  <style type="text/css">

  </style>

  <title></title>

<script type="text/javascript">
window.onload=function(){
var ctx = document.getElementById("myChart").getContext('2d');
var myChart = new Chart(ctx, {
  type: 'scatter',
  data: {
    datasets: [{
      label: 'power',
      data: %s
    }]
  },
  options: {
		scales: {
			xAxes: [{
				type: 'logarithmic',
				position: 'bottom'
			}],
			yAxes: [{
				type: 'logarithmic',
				position: 'left'
			}]
		}
  }
});

}

</script>
</head>

<body>
  <canvas id="myChart" width="400" height="400"></canvas>
  <script>
</script>
</body>
`

type pair struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func Pairs_json(data []pair) string {
	data_json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	data_string := fmt.Sprintf("%s", data_json)
	return data_string
}

func Power_json(data []power) string {
	data_json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	data_string := fmt.Sprintf("%s", data_json)
	return data_string
}

func OutputHtml(path string, json_data string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(f, "test data")
	fmt.Fprintf(f, template, json_data)

}
