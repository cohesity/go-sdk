# MongoDBSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | **[]string** | Specify the MongoS hosts for a sharded cluster and the MongoD hosts for a non-sharded cluster. You can specify a sub-set of the hosts. | 
**AuthType** | **NullableString** | MongoDB authentication type. | 
**Username** | Pointer to **NullableString** | Specifies the username of the MongoDB cluster. Should be set if &#39;authType&#39; is &#39;LDAP&#39; or &#39;SCRAM&#39;. | [optional] 
**Principal** | Pointer to **NullableString** | Specifies the principal name of the MongoDB cluster. Should be set if &#39;authType&#39; is &#39;KERBEROS&#39;. | [optional] 
**Password** | Pointer to **NullableString** | Specifies the password for the MongoDB cluster. Should be set if &#39;authType&#39; is &#39;LDAP&#39; or &#39;SCRAM&#39;. | [optional] 
**AuthenticatingDatabase** | Pointer to **NullableString** | Authenticating Database for this cluster. Should be set if &#39;authType&#39; is &#39;LDAP&#39; or &#39;SCRAM&#39;. | [optional] 
**IsSslRequired** | **bool** | Set to true if connection to MongoDB has to be over SSL. | 
**UseSecondaryForBackup** | **bool** | Set this to true if you want the system to peform backups from secondary nodes. | 
**SecondaryNodeTag** | Pointer to **string** | MongoDB Secondary node tag. Required only if &#39;useSecondaryForBackup&#39; is true.The system will use this to identify the secondary nodes for reading backup data. | [optional] 

## Methods

### NewMongoDBSourceRegistrationParams

`func NewMongoDBSourceRegistrationParams(hosts []string, authType NullableString, isSslRequired bool, useSecondaryForBackup bool, ) *MongoDBSourceRegistrationParams`

NewMongoDBSourceRegistrationParams instantiates a new MongoDBSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBSourceRegistrationParamsWithDefaults

`func NewMongoDBSourceRegistrationParamsWithDefaults() *MongoDBSourceRegistrationParams`

NewMongoDBSourceRegistrationParamsWithDefaults instantiates a new MongoDBSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *MongoDBSourceRegistrationParams) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *MongoDBSourceRegistrationParams) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *MongoDBSourceRegistrationParams) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetAuthType

`func (o *MongoDBSourceRegistrationParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MongoDBSourceRegistrationParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MongoDBSourceRegistrationParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### SetAuthTypeNil

`func (o *MongoDBSourceRegistrationParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *MongoDBSourceRegistrationParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetUsername

`func (o *MongoDBSourceRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MongoDBSourceRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MongoDBSourceRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MongoDBSourceRegistrationParams) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *MongoDBSourceRegistrationParams) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *MongoDBSourceRegistrationParams) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPrincipal

`func (o *MongoDBSourceRegistrationParams) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *MongoDBSourceRegistrationParams) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *MongoDBSourceRegistrationParams) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *MongoDBSourceRegistrationParams) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### SetPrincipalNil

`func (o *MongoDBSourceRegistrationParams) SetPrincipalNil(b bool)`

 SetPrincipalNil sets the value for Principal to be an explicit nil

### UnsetPrincipal
`func (o *MongoDBSourceRegistrationParams) UnsetPrincipal()`

UnsetPrincipal ensures that no value is present for Principal, not even an explicit nil
### GetPassword

`func (o *MongoDBSourceRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MongoDBSourceRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MongoDBSourceRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MongoDBSourceRegistrationParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MongoDBSourceRegistrationParams) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MongoDBSourceRegistrationParams) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetAuthenticatingDatabase

`func (o *MongoDBSourceRegistrationParams) GetAuthenticatingDatabase() string`

GetAuthenticatingDatabase returns the AuthenticatingDatabase field if non-nil, zero value otherwise.

### GetAuthenticatingDatabaseOk

`func (o *MongoDBSourceRegistrationParams) GetAuthenticatingDatabaseOk() (*string, bool)`

GetAuthenticatingDatabaseOk returns a tuple with the AuthenticatingDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatingDatabase

`func (o *MongoDBSourceRegistrationParams) SetAuthenticatingDatabase(v string)`

SetAuthenticatingDatabase sets AuthenticatingDatabase field to given value.

### HasAuthenticatingDatabase

`func (o *MongoDBSourceRegistrationParams) HasAuthenticatingDatabase() bool`

HasAuthenticatingDatabase returns a boolean if a field has been set.

### SetAuthenticatingDatabaseNil

`func (o *MongoDBSourceRegistrationParams) SetAuthenticatingDatabaseNil(b bool)`

 SetAuthenticatingDatabaseNil sets the value for AuthenticatingDatabase to be an explicit nil

### UnsetAuthenticatingDatabase
`func (o *MongoDBSourceRegistrationParams) UnsetAuthenticatingDatabase()`

UnsetAuthenticatingDatabase ensures that no value is present for AuthenticatingDatabase, not even an explicit nil
### GetIsSslRequired

`func (o *MongoDBSourceRegistrationParams) GetIsSslRequired() bool`

GetIsSslRequired returns the IsSslRequired field if non-nil, zero value otherwise.

### GetIsSslRequiredOk

`func (o *MongoDBSourceRegistrationParams) GetIsSslRequiredOk() (*bool, bool)`

GetIsSslRequiredOk returns a tuple with the IsSslRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSslRequired

`func (o *MongoDBSourceRegistrationParams) SetIsSslRequired(v bool)`

SetIsSslRequired sets IsSslRequired field to given value.


### GetUseSecondaryForBackup

`func (o *MongoDBSourceRegistrationParams) GetUseSecondaryForBackup() bool`

GetUseSecondaryForBackup returns the UseSecondaryForBackup field if non-nil, zero value otherwise.

### GetUseSecondaryForBackupOk

`func (o *MongoDBSourceRegistrationParams) GetUseSecondaryForBackupOk() (*bool, bool)`

GetUseSecondaryForBackupOk returns a tuple with the UseSecondaryForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSecondaryForBackup

`func (o *MongoDBSourceRegistrationParams) SetUseSecondaryForBackup(v bool)`

SetUseSecondaryForBackup sets UseSecondaryForBackup field to given value.


### GetSecondaryNodeTag

`func (o *MongoDBSourceRegistrationParams) GetSecondaryNodeTag() string`

GetSecondaryNodeTag returns the SecondaryNodeTag field if non-nil, zero value otherwise.

### GetSecondaryNodeTagOk

`func (o *MongoDBSourceRegistrationParams) GetSecondaryNodeTagOk() (*string, bool)`

GetSecondaryNodeTagOk returns a tuple with the SecondaryNodeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryNodeTag

`func (o *MongoDBSourceRegistrationParams) SetSecondaryNodeTag(v string)`

SetSecondaryNodeTag sets SecondaryNodeTag field to given value.

### HasSecondaryNodeTag

`func (o *MongoDBSourceRegistrationParams) HasSecondaryNodeTag() bool`

HasSecondaryNodeTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


