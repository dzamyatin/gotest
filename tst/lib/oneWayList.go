package lib

type OneWayList struct {
	root *NodeOfOneWayList
}

func (o *OneWayList) AddVal(val int) {
	o.Add(&NodeOfOneWayList{value: val})
}

func (o *OneWayList) Add(node *NodeOfOneWayList) {
	if o.root == nil {
		o.root = node
		return
	}

	curr := o.root

	for {
		if curr.Link() == nil {
			curr.Add(node)
			break
		}

		curr = curr.Link()
	}
}

func (o *OneWayList) ToSlice() []*NodeOfOneWayList {
	if o.root == nil {
		return nil
	}
	return o.root.ToSlice()
}

func (o *OneWayList) Flip() {
	if o.root == nil {
		return
	}

	s := o.root.ToSlice()

	var prev *NodeOfOneWayList
	for i := len(s) - 1; i >= 0; i-- {
		if prev != nil {
			prev.Add(s[i])
		}
		if prev == nil {
			o.root = s[i]
		}
		prev = s[i]
	}
	prev.Add(nil)
}

type NodeOfOneWayList struct {
	value int
	link  *NodeOfOneWayList
}

func (n *NodeOfOneWayList) Add(add *NodeOfOneWayList) {
	n.link = add
}

func (n *NodeOfOneWayList) Value() int {
	return n.value
}

func (n *NodeOfOneWayList) Link() *NodeOfOneWayList {
	return n.link
}

func (n *NodeOfOneWayList) ToSlice() []*NodeOfOneWayList {
	circuit := make(map[*NodeOfOneWayList]struct{})
	var res []*NodeOfOneWayList
	curr := n
	for {
		if _, ok := circuit[curr]; ok {
			break
		}

		res = append(res, curr)
		circuit[curr] = struct{}{}

		if curr.Link() == nil {
			break
		}

		curr = curr.Link()
	}

	return res
}
