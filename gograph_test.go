package gograph
import (
	"testing"
)

func TestA(t *testing.T){
	mygraph := MakeRandomGraph()
	t.Log(mygraph.ToString())
}

func TestB(t *testing.T){
	mygraph := MakeRandomGraph()
	nodeHashmap := mygraph.BuildNodeHashmap()
	t.Log(nodeHashmap["SampleNode1"].ToString())

}