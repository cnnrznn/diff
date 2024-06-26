package diff

import "fmt"

func Diff(s1, s2 string) Sequence {
	queue := []*Node{{nil, 0, 0, Equal, ""}}
	visited := map[string]struct{}{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		i, j := node.i, node.j

		if _, ok := visited[makeKey(i, j)]; ok {
			continue
		}
		visited[makeKey(i, j)] = struct{}{}

		// Break condition; we've found the shortest path to match
		if i == len(s1) && j == len(s2) {
			return Sequence{
				Steps: makeSteps(node),
			}
		}

		// follow diagonals
		// While s1[i], s2[j] are the same, increment i, j and create a new (free) node
		endi, endj := i, j
		for endi < len(s1) && endj < len(s2) && s1[endi] == s2[endj] {
			endi++
			endj++
		}
		if endi > i {
			queue = append(queue, &Node{
				parent: node,
				i:      endi,
				j:      endj,
				action: Equal,
				seq:    s1[i:endi],
			})
		}

		// add left and right
		// finally, add nodes that cost an extra step
		if i < len(s1) {
			queue = append(queue, &Node{
				parent: node,
				i:      i + 1,
				j:      j,
				action: Remove,
				seq:    s1[i : i+1],
			})
		}
		if j < len(s2) {
			queue = append(queue, &Node{
				parent: node,
				i:      i,
				j:      j + 1,
				action: Add,
				seq:    s2[j : j+1],
			})
		}
	}

	// ERROR, sequence not found
	return Sequence{}
}

type Node struct {
	parent *Node
	i, j   int
	action Action
	seq    string
}

type Sequence struct {
	Steps []Step
}

type Step struct {
	Action Action
	Seq    string
}

type Action int

const (
	Add Action = iota
	Remove
	Equal
)

// Construct steps slice from node traversal
func makeSteps(n *Node) []Step {
	result := []Step{}

	for n != nil {
		result = append([]Step{{
			Action: n.action,
			Seq:    n.seq,
		}}, result...)

		n = n.parent
	}

	return cleanSteps(result)
}

// Clean the steps slice for the end user
// - Merge adjacent steps with the same action
// - Remove start step (noop)
func cleanSteps(steps []Step) []Step {
	result := []Step{}
	curr := steps[1]

	for i := 2; i < len(steps); i++ {
		switch {
		case steps[i].Action == curr.Action:
			curr.Seq += steps[i].Seq
		default:
			result = append(result, curr)
			curr = steps[i]
		}
	}
	result = append(result, curr)

	return result
}

func makeKey(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}
