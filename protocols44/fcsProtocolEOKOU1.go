package protocols44

import "encoding/xml"

// Export was generated 2020-10-30 10:52:02 by DESKTOP-ES8AI3K\runov on DESKTOP-ES8AI3K.
type ProtocolEOKOU1 struct {
	XMLName          xml.Name `xml:"export"`
	Text             string   `xml:",chardata"`
	SchemaLocation   string   `xml:"schemaLocation,attr"`
	Xmlns            string   `xml:"xmlns,attr"`
	Ns6              string   `xml:"ns6,attr"`
	Ns5              string   `xml:"ns5,attr"`
	Ns8              string   `xml:"ns8,attr"`
	Ns7              string   `xml:"ns7,attr"`
	Ns9              string   `xml:"ns9,attr"`
	Ns11             string   `xml:"ns11,attr"`
	Ns10             string   `xml:"ns10,attr"`
	Ns2              string   `xml:"ns2,attr"`
	Ns4              string   `xml:"ns4,attr"`
	Ns3              string   `xml:"ns3,attr"`
	Xsi              string   `xml:"xsi,attr"`
	EpProtocolEOKOU1 struct {
		Text              string `xml:",chardata"`
		SchemeVersion     string `xml:"schemeVersion,attr"`
		ID                string `xml:"id"`
		ExternalId        string `xml:"externalId"`
		VersionNumber     string `xml:"versionNumber"`
		FoundationDocInfo struct {
			Text                        string `xml:",chardata"`
			FoundationDocNumber         string `xml:"foundationDocNumber"`
			FoundationDocNumberExternal string `xml:"foundationDocNumberExternal"`
		} `xml:"foundationDocInfo"`
		CommonInfo struct {
			Text              string `xml:",chardata"`
			PurchaseNumber    string `xml:"purchaseNumber"`
			DocNumber         string `xml:"docNumber"`
			PublishDTInEIS    string `xml:"publishDTInEIS"`
			Href              string `xml:"href"`
			DocNumberExternal string `xml:"docNumberExternal"`
			PublishDTInETP    string `xml:"publishDTInETP"`
			Place             string `xml:"place"`
			ProcedureDT       string `xml:"procedureDT"`
			SignDT            string `xml:"signDT"`
			HrefExternal      string `xml:"hrefExternal"`
		} `xml:"commonInfo"`
		ProtocolPublisherInfo struct {
			Text         string `xml:",chardata"`
			PublisherOrg struct {
				Text            string `xml:",chardata"`
				RegNum          string `xml:"regNum"`
				ConsRegistryNum string `xml:"consRegistryNum"`
				FullName        string `xml:"fullName"`
				PostAddress     string `xml:"postAddress"`
				FactAddress     string `xml:"factAddress"`
				INN             string `xml:"INN"`
				KPP             string `xml:"KPP"`
			} `xml:"publisherOrg"`
			PublisherRole string `xml:"publisherRole"`
		} `xml:"protocolPublisherInfo"`
		CommissionInfo struct {
			Text              string `xml:",chardata"`
			CommissionName    string `xml:"commissionName"`
			CommissionMembers struct {
				Text             string `xml:",chardata"`
				CommissionMember []struct {
					Text         string `xml:",chardata"`
					MemberNumber string `xml:"memberNumber"`
					NameInfo     struct {
						Text       string `xml:",chardata"`
						LastName   string `xml:"lastName"`
						FirstName  string `xml:"firstName"`
						MiddleName string `xml:"middleName"`
					} `xml:"nameInfo"`
					Role struct {
						Text      string `xml:",chardata"`
						Code      string `xml:"code"`
						Name      string `xml:"name"`
						RightVote string `xml:"rightVote"`
					} `xml:"role"`
				} `xml:"commissionMember"`
			} `xml:"commissionMembers"`
		} `xml:"commissionInfo"`
		PrintFormInfo struct {
			Text      string `xml:",chardata"`
			URL       string `xml:"url"`
			Signature []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"signature"`
		} `xml:"printFormInfo"`
		ExtPrintFormInfo struct {
			Text      string `xml:",chardata"`
			URL       string `xml:"url"`
			Signature string `xml:"signature"`
			FileType  string `xml:"fileType"`
		} `xml:"extPrintFormInfo"`
		AfterProlongation string `xml:"afterProlongation"`
		ProtocolInfo      struct {
			Text             string `xml:",chardata"`
			ApplicationsInfo struct {
				Text            string `xml:",chardata"`
				ApplicationInfo []struct {
					Text       string `xml:",chardata"`
					CommonInfo struct {
						Text                 string `xml:",chardata"`
						AppNumber            string `xml:"appNumber"`
						AppDT                string `xml:"appDT"`
						AdmissionResultsInfo struct {
							Text                string `xml:",chardata"`
							AdmissionResultInfo []struct {
								Text                 string `xml:",chardata"`
								CommissionMemberInfo struct {
									Text         string `xml:",chardata"`
									MemberNumber string `xml:"memberNumber"`
									NameInfo     struct {
										Text       string `xml:",chardata"`
										LastName   string `xml:"lastName"`
										FirstName  string `xml:"firstName"`
										MiddleName string `xml:"middleName"`
									} `xml:"nameInfo"`
									Role struct {
										Text      string `xml:",chardata"`
										Code      string `xml:"code"`
										Name      string `xml:"name"`
										RightVote string `xml:"rightVote"`
									} `xml:"role"`
								} `xml:"commissionMemberInfo"`
								Admitted string `xml:"admitted"`
							} `xml:"admissionResultInfo"`
						} `xml:"admissionResultsInfo"`
					} `xml:"commonInfo"`
					AdmittedInfo struct {
						Text            string `xml:",chardata"`
						AppAdmittedInfo struct {
							Text     string `xml:",chardata"`
							Admitted string `xml:"admitted"`
						} `xml:"appAdmittedInfo"`
					} `xml:"admittedInfo"`
				} `xml:"applicationInfo"`
			} `xml:"applicationsInfo"`
			Research string `xml:"research"`
		} `xml:"protocolInfo"`
	} `xml:"epProtocolEOKOU1"`
}

