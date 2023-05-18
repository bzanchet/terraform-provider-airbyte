// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum string

const (
	DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnumLoginPassword DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum = "login/password"
)

func (e DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum) ToPointer() *DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum {
	return &e
}

func (e *DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "login/password":
		*e = DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum: %v", v)
	}
}

// DestinationMongodbAuthorizationTypeLoginPassword - Login/Password.
type DestinationMongodbAuthorizationTypeLoginPassword struct {
	Authorization DestinationMongodbAuthorizationTypeLoginPasswordAuthorizationEnum `json:"authorization"`
	// Password associated with the username.
	Password string `json:"password"`
	// Username to use to access the database.
	Username string `json:"username"`
}

type DestinationMongodbAuthorizationTypeNoneAuthorizationEnum string

const (
	DestinationMongodbAuthorizationTypeNoneAuthorizationEnumNone DestinationMongodbAuthorizationTypeNoneAuthorizationEnum = "none"
)

func (e DestinationMongodbAuthorizationTypeNoneAuthorizationEnum) ToPointer() *DestinationMongodbAuthorizationTypeNoneAuthorizationEnum {
	return &e
}

func (e *DestinationMongodbAuthorizationTypeNoneAuthorizationEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = DestinationMongodbAuthorizationTypeNoneAuthorizationEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbAuthorizationTypeNoneAuthorizationEnum: %v", v)
	}
}

// DestinationMongodbAuthorizationTypeNone - None.
type DestinationMongodbAuthorizationTypeNone struct {
	Authorization DestinationMongodbAuthorizationTypeNoneAuthorizationEnum `json:"authorization"`
}

type DestinationMongodbAuthorizationTypeType string

const (
	DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeNone          DestinationMongodbAuthorizationTypeType = "destination-mongodb_Authorization type_None"
	DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeLoginPassword DestinationMongodbAuthorizationTypeType = "destination-mongodb_Authorization type_Login/Password"
)

type DestinationMongodbAuthorizationType struct {
	DestinationMongodbAuthorizationTypeNone          *DestinationMongodbAuthorizationTypeNone
	DestinationMongodbAuthorizationTypeLoginPassword *DestinationMongodbAuthorizationTypeLoginPassword

	Type DestinationMongodbAuthorizationTypeType
}

func CreateDestinationMongodbAuthorizationTypeDestinationMongodbAuthorizationTypeNone(destinationMongodbAuthorizationTypeNone DestinationMongodbAuthorizationTypeNone) DestinationMongodbAuthorizationType {
	typ := DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeNone

	return DestinationMongodbAuthorizationType{
		DestinationMongodbAuthorizationTypeNone: &destinationMongodbAuthorizationTypeNone,
		Type:                                    typ,
	}
}

func CreateDestinationMongodbAuthorizationTypeDestinationMongodbAuthorizationTypeLoginPassword(destinationMongodbAuthorizationTypeLoginPassword DestinationMongodbAuthorizationTypeLoginPassword) DestinationMongodbAuthorizationType {
	typ := DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeLoginPassword

	return DestinationMongodbAuthorizationType{
		DestinationMongodbAuthorizationTypeLoginPassword: &destinationMongodbAuthorizationTypeLoginPassword,
		Type: typ,
	}
}

func (u *DestinationMongodbAuthorizationType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbAuthorizationTypeNone := new(DestinationMongodbAuthorizationTypeNone)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbAuthorizationTypeNone); err == nil {
		u.DestinationMongodbAuthorizationTypeNone = destinationMongodbAuthorizationTypeNone
		u.Type = DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeNone
		return nil
	}

	destinationMongodbAuthorizationTypeLoginPassword := new(DestinationMongodbAuthorizationTypeLoginPassword)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbAuthorizationTypeLoginPassword); err == nil {
		u.DestinationMongodbAuthorizationTypeLoginPassword = destinationMongodbAuthorizationTypeLoginPassword
		u.Type = DestinationMongodbAuthorizationTypeTypeDestinationMongodbAuthorizationTypeLoginPassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbAuthorizationType) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbAuthorizationTypeNone != nil {
		return json.Marshal(u.DestinationMongodbAuthorizationTypeNone)
	}

	if u.DestinationMongodbAuthorizationTypeLoginPassword != nil {
		return json.Marshal(u.DestinationMongodbAuthorizationTypeLoginPassword)
	}

	return nil, nil
}

