package structs

import "encoding/xml"

// Запрос предложений в электронной форме
type NotificationZP504 struct {
	XMLName           xml.Name `xml:"export"`
	Text              string   `xml:",chardata"`
	Xmlns             string   `xml:"xmlns,attr"`
	Ns6               string   `xml:"ns6,attr"`
	Ns5               string   `xml:"ns5,attr"`
	Ns8               string   `xml:"ns8,attr"`
	Ns7               string   `xml:"ns7,attr"`
	Ns9               string   `xml:"ns9,attr"`
	Ns11              string   `xml:"ns11,attr"`
	Ns10              string   `xml:"ns10,attr"`
	Ns2               string   `xml:"ns2,attr"`
	Ns4               string   `xml:"ns4,attr"`
	Ns3               string   `xml:"ns3,attr"`
	EpNotificationEZP struct {
		Text          string `xml:",chardata"`
		SchemeVersion string `xml:"schemeVersion,attr"`
		ID            string `xml:"id"`
		ExternalId    string `xml:"externalId"`
		VersionNumber string `xml:"versionNumber"`
		CommonInfo    struct {
			Text               string `xml:",chardata"`
			PurchaseNumber     string `xml:"purchaseNumber"`
			DocNumber          string `xml:"docNumber"`
			PlannedPublishDate string `xml:"plannedPublishDate"`
			PublishDTInEIS     string `xml:"publishDTInEIS"`
			Href               string `xml:"href"`
			PurchaseObjectInfo string `xml:"purchaseObjectInfo"`
			PlacingWay         struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
			} `xml:"placingWay"`
			ETP struct {
				Text string `xml:",chardata"`
				Code string `xml:"code"`
				Name string `xml:"name"`
				URL  string `xml:"url"`
			} `xml:"ETP"`
			ContractConclusionOnSt83Ch2 string `xml:"contractConclusionOnSt83Ch2"`
		} `xml:"commonInfo"`
		PurchaseResponsibleInfo struct {
			Text               string `xml:",chardata"`
			ResponsibleOrgInfo struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				PostAddress     string `xml:"postAddress"`
				FactAddress     string `xml:"factAddress"`
				INN             string `xml:"INN"`
				KPP             string `xml:"KPP"`
			} `xml:"responsibleOrgInfo"`
			ResponsibleRole string `xml:"responsibleRole"`
			ResponsibleInfo struct {
				Text              string `xml:",chardata"`
				OrgPostAddress    string `xml:"orgPostAddress"`
				OrgFactAddress    string `xml:"orgFactAddress"`
				ContactPersonInfo struct {
					Text       string `xml:",chardata"`
					LastName   string `xml:"lastName"`
					FirstName  string `xml:"firstName"`
					MiddleName string `xml:"middleName"`
				} `xml:"contactPersonInfo"`
				ContactEMail string `xml:"contactEMail"`
				ContactPhone string `xml:"contactPhone"`
				ContactFax   string `xml:"contactFax"`
			} `xml:"responsibleInfo"`
		} `xml:"purchaseResponsibleInfo"`
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
			AttachmentInfo struct {
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
		NotificationInfo struct {
			Text          string `xml:",chardata"`
			ProcedureInfo struct {
				Text           string `xml:",chardata"`
				CollectingInfo struct {
					Text    string `xml:",chardata"`
					StartDT string `xml:"startDT"`
					EndDT   string `xml:"endDT"`
					Place   string `xml:"place"`
					Order   string `xml:"order"`
				} `xml:"collectingInfo"`
				ScoringInfo struct {
					Text    string `xml:",chardata"`
					EndDate string `xml:"endDate"`
				} `xml:"scoringInfo"`
			} `xml:"procedureInfo"`
			ContractConditionsInfo struct {
				Text         string `xml:",chardata"`
				MaxPriceInfo struct {
					Text     string `xml:",chardata"`
					MaxPrice string `xml:"maxPrice"`
					Currency struct {
						Text string `xml:",chardata"`
						Code string `xml:"code"`
						Name string `xml:"name"`
					} `xml:"currency"`
				} `xml:"maxPriceInfo"`
			} `xml:"contractConditionsInfo"`
			CustomerRequirementsInfo struct {
				Text                    string `xml:",chardata"`
				CustomerRequirementInfo struct {
					Text     string `xml:",chardata"`
					Customer struct {
						Text            string `xml:",chardata"`
						RegNum          string `xml:"regNum"`
						ConsRegistryNum string `xml:"consRegistryNum"`
						FullName        string `xml:"fullName"`
					} `xml:"customer"`
					ApplicationGuarantee struct {
						Text    string `xml:",chardata"`
						Amount  string `xml:"amount"`
						Account struct {
							Text              string `xml:",chardata"`
							Bik               string `xml:"bik"`
							SettlementAccount string `xml:"settlementAccount"`
							PersonalAccount   string `xml:"personalAccount"`
						} `xml:"account"`
						ProcedureInfo string `xml:"procedureInfo"`
						Part          string `xml:"part"`
					} `xml:"applicationGuarantee"`
					ContractGuarantee struct {
						Text    string `xml:",chardata"`
						Account struct {
							Text              string `xml:",chardata"`
							Bik               string `xml:"bik"`
							SettlementAccount string `xml:"settlementAccount"`
							PersonalAccount   string `xml:"personalAccount"`
						} `xml:"account"`
						ProcedureInfo string `xml:"procedureInfo"`
						Part          string `xml:"part"`
					} `xml:"contractGuarantee"`
					ContractConditionsInfo struct {
						Text         string `xml:",chardata"`
						MaxPriceInfo struct {
							Text     string `xml:",chardata"`
							MaxPrice string `xml:"maxPrice"`
						} `xml:"maxPriceInfo"`
						PurchaseCode string `xml:"purchaseCode"`
						IKZInfo      struct {
							Text        string `xml:",chardata"`
							PublishYear string `xml:"publishYear"`
							OKPD2Info   struct {
								Text      string `xml:",chardata"`
								Undefined string `xml:"undefined"`
							} `xml:"OKPD2Info"`
							KVRInfo struct {
								Text      string `xml:",chardata"`
								Undefined string `xml:"undefined"`
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
							KVRsInfo struct {
								Text        string `xml:",chardata"`
								CurrentYear string `xml:"currentYear"`
								KVRInfo     struct {
									Text string `xml:",chardata"`
									KVR  struct {
										Text string `xml:",chardata"`
										Code string `xml:"code"`
										Name string `xml:"name"`
									} `xml:"KVR"`
									KVRYearsInfo struct {
										Text        string `xml:",chardata"`
										Total       string `xml:"total"`
										CurrentYear string `xml:"currentYear"`
										FirstYear   string `xml:"firstYear"`
										SecondYear  string `xml:"secondYear"`
										SubsecYears string `xml:"subsecYears"`
									} `xml:"KVRYearsInfo"`
								} `xml:"KVRInfo"`
								TotalSum string `xml:"totalSum"`
							} `xml:"KVRsInfo"`
						} `xml:"contractExecutionPaymentPlan"`
						DeliveryPlacesInfo struct {
							Text              string `xml:",chardata"`
							DeliveryPlaceInfo []struct {
								Text          string `xml:",chardata"`
								DeliveryPlace string `xml:"deliveryPlace"`
							} `xml:"deliveryPlaceInfo"`
						} `xml:"deliveryPlacesInfo"`
						DeliveryTerm string `xml:"deliveryTerm"`
					} `xml:"contractConditionsInfo"`
				} `xml:"customerRequirementInfo"`
			} `xml:"customerRequirementsInfo"`
			PurchaseObjectsInfo struct {
				Text                       string `xml:",chardata"`
				NotDrugPurchaseObjectsInfo struct {
					Text           string `xml:",chardata"`
					PurchaseObject []struct {
						Text  string `xml:",chardata"`
						OKPD2 struct {
							Text     string `xml:",chardata"`
							OKPDCode string `xml:"OKPDCode"`
							OKPDName string `xml:"OKPDName"`
						} `xml:"OKPD2"`
						Name string `xml:"name"`
						OKEI struct {
							Text         string `xml:",chardata"`
							Code         string `xml:"code"`
							NationalCode string `xml:"nationalCode"`
							Name         string `xml:"name"`
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
					TotalSum          string `xml:"totalSum"`
					QuantityUndefined string `xml:"quantityUndefined"`
				} `xml:"notDrugPurchaseObjectsInfo"`
			} `xml:"purchaseObjectsInfo"`
			RequirementsInfo struct {
				Text            string `xml:",chardata"`
				RequirementInfo []struct {
					Text                      string `xml:",chardata"`
					PreferenseRequirementInfo struct {
						Text      string `xml:",chardata"`
						ShortName string `xml:"shortName"`
						Name      string `xml:"name"`
					} `xml:"preferenseRequirementInfo"`
				} `xml:"requirementInfo"`
			} `xml:"requirementsInfo"`
			RestrictionsInfo struct {
				Text            string `xml:",chardata"`
				RestrictionInfo struct {
					Text                      string `xml:",chardata"`
					PreferenseRequirementInfo struct {
						Text      string `xml:",chardata"`
						ShortName string `xml:"shortName"`
						Name      string `xml:"name"`
					} `xml:"preferenseRequirementInfo"`
				} `xml:"restrictionInfo"`
			} `xml:"restrictionsInfo"`
		} `xml:"notificationInfo"`
		Documentation struct {
			Text                     string `xml:",chardata"`
			DocumentRequirementsInfo struct {
				Text                    string `xml:",chardata"`
				DocumentRequirementInfo []struct {
					Text      string `xml:",chardata"`
					SId       string `xml:"sId"`
					Number    string `xml:"number"`
					Name      string `xml:"name"`
					Mandatory string `xml:"mandatory"`
				} `xml:"documentRequirementInfo"`
			} `xml:"documentRequirementsInfo"`
			Modifiable   string `xml:"modifiable"`
			CriterisInfo struct {
				Text          string `xml:",chardata"`
				CriterionInfo []struct {
					Text              string `xml:",chardata"`
					CostCriterionInfo struct {
						Text          string `xml:",chardata"`
						CriterionCode string `xml:"criterionCode"`
						Value         string `xml:"value"`
					} `xml:"costCriterionInfo"`
					QualitativeCriterionInfo struct {
						Text           string `xml:",chardata"`
						CriterionCode  string `xml:"criterionCode"`
						Value          string `xml:"value"`
						IndicatorsInfo struct {
							Text          string `xml:",chardata"`
							IndicatorInfo []struct {
								Text             string `xml:",chardata"`
								SId              string `xml:"sId"`
								Name             string `xml:"name"`
								Value            string `xml:"value"`
								MeasurementOrder string `xml:"measurementOrder"`
							} `xml:"indicatorInfo"`
						} `xml:"indicatorsInfo"`
					} `xml:"qualitativeCriterionInfo"`
				} `xml:"criterionInfo"`
			} `xml:"criterisInfo"`
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
		} `xml:"documentation"`
	} `xml:"epNotificationEZP"`
}
