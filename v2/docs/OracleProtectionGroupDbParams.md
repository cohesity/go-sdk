# OracleProtectionGroupDbParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseId** | Pointer to **NullableInt64** | Specifies the id of the Oracle database. | [optional] 
**DatabaseName** | Pointer to **NullableString** | Specifies the name of the Oracle database. | [optional] [readonly] 
**DbChannels** | Pointer to [**[]OracleDbChannel**](OracleDbChannel.md) | Specifies the Oracle database node channels info. If not specified, the default values assigned by the server are applied to all the databases. | [optional] 

## Methods

### NewOracleProtectionGroupDbParams

`func NewOracleProtectionGroupDbParams() *OracleProtectionGroupDbParams`

NewOracleProtectionGroupDbParams instantiates a new OracleProtectionGroupDbParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleProtectionGroupDbParamsWithDefaults

`func NewOracleProtectionGroupDbParamsWithDefaults() *OracleProtectionGroupDbParams`

NewOracleProtectionGroupDbParamsWithDefaults instantiates a new OracleProtectionGroupDbParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseId

`func (o *OracleProtectionGroupDbParams) GetDatabaseId() int64`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *OracleProtectionGroupDbParams) GetDatabaseIdOk() (*int64, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *OracleProtectionGroupDbParams) SetDatabaseId(v int64)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *OracleProtectionGroupDbParams) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### SetDatabaseIdNil

`func (o *OracleProtectionGroupDbParams) SetDatabaseIdNil(b bool)`

 SetDatabaseIdNil sets the value for DatabaseId to be an explicit nil

### UnsetDatabaseId
`func (o *OracleProtectionGroupDbParams) UnsetDatabaseId()`

UnsetDatabaseId ensures that no value is present for DatabaseId, not even an explicit nil
### GetDatabaseName

`func (o *OracleProtectionGroupDbParams) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *OracleProtectionGroupDbParams) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *OracleProtectionGroupDbParams) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *OracleProtectionGroupDbParams) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *OracleProtectionGroupDbParams) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *OracleProtectionGroupDbParams) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetDbChannels

`func (o *OracleProtectionGroupDbParams) GetDbChannels() []OracleDbChannel`

GetDbChannels returns the DbChannels field if non-nil, zero value otherwise.

### GetDbChannelsOk

`func (o *OracleProtectionGroupDbParams) GetDbChannelsOk() (*[]OracleDbChannel, bool)`

GetDbChannelsOk returns a tuple with the DbChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbChannels

`func (o *OracleProtectionGroupDbParams) SetDbChannels(v []OracleDbChannel)`

SetDbChannels sets DbChannels field to given value.

### HasDbChannels

`func (o *OracleProtectionGroupDbParams) HasDbChannels() bool`

HasDbChannels returns a boolean if a field has been set.

### SetDbChannelsNil

`func (o *OracleProtectionGroupDbParams) SetDbChannelsNil(b bool)`

 SetDbChannelsNil sets the value for DbChannels to be an explicit nil

### UnsetDbChannels
`func (o *OracleProtectionGroupDbParams) UnsetDbChannels()`

UnsetDbChannels ensures that no value is present for DbChannels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


