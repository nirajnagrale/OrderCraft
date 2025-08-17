package nodesip

type NodesIp map[string]string

func NewNodesIp() *NodesIp {
	return &NodesIp{}
}

func (n *NodesIp) AddNode(id string, ip string) {
	(*n)[id] = ip
}
