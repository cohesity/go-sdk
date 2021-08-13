# RecoverOracleGranularRestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GranularityType** | Pointer to **NullableString** | Specifies type of granular restore. | [optional] 
**PdbRestoreParams** | Pointer to [**OraclePdbRestoreParams**](OraclePdbRestoreParams.md) | Specifies information about the list of pdbs to be restored. | [optional] 

## Methods

### NewRecoverOracleGranularRestoreInfo

`func NewRecoverOracleGranularRestoreInfo() *RecoverOracleGranularRestoreInfo`

NewRecoverOracleGranularRestoreInfo instantiates a new RecoverOracleGranularRestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleGranularRestoreInfoWithDefaults

`func NewRecoverOracleGranularRestoreInfoWithDefaults() *RecoverOracleGranularRestoreInfo`

NewRecoverOracleGranularRestoreInfoWithDefaults instantiates a new RecoverOracleGranularRestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularityType

`func (o *RecoverOracleGranularRestoreInfo) GetGranularityType() string`

GetGranularityType returns the GranularityType field if non-nil, zero value otherwise.

### GetGranularityTypeOk

`func (o *RecoverOracleGranularRestoreInfo) GetGranularityTypeOk() (*string, bool)`

GetGranularityTypeOk returns a tuple with the GranularityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityType

`func (o *RecoverOracleGranularRestoreInfo) SetGranularityType(v string)`

SetGranularityType sets GranularityType field to given value.

### HasGranularityType

`func (o *RecoverOracleGranularRestoreInfo) HasGranularityType() bool`

HasGranularityType returns a boolean if a field has been set.

### SetGranularityTypeNil

`func (o *RecoverOracleGranularRestoreInfo) SetGranularityTypeNil(b bool)`

 SetGranularityTypeNil sets the value for GranularityType to be an explicit nil

### UnsetGranularityType
`func (o *RecoverOracleGranularRestoreInfo) UnsetGranularityType()`

UnsetGranularityType ensures that no value is present for GranularityType, not even an explicit nil
### GetPdbRestoreParams

`func (o *RecoverOracleGranularRestoreInfo) GetPdbRestoreParams() OraclePdbRestoreParams`

GetPdbRestoreParams returns the PdbRestoreParams field if non-nil, zero value otherwise.

### GetPdbRestoreParamsOk

`func (o *RecoverOracleGranularRestoreInfo) GetPdbRestoreParamsOk() (*OraclePdbRestoreParams, bool)`

GetPdbRestoreParamsOk returns a tuple with the PdbRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbRestoreParams

`func (o *RecoverOracleGranularRestoreInfo) SetPdbRestoreParams(v OraclePdbRestoreParams)`

SetPdbRestoreParams sets PdbRestoreParams field to given value.

### HasPdbRestoreParams

`func (o *RecoverOracleGranularRestoreInfo) HasPdbRestoreParams() bool`

HasPdbRestoreParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


