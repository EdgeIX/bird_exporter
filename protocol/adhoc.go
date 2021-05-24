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
	MasterOnly bool
}

var EdgeIX = 24224

func GetLargeCommunities() []LargeCommunity {
	LargeCommunities := []LargeCommunity{
                LargeCommunity{
                        Name: "RPKI Invalids",
                        ASN: EdgeIX,
                        First:  1101,
                        Last: 13,
			MasterOnly: false,
                },
                LargeCommunity{
                       Name: "RPKI Valid",
                       ASN: EdgeIX,
                       First: 1000,
                       Last: 1,
		       MasterOnly: true,
                },
                LargeCommunity{
                       Name: "RPKI Unknown",
                       ASN: EdgeIX,
                       First: 1000,
                       Last: 2,
		       MasterOnly: true,
                },
                LargeCommunity{
                       Name: "IRRDB Filtered Loose",
                       ASN: EdgeIX,
                       First: 1001,
                       Last: 1000,
		       MasterOnly: false,
                },
                LargeCommunity{
                       Name: "IRRDB Filtered Strict",
                       ASN: EdgeIX,
                       First: 1001,
                       Last: 1001,
		       MasterOnly: false,
                },
                LargeCommunity{
                       Name: "IRRDB Prefix Filtered",
                       ASN: EdgeIX,
                       First: 1101,
                       Last: 9,
		       MasterOnly: false,
                },
                LargeCommunity{
                       Name: "IRRDB Origin AS Filtered",
                       ASN: EdgeIX,
                       First: 1101,
                       Last: 10,
		       MasterOnly: false,
                },
                LargeCommunity{
                       Name: "Prefix not in Origin AS",
                       ASN: EdgeIX,
                       First: 1101,
                       Last: 11,
		       MasterOnly: false,
                },
        }
	return LargeCommunities
}

