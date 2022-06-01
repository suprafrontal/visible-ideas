package visibleIdeas

import "fmt"

// returns the index of a node with id id in the array of nodes nodes
func findFirstNodeIndexByID(nodes []Node, id string) int {
	for idx, node := range nodes {
		if node.ID == id {
			return idx
		}
	}
	return -1
}

// returns the index of a first node with identifier identifier in the array of nodes nodes
func findFirstNodeIndexByIdentifier(nodes []Node, identifier Identifier) int {
	for idx, node := range nodes {
		if node.Data.Identifier.ID == identifier.ID {
			return idx
		}
	}
	return -1
}

func sameConnectionAlreadyExists(edges []Edge, theEdge Edge) bool {
	for _, e := range edges {
		if e.SourceID == theEdge.SourceID && e.TargetID == theEdge.TargetID {
			return true
		}
	}
	return false
}

// MergeFlows Assumes the sensor nodes
func MergeFlows(newID string, newTitle string, flows ...DiagnosisFlow) (DiagnosisFlow, error) {
	r := DiagnosisFlow{
		ID:    newID,
		Title: newTitle,
		Nodes: []Node{},
		Edges: []Edge{},
	}

	// sanity checks
	// allNodeIDs := []string
	// a maps of node ids that points duplicate nodes to the nodes already in the resulting merged flow
	duplicateNodeIDMapping := map[string]string{}

	for _, f := range flows {
		// Do the nodes first -----------------------------------------------------------------
		for _, n := range f.Nodes {
			switch n.Type {
			case ContainerNode:
				// we only need only one set of container nodes, but for now let's ignore them

			case SensorNode, DDxNode:
				existingNodeID := findFirstNodeIndexByIdentifier(r.Nodes, n.Data.Identifier)
				if existingNodeID > -1 {
					duplicateNodeIDMapping[r.Nodes[existingNodeID].ID] = n.ID
					break
				}
				if findFirstNodeIndexByID(r.Nodes, n.ID) > -1 {
					return r, fmt.Errorf("two %s nodes with the exact same id but different identifiers detected which is unexpected: '%s'", n.Type, n.ID)
				}
				r.Nodes = append(r.Nodes, n)

			case FindingComputeNode, LogicNode, WeightNode:
				if findFirstNodeIndexByID(r.Nodes, n.ID) > -1 {
					fmt.Printf("two %s nodes with the exact same id detected which is must mean a real duplicate: '%s'", n.Type, n.ID)
					break
				}
				r.Nodes = append(r.Nodes, n)

			default:
				return r, fmt.Errorf("flow contains '%s' nodes which are not supported", n.Type)
			}
		}
		// Then do the edges -----------------------------------------------------------------
		for _, e := range f.Edges {
			if duplicateNodeIDMapping[e.SourceID] != "" {
				e.SourceID = duplicateNodeIDMapping[e.SourceID]
			}
			if duplicateNodeIDMapping[e.TargetID] != "" {
				e.TargetID = duplicateNodeIDMapping[e.TargetID]
			}
			if !sameConnectionAlreadyExists(r.Edges, e) {
				r.Edges = append(r.Edges, e)
			}
		}
	}
	return r, nil
}
