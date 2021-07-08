# CassandraBackupJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CassandraAdditionalInfo** | Pointer to [**CassandraAdditionalParams**](CassandraAdditionalParams.md) |  | [optional] 
**SelectedDataCenterVec** | Pointer to **[]string** | The data centers selected for backup. | [optional] 

## Methods

### NewCassandraBackupJobParams

`func NewCassandraBackupJobParams() *CassandraBackupJobParams`

NewCassandraBackupJobParams instantiates a new CassandraBackupJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCassandraBackupJobParamsWithDefaults

`func NewCassandraBackupJobParamsWithDefaults() *CassandraBackupJobParams`

NewCassandraBackupJobParamsWithDefaults instantiates a new CassandraBackupJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCassandraAdditionalInfo

`func (o *CassandraBackupJobParams) GetCassandraAdditionalInfo() CassandraAdditionalParams`

GetCassandraAdditionalInfo returns the CassandraAdditionalInfo field if non-nil, zero value otherwise.

### GetCassandraAdditionalInfoOk

`func (o *CassandraBackupJobParams) GetCassandraAdditionalInfoOk() (*CassandraAdditionalParams, bool)`

GetCassandraAdditionalInfoOk returns a tuple with the CassandraAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraAdditionalInfo

`func (o *CassandraBackupJobParams) SetCassandraAdditionalInfo(v CassandraAdditionalParams)`

SetCassandraAdditionalInfo sets CassandraAdditionalInfo field to given value.

### HasCassandraAdditionalInfo

`func (o *CassandraBackupJobParams) HasCassandraAdditionalInfo() bool`

HasCassandraAdditionalInfo returns a boolean if a field has been set.

### GetSelectedDataCenterVec

`func (o *CassandraBackupJobParams) GetSelectedDataCenterVec() []string`

GetSelectedDataCenterVec returns the SelectedDataCenterVec field if non-nil, zero value otherwise.

### GetSelectedDataCenterVecOk

`func (o *CassandraBackupJobParams) GetSelectedDataCenterVecOk() (*[]string, bool)`

GetSelectedDataCenterVecOk returns a tuple with the SelectedDataCenterVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDataCenterVec

`func (o *CassandraBackupJobParams) SetSelectedDataCenterVec(v []string)`

SetSelectedDataCenterVec sets SelectedDataCenterVec field to given value.

### HasSelectedDataCenterVec

`func (o *CassandraBackupJobParams) HasSelectedDataCenterVec() bool`

HasSelectedDataCenterVec returns a boolean if a field has been set.

### SetSelectedDataCenterVecNil

`func (o *CassandraBackupJobParams) SetSelectedDataCenterVecNil(b bool)`

 SetSelectedDataCenterVecNil sets the value for SelectedDataCenterVec to be an explicit nil

### UnsetSelectedDataCenterVec
`func (o *CassandraBackupJobParams) UnsetSelectedDataCenterVec()`

UnsetSelectedDataCenterVec ensures that no value is present for SelectedDataCenterVec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


