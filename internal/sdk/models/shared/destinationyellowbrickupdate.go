// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode string

const (
	DestinationYellowbrickUpdateSchemasSSLModeSSLModes6ModeVerifyFull DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode = "verify-full"
)

func (e DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode) ToPointer() *DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode: %v", v)
	}
}

// DestinationYellowbrickUpdateVerifyFull - Verify-full SSL mode.
type DestinationYellowbrickUpdateVerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                                  `json:"client_key_password,omitempty"`
	mode              *DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode `const:"verify-full" json:"mode"`
}

func (d DestinationYellowbrickUpdateVerifyFull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateVerifyFull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateVerifyFull) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *DestinationYellowbrickUpdateVerifyFull) GetClientCertificate() string {
	if o == nil {
		return ""
	}
	return o.ClientCertificate
}

func (o *DestinationYellowbrickUpdateVerifyFull) GetClientKey() string {
	if o == nil {
		return ""
	}
	return o.ClientKey
}

func (o *DestinationYellowbrickUpdateVerifyFull) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *DestinationYellowbrickUpdateVerifyFull) GetMode() *DestinationYellowbrickUpdateSchemasSSLModeSSLModes6Mode {
	return DestinationYellowbrickUpdateSchemasSSLModeSSLModes6ModeVerifyFull.ToPointer()
}

type DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode string

const (
	DestinationYellowbrickUpdateSchemasSSLModeSSLModes5ModeVerifyCa DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode = "verify-ca"
)

func (e DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode) ToPointer() *DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-ca":
		*e = DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode: %v", v)
	}
}

// DestinationYellowbrickUpdateVerifyCa - Verify-ca SSL mode.
type DestinationYellowbrickUpdateVerifyCa struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                                  `json:"client_key_password,omitempty"`
	mode              *DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode `const:"verify-ca" json:"mode"`
}

func (d DestinationYellowbrickUpdateVerifyCa) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateVerifyCa) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateVerifyCa) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *DestinationYellowbrickUpdateVerifyCa) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *DestinationYellowbrickUpdateVerifyCa) GetMode() *DestinationYellowbrickUpdateSchemasSSLModeSSLModes5Mode {
	return DestinationYellowbrickUpdateSchemasSSLModeSSLModes5ModeVerifyCa.ToPointer()
}

type DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode string

const (
	DestinationYellowbrickUpdateSchemasSSLModeSSLModesModeRequire DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode = "require"
)

func (e DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode) ToPointer() *DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		*e = DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode: %v", v)
	}
}

// DestinationYellowbrickUpdateRequire - Require SSL mode.
type DestinationYellowbrickUpdateRequire struct {
	mode *DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode `const:"require" json:"mode"`
}

func (d DestinationYellowbrickUpdateRequire) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateRequire) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateRequire) GetMode() *DestinationYellowbrickUpdateSchemasSSLModeSSLModesMode {
	return DestinationYellowbrickUpdateSchemasSSLModeSSLModesModeRequire.ToPointer()
}

type DestinationYellowbrickUpdateSchemasSslModeMode string

const (
	DestinationYellowbrickUpdateSchemasSslModeModePrefer DestinationYellowbrickUpdateSchemasSslModeMode = "prefer"
)

func (e DestinationYellowbrickUpdateSchemasSslModeMode) ToPointer() *DestinationYellowbrickUpdateSchemasSslModeMode {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasSslModeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prefer":
		*e = DestinationYellowbrickUpdateSchemasSslModeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasSslModeMode: %v", v)
	}
}

// DestinationYellowbrickUpdatePrefer - Prefer SSL mode.
type DestinationYellowbrickUpdatePrefer struct {
	mode *DestinationYellowbrickUpdateSchemasSslModeMode `const:"prefer" json:"mode"`
}

func (d DestinationYellowbrickUpdatePrefer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdatePrefer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdatePrefer) GetMode() *DestinationYellowbrickUpdateSchemasSslModeMode {
	return DestinationYellowbrickUpdateSchemasSslModeModePrefer.ToPointer()
}

