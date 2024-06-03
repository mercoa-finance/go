// This file was auto-generated by Fern from our API Definition.

package entity

import (
	json "encoding/json"
	fmt "fmt"
	mercoafinancego "github.com/mercoa-finance/go"
	core "github.com/mercoa-finance/go/core"
)

type FindEntities struct {
	// If true, will include entity payment methods as part of the response
	PaymentMethods *bool `json:"-" url:"paymentMethods,omitempty"`
	// If true, only entities with a direct relationship to the requesting organization will be returned. If false or not provided, all entities will be returned.
	IsCustomer *bool `json:"-" url:"isCustomer,omitempty"`
	// ID used to identify this entity in your system
	ForeignID []*string                       `json:"-" url:"foreignId,omitempty"`
	Status    []*mercoafinancego.EntityStatus `json:"-" url:"status,omitempty"`
	// If true, entities that are marked as payees will be returned.
	// If false or not provided, entities that are marked as payees will not be returned.
	IsPayee *bool `json:"-" url:"isPayee,omitempty"`
	// If true or not provided, entities that are marked as payors will be returned.
	// If false, entities that are marked as payors will not be returned.
	IsPayor *bool `json:"-" url:"isPayor,omitempty"`
	// Filter entities by name. Partial matches are supported.
	Name *string `json:"-" url:"name,omitempty"`
	// Number of entities to return. Limit can range between 1 and 100, and the default is 10.
	Limit *int `json:"-" url:"limit,omitempty"`
	// The ID of the entity to start after. If not provided, the first page of entities will be returned.
	StartingAfter *mercoafinancego.EntityID `json:"-" url:"startingAfter,omitempty"`
}

type GenerateOnboardingLink struct {
	// The type of onboarding link to generate. If not provided, the default is payee. The onboarding options are determined by your organization's onboarding configuration.
	Type mercoafinancego.EntityOnboardingLinkType `json:"-" url:"type"`
	// Expressed in seconds or a string describing a time span. The default is 24h.
	ExpiresIn *string `json:"-" url:"expiresIn,omitempty"`
	// The ID of the entity to connect to. If onboarding a payee, this should be the payor entity ID. If onboarding a payor, this should be the payee entity ID. If no connected entity ID is provided, the onboarding link will be for a standalone entity.
	ConnectedEntityID *mercoafinancego.EntityID `json:"-" url:"connectedEntityId,omitempty"`
}

type PlaidLinkTokenRequest struct {
	// ID of Bank Account to update
	PaymentMethodID *mercoafinancego.PaymentMethodID `json:"-" url:"paymentMethodId,omitempty"`
}

type SendOnboardingLink struct {
	// The type of onboarding link to generate. If not provided, the default is payee. The onboarding options are determined by your organization's onboarding configuration.
	Type mercoafinancego.EntityOnboardingLinkType `json:"-" url:"type"`
	// Expressed in seconds or a string describing a time span. The default is 7 days.
	ExpiresIn *string `json:"-" url:"expiresIn,omitempty"`
	// The ID of the entity to connect to. If onboarding a payee, this should be the payor entity ID. If onboarding a payor, this should be the payee entity ID. If no connected entity ID is provided, the onboarding link will be for a standalone entity.
	ConnectedEntityID *mercoafinancego.EntityID `json:"-" url:"connectedEntityId,omitempty"`
}

type CodatCompanyCreationRequest struct {
	// If the company already exists in Codat, provide the company ID to link the company to the entity. If the company does not exist, leave this field blank and Codat will create a new company.
	CompanyID *string `json:"companyId,omitempty" url:"companyId,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CodatCompanyCreationRequest) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CodatCompanyCreationRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CodatCompanyCreationRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CodatCompanyCreationRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CodatCompanyCreationRequest) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CodatCompanyResponse struct {
	CompanyID string `json:"companyId" url:"companyId"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CodatCompanyResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CodatCompanyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CodatCompanyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CodatCompanyResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CodatCompanyResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type RutterCompanyCreationRequest struct {
	// The access token for the existing Rutter connection. If the connection does not exist, leave this field blank and Rutter will create a new connection.
	AccessToken *string `json:"accessToken,omitempty" url:"accessToken,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RutterCompanyCreationRequest) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RutterCompanyCreationRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler RutterCompanyCreationRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RutterCompanyCreationRequest(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RutterCompanyCreationRequest) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type RutterCompanyResponse struct {
	AccessToken string `json:"accessToken" url:"accessToken"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RutterCompanyResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RutterCompanyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RutterCompanyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RutterCompanyResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RutterCompanyResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}
