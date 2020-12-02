package protocols223

import "encoding/xml"

// PurchaseProtocolOSZ was generated 2020-11-10 17:12:25 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type PurchaseProtocolOSZ struct {
	XMLName        xml.Name `xml:"purchaseProtocolOSZ"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Ns5            string   `xml:"ns5,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Ns3            string   `xml:"ns3,attr"`
	Ns4            string   `xml:"ns4,attr"`
	Header         struct {
		Text           string `xml:",chardata"`
		Guid           string `xml:"guid"`
		CreateDateTime string `xml:"createDateTime"`
	} `xml:"header"`
	Body struct {
		Text string `xml:",chardata"`
		Item struct {
			Text                    string `xml:",chardata"`
			Guid                    string `xml:"guid"`
			PurchaseProtocolOSZData struct {
				Text           string `xml:",chardata"`
				Guid           string `xml:"guid"`
				CreateDateTime string `xml:"createDateTime"`
				PurchaseInfo   struct {
					Text                 string `xml:",chardata"`
					PurchaseNoticeNumber string `xml:"purchaseNoticeNumber"`
					Name                 string `xml:"name"`
					PurchaseMethodCode   string `xml:"purchaseMethodCode"`
					PurchaseCodeName     string `xml:"purchaseCodeName"`
				} `xml:"purchaseInfo"`
				RegistrationNumber string `xml:"registrationNumber"`
				Placer             struct {
					Text     string `xml:",chardata"`
					MainInfo struct {
						Text          string `xml:",chardata"`
						FullName      string `xml:"fullName"`
						ShortName     string `xml:"shortName"`
						Inn           string `xml:"inn"`
						Kpp           string `xml:"kpp"`
						Ogrn          string `xml:"ogrn"`
						LegalAddress  string `xml:"legalAddress"`
						PostalAddress string `xml:"postalAddress"`
					} `xml:"mainInfo"`
				} `xml:"placer"`
				MissedContest       string `xml:"missedContest"`
				PublicationDateTime string `xml:"publicationDateTime"`
				Status              string `xml:"status"`
				Version             string `xml:"version"`
				Attachements        struct {
					Text     string `xml:",chardata"`
					Document struct {
						Text           string `xml:",chardata"`
						CreateDateTime string `xml:"createDateTime"`
						FileName       string `xml:"fileName"`
						Description    string `xml:"description"`
						URL            string `xml:"url"`
					} `xml:"document"`
				} `xml:"attachements"`
				ProcedureDate                string `xml:"procedureDate"`
				ProcedurePlace               string `xml:"procedurePlace"`
				CommissionNumber             string `xml:"commissionNumber"`
				CommissionName               string `xml:"commissionName"`
				CommissionResult             string `xml:"commissionResult"`
				ProtocolRZRegistrationNumber string `xml:"protocolRZRegistrationNumber"`
				ProtocolRZVersion            string `xml:"protocolRZVersion"`
				LotApplicationsList          struct {
					Text                    string `xml:",chardata"`
					ProtocolLotApplications struct {
						Text          string `xml:",chardata"`
						Guid          string `xml:"guid"`
						OrdinalNumber string `xml:"ordinalNumber"`
						Subject       string `xml:"subject"`
						Information   string `xml:"information"`
						Application   []struct {
							Text              string `xml:",chardata"`
							ApplicationNumber string `xml:"applicationNumber"`
							SupplierInfo      struct {
								Text string `xml:",chardata"`
								Name string `xml:"name"`
								Inn  string `xml:"inn"`
								Kpp  string `xml:"kpp"`
								Ogrn string `xml:"ogrn"`
								Type string `xml:"type"`
							} `xml:"supplierInfo"`
							ApplicationResult string `xml:"applicationResult"`
							ApplicationRate   string `xml:"applicationRate"`
							ApplicationPlace  string `xml:"applicationPlace"`
						} `xml:"application"`
					} `xml:"protocolLotApplications"`
				} `xml:"lotApplicationsList"`
			} `xml:"purchaseProtocolOSZData"`
		} `xml:"item"`
	} `xml:"body"`
}

