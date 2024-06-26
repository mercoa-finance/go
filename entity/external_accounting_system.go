// This file was auto-generated by Fern from our API Definition.

package entity

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mercoa-finance/go/core"
)

type SyncExternalSystemRequest struct {
	// Sync vendors from external accounting system. Default is to pull vendors from external system.
	Vendors *SyncType `json:"-" url:"vendors,omitempty"`
	// Sync bills from external accounting system. Default is to not sync bills. Invoices that already exist in both systems will not be updated, only new invoices not present in the other system will be created.
	Bills *SyncType `json:"-" url:"bills,omitempty"`
	// Sync GL accounts from external accounting system. Default is to pull GL accounts from external system. Pushing GL accounts is not supported.
	GlAccounts *SyncType `json:"-" url:"glAccounts,omitempty"`
}

type ExternalAccountingSystemCompanyCreationRequest struct {
	Type   string
	Codat  *CodatCompanyCreationRequest
	Rutter *RutterCompanyCreationRequest
}

func (e *ExternalAccountingSystemCompanyCreationRequest) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
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
	if e.Codat != nil {
		return core.MarshalJSONWithExtraProperty(e.Codat, "type", "codat")
	}
	if e.Rutter != nil {
		return core.MarshalJSONWithExtraProperty(e.Rutter, "type", "rutter")
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

type ExternalAccountingSystemCompanyResponse struct {
	Type   string
	Codat  *CodatCompanyResponse
	Rutter *RutterCompanyResponse
	None   *CodatCompanyResponse
}

func (e *ExternalAccountingSystemCompanyResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	e.Type = unmarshaler.Type
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
	if e.Codat != nil {
		return core.MarshalJSONWithExtraProperty(e.Codat, "type", "codat")
	}
	if e.Rutter != nil {
		return core.MarshalJSONWithExtraProperty(e.Rutter, "type", "rutter")
	}
	if e.None != nil {
		return core.MarshalJSONWithExtraProperty(e.None, "type", "none")
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
