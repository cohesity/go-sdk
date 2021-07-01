# MongoDBConnectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **NullableString** | Specifies whether authentication is configured on this MongoDB cluster. Specifies the type of an MongoDB source entity. &#39;SCRAM&#39; &#39;LDAP&#39; &#39;NONE&#39; | [optional] 
**AuthenticatingDatabaseName** | Pointer to **NullableString** | Specifies the Authenticating Database for this MongoDB cluster. | [optional] 
**RequiresSsl** | Pointer to **NullableBool** | Specifies whether connection is allowed through SSL only in this cluster. | [optional] 
**SecondaryNodeTag** | Pointer to **NullableString** | MongoDB Secondary node tag. Required only if &#39;useSecondaryForBackup&#39; is true. The system will use this to identify the secondary nodes for reading backup data. | [optional] 
**Seeds** | Pointer to **[]string** | Specifies the seeds of this MongoDB Cluster. | [optional] 
**UseSecondaryForBackup** | Pointer to **NullableBool** | Set this to true if you want the system to peform backups from secondary nodes. | [optional] 

## Methods

### NewMongoDBConnectParams

`func NewMongoDBConnectParams() *MongoDBConnectParams`

NewMongoDBConnectParams instantiates a new MongoDBConnectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBConnectParamsWithDefaults

`func NewMongoDBConnectParamsWithDefaults() *MongoDBConnectParams`

NewMongoDBConnectParamsWithDefaults instantiates a new MongoDBConnectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *MongoDBConnectParams) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MongoDBConnectParams) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MongoDBConnectParams) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *MongoDBConnectParams) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *MongoDBConnectParams) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *MongoDBConnectParams) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetAuthenticatingDatabaseName

`func (o *MongoDBConnectParams) GetAuthenticatingDatabaseName() string`

GetAuthenticatingDatabaseName returns the AuthenticatingDatabaseName field if non-nil, zero value otherwise.

### GetAuthenticatingDatabaseNameOk

`func (o *MongoDBConnectParams) GetAuthenticatingDatabaseNameOk() (*string, bool)`

GetAuthenticatingDatabaseNameOk returns a tuple with the AuthenticatingDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatingDatabaseName

`func (o *MongoDBConnectParams) SetAuthenticatingDatabaseName(v string)`

SetAuthenticatingDatabaseName sets AuthenticatingDatabaseName field to given value.

### HasAuthenticatingDatabaseName

`func (o *MongoDBConnectParams) HasAuthenticatingDatabaseName() bool`

HasAuthenticatingDatabaseName returns a boolean if a field has been set.

### SetAuthenticatingDatabaseNameNil

`func (o *MongoDBConnectParams) SetAuthenticatingDatabaseNameNil(b bool)`

 SetAuthenticatingDatabaseNameNil sets the value for AuthenticatingDatabaseName to be an explicit nil

### UnsetAuthenticatingDatabaseName
`func (o *MongoDBConnectParams) UnsetAuthenticatingDatabaseName()`

UnsetAuthenticatingDatabaseName ensures that no value is present for AuthenticatingDatabaseName, not even an explicit nil
### GetRequiresSsl

`func (o *MongoDBConnectParams) GetRequiresSsl() bool`

GetRequiresSsl returns the RequiresSsl field if non-nil, zero value otherwise.

### GetRequiresSslOk

`func (o *MongoDBConnectParams) GetRequiresSslOk() (*bool, bool)`

GetRequiresSslOk returns a tuple with the RequiresSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresSsl

`func (o *MongoDBConnectParams) SetRequiresSsl(v bool)`

SetRequiresSsl sets RequiresSsl field to given value.

### HasRequiresSsl

`func (o *MongoDBConnectParams) HasRequiresSsl() bool`

HasRequiresSsl returns a boolean if a field has been set.

### SetRequiresSslNil

`func (o *MongoDBConnectParams) SetRequiresSslNil(b bool)`

 SetRequiresSslNil sets the value for RequiresSsl to be an explicit nil

### UnsetRequiresSsl
`func (o *MongoDBConnectParams) UnsetRequiresSsl()`

UnsetRequiresSsl ensures that no value is present for RequiresSsl, not even an explicit nil
### GetSecondaryNodeTag

`func (o *MongoDBConnectParams) GetSecondaryNodeTag() string`

GetSecondaryNodeTag returns the SecondaryNodeTag field if non-nil, zero value otherwise.

### GetSecondaryNodeTagOk

`func (o *MongoDBConnectParams) GetSecondaryNodeTagOk() (*string, bool)`

GetSecondaryNodeTagOk returns a tuple with the SecondaryNodeTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryNodeTag

`func (o *MongoDBConnectParams) SetSecondaryNodeTag(v string)`

SetSecondaryNodeTag sets SecondaryNodeTag field to given value.

### HasSecondaryNodeTag

`func (o *MongoDBConnectParams) HasSecondaryNodeTag() bool`

HasSecondaryNodeTag returns a boolean if a field has been set.

### SetSecondaryNodeTagNil

`func (o *MongoDBConnectParams) SetSecondaryNodeTagNil(b bool)`

 SetSecondaryNodeTagNil sets the value for SecondaryNodeTag to be an explicit nil

### UnsetSecondaryNodeTag
`func (o *MongoDBConnectParams) UnsetSecondaryNodeTag()`

UnsetSecondaryNodeTag ensures that no value is present for SecondaryNodeTag, not even an explicit nil
### GetSeeds

`func (o *MongoDBConnectParams) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *MongoDBConnectParams) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *MongoDBConnectParams) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *MongoDBConnectParams) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.

### SetSeedsNil

`func (o *MongoDBConnectParams) SetSeedsNil(b bool)`

 SetSeedsNil sets the value for Seeds to be an explicit nil

### UnsetSeeds
`func (o *MongoDBConnectParams) UnsetSeeds()`

UnsetSeeds ensures that no value is present for Seeds, not even an explicit nil
### GetUseSecondaryForBackup

`func (o *MongoDBConnectParams) GetUseSecondaryForBackup() bool`

GetUseSecondaryForBackup returns the UseSecondaryForBackup field if non-nil, zero value otherwise.

### GetUseSecondaryForBackupOk

`func (o *MongoDBConnectParams) GetUseSecondaryForBackupOk() (*bool, bool)`

GetUseSecondaryForBackupOk returns a tuple with the UseSecondaryForBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSecondaryForBackup

`func (o *MongoDBConnectParams) SetUseSecondaryForBackup(v bool)`

SetUseSecondaryForBackup sets UseSecondaryForBackup field to given value.

### HasUseSecondaryForBackup

`func (o *MongoDBConnectParams) HasUseSecondaryForBackup() bool`

HasUseSecondaryForBackup returns a boolean if a field has been set.

### SetUseSecondaryForBackupNil

`func (o *MongoDBConnectParams) SetUseSecondaryForBackupNil(b bool)`

 SetUseSecondaryForBackupNil sets the value for UseSecondaryForBackup to be an explicit nil

### UnsetUseSecondaryForBackup
`func (o *MongoDBConnectParams) UnsetUseSecondaryForBackup()`

UnsetUseSecondaryForBackup ensures that no value is present for UseSecondaryForBackup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


