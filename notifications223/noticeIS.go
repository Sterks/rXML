package notifications223

import "encoding/xml"

// PurchaseNotice was generated 2020-11-09 23:33:28 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type PurchaseNoticeIS struct {
	XMLName        xml.Name `xml:"purchaseNotice"`
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
			Text               string `xml:",chardata"`
			Guid               string `xml:"guid"`
			PurchaseNoticeData struct {
				Text               string `xml:",chardata"`
				Guid               string `xml:"guid"`
				CreateDateTime     string `xml:"createDateTime"`
				RegistrationNumber string `xml:"registrationNumber"`
				Name               string `xml:"name"`
				Customer           struct {
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
				} `xml:"customer"`
				PurchaseMethodCode string `xml:"purchaseMethodCode"`
				PurchaseCodeName   string `xml:"purchaseCodeName"`
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
				Contact struct {
					Text         string `xml:",chardata"`
					FirstName    string `xml:"firstName"`
					MiddleName   string `xml:"middleName"`
					LastName     string `xml:"lastName"`
					Phone        string `xml:"phone"`
					Fax          string `xml:"fax"`
					Email        string `xml:"email"`
					Organization struct {
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
					} `xml:"organization"`
				} `xml:"contact"`
				SubmissionCloseDateTime string `xml:"submissionCloseDateTime"`
				PublicationPlannedDate  string `xml:"publicationPlannedDate"`
				PublicationDateTime     string `xml:"publicationDateTime"`
				DocumentationDelivery   struct {
					Text                  string `xml:",chardata"`
					DeliveryStartDateTime string `xml:"deliveryStartDateTime"`
					DeliveryEndDateTime   string `xml:"deliveryEndDateTime"`
					Place                 string `xml:"place"`
					Procedure             string `xml:"procedure"`
					URL                   string `xml:"url"`
				} `xml:"documentationDelivery"`
				Status  string `xml:"status"`
				Version string `xml:"version"`
				Lots    struct {
					Text string `xml:",chardata"`
					Lot  struct {
						Text          string `xml:",chardata"`
						OrdinalNumber string `xml:"ordinalNumber"`
						Subject       string `xml:"subject"`
						Currency      struct {
							Text        string `xml:",chardata"`
							Code        string `xml:"code"`
							DigitalCode string `xml:"digitalCode"`
							Name        string `xml:"name"`
						} `xml:"currency"`
						InitialSum    string `xml:"initialSum"`
						DeliveryPlace string `xml:"deliveryPlace"`
						LotItems      struct {
							Text    string `xml:",chardata"`
							LotItem struct {
								Text          string `xml:",chardata"`
								OrdinalNumber string `xml:"ordinalNumber"`
								Okdp          struct {
									Text string `xml:",chardata"`
									Code string `xml:"code"`
									Name string `xml:"name"`
								} `xml:"okdp"`
								Okved struct {
									Text string `xml:",chardata"`
									Code string `xml:"code"`
									Name string `xml:"name"`
								} `xml:"okved"`
								Okei struct {
									Text   string `xml:",chardata"`
									Code   string `xml:"code"`
									Name   string `xml:"name"`
									OkeiId string `xml:"okeiId"`
								} `xml:"okei"`
								AdditionalInfo string `xml:"additionalInfo"`
							} `xml:"lotItem"`
						} `xml:"lotItems"`
					} `xml:"lot"`
				} `xml:"lots"`
				Attachements struct {
					Text     string `xml:",chardata"`
					Document []struct {
						Text           string `xml:",chardata"`
						CreateDateTime string `xml:"createDateTime"`
						FileName       string `xml:"fileName"`
						Description    string `xml:"description"`
						URL            string `xml:"url"`
					} `xml:"document"`
				} `xml:"attachements"`
				NonelectronicPlaceInfo struct {
					Text                string `xml:",chardata"`
					SummarizingPlace    string `xml:"summarizingPlace"`
					SummarizingDateTime string `xml:"summarizingDateTime"`
				} `xml:"nonelectronicPlaceInfo"`
				ExaminationPlace    string `xml:"examinationPlace"`
				ExaminationDateTime string `xml:"examinationDateTime"`
			} `xml:"purchaseNoticeData"`
		} `xml:"item"`
	} `xml:"body"`
}
