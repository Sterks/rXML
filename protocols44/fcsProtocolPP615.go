package protocols44

import "encoding/xml"

// Export was generated 2020-10-30 10:53:50 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type ProtocolPP615 struct {
	XMLName            xml.Name `xml:"export"`
	Text               string   `xml:",chardata"`
	Xmlns              string   `xml:"xmlns,attr"`
	Ns6                string   `xml:"ns6,attr"`
	Ns5                string   `xml:"ns5,attr"`
	Ns8                string   `xml:"ns8,attr"`
	Ns7                string   `xml:"ns7,attr"`
	Ns9                string   `xml:"ns9,attr"`
	Ns11               string   `xml:"ns11,attr"`
	Ns10               string   `xml:"ns10,attr"`
	Ns2                string   `xml:"ns2,attr"`
	Ns4                string   `xml:"ns4,attr"`
	Ns3                string   `xml:"ns3,attr"`
	Pprf615ProtocolEF2 struct {
		Text          string `xml:",chardata"`
		SchemeVersion string `xml:"schemeVersion,attr"`
		ID            string `xml:"id"`
		ExternalId    string `xml:"externalId"`
		VersionNumber string `xml:"versionNumber"`
		CommonInfo    struct {
			Text              string `xml:",chardata"`
			PurchaseNumber    string `xml:"purchaseNumber"`
			DocNumber         string `xml:"docNumber"`
			DocPublishDate    string `xml:"docPublishDate"`
			Href              string `xml:"href"`
			DocNumberExternal string `xml:"docNumberExternal"`
			CreateDate        string `xml:"createDate"`
			SignDate          string `xml:"signDate"`
			HrefExternal      string `xml:"hrefExternal"`
		} `xml:"commonInfo"`
		PrintFormInfo struct {
			Text      string `xml:",chardata"`
			URL       string `xml:"url"`
			Signature []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"signature"`
		} `xml:"printFormInfo"`
		ExtPrintFormInfo struct {
			Text      string `xml:",chardata"`
			URL       string `xml:"url"`
			Signature string `xml:"signature"`
			FileType  string `xml:"fileType"`
		} `xml:"extPrintFormInfo"`
		ProtocolPublisherInfo struct {
			Text             string `xml:",chardata"`
			PublisherOrgInfo struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				Region          struct {
					Text      string `xml:",chardata"`
					KladrCode string `xml:"kladrCode"`
					FullName  string `xml:"fullName"`
				} `xml:"region"`
			} `xml:"publisherOrgInfo"`
			PublisherRole string `xml:"publisherRole"`
		} `xml:"protocolPublisherInfo"`
		ProtocolInfo struct {
			Text        string `xml:",chardata"`
			BiddingInfo struct {
				Text             string `xml:",chardata"`
				BiddingStartDate string `xml:"biddingStartDate"`
				BiddingEndDate   string `xml:"biddingEndDate"`
				MaxPrice         string `xml:"maxPrice"`
			} `xml:"biddingInfo"`
			Applications struct {
				Text        string `xml:",chardata"`
				Application struct {
					Text               string `xml:",chardata"`
					JournalNumber      string `xml:"journalNumber"`
					AppDate            string `xml:"appDate"`
					AppParticipantInfo struct {
						Text              string `xml:",chardata"`
						LegalEntityRFInfo struct {
							Text           string `xml:",chardata"`
							FullName       string `xml:"fullName"`
							INN            string `xml:"INN"`
							KPP            string `xml:"KPP"`
							OrgFactAddress string `xml:"orgFactAddress"`
							EMail          string `xml:"eMail"`
							Phone          string `xml:"phone"`
						} `xml:"legalEntityRFInfo"`
					} `xml:"appParticipantInfo"`
					LastOffer struct {
						Text  string `xml:",chardata"`
						Price string `xml:"price"`
						Date  string `xml:"date"`
					} `xml:"lastOffer"`
					AppRating string `xml:"appRating"`
				} `xml:"application"`
			} `xml:"applications"`
			AbandonedReason struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
			} `xml:"abandonedReason"`
		} `xml:"protocolInfo"`
	} `xml:"pprf615ProtocolEF2"`
}

