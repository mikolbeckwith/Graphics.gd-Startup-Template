package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
)

type Main struct {
	Node.Extension[Main] `gd:"MainScene"`
	Label                struct {
		Label.Instance
	}
}

func main() {
	classdb.Register[Main]()
	startup.Scene()
}
