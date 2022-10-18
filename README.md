[![Go Reference](https://pkg.go.dev/badge/github.com/UserJY/GoGraph#section-documentation.svg)](https://pkg.go.dev/github.com/UserJY/GoGraph#section-documentation)

# GoGraph

* graph data structure that is compatible with javascript's SigmaJS and Graphology
* Adding/Removing nodes and edges 
* Import/Export graph as JSON

# How to Use

```bash
go get -u github.com/UserJY/GoGraph
```

```bash
package main

import "github.com/UserJY/GoGraph" 

func main(){
	somegraph := gograph.MakeRandomGraph()
	
	somegraph.Print()
  	somegraph.ExportJSON("out.json")
}
```
