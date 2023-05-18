// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceDockerhubDockerhubEnum string

const (
	SourceDockerhubDockerhubEnumDockerhub SourceDockerhubDockerhubEnum = "dockerhub"
)

func (e SourceDockerhubDockerhubEnum) ToPointer() *SourceDockerhubDockerhubEnum {
	return &e
}

func (e *SourceDockerhubDockerhubEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dockerhub":
		*e = SourceDockerhubDockerhubEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDockerhubDockerhubEnum: %v", v)
	}
}

type SourceDockerhub struct {
	// Username of DockerHub person or organization (for https://hub.docker.com/v2/repositories/USERNAME/ API call)
	DockerUsername string                       `json:"docker_username"`
	SourceType     SourceDockerhubDockerhubEnum `json:"sourceType"`
}
