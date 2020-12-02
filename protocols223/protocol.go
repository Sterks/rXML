package protocols223

import "encoding/xml"

// PurchaseProtocol was generated 2020-11-10 16:16:41 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type PurchaseProtocol struct {
	XMLName        xml.Name `xml:"purchaseProtocol"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Ns4            string   `xml:"ns4,attr"`
	Ns3            string   `xml:"ns3,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Ns6            string   `xml:"ns6,attr"`
	Ns5            string   `xml:"ns5,attr"`
	Ns8            string   `xml:"ns8,attr"`
	Ns7            string   `xml:"ns7,attr"`
	Ns13           string   `xml:"ns13,attr"`
	Ns9            string   `xml:"ns9,attr"`
	Ns12           string   `xml:"ns12,attr"`
	Ns11           string   `xml:"ns11,attr"`
	Ns10           string   `xml:"ns10,attr"`
	Ns17           string   `xml:"ns17,attr"`
	Ns16           string   `xml:"ns16,attr"`
	Ns15           string   `xml:"ns15,attr"`
	Ns14           string   `xml:"ns14,attr"`
	Ns19           string   `xml:"ns19,attr"`
	Ns18           string   `xml:"ns18,attr"`
	Ns20           string   `xml:"ns20,attr"`
	Ns24           string   `xml:"ns24,attr"`
	Ns23           string   `xml:"ns23,attr"`
	Ns22           string   `xml:"ns22,attr"`
	Ns21           string   `xml:"ns21,attr"`
	Ns25           string   `xml:"ns25,attr"`
	Header         struct {
		Text           string `xml:",chardata"`
		Guid           string `xml:"guid"`
		CreateDateTime string `xml:"createDateTime"`
	} `xml:"header"`
	Body struct {
		Text string `xml:",chardata"`
		Item struct {
			Text                 string `xml:",chardata"`
			Guid                 string `xml:"guid"`
			PurchaseProtocolData struct {
				Text           string `xml:",chardata"`
				Guid           string `xml:"guid"`
				CreateDateTime string `xml:"createDateTime"`
				UrlEIS         string `xml:"urlEIS"`
				PurchaseInfo   struct {
					Text                 string `xml:",chardata"`
					PurchaseNoticeNumber string `xml:"purchaseNoticeNumber"`
					Name                 string `xml:"name"`
					PurchaseMethodCode   string `xml:"purchaseMethodCode"`
					PurchaseCodeName     string `xml:"purchaseCodeName"`
					Emergency            string `xml:"emergency"`
				} `xml:"purchaseInfo"`
				RegistrationNumber string `xml:"registrationNumber"`
				Placer             struct {
					Text     string `xml:",chardata"`
					MainInfo struct {
						Text          string `xml:",chardata"`
						FullName      string `xml:"fullName"`
						ShortName     string `xml:"shortName"`
						Iko           string `xml:"iko"`
						Inn           string `xml:"inn"`
						Kpp           string `xml:"kpp"`
						Ogrn          string `xml:"ogrn"`
						LegalAddress  string `xml:"legalAddress"`
						PostalAddress string `xml:"postalAddress"`
						Phone         string `xml:"phone"`
						Email         string `xml:"email"`
						Okato         string `xml:"okato"`
						Okopf         string `xml:"okopf"`
						OkopfName     string `xml:"okopfName"`
						Okpo          string `xml:"okpo"`
					} `xml:"mainInfo"`
				} `xml:"placer"`
				Customer struct {
					Text     string `xml:",chardata"`
					MainInfo struct {
						Text          string `xml:",chardata"`
						FullName      string `xml:"fullName"`
						ShortName     string `xml:"shortName"`
						Iko           string `xml:"iko"`
						Inn           string `xml:"inn"`
						Kpp           string `xml:"kpp"`
						Ogrn          string `xml:"ogrn"`
						LegalAddress  string `xml:"legalAddress"`
						PostalAddress string `xml:"postalAddress"`
						Phone         string `xml:"phone"`
						Email         string `xml:"email"`
						Okato         string `xml:"okato"`
						Okopf         string `xml:"okopf"`
						OkopfName     string `xml:"okopfName"`
						Okpo          string `xml:"okpo"`
					} `xml:"mainInfo"`
				} `xml:"customer"`
				MissedContest              string `xml:"missedContest"`
				MissedReason               string `xml:"missedReason"`
				PurchaseCancellationReason string `xml:"purchaseCancellationReason"`
				PublicationDateTime        string `xml:"publicationDateTime"`
				Status                     string `xml:"status"`
				Version                    string `xml:"version"`
				Attachments                struct {
					Text                string `xml:",chardata"`
					TotalDocumentsCount string `xml:"totalDocumentsCount"`
					Document            []struct {
						Text           string `xml:",chardata"`
						Guid           string `xml:"guid"`
						CreateDateTime string `xml:"createDateTime"`
						FileName       string `xml:"fileName"`
						Description    string `xml:"description"`
						URL            string `xml:"url"`
					} `xml:"document"`
				} `xml:"attachments"`
				Type                string `xml:"type"`
				TypeName            string `xml:"typeName"`
				ProcedureDate       string `xml:"procedureDate"`
				ProcedurePlace      string `xml:"procedurePlace"`
				LotApplicationsList struct {
					Text                    string `xml:",chardata"`
					ProtocolLotApplications struct {
						Text string `xml:",chardata"`
						Lot  struct {
							Text          string `xml:",chardata"`
							Guid          string `xml:"guid"`
							OrdinalNumber string `xml:"ordinalNumber"`
							Subject       string `xml:"subject"`
							Currency      struct {
								Text        string `xml:",chardata"`
								Code        string `xml:"code"`
								DigitalCode string `xml:"digitalCode"`
								Name        string `xml:"name"`
							} `xml:"currency"`
							InitialSum string `xml:"initialSum"`
						} `xml:"lot"`
						Application struct {
							Text              string `xml:",chardata"`
							ApplicationDate   string `xml:"applicationDate"`
							ApplicationNumber string `xml:"applicationNumber"`
							Price             string `xml:"price"`
							Currency          struct {
								Text        string `xml:",chardata"`
								Code        string `xml:"code"`
								DigitalCode string `xml:"digitalCode"`
								Name        string `xml:"name"`
							} `xml:"currency"`
							Accepted                string `xml:"accepted"`
							RejectionReasonCode     string `xml:"rejectionReasonCode"`
							RejectionReason         string `xml:"rejectionReason"`
							ConsiderationResultInfo struct {
								Text     string `xml:",chardata"`
								Accepted string `xml:"accepted"`
							} `xml:"considerationResultInfo"`
							ContractSigned          string `xml:"contractSigned"`
							CommissionDecisionPlace string `xml:"commissionDecisionPlace"`
						} `xml:"application"`
					} `xml:"protocolLotApplications"`
				} `xml:"lotApplicationsList"`
				ProtocolSignDate string `xml:"protocolSignDate"`
				IsLotOriented    string `xml:"isLotOriented"`
			} `xml:"purchaseProtocolData"`
		} `xml:"item"`
	} `xml:"body"`
}

