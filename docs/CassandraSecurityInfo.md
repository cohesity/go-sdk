# CassandraSecurityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraAuthRequired** | Pointer to **NullableBool** | Is Cassandra authentication required ? | [optional] 
**CassandraAuthType** | Pointer to **NullableString** | Cassandra Authentication type. Enum: [PASSWORD KERBEROS LDAP] Specifies the Cassandra auth type. &#39;PASSWORD&#39; &#39;KERBEROS&#39; &#39;LDAP&#39; | [optional] 
**CassandraAuthorizer** | Pointer to **NullableString** | Cassandra Authenticator/Authorizer. | [optional] 
**ClientEncryption** | Pointer to **NullableBool** | Is Client Encryption enabled for this cluster ? | [optional] 
**DseAuthorization** | Pointer to **NullableBool** | Is DSE Authorization enabled for this cluster ? | [optional] 
**ServerEncryptionReqClientAuth** | Pointer to **NullableBool** | Is &#39;Server encryption request client authentication&#39; enabled for this cluster ? | [optional] 
**ServerInternodeEncryptionType** | Pointer to **NullableString** | &#39;Server internal node Encryption&#39; type for this cluster. | [optional] 

## Methods

### NewCassandraSecurityInfo

`func NewCassandraSecurityInfo() *CassandraSecurityInfo`

NewCassandraSecurityInfo instantiates a new CassandraSecurityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraSecurityInfoWithDefaults

`func NewCassandraSecurityInfoWithDefaults() *CassandraSecurityInfo`

NewCassandraSecurityInfoWithDefaults instantiates a new CassandraSecurityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraAuthRequired

`func (o *CassandraSecurityInfo) GetCassandraAuthRequired() bool`

GetCassandraAuthRequired returns the CassandraAuthRequired field if non-nil, zero value otherwise.

### GetCassandraAuthRequiredOk

`func (o *CassandraSecurityInfo) GetCassandraAuthRequiredOk() (*bool, bool)`

GetCassandraAuthRequiredOk returns a tuple with the CassandraAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraAuthRequired

`func (o *CassandraSecurityInfo) SetCassandraAuthRequired(v bool)`

SetCassandraAuthRequired sets CassandraAuthRequired field to given value.

### HasCassandraAuthRequired

`func (o *CassandraSecurityInfo) HasCassandraAuthRequired() bool`

HasCassandraAuthRequired returns a boolean if a field has been set.

### SetCassandraAuthRequiredNil

`func (o *CassandraSecurityInfo) SetCassandraAuthRequiredNil(b bool)`

 SetCassandraAuthRequiredNil sets the value for CassandraAuthRequired to be an explicit nil

### UnsetCassandraAuthRequired
`func (o *CassandraSecurityInfo) UnsetCassandraAuthRequired()`

UnsetCassandraAuthRequired ensures that no value is present for CassandraAuthRequired, not even an explicit nil
### GetCassandraAuthType

`func (o *CassandraSecurityInfo) GetCassandraAuthType() string`

GetCassandraAuthType returns the CassandraAuthType field if non-nil, zero value otherwise.

### GetCassandraAuthTypeOk

`func (o *CassandraSecurityInfo) GetCassandraAuthTypeOk() (*string, bool)`

GetCassandraAuthTypeOk returns a tuple with the CassandraAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraAuthType

`func (o *CassandraSecurityInfo) SetCassandraAuthType(v string)`

SetCassandraAuthType sets CassandraAuthType field to given value.

### HasCassandraAuthType

`func (o *CassandraSecurityInfo) HasCassandraAuthType() bool`

HasCassandraAuthType returns a boolean if a field has been set.

### SetCassandraAuthTypeNil

`func (o *CassandraSecurityInfo) SetCassandraAuthTypeNil(b bool)`

 SetCassandraAuthTypeNil sets the value for CassandraAuthType to be an explicit nil

### UnsetCassandraAuthType
`func (o *CassandraSecurityInfo) UnsetCassandraAuthType()`

UnsetCassandraAuthType ensures that no value is present for CassandraAuthType, not even an explicit nil
### GetCassandraAuthorizer

`func (o *CassandraSecurityInfo) GetCassandraAuthorizer() string`

GetCassandraAuthorizer returns the CassandraAuthorizer field if non-nil, zero value otherwise.

### GetCassandraAuthorizerOk

`func (o *CassandraSecurityInfo) GetCassandraAuthorizerOk() (*string, bool)`

GetCassandraAuthorizerOk returns a tuple with the CassandraAuthorizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraAuthorizer

`func (o *CassandraSecurityInfo) SetCassandraAuthorizer(v string)`

SetCassandraAuthorizer sets CassandraAuthorizer field to given value.

### HasCassandraAuthorizer

`func (o *CassandraSecurityInfo) HasCassandraAuthorizer() bool`

HasCassandraAuthorizer returns a boolean if a field has been set.

### SetCassandraAuthorizerNil

`func (o *CassandraSecurityInfo) SetCassandraAuthorizerNil(b bool)`

 SetCassandraAuthorizerNil sets the value for CassandraAuthorizer to be an explicit nil

### UnsetCassandraAuthorizer
`func (o *CassandraSecurityInfo) UnsetCassandraAuthorizer()`

