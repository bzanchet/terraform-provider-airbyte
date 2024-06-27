// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

// SourceMailgunDomainRegionCode - Domain region code. 'EU' or 'US' are possible values. The default is 'US'.
type SourceMailgunDomainRegionCode string

const (
	SourceMailgunDomainRegionCodeUs SourceMailgunDomainRegionCode = "US"
	SourceMailgunDomainRegionCodeEu SourceMailgunDomainRegionCode = "EU"
)

func (e SourceMailgunDomainRegionCode) ToPointer() *SourceMailgunDomainRegionCode {
	return &e
}
func (e *SourceMailgunDomainRegionCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = SourceMailgunDomainRegionCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailgunDomainRegionCode: %v", v)
	}
}

type Mailgun string

const (
	MailgunMailgun Mailgun = "mailgun"
)

func (e Mailgun) ToPointer() *Mailgun {
	return &e
}
func (e *Mailgun) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mailgun":
		*e = Mailgun(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mailgun: %v", v)
	}
}

type SourceMailgun struct {
	// Domain region code. 'EU' or 'US' are possible values. The default is 'US'.
	DomainRegion *SourceMailgunDomainRegionCode `default:"US" json:"domain_region"`
	// Primary account API key to access your Mailgun data.
	PrivateKey string  `json:"private_key"`
	sourceType Mailgun `const:"mailgun" json:"sourceType"`
	// UTC date and time in the format 2020-10-01 00:00:00. Any data before this date will not be replicated. If omitted, defaults to 3 days ago.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceMailgun) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMailgun) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMailgun) GetDomainRegion() *SourceMailgunDomainRegionCode {
	if o == nil {
		return nil
	}
	return o.DomainRegion
}

func (o *SourceMailgun) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *SourceMailgun) GetSourceType() Mailgun {
	return MailgunMailgun
}

func (o *SourceMailgun) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
