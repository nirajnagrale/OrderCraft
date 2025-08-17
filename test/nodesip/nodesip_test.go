package nodesip_test

import (
	nodesip "ordercraft/pkg/Nodesip"
	"testing"
)

func TestNodesIp(t *testing.T) {
	nodesIp := nodesip.NewNodesIp()
	nodesIp.AddNode("node1", "192.168.1.1")
	nodesIp.AddNode("node2", "192.168.1.2")

	if ip, ok := (*nodesIp)["node1"]; !ok || ip != "192.168.1.1" {
		t.Errorf("Expected 192.168.1.1 for node1, got %v", ip)
	}

	if ip, ok := (*nodesIp)["node2"]; !ok || ip != "192.168.1.2" {
		t.Errorf("Expected 192.168.1.2 for node2, got %v", ip)
	}
}
