# OracleSourceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalOracleDbParamsVec** | Pointer to [**[]AdditionalOracleDBParams**](AdditionalOracleDBParams.md) | A vector of unique Oracle databases. Each vector entry represents the backup/restore parameters for one unique Oracle database. Uniqueness is determined by the database unique name. | [optional] 
**PersistMountpoints** | Pointer to **NullableBool** | This parameter indicates whether or not to persist mountpoints. Default is set to true, which was the behavior before this option. | [optional] 

## Methods

### NewOracleSourceParams

`func NewOracleSourceParams() *OracleSourceParams`

NewOracleSourceParams instantiates a new OracleSourceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleSourceParamsWithDefaults

`func NewOracleSourceParamsWithDefaults() *OracleSourceParams`

NewOracleSourceParamsWithDefaults instantiates a new OracleSourceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalOracleDbParamsVec

`func (o *OracleSourceParams) GetAdditionalOracleDbParamsVec() []AdditionalOracleDBParams`

GetAdditionalOracleDbParamsVec returns the AdditionalOracleDbParamsVec field if non-nil, zero value otherwise.

### GetAdditionalOracleDbParamsVecOk

`func (o *OracleSourceParams) GetAdditionalOracleDbParamsVecOk() (*[]AdditionalOracleDBParams, bool)`

GetAdditionalOracleDbParamsVecOk returns a tuple with the AdditionalOracleDbParamsVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalOracleDbParamsVec

`func (o *OracleSourceParams) SetAdditionalOracleDbParamsVec(v []AdditionalOracleDBParams)`

SetAdditionalOracleDbParamsVec sets AdditionalOracleDbParamsVec field to given value.

### HasAdditionalOracleDbParamsVec

`func (o *OracleSourceParams) HasAdditionalOracleDbParamsVec() bool`

HasAdditionalOracleDbParamsVec returns a boolean if a field has been set.

### SetAdditionalOracleDbParamsVecNil

`func (o *OracleSourceParams) SetAdditionalOracleDbParamsVecNil(b bool)`

 SetAdditionalOracleDbParamsVecNil sets the value for AdditionalOracleDbParamsVec to be an explicit nil

### UnsetAdditionalOracleDbParamsVec
`func (o *OracleSourceParams) UnsetAdditionalOracleDbParamsVec()`

UnsetAdditionalOracleDbParamsVec ensures that no value is present for AdditionalOracleDbParamsVec, not even an explicit nil
### GetPersistMountpoints

`func (o *OracleSourceParams) GetPersistMountpoints() bool`

GetPersistMountpoints returns the PersistMountpoints field if non-nil, zero value otherwise.

### GetPersistMountpointsOk

`func (o *OracleSourceParams) GetPersistMountpointsOk() (*bool, bool)`

GetPersistMountpointsOk returns a tuple with the PersistMountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistMountpoints

`func (o *OracleSourceParams) SetPersistMountpoints(v bool)`

SetPersistMountpoints sets PersistMountpoints field to given value.

### HasPersistMountpoints

`func (o *OracleSourceParams) HasPersistMountpoints() bool`

HasPersistMountpoints returns a boolean if a field has been set.

### SetPersistMountpointsNil

`func (o *OracleSourceParams) SetPersistMountpointsNil(b bool)`

 SetPersistMountpointsNil sets the value for PersistMountpoints to be an explicit nil

### UnsetPersistMountpoints
`func (o *OracleSourceParams) UnsetPersistMountpoints()`

UnsetPersistMountpoints ensures that no value is present for PersistMountpoints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


