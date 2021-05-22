package client

import "github.com/czerwonk/bird_exporter/protocol"

// Client retrieves information from Bird routing daemon
type Client interface {

	// GetProtocols retrieves protocol information and statistics from bird
	GetProtocols() ([]*protocol.Protocol, error)

	// Get Adhoc checks for the EdgeIX Implementation
	GetAdhoc(protocol *protocol.Protocol, community *protocol.LargeCommunity) ([]*protocol.Adhoc, error)

	// GetOSPFAreas retrieves OSPF specific information from bird
	GetOSPFAreas(protocol *protocol.Protocol) ([]*protocol.OspfArea, error)
}
