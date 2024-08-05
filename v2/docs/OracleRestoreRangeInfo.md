# OracleRestoreRangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScnRangeInfo** | Pointer to [**[]OracleRangeMetaInfo**](OracleRangeMetaInfo.md) | Specifies an array of SCN based oracle restore ranges. | [optional] 
**SequenceRangeInfo** | Pointer to [**[]OracleRangeMetaInfo**](OracleRangeMetaInfo.md) | Specifies an array of time based oracle restore ranges. | [optional] 
**TimeRangeInfo** | Pointer to [**[]OracleRangeMetaInfo**](OracleRangeMetaInfo.md) | Specifies an array of time based oracle restore ranges. | [optional] 

## Methods

### NewOracleRestoreRangeInfo

`func NewOracleRestoreRangeInfo() *OracleRestoreRangeInfo`

NewOracleRestoreRangeInfo instantiates a new OracleRestoreRangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleRestoreRangeInfoWithDefaults

`func NewOracleRestoreRangeInfoWithDefaults() *OracleRestoreRangeInfo`

NewOracleRestoreRangeInfoWithDefaults instantiates a new OracleRestoreRangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScnRangeInfo

`func (o *OracleRestoreRangeInfo) GetScnRangeInfo() []OracleRangeMetaInfo`

GetScnRangeInfo returns the ScnRangeInfo field if non-nil, zero value otherwise.

### GetScnRangeInfoOk

`func (o *OracleRestoreRangeInfo) GetScnRangeInfoOk() (*[]OracleRangeMetaInfo, bool)`

GetScnRangeInfoOk returns a tuple with the ScnRangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScnRangeInfo

`func (o *OracleRestoreRangeInfo) SetScnRangeInfo(v []OracleRangeMetaInfo)`

SetScnRangeInfo sets ScnRangeInfo field to given value.

### HasScnRangeInfo

`func (o *OracleRestoreRangeInfo) HasScnRangeInfo() bool`

HasScnRangeInfo returns a boolean if a field has been set.

### SetScnRangeInfoNil

`func (o *OracleRestoreRangeInfo) SetScnRangeInfoNil(b bool)`

 SetScnRangeInfoNil sets the value for ScnRangeInfo to be an explicit nil

### UnsetScnRangeInfo
`func (o *OracleRestoreRangeInfo) UnsetScnRangeInfo()`

UnsetScnRangeInfo ensures that no value is present for ScnRangeInfo, not even an explicit nil
### GetSequenceRangeInfo

`func (o *OracleRestoreRangeInfo) GetSequenceRangeInfo() []OracleRangeMetaInfo`

GetSequenceRangeInfo returns the SequenceRangeInfo field if non-nil, zero value otherwise.

### GetSequenceRangeInfoOk

`func (o *OracleRestoreRangeInfo) GetSequenceRangeInfoOk() (*[]OracleRangeMetaInfo, bool)`

GetSequenceRangeInfoOk returns a tuple with the SequenceRangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceRangeInfo

`func (o *OracleRestoreRangeInfo) SetSequenceRangeInfo(v []OracleRangeMetaInfo)`

SetSequenceRangeInfo sets SequenceRangeInfo field to given value.

### HasSequenceRangeInfo

`func (o *OracleRestoreRangeInfo) HasSequenceRangeInfo() bool`

HasSequenceRangeInfo returns a boolean if a field has been set.

### SetSequenceRangeInfoNil

`func (o *OracleRestoreRangeInfo) SetSequenceRangeInfoNil(b bool)`

 SetSequenceRangeInfoNil sets the value for SequenceRangeInfo to be an explicit nil

### UnsetSequenceRangeInfo
`func (o *OracleRestoreRangeInfo) UnsetSequenceRangeInfo()`

UnsetSequenceRangeInfo ensures that no value is present for SequenceRangeInfo, not even an explicit nil
### GetTimeRangeInfo

`func (o *OracleRestoreRangeInfo) GetTimeRangeInfo() []OracleRangeMetaInfo`

GetTimeRangeInfo returns the TimeRangeInfo field if non-nil, zero value otherwise.

### GetTimeRangeInfoOk

`func (o *OracleRestoreRangeInfo) GetTimeRangeInfoOk() (*[]OracleRangeMetaInfo, bool)`

GetTimeRangeInfoOk returns a tuple with the TimeRangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeInfo

`func (o *OracleRestoreRangeInfo) SetTimeRangeInfo(v []OracleRangeMetaInfo)`

SetTimeRangeInfo sets TimeRangeInfo field to given value.

### HasTimeRangeInfo

`func (o *OracleRestoreRangeInfo) HasTimeRangeInfo() bool`

HasTimeRangeInfo returns a boolean if a field has been set.

### SetTimeRangeInfoNil

`func (o *OracleRestoreRangeInfo) SetTimeRangeInfoNil(b bool)`

 SetTimeRangeInfoNil sets the value for TimeRangeInfo to be an explicit nil

### UnsetTimeRangeInfo
`func (o *OracleRestoreRangeInfo) UnsetTimeRangeInfo()`

UnsetTimeRangeInfo ensures that no value is present for TimeRangeInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


