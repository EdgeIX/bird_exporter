package protocol


type Adhoc struct {
	Name	string
	Matched	int
	Total	int
	Community *LargeCommunity
}

type LargeCommunity struct {
	Name	string
	ASN	int
	First	int
	Last	int
}

var EdgeIX = 24224

func GetLargeCommunities() []LargeCommunity {
	LargeCommunities := []LargeCommunity{
                LargeCommunity{
                        Name: "RPKI Invalids",
                        ASN: EdgeIX,
                        First:  1101,
                        Last: 13,
                },
		LargeCommunity{
			Name: "IRRDB Filtered Strict",
			ASN: EdgeIX,
			First: 1001,
			Last: 1001,
		},
        }
	return LargeCommunities
}
