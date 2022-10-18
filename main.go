package gograph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"log"
	"strconv"
	
)

type Node struct {
	Key string  `json:"key"`
	Attributes NodeAttr `json:"attributes"`
}

type NodeAttr struct {
	X int `json:"x"`
	Y int  `json:"y"`
	Size uint `json:"size,omitempty"`
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
}

type Edge struct {
	Key string `json:"key,omitempty"` 
	Source string `json:"source"`
	Target string `json:"target"`
	Attributes EdgeAttr `json:"attributes,omitempty"`
}

type EdgeAttr struct {
	Size int `json:"size,omitempty"`
	Type string `json:"type,omitempty"`
}

type SigmaGraph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges,omitempty"`
}

func (mygraph *SigmaGraph) AddNode(myinput string, nX int, nY int, nSize uint){
	mygraph.Nodes = append(mygraph.Nodes, Node{
		Key: myinput,
		Attributes: NodeAttr{
			X: nX,
			Y: nY,
			Size: nSize,
			Label: myinput,				
		}})
	}
func (mygraph *SigmaGraph) AddEdge(src string, tgt string){
	mygraph.Edges = append(mygraph.Edges, Edge{
		Source: src,
		Target: tgt,
		})
	}

func (mygraph *SigmaGraph) AddRandomNode(myinput string){
	mygraph.Nodes = append(mygraph.Nodes, Node{
		Key: myinput,
		Attributes: NodeAttr{
			X: rand.Intn(100),
			Y: rand.Intn(100),
			Label: myinput,				
		}})
	}

func (mygraph *SigmaGraph) ExportJSON(filepath string){
	out, err := json.MarshalIndent(mygraph, "", "  ")
	if err != nil {
		log.Fatal("Could not write SigmaGraph to JSON file: ", err)
	}
	_ = ioutil.WriteFile(filepath, out, 0644)
}

func (mygraph *SigmaGraph) ImportJSON(filepath string){
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
    err = json.Unmarshal([]byte(content), mygraph)
	if err != nil{
		log.Fatal("Invalid JSON properties: ", err)
	}

} 

func (mygraph *SigmaGraph) BuildNodeHashmap() map[string]*Node{
	nodeHashmap := make(map[string]*Node)
	for i := range mygraph.Nodes{
		nodeHashmap[mygraph.Nodes[i].Key] = &mygraph.Nodes[i]
	}
	return nodeHashmap
}

func ImportJSON(filepath string) SigmaGraph{
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
	var mygraph SigmaGraph
    err = json.Unmarshal([]byte(content), &mygraph)
	if err != nil{
		log.Fatal("Invalid JSON properties: ", err)
	}
	return mygraph
} 


func (mygraph *SigmaGraph) Print(){
	out, _ := json.MarshalIndent(mygraph, "", "  ")
	fmt.Println(string(out))
}

func (mygraph *SigmaGraph) ToAOB() []byte{
	out, _ := json.MarshalIndent(mygraph, "", "  ")
	return out
}

func (mygraph *SigmaGraph) ToString() string{
	out, _ := json.MarshalIndent(mygraph, "", "  ")
	return string(out)
}

func (node *Node) Print(){
	out, _ := json.MarshalIndent(node, "", "  ")
	fmt.Println(string(out))
}
func (node *Node) ToAOB() []byte{
	out, _ := json.MarshalIndent(node, "", "  ")
	return out
}
func (node *Node) ToString() string{
	out, _ := json.MarshalIndent(node, "", "  ")
	return string(out)
}

func (edge *Edge) Print(){
	out, _ := json.MarshalIndent(edge, "", "  ")
	fmt.Println(string(out))
}
func (edge *Edge) ToAOB() []byte{
	out, _ := json.MarshalIndent(edge, "", "  ")
	return out
}
func (edge *Edge) ToString() string{
	out, _ := json.MarshalIndent(edge, "", "  ")
	return string(out)
}

func MakeRandomGraph() SigmaGraph{
	var mygraph SigmaGraph
	for i := 1; i <= 50; i++ {
		mygraph.AddRandomNode("SampleNode"+strconv.Itoa(i))
	} 
	for i := 1; i <= 20; i++ {
		mygraph.AddEdge(mygraph.Nodes[rand.Intn(49)].Key,mygraph.Nodes[rand.Intn(49)].Key)
	}
	return mygraph
}