type DestinationYellowbrickUpdateSchemasMode string

const (
	DestinationYellowbrickUpdateSchemasModeAllow DestinationYellowbrickUpdateSchemasMode = "allow"
)

func (e DestinationYellowbrickUpdateSchemasMode) ToPointer() *DestinationYellowbrickUpdateSchemasMode {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		*e = DestinationYellowbrickUpdateSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasMode: %v", v)
	}
}

// DestinationYellowbrickUpdateAllow - Allow SSL mode.
type DestinationYellowbrickUpdateAllow struct {
	mode *DestinationYellowbrickUpdateSchemasMode `const:"allow" json:"mode"`
}

func (d DestinationYellowbrickUpdateAllow) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateAllow) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateAllow) GetMode() *DestinationYellowbrickUpdateSchemasMode {
	return DestinationYellowbrickUpdateSchemasModeAllow.ToPointer()
}

type DestinationYellowbrickUpdateMode string

const (
	DestinationYellowbrickUpdateModeDisable DestinationYellowbrickUpdateMode = "disable"
)

func (e DestinationYellowbrickUpdateMode) ToPointer() *DestinationYellowbrickUpdateMode {
	return &e
}
func (e *DestinationYellowbrickUpdateMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = DestinationYellowbrickUpdateMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateMode: %v", v)
	}
}

// DestinationYellowbrickUpdateDisable - Disable SSL.
type DestinationYellowbrickUpdateDisable struct {
	mode *DestinationYellowbrickUpdateMode `const:"disable" json:"mode"`
}

func (d DestinationYellowbrickUpdateDisable) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateDisable) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateDisable) GetMode() *DestinationYellowbrickUpdateMode {
	return DestinationYellowbrickUpdateModeDisable.ToPointer()
}

type DestinationYellowbrickUpdateSSLModesType string

const (
	DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateDisable    DestinationYellowbrickUpdateSSLModesType = "destination-yellowbrick-update_disable"
	DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateAllow      DestinationYellowbrickUpdateSSLModesType = "destination-yellowbrick-update_allow"
	DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdatePrefer     DestinationYellowbrickUpdateSSLModesType = "destination-yellowbrick-update_prefer"
	DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateRequire    DestinationYellowbrickUpdateSSLModesType = "destination-yellowbrick-update_require"
	DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateVerifyCa   DestinationYellowbrickUpdateSSLModesType = "destination-yellowbrick-update_verify-ca"
	DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateVerifyFull DestinationYellowbrickUpdateSSLModesType = "destination-yellowbrick-update_verify-full"
)

// DestinationYellowbrickUpdateSSLModes - SSL connection modes.
//
//	<b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
//	<b>allow</b> - Chose this mode to enable encryption only when required by the source database
//	<b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
//	<b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
//	 <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
//	 <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
//	See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
type DestinationYellowbrickUpdateSSLModes struct {
	DestinationYellowbrickUpdateDisable    *DestinationYellowbrickUpdateDisable
	DestinationYellowbrickUpdateAllow      *DestinationYellowbrickUpdateAllow
	DestinationYellowbrickUpdatePrefer     *DestinationYellowbrickUpdatePrefer
	DestinationYellowbrickUpdateRequire    *DestinationYellowbrickUpdateRequire
	DestinationYellowbrickUpdateVerifyCa   *DestinationYellowbrickUpdateVerifyCa
	DestinationYellowbrickUpdateVerifyFull *DestinationYellowbrickUpdateVerifyFull

	Type DestinationYellowbrickUpdateSSLModesType
}

func CreateDestinationYellowbrickUpdateSSLModesDestinationYellowbrickUpdateDisable(destinationYellowbrickUpdateDisable DestinationYellowbrickUpdateDisable) DestinationYellowbrickUpdateSSLModes {
	typ := DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateDisable

	return DestinationYellowbrickUpdateSSLModes{
		DestinationYellowbrickUpdateDisable: &destinationYellowbrickUpdateDisable,
		Type:                                typ,
	}
}

