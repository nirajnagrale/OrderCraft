package nodeRegistry

type nodeRegistry struct {
	registry map[string]string
}

func (node *nodeRegistry) AddRegistry(id string, address string) {
	node.registry[id] = address
}
