# OracleArchiveLogInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLogRestoreDest** | Pointer to **NullableString** | Specifies destination where archive logs are to be restored. | [optional] 
**RangeInfoVec** | Pointer to [**[]OracleRangeMetaInfo**](OracleRangeMetaInfo.md) | Specifies an array of oracle restore ranges. | [optional] 
**RangeType** | Pointer to **NullableString** | Specifies the type of range. | [optional] 

## Methods

### NewOracleArchiveLogInfo

`func NewOracleArchiveLogInfo() *OracleArchiveLogInfo`

NewOracleArchiveLogInfo instantiates a new OracleArchiveLogInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleArchiveLogInfoWithDefaults

`func NewOracleArchiveLogInfoWithDefaults() *OracleArchiveLogInfo`

NewOracleArchiveLogInfoWithDefaults instantiates a new OracleArchiveLogInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLogRestoreDest

`func (o *OracleArchiveLogInfo) GetArchiveLogRestoreDest() string`

GetArchiveLogRestoreDest returns the ArchiveLogRestoreDest field if non-nil, zero value otherwise.

### GetArchiveLogRestoreDestOk

`func (o *OracleArchiveLogInfo) GetArchiveLogRestoreDestOk() (*string, bool)`

GetArchiveLogRestoreDestOk returns a tuple with the ArchiveLogRestoreDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogRestoreDest

`func (o *OracleArchiveLogInfo) SetArchiveLogRestoreDest(v string)`

SetArchiveLogRestoreDest sets ArchiveLogRestoreDest field to given value.

### HasArchiveLogRestoreDest

`func (o *OracleArchiveLogInfo) HasArchiveLogRestoreDest() bool`

HasArchiveLogRestoreDest returns a boolean if a field has been set.

### SetArchiveLogRestoreDestNil

`func (o *OracleArchiveLogInfo) SetArchiveLogRestoreDestNil(b bool)`

 SetArchiveLogRestoreDestNil sets the value for ArchiveLogRestoreDest to be an explicit nil

### UnsetArchiveLogRestoreDest
`func (o *OracleArchiveLogInfo) UnsetArchiveLogRestoreDest()`

UnsetArchiveLogRestoreDest ensures that no value is present for ArchiveLogRestoreDest, not even an explicit nil
### GetRangeInfoVec

`func (o *OracleArchiveLogInfo) GetRangeInfoVec() []OracleRangeMetaInfo`

GetRangeInfoVec returns the RangeInfoVec field if non-nil, zero value otherwise.

### GetRangeInfoVecOk

`func (o *OracleArchiveLogInfo) GetRangeInfoVecOk() (*[]OracleRangeMetaInfo, bool)`

GetRangeInfoVecOk returns a tuple with the RangeInfoVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeInfoVec

`func (o *OracleArchiveLogInfo) SetRangeInfoVec(v []OracleRangeMetaInfo)`

SetRangeInfoVec sets RangeInfoVec field to given value.

### HasRangeInfoVec

`func (o *OracleArchiveLogInfo) HasRangeInfoVec() bool`

HasRangeInfoVec returns a boolean if a field has been set.

### SetRangeInfoVecNil

`func (o *OracleArchiveLogInfo) SetRangeInfoVecNil(b bool)`

 SetRangeInfoVecNil sets the value for RangeInfoVec to be an explicit nil

### UnsetRangeInfoVec
`func (o *OracleArchiveLogInfo) UnsetRangeInfoVec()`

UnsetRangeInfoVec ensures that no value is present for RangeInfoVec, not even an explicit nil
### GetRangeType

`func (o *OracleArchiveLogInfo) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *OracleArchiveLogInfo) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *OracleArchiveLogInfo) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.

### HasRangeType

`func (o *OracleArchiveLogInfo) HasRangeType() bool`

HasRangeType returns a boolean if a field has been set.

### SetRangeTypeNil

`func (o *OracleArchiveLogInfo) SetRangeTypeNil(b bool)`

 SetRangeTypeNil sets the value for RangeType to be an explicit nil

### UnsetRangeType
`func (o *OracleArchiveLogInfo) UnsetRangeType()`

UnsetRangeType ensures that no value is present for RangeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


