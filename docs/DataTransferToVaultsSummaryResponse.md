# DataTransferToVaultsSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTransferSummary** | Pointer to [**[]DataTransferToVaultSummary**](DataTransferToVaultSummary.md) | Array of Summary Data Transfer Statistics.  Specifies summary statistics about the transfer of data from the Cohesity Cluster to Vaults. | [optional] 

## Methods

### NewDataTransferToVaultsSummaryResponse

`func NewDataTransferToVaultsSummaryResponse() *DataTransferToVaultsSummaryResponse`

NewDataTransferToVaultsSummaryResponse instantiates a new DataTransferToVaultsSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferToVaultsSummaryResponseWithDefaults

`func NewDataTransferToVaultsSummaryResponseWithDefaults() *DataTransferToVaultsSummaryResponse`

NewDataTransferToVaultsSummaryResponseWithDefaults instantiates a new DataTransferToVaultsSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTransferSummary

`func (o *DataTransferToVaultsSummaryResponse) GetDataTransferSummary() []DataTransferToVaultSummary`

GetDataTransferSummary returns the DataTransferSummary field if non-nil, zero value otherwise.

### GetDataTransferSummaryOk

`func (o *DataTransferToVaultsSummaryResponse) GetDataTransferSummaryOk() (*[]DataTransferToVaultSummary, bool)`

GetDataTransferSummaryOk returns a tuple with the DataTransferSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSummary

`func (o *DataTransferToVaultsSummaryResponse) SetDataTransferSummary(v []DataTransferToVaultSummary)`

SetDataTransferSummary sets DataTransferSummary field to given value.

### HasDataTransferSummary

`func (o *DataTransferToVaultsSummaryResponse) HasDataTransferSummary() bool`

HasDataTransferSummary returns a boolean if a field has been set.

### SetDataTransferSummaryNil

`func (o *DataTransferToVaultsSummaryResponse) SetDataTransferSummaryNil(b bool)`

 SetDataTransferSummaryNil sets the value for DataTransferSummary to be an explicit nil

### UnsetDataTransferSummary
`func (o *DataTransferToVaultsSummaryResponse) UnsetDataTransferSummary()`

UnsetDataTransferSummary ensures that no value is present for DataTransferSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


