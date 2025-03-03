// This file was auto-generated by Fern from our API Definition.

package entity

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/mercoa-finance/go/internal"
)

type SyncExternalSystemRequest struct {
	// Sync vendors from external accounting system. Default is to pull vendors from external system.
	Vendors *SyncType `json:"-" url:"vendors,omitempty"`
	// Sync bills from external accounting system. Default is to not sync bills. Invoices that already exist in both systems will not be updated, only new invoices not present in the other system will be created.
	Bills *SyncType `json:"-" url:"bills,omitempty"`
	// Sync GL accounts from external accounting system. Default is to pull GL accounts from external system. Pushing GL accounts is not supported.
	GlAccounts *SyncType `json:"-" url:"glAccounts,omitempty"`
}

type CodatCompanyCreationRequest struct {
	// If the company already exists in Codat, provide the company ID to link the company to the entity. If the company does not exist, leave this field blank and Codat will create a new company.
	CompanyID *string `json:"companyId,omitempty" url:"companyId,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CodatCompanyCreationRequest) GetCompanyID() *string {
	if c == nil {
		return nil
	}
	return c.CompanyID
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
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CodatCompanyCreationRequest) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CodatCompanyResponse struct {
	CompanyID string `json:"companyId" url:"companyId"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CodatCompanyResponse) GetCompanyID() string {
	if c == nil {
		return ""
	}
	return c.CompanyID
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
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CodatCompanyResponse) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ExternalAccountingSystemCompanyCreationRequest struct {
	Type   string
	Codat  *CodatCompanyCreationRequest
	Rutter *RutterCompanyCreationRequest
}

func (e *ExternalAccountingSystemCompanyCreationRequest) GetType() string {
	if e == nil {
		return ""
	}
	return e.Type
}

func (e *ExternalAccountingSystemCompanyCreationRequest) GetCodat() *CodatCompanyCreationRequest {
	if e == nil {
		return nil
	}
	return e.Codat
}

func (e *ExternalAccountingSystemCompanyCreationRequest) GetRutter() *RutterCompanyCreationRequest {
	if e == nil {
		return nil
	}
	return e.Rutter
}

func (e *ExternalAccountingSystemCompanyCreationRequest) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", e)
	}
	switch unmarshaler.Type {
	case "codat":
		value := new(CodatCompanyCreationRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Codat = value
	case "rutter":
		value := new(RutterCompanyCreationRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Rutter = value
	}
	return nil
}

func (e ExternalAccountingSystemCompanyCreationRequest) MarshalJSON() ([]byte, error) {
	if err := e.validate(); err != nil {
		return nil, err
	}
	if e.Codat != nil {
		return internal.MarshalJSONWithExtraProperty(e.Codat, "type", "codat")
	}
	if e.Rutter != nil {
		return internal.MarshalJSONWithExtraProperty(e.Rutter, "type", "rutter")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", e)
}

type ExternalAccountingSystemCompanyCreationRequestVisitor interface {
	VisitCodat(*CodatCompanyCreationRequest) error
	VisitRutter(*RutterCompanyCreationRequest) error
}

func (e *ExternalAccountingSystemCompanyCreationRequest) Accept(visitor ExternalAccountingSystemCompanyCreationRequestVisitor) error {
	if e.Codat != nil {
		return visitor.VisitCodat(e.Codat)
	}
	if e.Rutter != nil {
		return visitor.VisitRutter(e.Rutter)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", e)
}

func (e *ExternalAccountingSystemCompanyCreationRequest) validate() error {
	if e == nil {
		return fmt.Errorf("type %T is nil", e)
	}
	var fields []string
	if e.Codat != nil {
		fields = append(fields, "codat")
	}
	if e.Rutter != nil {
		fields = append(fields, "rutter")
	}
	if len(fields) == 0 {
		if e.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", e, e.Type)
		}
		return fmt.Errorf("type %T is empty", e)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", e, fields)
	}
	if e.Type != "" {
		field := fields[0]
		if e.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				e,
				e.Type,
				e,
			)
		}
	}
	return nil
}

type ExternalAccountingSystemCompanyResponse struct {
	Type   string
	Codat  *CodatCompanyResponse
	Rutter *RutterCompanyResponse
	None   *CodatCompanyResponse
}

