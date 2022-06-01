package visibleIdeas

type FlowchartNodeData struct {
	Label string
}

type FlowchartPosition struct {
	X float64
	Y float64
}

type FlowchartNode struct {
	ID       string
	Data     FlowchartNodeData
	Position FlowchartPosition
}

type FlowchartEdge struct {
	ID     string
	Source string
	Target string
	Label  string
}

type Flowchart struct {
	ID    string
	Nodes []FlowchartNode
	Edges []FlowchartEdge
}