type DestinationMongodbMongodbEnum string

const (
	DestinationMongodbMongodbEnumMongodb DestinationMongodbMongodbEnum = "mongodb"
)

func (e DestinationMongodbMongodbEnum) ToPointer() *DestinationMongodbMongodbEnum {
	return &e
}

func (e *DestinationMongodbMongodbEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mongodb":
		*e = DestinationMongodbMongodbEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongodbEnum: %v", v)
	}
}

type DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum string

const (
	DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnumAtlas DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum = "atlas"
)

func (e DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum) ToPointer() *DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum {
	return &e
}

func (e *DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "atlas":
		*e = DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum: %v", v)
	}
}

// DestinationMongodbMongoDBInstanceTypeMongoDBAtlas - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbMongoDBInstanceTypeMongoDBAtlas struct {
	// URL of a cluster to connect to.
	ClusterURL string                                                        `json:"cluster_url"`
	Instance   DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstanceEnum `json:"instance"`
}

type DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum string

const (
	DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnumReplica DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum = "replica"
)

func (e DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum) ToPointer() *DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum {
	return &e
}

func (e *DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "replica":
		*e = DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum: %v", v)
	}
}

// DestinationMongodbMongoDbInstanceTypeReplicaSet - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbMongoDbInstanceTypeReplicaSet struct {
	Instance DestinationMongodbMongoDbInstanceTypeReplicaSetInstanceEnum `json:"instance"`
	// A replica set name.
	ReplicaSet *string `json:"replica_set,omitempty"`
	// The members of a replica set. Please specify `host`:`port` of each member seperated by comma.
	ServerAddresses string `json:"server_addresses"`
}

type DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum string

const (
	DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnumStandalone DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum = "standalone"
)

func (e DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum) ToPointer() *DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum {
	return &e
}

func (e *DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standalone":
		*e = DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum: %v", v)
	}
}

// DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance struct {
	// The Host of a Mongo database to be replicated.
	Host     string                                                                     `json:"host"`
	Instance DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstanceEnum `json:"instance"`
	// The Port of a Mongo database to be replicated.
	Port int64 `json:"port"`
}

type DestinationMongodbMongoDbInstanceTypeType string

const (
	DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbMongoDbInstanceTypeType = "destination-mongodb_MongoDb Instance Type_Standalone MongoDb Instance"
	DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeReplicaSet                DestinationMongodbMongoDbInstanceTypeType = "destination-mongodb_MongoDb Instance Type_Replica Set"
	DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas              DestinationMongodbMongoDbInstanceTypeType = "destination-mongodb_MongoDb Instance Type_MongoDB Atlas"
)

type DestinationMongodbMongoDbInstanceType struct {
	DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance *DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
	DestinationMongodbMongoDbInstanceTypeReplicaSet                *DestinationMongodbMongoDbInstanceTypeReplicaSet
	DestinationMongodbMongoDBInstanceTypeMongoDBAtlas              *DestinationMongodbMongoDBInstanceTypeMongoDBAtlas

	Type DestinationMongodbMongoDbInstanceTypeType
}

func CreateDestinationMongodbMongoDbInstanceTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance(destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance) DestinationMongodbMongoDbInstanceType {
	typ := DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance

	return DestinationMongodbMongoDbInstanceType{
		DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance: &destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance,
		Type: typ,
	}
}