func (e *ExternalAccountingSystemCompanyResponse) GetType() string {
	if e == nil {
		return ""
	}
	return e.Type
}

func (e *ExternalAccountingSystemCompanyResponse) GetCodat() *CodatCompanyResponse {
	if e == nil {
		return nil
	}
	return e.Codat
}

func (e *ExternalAccountingSystemCompanyResponse) GetRutter() *RutterCompanyResponse {
	if e == nil {
		return nil
	}
	return e.Rutter
}

func (e *ExternalAccountingSystemCompanyResponse) GetNone() *CodatCompanyResponse {
	if e == nil {
		return nil
	}
	return e.None
}

func (e *ExternalAccountingSystemCompanyResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", e)
	}
	switch unmarshaler.Type {
	case "codat":
		value := new(CodatCompanyResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Codat = value
	case "rutter":
		value := new(RutterCompanyResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.Rutter = value
	case "none":
		value := new(CodatCompanyResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		e.None = value
	}
	return nil
}

func (e ExternalAccountingSystemCompanyResponse) MarshalJSON() ([]byte, error) {
	if err := e.validate(); err != nil {
		return nil, err
	}
	if e.Codat != nil {
		return internal.MarshalJSONWithExtraProperty(e.Codat, "type", "codat")
	}
	if e.Rutter != nil {
		return internal.MarshalJSONWithExtraProperty(e.Rutter, "type", "rutter")
	}
	if e.None != nil {
		return internal.MarshalJSONWithExtraProperty(e.None, "type", "none")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", e)
}

type ExternalAccountingSystemCompanyResponseVisitor interface {
	VisitCodat(*CodatCompanyResponse) error
	VisitRutter(*RutterCompanyResponse) error
	VisitNone(*CodatCompanyResponse) error
}

func (e *ExternalAccountingSystemCompanyResponse) Accept(visitor ExternalAccountingSystemCompanyResponseVisitor) error {
	if e.Codat != nil {
		return visitor.VisitCodat(e.Codat)
	}
	if e.Rutter != nil {
		return visitor.VisitRutter(e.Rutter)
	}
	if e.None != nil {
		return visitor.VisitNone(e.None)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", e)
}

func (e *ExternalAccountingSystemCompanyResponse) validate() error {
	if e == nil {
		return fmt.Errorf("type %T is nil", e)
	}
	var fields []string
	if e.Codat != nil {
		fields = append(fields, "codat")
	}
	if e.Rutter != nil {
		fields = append(fields, "rutter")
	}
	if e.None != nil {
		fields = append(fields, "none")
	}
	if len(fields) == 0 {
		if e.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", e, e.Type)
		}
		return fmt.Errorf("type %T is empty", e)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", e, fields)
	}
	if e.Type != "" {
		field := fields[0]
		if e.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				e,
				e.Type,
				e,
			)
		}
	}
	return nil
}

type RutterCompanyCreationRequest struct {
	// The access token for the existing Rutter connection. If the connection does not exist, leave this field blank and Rutter will create a new connection.
	AccessToken *string `json:"accessToken,omitempty" url:"accessToken,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RutterCompanyCreationRequest) GetAccessToken() *string {
	if r == nil {
		return nil
	}
	return r.AccessToken
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
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RutterCompanyCreationRequest) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type RutterCompanyResponse struct {
	AccessToken string `json:"accessToken" url:"accessToken"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (r *RutterCompanyResponse) GetAccessToken() string {
	if r == nil {
		return ""
	}
	return r.AccessToken
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
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	r.rawJSON = json.RawMessage(data)
	return nil
}

func (r *RutterCompanyResponse) String() string {
	if len(r.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(r.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type SyncType string

const (
	SyncTypeNone SyncType = "none"
	SyncTypePush SyncType = "push"
	SyncTypePull SyncType = "pull"
	SyncTypeBoth SyncType = "both"
)

func NewSyncTypeFromString(s string) (SyncType, error) {
	switch s {
	case "none":
		return SyncTypeNone, nil
	case "push":
		return SyncTypePush, nil
	case "pull":
		return SyncTypePull, nil
	case "both":
		return SyncTypeBoth, nil
	}
	var t SyncType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SyncType) Ptr() *SyncType {
	return &s
}
