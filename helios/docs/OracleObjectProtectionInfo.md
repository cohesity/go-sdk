# OracleObjectProtectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **NullableInt64** | Specifies the id of the host on which databases are hosted | 
**DbParams** | Pointer to [**[]OracleProtectionGroupDbParams**](OracleProtectionGroupDbParams.md) | Specifies the properties of the Oracle databases. | [optional] 

## Methods

### NewOracleObjectProtectionInfo

`func NewOracleObjectProtectionInfo(objectId NullableInt64, ) *OracleObjectProtectionInfo`

NewOracleObjectProtectionInfo instantiates a new OracleObjectProtectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleObjectProtectionInfoWithDefaults

`func NewOracleObjectProtectionInfoWithDefaults() *OracleObjectProtectionInfo`

NewOracleObjectProtectionInfoWithDefaults instantiates a new OracleObjectProtectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *OracleObjectProtectionInfo) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *OracleObjectProtectionInfo) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *OracleObjectProtectionInfo) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### SetObjectIdNil

`func (o *OracleObjectProtectionInfo) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *OracleObjectProtectionInfo) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetDbParams

`func (o *OracleObjectProtectionInfo) GetDbParams() []OracleProtectionGroupDbParams`

GetDbParams returns the DbParams field if non-nil, zero value otherwise.

### GetDbParamsOk

`func (o *OracleObjectProtectionInfo) GetDbParamsOk() (*[]OracleProtectionGroupDbParams, bool)`

GetDbParamsOk returns a tuple with the DbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbParams

`func (o *OracleObjectProtectionInfo) SetDbParams(v []OracleProtectionGroupDbParams)`

SetDbParams sets DbParams field to given value.

### HasDbParams

`func (o *OracleObjectProtectionInfo) HasDbParams() bool`

HasDbParams returns a boolean if a field has been set.

### SetDbParamsNil

`func (o *OracleObjectProtectionInfo) SetDbParamsNil(b bool)`

 SetDbParamsNil sets the value for DbParams to be an explicit nil

### UnsetDbParams
`func (o *OracleObjectProtectionInfo) UnsetDbParams()`

UnsetDbParams ensures that no value is present for DbParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


