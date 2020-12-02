package protocols44

import "encoding/xml"

// Export was generated 2020-10-30 10:47:51 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type ProtocolEF2 struct {
	XMLName        xml.Name `xml:"export"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Ns6            string   `xml:"ns6,attr"`
	Ns5            string   `xml:"ns5,attr"`
	Ns8            string   `xml:"ns8,attr"`
	Ns7            string   `xml:"ns7,attr"`
	Ns9            string   `xml:"ns9,attr"`
	Ns11           string   `xml:"ns11,attr"`
	Ns10           string   `xml:"ns10,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Ns4            string   `xml:"ns4,attr"`
	Ns3            string   `xml:"ns3,attr"`
	FcsProtocolEF2 struct {
		Text                     string `xml:",chardata"`
		SchemeVersion            string `xml:"schemeVersion,attr"`
		ID                       string `xml:"id"`
		ExternalId               string `xml:"externalId"`
		PurchaseNumber           string `xml:"purchaseNumber"`
		ProtocolNumber           string `xml:"protocolNumber"`
		FoundationProtocolNumber string `xml:"foundationProtocolNumber"`
		Place                    string `xml:"place"`
		ProtocolDate             string `xml:"protocolDate"`
		SignDate                 string `xml:"signDate"`
		PublishDate              string `xml:"publishDate"`
		Href                     string `xml:"href"`
		PrintForm                struct {
			Text string `xml:",chardata"`
			URL  string `xml:"url"`
		} `xml:"printForm"`
		ProtocolPublisher struct {
			Text         string `xml:",chardata"`
			PublisherOrg struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
			} `xml:"publisherOrg"`
			PublisherRole string `xml:"publisherRole"`
		} `xml:"protocolPublisher"`
		Attachments struct {
			Text       string `xml:",chardata"`
			Attachment struct {
				Text           string `xml:",chardata"`
				FileName       string `xml:"fileName"`
				FileSize       string `xml:"fileSize"`
				DocDescription string `xml:"docDescription"`
				URL            string `xml:"url"`
			} `xml:"attachment"`
		} `xml:"attachments"`
		ProtocolLot struct {
			Text         string `xml:",chardata"`
			Applications struct {
				Text        string `xml:",chardata"`
				Application []struct {
					Text          string `xml:",chardata"`
					JournalNumber string `xml:"journalNumber"`
					PriceOffers   struct {
						Text       string `xml:",chardata"`
						FirstOffer struct {
							Text  string `xml:",chardata"`
							Price string `xml:"price"`
							Date  string `xml:"date"`
						} `xml:"firstOffer"`
						LastOffer struct {
							Text  string `xml:",chardata"`
							Price string `xml:"price"`
							Date  string `xml:"date"`
						} `xml:"lastOffer"`
						OffersQuantity string `xml:"offersQuantity"`
					} `xml:"priceOffers"`
				} `xml:"application"`
			} `xml:"applications"`
		} `xml:"protocolLot"`
	} `xml:"fcsProtocolEF2"`
}