func CreateDestinationYellowbrickUpdateSSLModesDestinationYellowbrickUpdateAllow(destinationYellowbrickUpdateAllow DestinationYellowbrickUpdateAllow) DestinationYellowbrickUpdateSSLModes {
	typ := DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateAllow

	return DestinationYellowbrickUpdateSSLModes{
		DestinationYellowbrickUpdateAllow: &destinationYellowbrickUpdateAllow,
		Type:                              typ,
	}
}

func CreateDestinationYellowbrickUpdateSSLModesDestinationYellowbrickUpdatePrefer(destinationYellowbrickUpdatePrefer DestinationYellowbrickUpdatePrefer) DestinationYellowbrickUpdateSSLModes {
	typ := DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdatePrefer

	return DestinationYellowbrickUpdateSSLModes{
		DestinationYellowbrickUpdatePrefer: &destinationYellowbrickUpdatePrefer,
		Type:                               typ,
	}
}

func CreateDestinationYellowbrickUpdateSSLModesDestinationYellowbrickUpdateRequire(destinationYellowbrickUpdateRequire DestinationYellowbrickUpdateRequire) DestinationYellowbrickUpdateSSLModes {
	typ := DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateRequire

	return DestinationYellowbrickUpdateSSLModes{
		DestinationYellowbrickUpdateRequire: &destinationYellowbrickUpdateRequire,
		Type:                                typ,
	}
}

func CreateDestinationYellowbrickUpdateSSLModesDestinationYellowbrickUpdateVerifyCa(destinationYellowbrickUpdateVerifyCa DestinationYellowbrickUpdateVerifyCa) DestinationYellowbrickUpdateSSLModes {
	typ := DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateVerifyCa

	return DestinationYellowbrickUpdateSSLModes{
		DestinationYellowbrickUpdateVerifyCa: &destinationYellowbrickUpdateVerifyCa,
		Type:                                 typ,
	}
}

func CreateDestinationYellowbrickUpdateSSLModesDestinationYellowbrickUpdateVerifyFull(destinationYellowbrickUpdateVerifyFull DestinationYellowbrickUpdateVerifyFull) DestinationYellowbrickUpdateSSLModes {
	typ := DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateVerifyFull

	return DestinationYellowbrickUpdateSSLModes{
		DestinationYellowbrickUpdateVerifyFull: &destinationYellowbrickUpdateVerifyFull,
		Type:                                   typ,
	}
}

func (u *DestinationYellowbrickUpdateSSLModes) UnmarshalJSON(data []byte) error {

	var destinationYellowbrickUpdateDisable DestinationYellowbrickUpdateDisable = DestinationYellowbrickUpdateDisable{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateDisable, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateDisable = &destinationYellowbrickUpdateDisable
		u.Type = DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateDisable
		return nil
	}

	var destinationYellowbrickUpdateAllow DestinationYellowbrickUpdateAllow = DestinationYellowbrickUpdateAllow{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateAllow, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateAllow = &destinationYellowbrickUpdateAllow
		u.Type = DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateAllow
		return nil
	}

	var destinationYellowbrickUpdatePrefer DestinationYellowbrickUpdatePrefer = DestinationYellowbrickUpdatePrefer{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdatePrefer, "", true, true); err == nil {
		u.DestinationYellowbrickUpdatePrefer = &destinationYellowbrickUpdatePrefer
		u.Type = DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdatePrefer
		return nil
	}

	var destinationYellowbrickUpdateRequire DestinationYellowbrickUpdateRequire = DestinationYellowbrickUpdateRequire{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateRequire, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateRequire = &destinationYellowbrickUpdateRequire
		u.Type = DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateRequire
		return nil
	}

	var destinationYellowbrickUpdateVerifyCa DestinationYellowbrickUpdateVerifyCa = DestinationYellowbrickUpdateVerifyCa{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateVerifyCa, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateVerifyCa = &destinationYellowbrickUpdateVerifyCa
		u.Type = DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateVerifyCa
		return nil
	}

	var destinationYellowbrickUpdateVerifyFull DestinationYellowbrickUpdateVerifyFull = DestinationYellowbrickUpdateVerifyFull{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateVerifyFull, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateVerifyFull = &destinationYellowbrickUpdateVerifyFull
		u.Type = DestinationYellowbrickUpdateSSLModesTypeDestinationYellowbrickUpdateVerifyFull
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationYellowbrickUpdateSSLModes", string(data))
}

func (u DestinationYellowbrickUpdateSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationYellowbrickUpdateDisable != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateDisable, "", true)
	}

	if u.DestinationYellowbrickUpdateAllow != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateAllow, "", true)
	}

	if u.DestinationYellowbrickUpdatePrefer != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdatePrefer, "", true)
	}

	if u.DestinationYellowbrickUpdateRequire != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateRequire, "", true)
	}

	if u.DestinationYellowbrickUpdateVerifyCa != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateVerifyCa, "", true)
	}

	if u.DestinationYellowbrickUpdateVerifyFull != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateVerifyFull, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationYellowbrickUpdateSSLModes: all fields are null")
}

// DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationYellowbrickUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationYellowbrickUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationYellowbrickUpdatePasswordAuthentication) GetTunnelMethod() DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationYellowbrickUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationYellowbrickUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationYellowbrickUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationYellowbrickUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationYellowbrickUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationYellowbrickUpdateSchemasTunnelMethod string

const (
	DestinationYellowbrickUpdateSchemasTunnelMethodSSHKeyAuth DestinationYellowbrickUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationYellowbrickUpdateSchemasTunnelMethod) ToPointer() *DestinationYellowbrickUpdateSchemasTunnelMethod {
	return &e
}
func (e *DestinationYellowbrickUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationYellowbrickUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateSchemasTunnelMethod: %v", v)
	}
}

type DestinationYellowbrickUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationYellowbrickUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationYellowbrickUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationYellowbrickUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationYellowbrickUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationYellowbrickUpdateSchemasTunnelMethod {
	return DestinationYellowbrickUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationYellowbrickUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationYellowbrickUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationYellowbrickUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationYellowbrickUpdateTunnelMethod string

const (
	DestinationYellowbrickUpdateTunnelMethodNoTunnel DestinationYellowbrickUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationYellowbrickUpdateTunnelMethod) ToPointer() *DestinationYellowbrickUpdateTunnelMethod {
	return &e
}
func (e *DestinationYellowbrickUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationYellowbrickUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationYellowbrickUpdateTunnelMethod: %v", v)
	}
}

type DestinationYellowbrickUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationYellowbrickUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationYellowbrickUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdateNoTunnel) GetTunnelMethod() DestinationYellowbrickUpdateTunnelMethod {
	return DestinationYellowbrickUpdateTunnelMethodNoTunnel
}

type DestinationYellowbrickUpdateSSHTunnelMethodType string

const (
	DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdateNoTunnel               DestinationYellowbrickUpdateSSHTunnelMethodType = "destination-yellowbrick-update_No Tunnel"
	DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdateSSHKeyAuthentication   DestinationYellowbrickUpdateSSHTunnelMethodType = "destination-yellowbrick-update_SSH Key Authentication"
	DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdatePasswordAuthentication DestinationYellowbrickUpdateSSHTunnelMethodType = "destination-yellowbrick-update_Password Authentication"
)

// DestinationYellowbrickUpdateSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationYellowbrickUpdateSSHTunnelMethod struct {
	DestinationYellowbrickUpdateNoTunnel               *DestinationYellowbrickUpdateNoTunnel
	DestinationYellowbrickUpdateSSHKeyAuthentication   *DestinationYellowbrickUpdateSSHKeyAuthentication
	DestinationYellowbrickUpdatePasswordAuthentication *DestinationYellowbrickUpdatePasswordAuthentication

	Type DestinationYellowbrickUpdateSSHTunnelMethodType
}