UnsetCassandraAuthorizer ensures that no value is present for CassandraAuthorizer, not even an explicit nil
### GetClientEncryption

`func (o *CassandraSecurityInfo) GetClientEncryption() bool`

GetClientEncryption returns the ClientEncryption field if non-nil, zero value otherwise.

### GetClientEncryptionOk

`func (o *CassandraSecurityInfo) GetClientEncryptionOk() (*bool, bool)`

GetClientEncryptionOk returns a tuple with the ClientEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEncryption

`func (o *CassandraSecurityInfo) SetClientEncryption(v bool)`

SetClientEncryption sets ClientEncryption field to given value.

### HasClientEncryption

`func (o *CassandraSecurityInfo) HasClientEncryption() bool`

HasClientEncryption returns a boolean if a field has been set.

### SetClientEncryptionNil

`func (o *CassandraSecurityInfo) SetClientEncryptionNil(b bool)`

 SetClientEncryptionNil sets the value for ClientEncryption to be an explicit nil

### UnsetClientEncryption
`func (o *CassandraSecurityInfo) UnsetClientEncryption()`

UnsetClientEncryption ensures that no value is present for ClientEncryption, not even an explicit nil
### GetDseAuthorization

`func (o *CassandraSecurityInfo) GetDseAuthorization() bool`

GetDseAuthorization returns the DseAuthorization field if non-nil, zero value otherwise.

### GetDseAuthorizationOk

`func (o *CassandraSecurityInfo) GetDseAuthorizationOk() (*bool, bool)`

GetDseAuthorizationOk returns a tuple with the DseAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDseAuthorization

`func (o *CassandraSecurityInfo) SetDseAuthorization(v bool)`

SetDseAuthorization sets DseAuthorization field to given value.

### HasDseAuthorization

`func (o *CassandraSecurityInfo) HasDseAuthorization() bool`

HasDseAuthorization returns a boolean if a field has been set.

### SetDseAuthorizationNil

`func (o *CassandraSecurityInfo) SetDseAuthorizationNil(b bool)`

 SetDseAuthorizationNil sets the value for DseAuthorization to be an explicit nil

### UnsetDseAuthorization
`func (o *CassandraSecurityInfo) UnsetDseAuthorization()`

UnsetDseAuthorization ensures that no value is present for DseAuthorization, not even an explicit nil
### GetServerEncryptionReqClientAuth

`func (o *CassandraSecurityInfo) GetServerEncryptionReqClientAuth() bool`

GetServerEncryptionReqClientAuth returns the ServerEncryptionReqClientAuth field if non-nil, zero value otherwise.

### GetServerEncryptionReqClientAuthOk

`func (o *CassandraSecurityInfo) GetServerEncryptionReqClientAuthOk() (*bool, bool)`

GetServerEncryptionReqClientAuthOk returns a tuple with the ServerEncryptionReqClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEncryptionReqClientAuth

`func (o *CassandraSecurityInfo) SetServerEncryptionReqClientAuth(v bool)`

SetServerEncryptionReqClientAuth sets ServerEncryptionReqClientAuth field to given value.

### HasServerEncryptionReqClientAuth

`func (o *CassandraSecurityInfo) HasServerEncryptionReqClientAuth() bool`

HasServerEncryptionReqClientAuth returns a boolean if a field has been set.

### SetServerEncryptionReqClientAuthNil

`func (o *CassandraSecurityInfo) SetServerEncryptionReqClientAuthNil(b bool)`

 SetServerEncryptionReqClientAuthNil sets the value for ServerEncryptionReqClientAuth to be an explicit nil

### UnsetServerEncryptionReqClientAuth
`func (o *CassandraSecurityInfo) UnsetServerEncryptionReqClientAuth()`

UnsetServerEncryptionReqClientAuth ensures that no value is present for ServerEncryptionReqClientAuth, not even an explicit nil
### GetServerInternodeEncryptionType

`func (o *CassandraSecurityInfo) GetServerInternodeEncryptionType() string`

GetServerInternodeEncryptionType returns the ServerInternodeEncryptionType field if non-nil, zero value otherwise.

### GetServerInternodeEncryptionTypeOk

`func (o *CassandraSecurityInfo) GetServerInternodeEncryptionTypeOk() (*string, bool)`

GetServerInternodeEncryptionTypeOk returns a tuple with the ServerInternodeEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInternodeEncryptionType

`func (o *CassandraSecurityInfo) SetServerInternodeEncryptionType(v string)`

SetServerInternodeEncryptionType sets ServerInternodeEncryptionType field to given value.

### HasServerInternodeEncryptionType

`func (o *CassandraSecurityInfo) HasServerInternodeEncryptionType() bool`

HasServerInternodeEncryptionType returns a boolean if a field has been set.

### SetServerInternodeEncryptionTypeNil

`func (o *CassandraSecurityInfo) SetServerInternodeEncryptionTypeNil(b bool)`

 SetServerInternodeEncryptionTypeNil sets the value for ServerInternodeEncryptionType to be an explicit nil

### UnsetServerInternodeEncryptionType
`func (o *CassandraSecurityInfo) UnsetServerInternodeEncryptionType()`

UnsetServerInternodeEncryptionType ensures that no value is present for ServerInternodeEncryptionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


