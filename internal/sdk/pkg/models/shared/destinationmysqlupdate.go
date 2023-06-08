// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationMysqlUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod `json:"tunnel_method"`
}

type DestinationMysqlUpdateSSHTunnelMethodType string

const (
	DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodNoTunnel               DestinationMysqlUpdateSSHTunnelMethodType = "destination-mysql-update_SSH Tunnel Method_No Tunnel"
	DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication   DestinationMysqlUpdateSSHTunnelMethodType = "destination-mysql-update_SSH Tunnel Method_SSH Key Authentication"
	DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication DestinationMysqlUpdateSSHTunnelMethodType = "destination-mysql-update_SSH Tunnel Method_Password Authentication"
)

type DestinationMysqlUpdateSSHTunnelMethod struct {
	DestinationMysqlUpdateSSHTunnelMethodNoTunnel               *DestinationMysqlUpdateSSHTunnelMethodNoTunnel
	DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
	DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication *DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication

	Type DestinationMysqlUpdateSSHTunnelMethodType
}

func CreateDestinationMysqlUpdateSSHTunnelMethodDestinationMysqlUpdateSSHTunnelMethodNoTunnel(destinationMysqlUpdateSSHTunnelMethodNoTunnel DestinationMysqlUpdateSSHTunnelMethodNoTunnel) DestinationMysqlUpdateSSHTunnelMethod {
	typ := DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodNoTunnel

	return DestinationMysqlUpdateSSHTunnelMethod{
		DestinationMysqlUpdateSSHTunnelMethodNoTunnel: &destinationMysqlUpdateSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationMysqlUpdateSSHTunnelMethodDestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication(destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) DestinationMysqlUpdateSSHTunnelMethod {
	typ := DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication

	return DestinationMysqlUpdateSSHTunnelMethod{
		DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication: &destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMysqlUpdateSSHTunnelMethodDestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication(destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication) DestinationMysqlUpdateSSHTunnelMethod {
	typ := DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication

	return DestinationMysqlUpdateSSHTunnelMethod{
		DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication: &destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMysqlUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMysqlUpdateSSHTunnelMethodNoTunnel := new(DestinationMysqlUpdateSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMysqlUpdateSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationMysqlUpdateSSHTunnelMethodNoTunnel = destinationMysqlUpdateSSHTunnelMethodNoTunnel
		u.Type = DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication := new(DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication = destinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication := new(DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication = destinationMysqlUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMysqlUpdateSSHTunnelMethodTypeDestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMysqlUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMysqlUpdateSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationMysqlUpdateSSHTunnelMethodNoTunnel)
	}

	if u.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationMysqlUpdateSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationMysqlUpdateSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationMysqlUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port int64 `json:"port"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMysqlUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}