func CreateDestinationMongodbMongoDbInstanceTypeDestinationMongodbMongoDbInstanceTypeReplicaSet(destinationMongodbMongoDbInstanceTypeReplicaSet DestinationMongodbMongoDbInstanceTypeReplicaSet) DestinationMongodbMongoDbInstanceType {
	typ := DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeReplicaSet

	return DestinationMongodbMongoDbInstanceType{
		DestinationMongodbMongoDbInstanceTypeReplicaSet: &destinationMongodbMongoDbInstanceTypeReplicaSet,
		Type: typ,
	}
}

func CreateDestinationMongodbMongoDbInstanceTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas(destinationMongodbMongoDBInstanceTypeMongoDBAtlas DestinationMongodbMongoDBInstanceTypeMongoDBAtlas) DestinationMongodbMongoDbInstanceType {
	typ := DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas

	return DestinationMongodbMongoDbInstanceType{
		DestinationMongodbMongoDBInstanceTypeMongoDBAtlas: &destinationMongodbMongoDBInstanceTypeMongoDBAtlas,
		Type: typ,
	}
}

func (u *DestinationMongodbMongoDbInstanceType) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance := new(DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance); err == nil {
		u.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance = destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
		u.Type = DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
		return nil
	}

	destinationMongodbMongoDbInstanceTypeReplicaSet := new(DestinationMongodbMongoDbInstanceTypeReplicaSet)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbMongoDbInstanceTypeReplicaSet); err == nil {
		u.DestinationMongodbMongoDbInstanceTypeReplicaSet = destinationMongodbMongoDbInstanceTypeReplicaSet
		u.Type = DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDbInstanceTypeReplicaSet
		return nil
	}

	destinationMongodbMongoDBInstanceTypeMongoDBAtlas := new(DestinationMongodbMongoDBInstanceTypeMongoDBAtlas)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbMongoDBInstanceTypeMongoDBAtlas); err == nil {
		u.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas = destinationMongodbMongoDBInstanceTypeMongoDBAtlas
		u.Type = DestinationMongodbMongoDbInstanceTypeTypeDestinationMongodbMongoDBInstanceTypeMongoDBAtlas
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbMongoDbInstanceType) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		return json.Marshal(u.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance)
	}

	if u.DestinationMongodbMongoDbInstanceTypeReplicaSet != nil {
		return json.Marshal(u.DestinationMongodbMongoDbInstanceTypeReplicaSet)
	}

	if u.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
		return json.Marshal(u.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas)
	}

	return nil, nil
}

// DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and password authentication
type DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum string

const (
	DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnumSSHPasswordAuth DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum = "SSH_PASSWORD_AUTH"
)

func (e DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) ToPointer() *DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum {
	return &e
}

func (e *DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum: %v", v)
	}
}

// DestinationMongodbSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	TunnelMethod DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

// DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum - Connect through a jump server tunnel host using username and ssh key
type DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum string

const (
	DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnumSSHKeyAuth DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum = "SSH_KEY_AUTH"
)

func (e DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) ToPointer() *DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum {
	return &e
}

func (e *DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum: %v", v)
	}
}

// DestinationMongodbSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	TunnelMethod DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethodEnum `json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort int64 `json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

// DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum - No ssh tunnel needed to connect to database
type DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum string

const (
	DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnumNoTunnel DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum = "NO_TUNNEL"
)

func (e DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum) ToPointer() *DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum {
	return &e
}

func (e *DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum: %v", v)
	}
}

// DestinationMongodbSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	TunnelMethod DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethodEnum `json:"tunnel_method"`
}

type DestinationMongodbSSHTunnelMethodType string

const (
	DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodNoTunnel               DestinationMongodbSSHTunnelMethodType = "destination-mongodb_SSH Tunnel Method_No Tunnel"
	DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodSSHKeyAuthentication   DestinationMongodbSSHTunnelMethodType = "destination-mongodb_SSH Tunnel Method_SSH Key Authentication"
	DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodPasswordAuthentication DestinationMongodbSSHTunnelMethodType = "destination-mongodb_SSH Tunnel Method_Password Authentication"
)

