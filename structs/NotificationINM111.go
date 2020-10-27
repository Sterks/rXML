package structs

import "encoding/xml"

// NotificationNM111 ...
type NotificationNM111 struct {
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
	FcsNotification111 struct {
		Text           string `xml:",chardata"`
		SchemeVersion  string `xml:"schemeVersion,attr"`
		ID             string `xml:"id"`
		PurchaseNumber string `xml:"purchaseNumber"`
		DocPublishDate string `xml:"docPublishDate"`
		DocNumber      string `xml:"docNumber"`
		Href           string `xml:"href"`
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
		} `xml:"purchaseResponsible"`
		PlacingWay struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"placingWay"`
		ProcedureInfo struct {
			Text              string `xml:",chardata"`
			CollectingEndDate string `xml:"collectingEndDate"`
		} `xml:"procedureInfo"`
		Lots struct {
			Text string `xml:",chardata"`
			Lot  struct {
				Text      string `xml:",chardata"`
				LotNumber string `xml:"lotNumber"`
				MaxPrice  string `xml:"maxPrice"`
				Currency  struct {
					Text string `xml:",chardata"`
					Code string `xml:"code"`
					Name string `xml:"name"`
				} `xml:"currency"`
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
					} `xml:"purchaseObject"`
				} `xml:"purchaseObjects"`
				Preferenses struct {
					Text       string `xml:",chardata"`
					Preferense []struct {
						Text      string `xml:",chardata"`
						ShortName string `xml:"shortName"`
						Name      string `xml:"name"`
						PrefValue string `xml:"prefValue"`
					} `xml:"preferense"`
				} `xml:"preferenses"`
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
						Text             string `xml:",chardata"`
						ShortName        string `xml:"shortName"`
						Name             string `xml:"name"`
						RestrictionsSt14 struct {
							Text            string `xml:",chardata"`
							RestrictionSt14 []struct {
								Text             string `xml:",chardata"`
								RequirementsType struct {
									Text            string `xml:",chardata"`
									RequirementType []struct {
										Text string `xml:",chardata"`
										Type string `xml:"type"`
									} `xml:"requirementType"`
								} `xml:"requirementsType"`
								NPAInfo struct {
									Text      string `xml:",chardata"`
									Code      string `xml:"code"`
									Name      string `xml:"name"`
									ShortName string `xml:"shortName"`
								} `xml:"NPAInfo"`
							} `xml:"restrictionSt14"`
						} `xml:"restrictionsSt14"`
					} `xml:"restriction"`
				} `xml:"restrictions"`
				PurchaseCode string `xml:"purchaseCode"`
				IKZInfo      struct {
					Text        string `xml:",chardata"`
					PublishYear string `xml:"publishYear"`
					OKPD2Info   struct {
						Text  string `xml:",chardata"`
						OKPD2 struct {
							Text     string `xml:",chardata"`
							OKPDCode string `xml:"OKPDCode"`
							OKPDName string `xml:"OKPDName"`
						} `xml:"OKPD2"`
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
				ContractExecutionPaymentPlan struct {
					Text                 string `xml:",chardata"`
					FinancingSourcesInfo struct {
						Text            string `xml:",chardata"`
						FinancingSource string `xml:"financingSource"`
						CurrentYear     string `xml:"currentYear"`
						FinanceInfo     struct {
							Text        string `xml:",chardata"`
							Total       string `xml:"total"`
							CurrentYear string `xml:"currentYear"`
							FirstYear   string `xml:"firstYear"`
							SecondYear  string `xml:"secondYear"`
							SubsecYears string `xml:"subsecYears"`
						} `xml:"financeInfo"`
					} `xml:"financingSourcesInfo"`
					BudgetFinancings struct {
						Text            string `xml:",chardata"`
						CurrentYear     string `xml:"currentYear"`
						BudgetFinancing struct {
							Text         string `xml:",chardata"`
							KBKCode      string `xml:"KBKCode"`
							KBKYearsInfo struct {
								Text        string `xml:",chardata"`
								Total       string `xml:"total"`
								CurrentYear string `xml:"currentYear"`
								FirstYear   string `xml:"firstYear"`
								SecondYear  string `xml:"secondYear"`
								SubsecYears string `xml:"subsecYears"`
							} `xml:"KBKYearsInfo"`
						} `xml:"budgetFinancing"`
						TotalSum string `xml:"totalSum"`
					} `xml:"budgetFinancings"`
				} `xml:"contractExecutionPaymentPlan"`
				BOInfo struct {
					Text         string `xml:",chardata"`
					BONumber     string `xml:"BONumber"`
					BODate       string `xml:"BODate"`
					InputBOFlag  string `xml:"inputBOFlag"`
					BORegistered string `xml:"BORegistered"`
				} `xml:"BOInfo"`
				MustPublicDiscussion string `xml:"mustPublicDiscussion"`
			} `xml:"lot"`
		} `xml:"lots"`
		ParticularsActProcurement string `xml:"particularsActProcurement"`
	} `xml:"fcsNotification111"`
}
