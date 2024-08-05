# ConnectorAzureMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureVmId** | **NullableString** | Azure VM id. | 
**MaxDataDiskCount** | **NullableInt64** | Maximum number of data disks which could be attached to rigel VM. | 

## Methods

### NewConnectorAzureMetadata

`func NewConnectorAzureMetadata(azureVmId NullableString, maxDataDiskCount NullableInt64, ) *ConnectorAzureMetadata`

NewConnectorAzureMetadata instantiates a new ConnectorAzureMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAzureMetadataWithDefaults

`func NewConnectorAzureMetadataWithDefaults() *ConnectorAzureMetadata`

NewConnectorAzureMetadataWithDefaults instantiates a new ConnectorAzureMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureVmId

`func (o *ConnectorAzureMetadata) GetAzureVmId() string`

GetAzureVmId returns the AzureVmId field if non-nil, zero value otherwise.

### GetAzureVmIdOk

`func (o *ConnectorAzureMetadata) GetAzureVmIdOk() (*string, bool)`

GetAzureVmIdOk returns a tuple with the AzureVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureVmId

`func (o *ConnectorAzureMetadata) SetAzureVmId(v string)`

SetAzureVmId sets AzureVmId field to given value.


### SetAzureVmIdNil

`func (o *ConnectorAzureMetadata) SetAzureVmIdNil(b bool)`

 SetAzureVmIdNil sets the value for AzureVmId to be an explicit nil

### UnsetAzureVmId
`func (o *ConnectorAzureMetadata) UnsetAzureVmId()`

UnsetAzureVmId ensures that no value is present for AzureVmId, not even an explicit nil
### GetMaxDataDiskCount

`func (o *ConnectorAzureMetadata) GetMaxDataDiskCount() int64`

GetMaxDataDiskCount returns the MaxDataDiskCount field if non-nil, zero value otherwise.

### GetMaxDataDiskCountOk

`func (o *ConnectorAzureMetadata) GetMaxDataDiskCountOk() (*int64, bool)`

GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataDiskCount

`func (o *ConnectorAzureMetadata) SetMaxDataDiskCount(v int64)`

SetMaxDataDiskCount sets MaxDataDiskCount field to given value.


### SetMaxDataDiskCountNil

`func (o *ConnectorAzureMetadata) SetMaxDataDiskCountNil(b bool)`

 SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil

### UnsetMaxDataDiskCount
`func (o *ConnectorAzureMetadata) UnsetMaxDataDiskCount()`

UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


