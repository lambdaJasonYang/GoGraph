# GoGraph

* graph data structure that is compatible with javascript's SigmaJS and Graphology
* Adding/Removing nodes and edges 
* Import/Export graph as JSON

# How to Use

```bash
go get -u github.com/UserJY/GoSigma
```

```bash
package main

import "github.com/UserJY/GoSigma" 

func main(){
	somegraph := gosigma.MakeRandomGraph()
	somegraph.Print()
  somegraph.ExportJSON("out.json")
}
```
