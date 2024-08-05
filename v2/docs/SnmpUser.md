# SnmpUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPassword** | Pointer to **NullableString** | AuthPassword is the authentication password for SNMP V3 users. | [optional] 
**AuthProtocol** | Pointer to **NullableString** | AuthPrototol is the authentication protocol for SNMP V3 users. | [optional] 
**ContextName** | Pointer to **NullableString** | ContextName is the context used for SNMP V3. | [optional] 
**EngineId** | Pointer to **NullableString** | EngineID is the SNMP V3 Engine ID used by trap users. | [optional] 
**PrivPassword** | Pointer to **NullableString** | PrivPassword is the privacy password for SNMP V3 users. | [optional] 
**PrivProtocol** | Pointer to **NullableString** | PrivPrototol is the privacy protocol for SNMP V3 users. | [optional] 
**SecurityLevel** | Pointer to **NullableString** | SecurityLevel is the SNMP V3 security level. It can be authNoPriv, noPriv, and authPriv. | [optional] 
**UserName** | Pointer to **NullableString** | UserName is the user name to access SNMP V2 or SNMP V3 agent. | [optional] 
**UserType** | Pointer to **NullableString** | UserType is the SNMP user type, can be read-only user, read/write user, or trap users. | [optional] 

## Methods

### NewSnmpUser

`func NewSnmpUser() *SnmpUser`

NewSnmpUser instantiates a new SnmpUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpUserWithDefaults

`func NewSnmpUserWithDefaults() *SnmpUser`

NewSnmpUserWithDefaults instantiates a new SnmpUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPassword

`func (o *SnmpUser) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SnmpUser) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SnmpUser) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SnmpUser) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### SetAuthPasswordNil

`func (o *SnmpUser) SetAuthPasswordNil(b bool)`

 SetAuthPasswordNil sets the value for AuthPassword to be an explicit nil

### UnsetAuthPassword
`func (o *SnmpUser) UnsetAuthPassword()`

UnsetAuthPassword ensures that no value is present for AuthPassword, not even an explicit nil
### GetAuthProtocol

`func (o *SnmpUser) GetAuthProtocol() string`

GetAuthProtocol returns the AuthProtocol field if non-nil, zero value otherwise.

### GetAuthProtocolOk

`func (o *SnmpUser) GetAuthProtocolOk() (*string, bool)`

GetAuthProtocolOk returns a tuple with the AuthProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProtocol

`func (o *SnmpUser) SetAuthProtocol(v string)`

SetAuthProtocol sets AuthProtocol field to given value.

### HasAuthProtocol

`func (o *SnmpUser) HasAuthProtocol() bool`

HasAuthProtocol returns a boolean if a field has been set.

### SetAuthProtocolNil

`func (o *SnmpUser) SetAuthProtocolNil(b bool)`

 SetAuthProtocolNil sets the value for AuthProtocol to be an explicit nil

### UnsetAuthProtocol
`func (o *SnmpUser) UnsetAuthProtocol()`

UnsetAuthProtocol ensures that no value is present for AuthProtocol, not even an explicit nil
### GetContextName

`func (o *SnmpUser) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *SnmpUser) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *SnmpUser) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *SnmpUser) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### SetContextNameNil

`func (o *SnmpUser) SetContextNameNil(b bool)`

 SetContextNameNil sets the value for ContextName to be an explicit nil

### UnsetContextName
`func (o *SnmpUser) UnsetContextName()`

UnsetContextName ensures that no value is present for ContextName, not even an explicit nil
### GetEngineId

`func (o *SnmpUser) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *SnmpUser) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *SnmpUser) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *SnmpUser) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### SetEngineIdNil

`func (o *SnmpUser) SetEngineIdNil(b bool)`

 SetEngineIdNil sets the value for EngineId to be an explicit nil

### UnsetEngineId
`func (o *SnmpUser) UnsetEngineId()`

UnsetEngineId ensures that no value is present for EngineId, not even an explicit nil
### GetPrivPassword

`func (o *SnmpUser) GetPrivPassword() string`

GetPrivPassword returns the PrivPassword field if non-nil, zero value otherwise.

### GetPrivPasswordOk

`func (o *SnmpUser) GetPrivPasswordOk() (*string, bool)`

GetPrivPasswordOk returns a tuple with the PrivPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivPassword

`func (o *SnmpUser) SetPrivPassword(v string)`

SetPrivPassword sets PrivPassword field to given value.

### HasPrivPassword

`func (o *SnmpUser) HasPrivPassword() bool`

HasPrivPassword returns a boolean if a field has been set.

### SetPrivPasswordNil

`func (o *SnmpUser) SetPrivPasswordNil(b bool)`

 SetPrivPasswordNil sets the value for PrivPassword to be an explicit nil

### UnsetPrivPassword
`func (o *SnmpUser) UnsetPrivPassword()`

UnsetPrivPassword ensures that no value is present for PrivPassword, not even an explicit nil
### GetPrivProtocol

`func (o *SnmpUser) GetPrivProtocol() string`

GetPrivProtocol returns the PrivProtocol field if non-nil, zero value otherwise.

### GetPrivProtocolOk

`func (o *SnmpUser) GetPrivProtocolOk() (*string, bool)`

GetPrivProtocolOk returns a tuple with the PrivProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivProtocol

`func (o *SnmpUser) SetPrivProtocol(v string)`

SetPrivProtocol sets PrivProtocol field to given value.

### HasPrivProtocol

`func (o *SnmpUser) HasPrivProtocol() bool`

HasPrivProtocol returns a boolean if a field has been set.

### SetPrivProtocolNil

`func (o *SnmpUser) SetPrivProtocolNil(b bool)`

 SetPrivProtocolNil sets the value for PrivProtocol to be an explicit nil

### UnsetPrivProtocol
`func (o *SnmpUser) UnsetPrivProtocol()`

UnsetPrivProtocol ensures that no value is present for PrivProtocol, not even an explicit nil
### GetSecurityLevel

`func (o *SnmpUser) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpUser) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpUser) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpUser) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### SetSecurityLevelNil

`func (o *SnmpUser) SetSecurityLevelNil(b bool)`

 SetSecurityLevelNil sets the value for SecurityLevel to be an explicit nil

### UnsetSecurityLevel
`func (o *SnmpUser) UnsetSecurityLevel()`

UnsetSecurityLevel ensures that no value is present for SecurityLevel, not even an explicit nil
### GetUserName

`func (o *SnmpUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SnmpUser) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SnmpUser) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SnmpUser) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SnmpUser) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SnmpUser) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUserType

`func (o *SnmpUser) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *SnmpUser) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *SnmpUser) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *SnmpUser) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *SnmpUser) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *SnmpUser) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


