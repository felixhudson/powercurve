package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
<div style="width:600px">
  <canvas id="myChart" width="400" height="400"></canvas>
</div>
	<script>
</script>
</body>
`

type pair struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// PairsJSON will take a pair of values and output as json
func PairsJSON(data []pair) string {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	dataString := fmt.Sprintf("%s", dataJSON)
	return dataString
}

// PowerJSON will reformat the data
func PowerJSON(data []power) string {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	dataString := fmt.Sprintf("%s", dataJSON)
	return dataString
}

// OutputHTML will write the output graph
func OutputHTML(dir string, path string, jsonData string) {
	f, err := os.Create(filepath.Join(dir, path))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(f, "test data")
	fmt.Fprintf(f, template, jsonData)

}
