# OracleProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbParams** | Pointer to [**[]OracleProtectionGroupDbParams**](OracleProtectionGroupDbParams.md) | Specifies the properties of the Oracle databases. | [optional] 
**SourceId** | **NullableInt64** | Specifies the id of the host on which databases are hosted. | 
**SourceName** | Pointer to **NullableString** | Specifies the name of the host on which databases are hosted. | [optional] [readonly] 

## Methods

### NewOracleProtectionGroupObjectParams

`func NewOracleProtectionGroupObjectParams(sourceId NullableInt64, ) *OracleProtectionGroupObjectParams`

NewOracleProtectionGroupObjectParams instantiates a new OracleProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleProtectionGroupObjectParamsWithDefaults

`func NewOracleProtectionGroupObjectParamsWithDefaults() *OracleProtectionGroupObjectParams`

NewOracleProtectionGroupObjectParamsWithDefaults instantiates a new OracleProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbParams

`func (o *OracleProtectionGroupObjectParams) GetDbParams() []OracleProtectionGroupDbParams`

GetDbParams returns the DbParams field if non-nil, zero value otherwise.

### GetDbParamsOk

`func (o *OracleProtectionGroupObjectParams) GetDbParamsOk() (*[]OracleProtectionGroupDbParams, bool)`

GetDbParamsOk returns a tuple with the DbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbParams

`func (o *OracleProtectionGroupObjectParams) SetDbParams(v []OracleProtectionGroupDbParams)`

SetDbParams sets DbParams field to given value.

### HasDbParams

`func (o *OracleProtectionGroupObjectParams) HasDbParams() bool`

HasDbParams returns a boolean if a field has been set.

### SetDbParamsNil

`func (o *OracleProtectionGroupObjectParams) SetDbParamsNil(b bool)`

 SetDbParamsNil sets the value for DbParams to be an explicit nil

### UnsetDbParams
`func (o *OracleProtectionGroupObjectParams) UnsetDbParams()`

UnsetDbParams ensures that no value is present for DbParams, not even an explicit nil
### GetSourceId

`func (o *OracleProtectionGroupObjectParams) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *OracleProtectionGroupObjectParams) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *OracleProtectionGroupObjectParams) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *OracleProtectionGroupObjectParams) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *OracleProtectionGroupObjectParams) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSourceName

`func (o *OracleProtectionGroupObjectParams) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *OracleProtectionGroupObjectParams) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *OracleProtectionGroupObjectParams) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *OracleProtectionGroupObjectParams) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *OracleProtectionGroupObjectParams) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *OracleProtectionGroupObjectParams) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


