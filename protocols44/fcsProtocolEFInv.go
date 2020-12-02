package protocols44

import "encoding/xml"

// Export was generated 2020-10-30 10:49:02 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type ProtocolEFInv struct {
	XMLName                   xml.Name `xml:"export"`
	Text                      string   `xml:",chardata"`
	Xmlns                     string   `xml:"xmlns,attr"`
	Ns6                       string   `xml:"ns6,attr"`
	Ns5                       string   `xml:"ns5,attr"`
	Ns8                       string   `xml:"ns8,attr"`
	Ns7                       string   `xml:"ns7,attr"`
	Ns9                       string   `xml:"ns9,attr"`
	Ns11                      string   `xml:"ns11,attr"`
	Ns10                      string   `xml:"ns10,attr"`
	Ns2                       string   `xml:"ns2,attr"`
	Ns4                       string   `xml:"ns4,attr"`
	Ns3                       string   `xml:"ns3,attr"`
	FcsProtocolEFInvalidation struct {
		Text           string `xml:",chardata"`
		SchemeVersion  string `xml:"schemeVersion,attr"`
		ID             string `xml:"id"`
		PurchaseNumber string `xml:"purchaseNumber"`
		ProtocolNumber string `xml:"protocolNumber"`
		Place          string `xml:"place"`
		ProtocolDate   string `xml:"protocolDate"`
		SignDate       string `xml:"signDate"`
		PublishDate    string `xml:"publishDate"`
		Href           string `xml:"href"`
		PrintForm      struct {
			Text string `xml:",chardata"`
			URL  string `xml:"url"`
		} `xml:"printForm"`
		Attachments struct {
			Text       string `xml:",chardata"`
			Attachment struct {
				Text               string `xml:",chardata"`
				PublishedContentId string `xml:"publishedContentId"`
				FileName           string `xml:"fileName"`
				FileSize           string `xml:"fileSize"`
				DocDescription     string `xml:"docDescription"`
				URL                string `xml:"url"`
			} `xml:"attachment"`
		} `xml:"attachments"`
		ProtocolLot struct {
			Text            string `xml:",chardata"`
			AbandonedReason struct {
				Text       string `xml:",chardata"`
				Code       string `xml:"code"`
				ObjectName string `xml:"objectName"`
				Name       string `xml:"name"`
				Type       string `xml:"type"`
			} `xml:"abandonedReason"`
		} `xml:"protocolLot"`
	} `xml:"fcsProtocolEFInvalidation"`
}

