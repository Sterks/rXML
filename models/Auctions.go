package models

import "time"

// Auctions ...
type Auctions struct {
	ID                         int       `json:"tA_ID"`
	RegistryNumber             string    `json:"tA_RegistryNumber"`
	Lot                        string    `json:"tA_Lot"`
	NameAu                     string    `json:"tA_Name"`
	PriceStart                 float64   `json:"tA_PriceStart"`
	PriceEnd                   float64     `json:"tA_PriceEnd"`
	Provision                  float64     `json:"tA_Provision"`
	DateFactText               string    `json:"tA_DateFactText"`
	DateFact                   time.Time `json:"tA_DateFact"`
	Type                       int       `json:"tA_Type"`
	Site                       string    `json:"tA_Site"`
	DateProtocol               time.Time `json:"tA_DateProtocol"`
	DateProtocolPosting        time.Time `json:tA_DateProtocolPosting"`
	ProtocolURL                string    `json:"ProtocolURL"`
	ZakupkiURL                 string    `json:"ZakupkiURL"`
	Result                     bool      `json:"tA_Result"`
	TypeDoc                    int       `json:"tA_TypeDoc"`
	TypeContent                int       `json:"tA_TypeContent"`
	DateCreate                 time.Time `json:"tA_DateCreate"`
	DateParser                 time.Time `json:"tA_DateParser"`
	WorkDateText               string    `json:"tA_WorkDateText"`
	SiteText                   string    `json:"tA_SiteText"`
	DateProtocolText           string    `json:"tA_DateProtokolText"`
	Currency                   int       `json:"tA_Currency"`
	ConsumerText               string    `json:tA_ConsumerText"`
	ProvisionIsCalculate       bool      `json:"tA_ProvisionIsCalculate"`
	ProvisionText              string    `json:"tA_ProvisionText"`
	IsPriorityRegion           bool      `json:"tA_IsPriorityRegion"`
	IsJoint                    bool      `json:"tA_IsJoint"`
	DatePublication            time.Time `json:"tA_DatePublication"`
	DateParticipationStart     time.Time `json:"tA_DateParticipationStart"`
	DateParticipationEnd       time.Time `json:"tA_DateParticipationEnd"`
	DateScoring                time.Time `json:"tA_DateScoring"`
	DatePlanProtocol           time.Time `json:"tA_DatePlanProtocol"`
	ProvisionParticipation     int       `json:"tA_ProvisionParticipation"`
	IsSmallBuisnessOnly        bool      `json:"tA_IsSmallBuisnessOnly"`
	DateUpdateFromNotification time.Time `json:"tA_DateUpdateFromNotification"`
	DateUpdateFromProtocol     time.Time `json:"tA_DateUpdateFromProtocol"`
	IsSmallBuisnessAdvantage   bool      `json:"tA_IsSmallBuisnessAdvantage"`
	ProvisionAsPercent         float32       `json:"tA_ProvisionAsPercent"`
	PriceForUnit               int64     `json:"tA_PriceForUnit"`
	CountForUnit               int       `json:"tA_CountForUnit"`
	TypeExt                    int       `json:"tA_TypeExt"`
	ProvisionWarranty          int       `json:"tA_ProvisionWarranty"`
	ProvisionWarrantyAsPercent int       `json:"tA_ProvisionWarrantyAsPercent"`
	Advance                    int       `json:"tA_Advance"`
	AdvanceAsPercent           int       `json:"tA_AdvanceAsPercent"`
	Comment                    string    `json:"tA_Comment"`
	DateWorkStart              time.Time `json:"tA_DateWorkStart"`
	DateWorkEnd                time.Time `json:"tA_DateWorkEnd"`
}

// Если обеспечение отсутствует по считать процент от НМЦ
func (a Auctions)SetProvision(){
	a.Provision = a.PriceStart * float64(a.ProvisionAsPercent)
}