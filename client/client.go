package client

import "github.com/czerwonk/bird_exporter/protocol"
import "github.com/czerwonk/bird_exporter/edgeix"

// Client retrieves information from Bird routing daemon
type Client interface {

	// GetProtocols retrieves protocol information and statistics from bird
	GetProtocols() ([]*protocol.Protocol, error)

	// EdgeIX Testing - GetRPKI
	GetEdgeIX() ([]*edgeix.Table, error)

	// GetOSPFAreas retrieves OSPF specific information from bird
	GetOSPFAreas(protocol *protocol.Protocol) ([]*protocol.OspfArea, error)
}