func CreateDestinationYellowbrickUpdateSSHTunnelMethodDestinationYellowbrickUpdateNoTunnel(destinationYellowbrickUpdateNoTunnel DestinationYellowbrickUpdateNoTunnel) DestinationYellowbrickUpdateSSHTunnelMethod {
	typ := DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdateNoTunnel

	return DestinationYellowbrickUpdateSSHTunnelMethod{
		DestinationYellowbrickUpdateNoTunnel: &destinationYellowbrickUpdateNoTunnel,
		Type:                                 typ,
	}
}

func CreateDestinationYellowbrickUpdateSSHTunnelMethodDestinationYellowbrickUpdateSSHKeyAuthentication(destinationYellowbrickUpdateSSHKeyAuthentication DestinationYellowbrickUpdateSSHKeyAuthentication) DestinationYellowbrickUpdateSSHTunnelMethod {
	typ := DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdateSSHKeyAuthentication

	return DestinationYellowbrickUpdateSSHTunnelMethod{
		DestinationYellowbrickUpdateSSHKeyAuthentication: &destinationYellowbrickUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationYellowbrickUpdateSSHTunnelMethodDestinationYellowbrickUpdatePasswordAuthentication(destinationYellowbrickUpdatePasswordAuthentication DestinationYellowbrickUpdatePasswordAuthentication) DestinationYellowbrickUpdateSSHTunnelMethod {
	typ := DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdatePasswordAuthentication

	return DestinationYellowbrickUpdateSSHTunnelMethod{
		DestinationYellowbrickUpdatePasswordAuthentication: &destinationYellowbrickUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationYellowbrickUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationYellowbrickUpdateNoTunnel DestinationYellowbrickUpdateNoTunnel = DestinationYellowbrickUpdateNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateNoTunnel = &destinationYellowbrickUpdateNoTunnel
		u.Type = DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdateNoTunnel
		return nil
	}

	var destinationYellowbrickUpdateSSHKeyAuthentication DestinationYellowbrickUpdateSSHKeyAuthentication = DestinationYellowbrickUpdateSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationYellowbrickUpdateSSHKeyAuthentication = &destinationYellowbrickUpdateSSHKeyAuthentication
		u.Type = DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdateSSHKeyAuthentication
		return nil
	}

	var destinationYellowbrickUpdatePasswordAuthentication DestinationYellowbrickUpdatePasswordAuthentication = DestinationYellowbrickUpdatePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationYellowbrickUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationYellowbrickUpdatePasswordAuthentication = &destinationYellowbrickUpdatePasswordAuthentication
		u.Type = DestinationYellowbrickUpdateSSHTunnelMethodTypeDestinationYellowbrickUpdatePasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationYellowbrickUpdateSSHTunnelMethod", string(data))
}

func (u DestinationYellowbrickUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationYellowbrickUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateNoTunnel, "", true)
	}

	if u.DestinationYellowbrickUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationYellowbrickUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationYellowbrickUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationYellowbrickUpdateSSHTunnelMethod: all fields are null")
}

type DestinationYellowbrickUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"5432" json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// Encrypt data using SSL. When activating SSL, please select one of the connection modes.
	Ssl *bool `default:"false" json:"ssl"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the source database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
	//  See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
	SslMode *DestinationYellowbrickUpdateSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationYellowbrickUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationYellowbrickUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationYellowbrickUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationYellowbrickUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationYellowbrickUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationYellowbrickUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationYellowbrickUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationYellowbrickUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationYellowbrickUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationYellowbrickUpdate) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *DestinationYellowbrickUpdate) GetSslMode() *DestinationYellowbrickUpdateSSLModes {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *DestinationYellowbrickUpdate) GetTunnelMethod() *DestinationYellowbrickUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationYellowbrickUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
