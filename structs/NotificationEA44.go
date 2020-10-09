package structs

import "encoding/xml"

// Export was generated 2020-07-13 16:38:11 by SUPERVISOR-IT\SuperAdmin on Supervisor-IT.
type NotificationEA44 struct {
	XMLName           xml.Name `xml:"export" json:"export"`
	Text              string   `xml:",chardata" json:"chardata"`
	Xmlns             string   `xml:"xmlns,attr" json:"xmlns"`
	Ns6               string   `xml:"ns6,attr" json:"ns6"`
	Ns5               string   `xml:"ns5,attr" json:"ns5"`
	Ns8               string   `xml:"ns8,attr" json:"ns8"`
	Ns7               string   `xml:"ns7,attr" json:"ns7"`
	Ns9               string   `xml:"ns9,attr" json:"ns9"`
	Ns11              string   `xml:"ns11,attr" json:"ns11"`
	Ns10              string   `xml:"ns10,attr" json:"ns10"`
	Ns2               string   `xml:"ns2,attr" json:"ns2"`
	Ns4               string   `xml:"ns4,attr" json:"ns4"`
	Ns3               string   `xml:"ns3,attr"json:"ns3"`
	FcsNotificationEF struct {
		Text           string `xml:",chardata" json:"text"`
		SchemeVersion  string `xml:"schemeVersion,attr" json:"schemeVersion"`
		ID             string `xml:"id" json:"id"`
		ExternalId     string `xml:"externalId" json:"externalId"`
		PurchaseNumber string `xml:"purchaseNumber" json:"purchaseNumber"`
		DocPublishDate string `xml:"docPublishDate" json:"docPublishDate"`
		DocNumber      string `xml:"docNumber" json:"docNumber"`
		Href           string `xml:"href" json:"href"`
		PrintForm      struct {
			Text      string `xml:",chardata"`
			URL       string `xml:"url"`
			Signature struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"signature"`
		} `xml:"printForm"`
		PurchaseObjectInfo  string `xml:"purchaseObjectInfo"`
		PurchaseResponsible struct {
			Text           string `xml:",chardata"`
			ResponsibleOrg struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				PostAddress     string `xml:"postAddress"`
				FactAddress     string `xml:"factAddress"`
				INN             string `xml:"INN"`
				KPP             string `xml:"KPP"`
			} `xml:"responsibleOrg"`
			ResponsibleRole string `xml:"responsibleRole"`
			ResponsibleInfo struct {
				Text           string `xml:",chardata"`
				OrgPostAddress string `xml:"orgPostAddress"`
				OrgFactAddress string `xml:"orgFactAddress"`
				ContactPerson  struct {
					Text       string `xml:",chardata"`
					LastName   string `xml:"lastName"`
					FirstName  string `xml:"firstName"`
					MiddleName string `xml:"middleName"`
				} `xml:"contactPerson"`
				ContactEMail string `xml:"contactEMail"`
				ContactPhone string `xml:"contactPhone"`
				ContactFax   string `xml:"contactFax"`
			} `xml:"responsibleInfo"`
		} `xml:"purchaseResponsible"`
		PlacingWay struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"placingWay"`
		ContractConclusionOnSt83Ch2 string `xml:"contractConclusionOnSt83Ch2"`
		ETP                         struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"ETP"`
		ProcedureInfo struct {
			Text       string `xml:",chardata"`
			Collecting struct {
				Text      string `xml:",chardata"`
				StartDate string `xml:"startDate"`
				Place     string `xml:"place"`
				Order     string `xml:"order"`
				EndDate   string `xml:"endDate"`
			} `xml:"collecting"`
			Scoring struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"scoring"`
			Bidding struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"bidding"`
		} `xml:"procedureInfo"`
		Lot struct {
			Text     string `xml:",chardata"`
			MaxPrice string `xml:"maxPrice"`
			Currency struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
			} `xml:"currency"`
			QuantityUndefined    string `xml:"quantityUndefined"`
			CustomerRequirements struct {
				Text                string `xml:",chardata"`
				CustomerRequirement struct {
					Text     string `xml:",chardata"`
					Customer struct {
						Text            string `xml:",chardata"`
						RegNum          string `xml:"regNum"`
						ConsRegistryNum string `xml:"consRegistryNum"`
						FullName        string `xml:"fullName"`
					} `xml:"customer"`
					MaxPrice             string `xml:"maxPrice"`
					MustPublicDiscussion string `xml:"mustPublicDiscussion"`
					PublicDiscussionInfo struct {
						Text                     string `xml:",chardata"`
						PublicDiscussionNotInEIS string `xml:"publicDiscussionNotInEIS"`
					} `xml:"publicDiscussionInfo"`
					KladrPlaces struct {
						Text       string `xml:",chardata"`
						KladrPlace struct {
							Text          string `xml:",chardata"`
							DeliveryPlace string `xml:"deliveryPlace"`
						} `xml:"kladrPlace"`
					} `xml:"kladrPlaces"`
					DeliveryTerm      string `xml:"deliveryTerm"`
					ContractGuarantee struct {
						Text              string  `xml:",chardata"`
						Part              float32 `xml:"part"`
						ProcedureInfo     string  `xml:"procedureInfo"`
						SettlementAccount string  `xml:"settlementAccount"`
						PersonalAccount   string  `xml:"personalAccount"`
						Bik               string  `xml:"bik"`
						Amount            float64 `xml:"amount"`
					} `xml:"contractGuarantee"`
					PurchaseCode string `xml:"purchaseCode"`
					IKZInfo      struct {
						Text        string `xml:",chardata"`
						PublishYear string `xml:"publishYear"`
						OKPD2Info   struct {
							Text      string `xml:",chardata"`
							Undefined string `xml:"undefined"`
						} `xml:"OKPD2Info"`
						KVRInfo struct {
							Text string `xml:",chardata"`
							KVR  struct {
								Text string `xml:",chardata"`
								Code string `xml:"code"`
								Name string `xml:"name"`
							} `xml:"KVR"`
						} `xml:"KVRInfo"`
						CustomerCode        string `xml:"customerCode"`
						PurchaseNumber      string `xml:"purchaseNumber"`
						PurchaseOrderNumber string `xml:"purchaseOrderNumber"`
					} `xml:"IKZInfo"`
					TenderPlan2020Info struct {
						Text               string `xml:",chardata"`
						Plan2020Number     string `xml:"plan2020Number"`
						Position2020Number string `xml:"position2020Number"`
					} `xml:"tenderPlan2020Info"`
					ContractExecutionPaymentPlan struct {
						Text                 string `xml:",chardata"`
						FinancingSourcesInfo struct {
							Text        string `xml:",chardata"`
							CurrentYear string `xml:"currentYear"`
							FinanceInfo struct {
								Text        string `xml:",chardata"`
								Total       string `xml:"total"`
								CurrentYear string `xml:"currentYear"`
								FirstYear   string `xml:"firstYear"`
								SecondYear  string `xml:"secondYear"`
								SubsecYears string `xml:"subsecYears"`
							} `xml:"financeInfo"`
						} `xml:"financingSourcesInfo"`
					} `xml:"contractExecutionPaymentPlan"`
				} `xml:"customerRequirement"`
			} `xml:"customerRequirements"`
			PurchaseObjects struct {
				Text           string `xml:",chardata"`
				PurchaseObject struct {
					Text  string `xml:",chardata"`
					OKPD2 struct {
						Text string `xml:",chardata"`
						Code string `xml:"code"`
						Name string `xml:"name"`
					} `xml:"OKPD2"`
					Name string `xml:"name"`
					OKEI struct {
						Text         string `xml:",chardata"`
						Code         string `xml:"code"`
						NationalCode string `xml:"nationalCode"`
						FullName     string `xml:"fullName"`
					} `xml:"OKEI"`
					CustomerQuantities struct {
						Text             string `xml:",chardata"`
						CustomerQuantity struct {
							Text     string `xml:",chardata"`
							Customer struct {
								Text            string `xml:",chardata"`
								RegNum          string `xml:"regNum"`
								ConsRegistryNum string `xml:"consRegistryNum"`
								FullName        string `xml:"fullName"`
							} `xml:"customer"`
							Quantity string `xml:"quantity"`
						} `xml:"customerQuantity"`
					} `xml:"customerQuantities"`
					Price    string `xml:"price"`
					Quantity struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"quantity"`
					Sum string `xml:"sum"`
				} `xml:"purchaseObject"`
				TotalSum string `xml:"totalSum"`
			} `xml:"purchaseObjects"`
			Requirements struct {
				Text        string `xml:",chardata"`
				Requirement []struct {
					Text      string `xml:",chardata"`
					ShortName string `xml:"shortName"`
					Name      string `xml:"name"`
				} `xml:"requirement"`
			} `xml:"requirements"`
			Restrictions struct {
				Text        string `xml:",chardata"`
				Restriction struct {
					Text      string `xml:",chardata"`
					ShortName string `xml:"shortName"`
					Name      string `xml:"name"`
				} `xml:"restriction"`
			} `xml:"restrictions"`
		} `xml:"lot"`
		Attachments struct {
			Text       string `xml:",chardata"`
			Attachment []struct {
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
			} `xml:"attachment"`
		} `xml:"attachments"`
		Documentation struct {
			Text                   string `xml:",chardata"`
			PurchaseObjectsCh9St37 string `xml:"purchaseObjectsCh9St37"`
			Modifiable             string `xml:"modifiable"`
			ClarificationInfo      struct {
				Text                    string `xml:",chardata"`
				StartDate               string `xml:"startDate"`
				FilledManuallyStartDate string `xml:"filledManuallyStartDate"`
				EndDate                 string `xml:"endDate"`
				DeliveryProcedure       string `xml:"deliveryProcedure"`
			} `xml:"clarificationInfo"`
			OnesideRejectionCh9St95 string `xml:"onesideRejectionCh9St95"`
			PrintFormInfo           struct {
				Text      string `xml:",chardata"`
				URL       string `xml:"url"`
				Signature struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"signature"`
			} `xml:"printFormInfo"`
		} `xml:"documentation"`
	} `xml:"fcsNotificationEF"`
}
