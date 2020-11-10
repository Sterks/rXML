package notifications223

import "encoding/xml"

// PurchaseNoticeAESMBO was generated 2020-11-09 23:29:18 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type PurchaseNoticeAESMBO struct {
	XMLName        xml.Name `xml:"purchaseNoticeAESMBO"`
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
			Text                     string `xml:",chardata"`
			Guid                     string `xml:"guid"`
			PurchaseNoticeAESMBOData struct {
				Text               string `xml:",chardata"`
				Guid               string `xml:"guid"`
				CreateDateTime     string `xml:"createDateTime"`
				UrlEIS             string `xml:"urlEIS"`
				UrlVSRZ            string `xml:"urlVSRZ"`
				RegistrationNumber string `xml:"registrationNumber"`
				Name               string `xml:"name"`
				Customer           struct {
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
						TimeZone      struct {
							Text   string `xml:",chardata"`
							Offset string `xml:"offset"`
							Name   string `xml:"name"`
						} `xml:"timeZone"`
						Region string `xml:"region"`
					} `xml:"mainInfo"`
				} `xml:"customer"`
				PurchaseMethodCode string `xml:"purchaseMethodCode"`
				PurchaseCodeName   string `xml:"purchaseCodeName"`
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
				Contact struct {
					Text       string `xml:",chardata"`
					FirstName  string `xml:"firstName"`
					MiddleName string `xml:"middleName"`
					LastName   string `xml:"lastName"`
					Phone      string `xml:"phone"`
					Email      string `xml:"email"`
				} `xml:"contact"`
				PublicationDateTime   string `xml:"publicationDateTime"`
				DocumentationDelivery struct {
					Text                  string `xml:",chardata"`
					DeliveryStartDateTime string `xml:"deliveryStartDateTime"`
					DeliveryEndDateTime   string `xml:"deliveryEndDateTime"`
					Place                 string `xml:"place"`
					Procedure             string `xml:"procedure"`
				} `xml:"documentationDelivery"`
				Status       string `xml:"status"`
				Version      string `xml:"version"`
				NotDishonest string `xml:"notDishonest"`
				Attachments  struct {
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
				ModificationDate          string `xml:"modificationDate"`
				SaveUserId                string `xml:"saveUserId"`
				DeliveryPlaceIndication   string `xml:"deliveryPlaceIndication"`
				Emergency                 string `xml:"emergency"`
				JointPurchase             string `xml:"jointPurchase"`
				ForSmallOrMiddle          string `xml:"forSmallOrMiddle"`
				AntimonopolyDecisionTaken string `xml:"antimonopolyDecisionTaken"`
				ApplSubmisionStartDate    string `xml:"applSubmisionStartDate"`
				ApplSubmisionOrder        string `xml:"applSubmisionOrder"`
				SummingupOrder            string `xml:"summingupOrder"`
				SignatureAuthorizedBody   string `xml:"signatureAuthorizedBody"`
				ElectronicPlaceInfo       struct {
					Text              string `xml:",chardata"`
					Name              string `xml:"name"`
					URL               string `xml:"url"`
					ElectronicPlaceId string `xml:"electronicPlaceId"`
				} `xml:"electronicPlaceInfo"`
				PlacingProcedure struct {
					Text              string `xml:",chardata"`
					SummingupDateTime string `xml:"summingupDateTime"`
					SummingupPlace    string `xml:"summingupPlace"`
				} `xml:"placingProcedure"`
				SubmissionCloseDateTime string `xml:"submissionCloseDateTime"`
				PublicationPlannedDate  string `xml:"publicationPlannedDate"`
				TemplateVersion         string `xml:"templateVersion"`
				TemplateStructure       struct {
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
				} `xml:"templateStructure"`
				ExtendFields struct {
					Text              string `xml:",chardata"`
					NoticeExtendField []struct {
						Text        string `xml:",chardata"`
						ExtendField []struct {
							Text        string `xml:",chardata"`
							IntegrCode  string `xml:"integrCode"`
							Description string `xml:"description"`
							Type        string `xml:"type"`
							Value       struct {
								Chardata string `xml:",chardata"`
								Text     string `xml:"text"`
								Date     string `xml:"date"`
								Boolean  string `xml:"boolean"`
								Time     string `xml:"time"`
								DateTime string `xml:"dateTime"`
							} `xml:"value"`
						} `xml:"extendField"`
						Position struct {
							Text           string `xml:",chardata"`
							TabOrdinal     string `xml:"tabOrdinal"`
							TabName        string `xml:"tabName"`
							SectionOrdinal string `xml:"sectionOrdinal"`
							SectionName    string `xml:"sectionName"`
						} `xml:"position"`
					} `xml:"noticeExtendField"`
				} `xml:"extendFields"`
				Lots struct {
					Text string `xml:",chardata"`
					Lot  struct {
						Text           string `xml:",chardata"`
						Guid           string `xml:"guid"`
						OrdinalNumber  string `xml:"ordinalNumber"`
						LotEditEnabled string `xml:"lotEditEnabled"`
						LotData        struct {
							Text     string `xml:",chardata"`
							Subject  string `xml:"subject"`
							Currency struct {
								Text        string `xml:",chardata"`
								Code        string `xml:"code"`
								DigitalCode string `xml:"digitalCode"`
								Name        string `xml:"name"`
							} `xml:"currency"`
							InitialSum    string `xml:"initialSum"`
							OrderPricing  string `xml:"orderPricing"`
							DeliveryPlace struct {
								Text    string `xml:",chardata"`
								Address string `xml:"address"`
							} `xml:"deliveryPlace"`
							LotItems struct {
								Text     string `xml:",chardata"`
								NewCodes string `xml:"newCodes"`
								LotItem  struct {
									Text          string `xml:",chardata"`
									Guid          string `xml:"guid"`
									OrdinalNumber string `xml:"ordinalNumber"`
									Okpd2         struct {
										Text string `xml:",chardata"`
										Code string `xml:"code"`
										Name string `xml:"name"`
									} `xml:"okpd2"`
									Okved2 struct {
										Text string `xml:",chardata"`
										Code string `xml:"code"`
										Name string `xml:"name"`
									} `xml:"okved2"`
									Okei struct {
										Text string `xml:",chardata"`
										Code string `xml:"code"`
										Name string `xml:"name"`
									} `xml:"okei"`
									Qty string `xml:"qty"`
								} `xml:"lotItem"`
							} `xml:"lotItems"`
							ForSmallOrMiddle          string `xml:"forSmallOrMiddle"`
							ExcludePurchaseFromPlan   string `xml:"excludePurchaseFromPlan"`
							SubcontractorsRequirement string `xml:"subcontractorsRequirement"`
							IgnoredPurchase           string `xml:"ignoredPurchase"`
							Centralized               string `xml:"centralized"`
						} `xml:"lotData"`
						DeliveryPlaceIndication string `xml:"deliveryPlaceIndication"`
						JointLotData            struct {
							Text     string `xml:",chardata"`
							JointLot string `xml:"jointLot"`
						} `xml:"jointLotData"`
						LotPlanInfo struct {
							Text                   string `xml:",chardata"`
							PlanRegistrationNumber string `xml:"planRegistrationNumber"`
							PlanGuid               string `xml:"planGuid"`
							PositionNumber         string `xml:"positionNumber"`
							LotPlanPosition        string `xml:"lotPlanPosition"`
							PositionGuid           string `xml:"positionGuid"`
						} `xml:"lotPlanInfo"`
						Cancelled string `xml:"cancelled"`
					} `xml:"lot"`
				} `xml:"lots"`
			} `xml:"purchaseNoticeAESMBOData"`
		} `xml:"item"`
	} `xml:"body"`
}
