package protocols223

import "encoding/xml"

// PurchaseProtocolRZ2AESMBO was generated 2020-11-10 17:44:02 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type PurchaseProtocolRZ2AESMBO struct {
	XMLName        xml.Name `xml:"purchaseProtocolRZ2AESMBO"`
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
			Text                          string `xml:",chardata"`
			Guid                          string `xml:"guid"`
			PurchaseProtocolRZ2AESMBOData struct {
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
						Fax           string `xml:"fax"`
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
						Fax           string `xml:"fax"`
						Email         string `xml:"email"`
						Okato         string `xml:"okato"`
						Okopf         string `xml:"okopf"`
						OkopfName     string `xml:"okopfName"`
						Okpo          string `xml:"okpo"`
					} `xml:"mainInfo"`
				} `xml:"customer"`
				MissedContest              string `xml:"missedContest"`
				PurchaseCancellationReason string `xml:"purchaseCancellationReason"`
				PublicationDateTime        string `xml:"publicationDateTime"`
				Status                     string `xml:"status"`
				Version                    string `xml:"version"`
				Attachments                struct {
					Text                string `xml:",chardata"`
					TotalDocumentsCount string `xml:"totalDocumentsCount"`
					Document            struct {
						Text           string `xml:",chardata"`
						Guid           string `xml:"guid"`
						CreateDateTime string `xml:"createDateTime"`
						FileName       string `xml:"fileName"`
						Description    string `xml:"description"`
						URL            string `xml:"url"`
					} `xml:"document"`
				} `xml:"attachments"`
				SignatureAuthorizedBody string `xml:"signatureAuthorizedBody"`
				Type                    string `xml:"type"`
				TypeName                string `xml:"typeName"`
				TargetPhaseCode         string `xml:"targetPhaseCode"`
				ProcedureDate           string `xml:"procedureDate"`
				ProcedurePlace          string `xml:"procedurePlace"`
				LotApplicationsList     struct {
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
							Accepted          string `xml:"accepted"`
							WinnerIndication  string `xml:"winnerIndication"`
						} `xml:"application"`
					} `xml:"protocolLotApplications"`
				} `xml:"lotApplicationsList"`
				ProtocolSignDate  string `xml:"protocolSignDate"`
				TemplateVersion   string `xml:"templateVersion"`
				TemplateStructure struct {
					Text  string `xml:",chardata"`
					Level []struct {
						Text     string `xml:",chardata"`
						TabLevel string `xml:"tabLevel"`
						Tab      []struct {
							Text       string `xml:",chardata"`
							TabOrdinal string `xml:"tabOrdinal"`
							TabName    string `xml:"tabName"`
							FixedName  string `xml:"fixedName"`
							Section    []struct {
								Text           string `xml:",chardata"`
								SectionOrdinal string `xml:"sectionOrdinal"`
								SectionName    string `xml:"sectionName"`
								FixedName      string `xml:"fixedName"`
							} `xml:"section"`
						} `xml:"tab"`
					} `xml:"level"`
					HiddenProtocolFields struct {
						Text string   `xml:",chardata"`
						Code []string `xml:"code"`
					} `xml:"hiddenProtocolFields"`
				} `xml:"templateStructure"`
				TemplateBlocks struct {
					Text                   string `xml:",chardata"`
					HideCommDecision       string `xml:"hideCommDecision"`
					HideCommDecisionAccess string `xml:"hideCommDecisionAccess"`
					HideCommDecisionResult string `xml:"hideCommDecisionResult"`
					HideProcedure          string `xml:"hideProcedure"`
					HideCancellation       string `xml:"hideCancellation"`
				} `xml:"templateBlocks"`
				ExtendFields struct {
					Text                string `xml:",chardata"`
					ProtocolExtendField struct {
						Text        string `xml:",chardata"`
						ExtendField []struct {
							Text        string `xml:",chardata"`
							IntegrCode  string `xml:"integrCode"`
							Description string `xml:"description"`
							Type        string `xml:"type"`
							Value       struct {
								Text     string `xml:",chardata"`
								DateTime string `xml:"dateTime"`
								URL      string `xml:"url"`
							} `xml:"value"`
						} `xml:"extendField"`
						Position struct {
							Text           string `xml:",chardata"`
							TabOrdinal     string `xml:"tabOrdinal"`
							TabName        string `xml:"tabName"`
							SectionOrdinal string `xml:"sectionOrdinal"`
							SectionName    string `xml:"sectionName"`
						} `xml:"position"`
					} `xml:"protocolExtendField"`
				} `xml:"extendFields"`
			} `xml:"purchaseProtocolRZ2AESMBOData"`
		} `xml:"item"`
	} `xml:"body"`
}

