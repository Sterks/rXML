package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant44FZ struct {
	ID               primitive.ObjectID
	Reest            string
	Lot 			 string
	Application struct {
		Text             string `xml:",chardata"`
		JournalNumber    string `xml:"journalNumber"`
		AppDate          string `xml:"appDate"`
		AdmissionResults struct {
			Text            string `xml:",chardata"`
			AdmissionResult []struct {
				Text                     string `xml:",chardata"`
				ProtocolCommissionMember struct {
					Text         string `xml:",chardata"`
					MemberNumber string `xml:"memberNumber"`
				} `xml:"protocolCommissionMember"`
				Admitted string `xml:"admitted"`
			} `xml:"admissionResult"`
		} `xml:"admissionResults"`
		Admitted string `xml:"admitted"`
	}
}
