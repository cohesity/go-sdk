# DataTransferFromVaultsSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTransferSummary** | Pointer to [**[]DataTransferFromVaultSummary**](DataTransferFromVaultSummary.md) | Array of Summary Data Transfer Statistics.  Specifies summary statistics about the transfer of data from Vaults to the Cohesity Cluster. | [optional] 

## Methods

### NewDataTransferFromVaultsSummaryResponse

`func NewDataTransferFromVaultsSummaryResponse() *DataTransferFromVaultsSummaryResponse`

NewDataTransferFromVaultsSummaryResponse instantiates a new DataTransferFromVaultsSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferFromVaultsSummaryResponseWithDefaults

`func NewDataTransferFromVaultsSummaryResponseWithDefaults() *DataTransferFromVaultsSummaryResponse`

NewDataTransferFromVaultsSummaryResponseWithDefaults instantiates a new DataTransferFromVaultsSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTransferSummary

`func (o *DataTransferFromVaultsSummaryResponse) GetDataTransferSummary() []DataTransferFromVaultSummary`

GetDataTransferSummary returns the DataTransferSummary field if non-nil, zero value otherwise.

### GetDataTransferSummaryOk

`func (o *DataTransferFromVaultsSummaryResponse) GetDataTransferSummaryOk() (*[]DataTransferFromVaultSummary, bool)`

GetDataTransferSummaryOk returns a tuple with the DataTransferSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSummary

`func (o *DataTransferFromVaultsSummaryResponse) SetDataTransferSummary(v []DataTransferFromVaultSummary)`

SetDataTransferSummary sets DataTransferSummary field to given value.

### HasDataTransferSummary

`func (o *DataTransferFromVaultsSummaryResponse) HasDataTransferSummary() bool`

HasDataTransferSummary returns a boolean if a field has been set.

### SetDataTransferSummaryNil

`func (o *DataTransferFromVaultsSummaryResponse) SetDataTransferSummaryNil(b bool)`

 SetDataTransferSummaryNil sets the value for DataTransferSummary to be an explicit nil

### UnsetDataTransferSummary
`func (o *DataTransferFromVaultsSummaryResponse) UnsetDataTransferSummary()`

UnsetDataTransferSummary ensures that no value is present for DataTransferSummary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


