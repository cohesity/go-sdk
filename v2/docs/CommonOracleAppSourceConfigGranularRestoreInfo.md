# CommonOracleAppSourceConfigGranularRestoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GranularityType** | Pointer to **NullableString** | Specifies type of granular restore. | [optional] 
**PdbRestoreParams** | Pointer to [**RecoverOracleGranularRestoreInfoPdbRestoreParams**](RecoverOracleGranularRestoreInfoPdbRestoreParams.md) |  | [optional] 

## Methods

### NewCommonOracleAppSourceConfigGranularRestoreInfo

`func NewCommonOracleAppSourceConfigGranularRestoreInfo() *CommonOracleAppSourceConfigGranularRestoreInfo`

NewCommonOracleAppSourceConfigGranularRestoreInfo instantiates a new CommonOracleAppSourceConfigGranularRestoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonOracleAppSourceConfigGranularRestoreInfoWithDefaults

`func NewCommonOracleAppSourceConfigGranularRestoreInfoWithDefaults() *CommonOracleAppSourceConfigGranularRestoreInfo`

NewCommonOracleAppSourceConfigGranularRestoreInfoWithDefaults instantiates a new CommonOracleAppSourceConfigGranularRestoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularityType

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) GetGranularityType() string`

GetGranularityType returns the GranularityType field if non-nil, zero value otherwise.

### GetGranularityTypeOk

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) GetGranularityTypeOk() (*string, bool)`

GetGranularityTypeOk returns a tuple with the GranularityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularityType

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) SetGranularityType(v string)`

SetGranularityType sets GranularityType field to given value.

### HasGranularityType

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) HasGranularityType() bool`

HasGranularityType returns a boolean if a field has been set.

### SetGranularityTypeNil

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) SetGranularityTypeNil(b bool)`

 SetGranularityTypeNil sets the value for GranularityType to be an explicit nil

### UnsetGranularityType
`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) UnsetGranularityType()`

UnsetGranularityType ensures that no value is present for GranularityType, not even an explicit nil
### GetPdbRestoreParams

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) GetPdbRestoreParams() RecoverOracleGranularRestoreInfoPdbRestoreParams`

GetPdbRestoreParams returns the PdbRestoreParams field if non-nil, zero value otherwise.

### GetPdbRestoreParamsOk

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) GetPdbRestoreParamsOk() (*RecoverOracleGranularRestoreInfoPdbRestoreParams, bool)`

GetPdbRestoreParamsOk returns a tuple with the PdbRestoreParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbRestoreParams

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) SetPdbRestoreParams(v RecoverOracleGranularRestoreInfoPdbRestoreParams)`

SetPdbRestoreParams sets PdbRestoreParams field to given value.

### HasPdbRestoreParams

`func (o *CommonOracleAppSourceConfigGranularRestoreInfo) HasPdbRestoreParams() bool`

HasPdbRestoreParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


