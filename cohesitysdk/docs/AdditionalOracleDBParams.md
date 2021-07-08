# AdditionalOracleDBParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEntityId** | Pointer to **NullableInt64** | Database app id. | [optional] 
**DbInfoChannelVec** | Pointer to [**[]OracleDBChannelInfo**](OracleDBChannelInfo.md) | The following proto message should be renamed to a more general message to represent parameters pertaining to a single unique Oracle database. Uniqueness of an Oracle database should be strictly determined by its database unique name. i.e. all backup parameters needed for a unique Oracle database should be expressed in the following proto.  It is a vector for future support of backing up Data Guard sources. We may or may not need this to be vector to support Data Guard sources. For now, the size of this vector is always 1. When we rename this proto in the future, if we determine there is no need to use a vector for Data Guard support, we can choose to remove &#39;repeated&#39;. | [optional] 

## Methods

### NewAdditionalOracleDBParams

`func NewAdditionalOracleDBParams() *AdditionalOracleDBParams`

NewAdditionalOracleDBParams instantiates a new AdditionalOracleDBParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalOracleDBParamsWithDefaults

`func NewAdditionalOracleDBParamsWithDefaults() *AdditionalOracleDBParams`

NewAdditionalOracleDBParamsWithDefaults instantiates a new AdditionalOracleDBParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEntityId

`func (o *AdditionalOracleDBParams) GetAppEntityId() int64`

GetAppEntityId returns the AppEntityId field if non-nil, zero value otherwise.

### GetAppEntityIdOk

`func (o *AdditionalOracleDBParams) GetAppEntityIdOk() (*int64, bool)`

GetAppEntityIdOk returns a tuple with the AppEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEntityId

`func (o *AdditionalOracleDBParams) SetAppEntityId(v int64)`

SetAppEntityId sets AppEntityId field to given value.

### HasAppEntityId

`func (o *AdditionalOracleDBParams) HasAppEntityId() bool`

HasAppEntityId returns a boolean if a field has been set.

### SetAppEntityIdNil

`func (o *AdditionalOracleDBParams) SetAppEntityIdNil(b bool)`

 SetAppEntityIdNil sets the value for AppEntityId to be an explicit nil

### UnsetAppEntityId
`func (o *AdditionalOracleDBParams) UnsetAppEntityId()`

UnsetAppEntityId ensures that no value is present for AppEntityId, not even an explicit nil
### GetDbInfoChannelVec

`func (o *AdditionalOracleDBParams) GetDbInfoChannelVec() []OracleDBChannelInfo`

GetDbInfoChannelVec returns the DbInfoChannelVec field if non-nil, zero value otherwise.

### GetDbInfoChannelVecOk

`func (o *AdditionalOracleDBParams) GetDbInfoChannelVecOk() (*[]OracleDBChannelInfo, bool)`

GetDbInfoChannelVecOk returns a tuple with the DbInfoChannelVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInfoChannelVec

`func (o *AdditionalOracleDBParams) SetDbInfoChannelVec(v []OracleDBChannelInfo)`

SetDbInfoChannelVec sets DbInfoChannelVec field to given value.

### HasDbInfoChannelVec

`func (o *AdditionalOracleDBParams) HasDbInfoChannelVec() bool`

HasDbInfoChannelVec returns a boolean if a field has been set.

### SetDbInfoChannelVecNil

`func (o *AdditionalOracleDBParams) SetDbInfoChannelVecNil(b bool)`

 SetDbInfoChannelVecNil sets the value for DbInfoChannelVec to be an explicit nil

### UnsetDbInfoChannelVec
`func (o *AdditionalOracleDBParams) UnsetDbInfoChannelVec()`

UnsetDbInfoChannelVec ensures that no value is present for DbInfoChannelVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


