package models

type TransactionWithDonationObject struct {
	Id             *string         `json:"id"`
	Type           *string         `json:"type"`
	RefundedAmount *int            `json:"refundedAmount"`
	Donation       *DonationObject `json:"donation"`
	Typename       *string         `json:"__typename"`
}

type DonationObject struct {
	Id              *string `json:"id"`
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	CreatedAtUtc    *int    `json:"createdAtUtc"`
	Amount          *int    `json:"amount"`
	ThankYouComment *string `json:"thankYouComment"`
	IsAnonymous     *bool   `json:"isAnonymous"`
	CompanyName     *string `json:"companyName"`
	Typename        *string `json:"__typename"`
}
