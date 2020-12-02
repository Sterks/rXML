package protocols44

import "encoding/xml"

// Export was generated 2020-10-30 10:55:07 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type ProtocolVPP struct {
	XMLName               xml.Name `xml:"export"`
	Text                  string   `xml:",chardata"`
	SchemaLocation        string   `xml:"schemaLocation,attr"`
	Xmlns                 string   `xml:"xmlns,attr"`
	Ns6                   string   `xml:"ns6,attr"`
	Ns5                   string   `xml:"ns5,attr"`
	Ns8                   string   `xml:"ns8,attr"`
	Ns7                   string   `xml:"ns7,attr"`
	Ns9                   string   `xml:"ns9,attr"`
	Ns11                  string   `xml:"ns11,attr"`
	Ns10                  string   `xml:"ns10,attr"`
	Ns2                   string   `xml:"ns2,attr"`
	Ns4                   string   `xml:"ns4,attr"`
	Ns3                   string   `xml:"ns3,attr"`
	Xsi                   string   `xml:"xsi,attr"`
	EpProtocolEZP1Extract struct {
		Text              string `xml:",chardata"`
		SchemeVersion     string `xml:"schemeVersion,attr"`
		ID                string `xml:"id"`
		ExternalId        string `xml:"externalId"`
		VersionNumber     string `xml:"versionNumber"`
		FoundationDocInfo struct {
			Text                        string `xml:",chardata"`
			FoundationDocNumber         string `xml:"foundationDocNumber"`
			FoundationDocNumberExternal string `xml:"foundationDocNumberExternal"`
		} `xml:"foundationDocInfo"`
		CommonInfo struct {
			Text              string `xml:",chardata"`
			PurchaseNumber    string `xml:"purchaseNumber"`
			DocNumber         string `xml:"docNumber"`
			PublishDTInEIS    string `xml:"publishDTInEIS"`
			Href              string `xml:"href"`
			DocNumberExternal string `xml:"docNumberExternal"`
			PublishDTInETP    string `xml:"publishDTInETP"`
			Place             string `xml:"place"`
			ProcedureDT       string `xml:"procedureDT"`
			SignDT            string `xml:"signDT"`
			HrefExternal      string `xml:"hrefExternal"`
		} `xml:"commonInfo"`
		ProtocolPublisherInfo struct {
			Text         string `xml:",chardata"`
			PublisherOrg struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				PostAddress     string `xml:"postAddress"`
				FactAddress     string `xml:"factAddress"`
				INN             string `xml:"INN"`
				KPP             string `xml:"KPP"`
			} `xml:"publisherOrg"`
			PublisherRole string `xml:"publisherRole"`
		} `xml:"protocolPublisherInfo"`
		CommissionName string `xml:"commissionName"`
		PrintFormInfo  struct {
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
		ProtocolInfo struct {
			Text            string `xml:",chardata"`
			BestPrice       string `xml:"bestPrice"`
			AbandonedReason struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
			} `xml:"abandonedReason"`
		} `xml:"protocolInfo"`
	} `xml:"epProtocolEZP1Extract"`
}

