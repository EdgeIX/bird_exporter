package edgeix

type Table struct {
	MatchingRoutes	int
	TotalRoutes     int
	TableName	string
}

type LargeCommunity struct {
	ASN		int
	FirstThree	int
	LastThree	int
}

func NewLargeCommunity(asn int, firstThree int, lastThree int) *LargeCommunity {
	return &LargeCommunity{ASN: asn, FirstThree: firstThree, LastThree: lastThree}
}
