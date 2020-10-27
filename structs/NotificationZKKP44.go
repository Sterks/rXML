package structs

import "encoding/xml"

// NotificationZKKP44 Закрытый конкурс
type NotificationZKKP44 struct {
	XMLName             xml.Name `xml:"export"`
	Text                string   `xml:",chardata"`
	Xmlns               string   `xml:"xmlns,attr"`
	Ns6                 string   `xml:"ns6,attr"`
	Ns5                 string   `xml:"ns5,attr"`
	Ns8                 string   `xml:"ns8,attr"`
	Ns7                 string   `xml:"ns7,attr"`
	Ns9                 string   `xml:"ns9,attr"`
	Ns11                string   `xml:"ns11,attr"`
	Ns10                string   `xml:"ns10,attr"`
	Ns2                 string   `xml:"ns2,attr"`
	Ns4                 string   `xml:"ns4,attr"`
	Ns3                 string   `xml:"ns3,attr"`
	FcsNotificationZakK struct {
		Text           string `xml:",chardata"`
		SchemeVersion  string `xml:"schemeVersion,attr"`
		ID             string `xml:"id"`
		PurchaseNumber string `xml:"purchaseNumber"`
		DirectDate     string `xml:"directDate"`
		DocPublishDate string `xml:"docPublishDate"`
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
		IsGOZ               string `xml:"isGOZ"`
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
					Text      string `xml:",chardata"`
					LastName  string `xml:"lastName"`
					FirstName string `xml:"firstName"`
				} `xml:"contactPerson"`
				ContactEMail string `xml:"contactEMail"`
				ContactPhone string `xml:"contactPhone"`
			} `xml:"responsibleInfo"`
		} `xml:"purchaseResponsible"`
		PlacingWay struct {
			Text string `xml:",chardata"`
			Code string `xml:"code"`
			Name string `xml:"name"`
		} `xml:"placingWay"`
		ContractConclusionOnSt83Ch2 string `xml:"contractConclusionOnSt83Ch2"`
		PurchaseDocumentation       struct {
			Text           string `xml:",chardata"`
			GrantStartDate string `xml:"grantStartDate"`
			GrantPlace     string `xml:"grantPlace"`
			GrantOrder     string `xml:"grantOrder"`
			Languages      string `xml:"languages"`
			GrantMeans     string `xml:"grantMeans"`
			GrantEndDate   string `xml:"grantEndDate"`
		} `xml:"purchaseDocumentation"`
		ProcedureInfo struct {
			Text       string `xml:",chardata"`
			Collecting struct {
				Text      string `xml:",chardata"`
				StartDate string `xml:"startDate"`
				Place     string `xml:"place"`
				Order     string `xml:"order"`
				EndDate   string `xml:"endDate"`
			} `xml:"collecting"`
			Opening struct {
				Text    string `xml:",chardata"`
				Date    string `xml:"date"`
				Place   string `xml:"place"`
				AddInfo string `xml:"addInfo"`
			} `xml:"opening"`
			Scoring struct {
				Text  string `xml:",chardata"`
				Date  string `xml:"date"`
				Place string `xml:"place"`
			} `xml:"scoring"`
		} `xml:"procedureInfo"`
		Lots struct {
			Text string `xml:",chardata"`
			Lot  struct {
				Text          string `xml:",chardata"`
				LotNumber     string `xml:"lotNumber"`
				LotObjectInfo string `xml:"lotObjectInfo"`
				MaxPrice      string `xml:"maxPrice"`
				Currency      struct {
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
						AdvancePaymentSum    struct {
							Text          string `xml:",chardata"`
							SumInPercents string `xml:"sumInPercents"`
						} `xml:"advancePaymentSum"`
						KladrPlaces struct {
							Text       string `xml:",chardata"`
							KladrPlace struct {
								Text          string `xml:",chardata"`
								DeliveryPlace string `xml:"deliveryPlace"`
							} `xml:"kladrPlace"`
						} `xml:"kladrPlaces"`
						DeliveryTerm         string `xml:"deliveryTerm"`
						ApplicationGuarantee struct {
							Text              string `xml:",chardata"`
							Amount            string `xml:"amount"`
							Part              string `xml:"part"`
							ProcedureInfo     string `xml:"procedureInfo"`
							SettlementAccount string `xml:"settlementAccount"`
							PersonalAccount   string `xml:"personalAccount"`
							Bik               string `xml:"bik"`
						} `xml:"applicationGuarantee"`
						ProvisionWarranty struct {
							Text          string `xml:",chardata"`
							Amount        string `xml:"amount"`
							Part          string `xml:"part"`
							ProcedureInfo string `xml:"procedureInfo"`
							Account       struct {
								Text              string `xml:",chardata"`
								Bik               string `xml:"bik"`
								SettlementAccount string `xml:"settlementAccount"`
								PersonalAccount   string `xml:"personalAccount"`
							} `xml:"account"`
						} `xml:"provisionWarranty"`
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
						BankSupportContractRequiredInfo struct {
							Text                        string `xml:",chardata"`
							BankSupportContractRequired string `xml:"bankSupportContractRequired"`
						} `xml:"bankSupportContractRequiredInfo"`
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
						Content   string `xml:"content"`
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
									RequirementType struct {
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
			} `xml:"lot"`
		} `xml:"lots"`
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
		Modification struct {
			Text               string `xml:",chardata"`
			ModificationNumber string `xml:"modificationNumber"`
			Info               string `xml:"info"`
			Reason             struct {
				Text                string `xml:",chardata"`
				ResponsibleDecision struct {
					Text         string `xml:",chardata"`
					DecisionDate string `xml:"decisionDate"`
				} `xml:"responsibleDecision"`
			} `xml:"reason"`
		} `xml:"modification"`
	} `xml:"fcsNotificationZakK"`
}
