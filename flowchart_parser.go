package visibleIdeas

import (
	"encoding/json"
	"fmt"

	"github.com/alimoeeny/ideas"
)

type Identifier struct {
	ID     string `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
}

type NodeType string

const (
	ContainerNode      = NodeType("ContainerNode")
	DDxNode            = NodeType("DDxNode")
	FindingComputeNode = NodeType("FindingComputeNode")
	LogicNode          = NodeType("LogicNode")
	SensorNode         = NodeType("SensorNode")
	WeightNode         = NodeType("WeightNode")
)

type ComputeNodeFunctionality string

const (
	Reported = ComputeNodeFunctionality("reported")
)

type ComputeNodeFunctionalityParam struct {
	Value interface{} `json:"value,omitempty"`
	Units ideas.Unit  `json:"units,omitempty"`
}

type NodePosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type NodeData struct {
	Label     string  `json:"label,omitempty"`
	Short     string  `json:"short,omitempty"`
	Weight    float64 `json:"weight,omitempty"`
	Operation struct {
		Title         string                          `json:"title,omitempty"`
		Functionality ComputeNodeFunctionality        `json:"functionality,omitempty"`
		Params        []ComputeNodeFunctionalityParam `json:"params,omitempty"`
	} `json:"operation,omitempty"`
	Energized  string     `json:"energized,omitempty"`
	NodeID     string     `json:"node_id,omitempty"`
	Identifier Identifier `json:"identifier,omitempty"`
}

type Node struct {
	ID               string                 `json:"id,omitempty"`
	Type             NodeType               `json:"type,omitempty"`
	Position         NodePosition           `json:"position,omitempty"`
	ClassName        string                 `json:"className,omitempty"`
	Style            map[string]interface{} `json:"style"`
	Draggable        bool                   `json:"draggable,omitempty"`
	Data             NodeData               `json:"data,omitempty"`
	ParentNodeID     string                 `json:"parentNode,omitempty"`
	Extent           string                 `json:"extent,omitempty"`
	PositionAbsolute NodePosition           `json:"positionAbsolute,omitempty"`
	ZIndex           int                    `json:"z,omitempty"`
	Width            int                    `json:"width,omitempty"`
	Height           int                    `json:"height,omitempty"`
	IsPaternt        bool                   `json:"isPaternt,omitempty"`
}

type Edge struct {
	ID       string `json:"id,omitempty"`
	SourceID string `json:"source,omitempty"`
	TargetID string `json:"target,omitempty"`
}

type DiagnosisFlow struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Nodes []Node `json:"nodes,omitempty"`
	Edges []Edge `json:"edges,omitempty"`
}

func flowchartJsonStringToWorkflow(s string) (*DiagnosisFlow, error) {
	var fc DiagnosisFlow
	err := json.Unmarshal([]byte(s), &fc)
	if err != nil {
		return nil, fmt.Errorf("not a valid json string: %s", err.Error())
	}

	return &fc, err
}

// func flowchartToWorkflow(flowchart Flowchart) (*ideas.Workflow, error) {
// 	return nil, nil
// }
