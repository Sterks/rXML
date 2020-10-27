package structs

import "encoding/xml"

// NotificationEA615 ...
type NotificationEA615 struct {
	XMLName               xml.Name `xml:"export"`
	Text                  string   `xml:",chardata"`
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
	Pprf615NotificationEF struct {
		Text          string `xml:",chardata"`
		SchemeVersion string `xml:"schemeVersion,attr"`
		ID            string `xml:"id"`
		VersionNumber string `xml:"versionNumber"`
		CommonInfo    struct {
			Text               string `xml:",chardata"`
			PurchaseNumber     string `xml:"purchaseNumber"`
			DocNumber          string `xml:"docNumber"`
			PurchaseObjectInfo string `xml:"purchaseObjectInfo"`
			DocPublishDate     string `xml:"docPublishDate"`
			Href               string `xml:"href"`
		} `xml:"commonInfo"`
		PlacingWayInfo struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"placingWayInfo"`
		PrintFormInfo struct {
			Text      string `xml:",chardata"`
			URL       string `xml:"url"`
			Signature struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"signature"`
		} `xml:"printFormInfo"`
		AttachmentsInfo struct {
			Text           string `xml:",chardata"`
			AttachmentInfo []struct {
				Text               string `xml:",chardata"`
				PublishedContentId string `xml:"publishedContentId"`
				FileName           string `xml:"fileName"`
				FileSize           string `xml:"fileSize"`
				DocDescription     string `xml:"docDescription"`
				URL                string `xml:"url"`
				CryptoSigns        struct {
					Text      string `xml:",chardata"`
					Signature struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"signature"`
				} `xml:"cryptoSigns"`
			} `xml:"attachmentInfo"`
		} `xml:"attachmentsInfo"`
		ModificationInfo struct {
			Text       string `xml:",chardata"`
			Info       string `xml:"info"`
			ReasonInfo struct {
				Text                    string `xml:",chardata"`
				ResponsibleDecisionInfo struct {
					Text         string `xml:",chardata"`
					DecisionDate string `xml:"decisionDate"`
				} `xml:"responsibleDecisionInfo"`
			} `xml:"reasonInfo"`
		} `xml:"modificationInfo"`
		PurchaseResponsibleInfo struct {
			Text               string `xml:",chardata"`
			PublisherRole      string `xml:"publisherRole"`
			ResponsibleOrgInfo struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				PostAddress     string `xml:"postAddress"`
				INN             string `xml:"INN"`
				KPP             string `xml:"KPP"`
				Region          struct {
					Text      string `xml:",chardata"`
					KladrCode string `xml:"kladrCode"`
					FullName  string `xml:"fullName"`
				} `xml:"region"`
			} `xml:"responsibleOrgInfo"`
			ResponsibleInfo struct {
				Text          string `xml:",chardata"`
				ContactPerson struct {
					Text       string `xml:",chardata"`
					LastName   string `xml:"lastName"`
					FirstName  string `xml:"firstName"`
					MiddleName string `xml:"middleName"`
				} `xml:"contactPerson"`
				ContactEMail string `xml:"contactEMail"`
				ContactPhone string `xml:"contactPhone"`
			} `xml:"responsibleInfo"`
		} `xml:"purchaseResponsibleInfo"`
		NotificationInfo struct {
			Text    string `xml:",chardata"`
			ETPInfo struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
				URL  string `xml:"url"`
			} `xml:"ETPInfo"`
			ProcedureInfo struct {
				Text              string `xml:",chardata"`
				CollectingEndDate string `xml:"collectingEndDate"`
				ScoringDate       string `xml:"scoringDate"`
				BiddingDate       string `xml:"biddingDate"`
			} `xml:"procedureInfo"`
			ContractCondition struct {
				Text         string `xml:",chardata"`
				MaxPriceInfo struct {
					Text                 string `xml:",chardata"`
					MaxPrice             string `xml:"maxPrice"`
					ApplicationGuarantee struct {
						Text    string `xml:",chardata"`
						Amount  string `xml:"amount"`
						Account struct {
							Text              string `xml:",chardata"`
							Bik               string `xml:"bik"`
							SettlementAccount string `xml:"settlementAccount"`
							PersonalAccount   string `xml:"personalAccount"`
						} `xml:"account"`
					} `xml:"applicationGuarantee"`
					ContractGuarantee struct {
						Text    string `xml:",chardata"`
						Amount  string `xml:"amount"`
						Account struct {
							Text              string `xml:",chardata"`
							Bik               string `xml:"bik"`
							SettlementAccount string `xml:"settlementAccount"`
							PersonalAccount   string `xml:"personalAccount"`
						} `xml:"account"`
					} `xml:"contractGuarantee"`
				} `xml:"maxPriceInfo"`
				KladrPlacesInfo struct {
					Text       string `xml:",chardata"`
					KladrPlace struct {
						Text  string `xml:",chardata"`
						Kladr struct {
							Text      string `xml:",chardata"`
							KladrCode string `xml:"kladrCode"`
							FullName  string `xml:"fullName"`
						} `xml:"kladr"`
						DeliveryPlace string `xml:"deliveryPlace"`
					} `xml:"kladrPlace"`
				} `xml:"kladrPlacesInfo"`
				DeliveryTerm       string `xml:"deliveryTerm"`
				DeliveryConditions string `xml:"deliveryConditions"`
			} `xml:"contractCondition"`
			PurchaseObjectsInfo struct {
				Text                     string `xml:",chardata"`
				ServicesWorksKindNPAInfo struct {
					Text                              string `xml:",chardata"`
					ServicesWorksKindSubjectRFNPAInfo struct {
						Text                        string `xml:",chardata"`
						ServiceWorkKindSubjectRFNPA string `xml:"serviceWorkKindSubjectRFNPA"`
					} `xml:"servicesWorksKindSubjectRFNPAInfo"`
				} `xml:"servicesWorksKindNPAInfo"`
			} `xml:"purchaseObjectsInfo"`
			PurchaseSubjectInfo struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
			} `xml:"purchaseSubjectInfo"`
		} `xml:"notificationInfo"`
	} `xml:"pprf615NotificationEF"`
}