type DestinationMongodbSSHTunnelMethod struct {
	DestinationMongodbSSHTunnelMethodNoTunnel               *DestinationMongodbSSHTunnelMethodNoTunnel
	DestinationMongodbSSHTunnelMethodSSHKeyAuthentication   *DestinationMongodbSSHTunnelMethodSSHKeyAuthentication
	DestinationMongodbSSHTunnelMethodPasswordAuthentication *DestinationMongodbSSHTunnelMethodPasswordAuthentication

	Type DestinationMongodbSSHTunnelMethodType
}

func CreateDestinationMongodbSSHTunnelMethodDestinationMongodbSSHTunnelMethodNoTunnel(destinationMongodbSSHTunnelMethodNoTunnel DestinationMongodbSSHTunnelMethodNoTunnel) DestinationMongodbSSHTunnelMethod {
	typ := DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodNoTunnel

	return DestinationMongodbSSHTunnelMethod{
		DestinationMongodbSSHTunnelMethodNoTunnel: &destinationMongodbSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationMongodbSSHTunnelMethodDestinationMongodbSSHTunnelMethodSSHKeyAuthentication(destinationMongodbSSHTunnelMethodSSHKeyAuthentication DestinationMongodbSSHTunnelMethodSSHKeyAuthentication) DestinationMongodbSSHTunnelMethod {
	typ := DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodSSHKeyAuthentication

	return DestinationMongodbSSHTunnelMethod{
		DestinationMongodbSSHTunnelMethodSSHKeyAuthentication: &destinationMongodbSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMongodbSSHTunnelMethodDestinationMongodbSSHTunnelMethodPasswordAuthentication(destinationMongodbSSHTunnelMethodPasswordAuthentication DestinationMongodbSSHTunnelMethodPasswordAuthentication) DestinationMongodbSSHTunnelMethod {
	typ := DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodPasswordAuthentication

	return DestinationMongodbSSHTunnelMethod{
		DestinationMongodbSSHTunnelMethodPasswordAuthentication: &destinationMongodbSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMongodbSSHTunnelMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	destinationMongodbSSHTunnelMethodNoTunnel := new(DestinationMongodbSSHTunnelMethodNoTunnel)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbSSHTunnelMethodNoTunnel); err == nil {
		u.DestinationMongodbSSHTunnelMethodNoTunnel = destinationMongodbSSHTunnelMethodNoTunnel
		u.Type = DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMongodbSSHTunnelMethodSSHKeyAuthentication := new(DestinationMongodbSSHTunnelMethodSSHKeyAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbSSHTunnelMethodSSHKeyAuthentication); err == nil {
		u.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication = destinationMongodbSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMongodbSSHTunnelMethodPasswordAuthentication := new(DestinationMongodbSSHTunnelMethodPasswordAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&destinationMongodbSSHTunnelMethodPasswordAuthentication); err == nil {
		u.DestinationMongodbSSHTunnelMethodPasswordAuthentication = destinationMongodbSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMongodbSSHTunnelMethodTypeDestinationMongodbSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbSSHTunnelMethodNoTunnel != nil {
		return json.Marshal(u.DestinationMongodbSSHTunnelMethodNoTunnel)
	}

	if u.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication != nil {
		return json.Marshal(u.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication)
	}

	if u.DestinationMongodbSSHTunnelMethodPasswordAuthentication != nil {
		return json.Marshal(u.DestinationMongodbSSHTunnelMethodPasswordAuthentication)
	}

	return nil, nil
}

type DestinationMongodb struct {
	// Authorization type.
	AuthType DestinationMongodbAuthorizationType `json:"auth_type"`
	// Name of the database.
	Database        string                        `json:"database"`
	DestinationType DestinationMongodbMongodbEnum `json:"destinationType"`
	// MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
	InstanceType *DestinationMongodbMongoDbInstanceType `json:"instance_type,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMongodbSSHTunnelMethod `json:"tunnel_method,omitempty"`
}
