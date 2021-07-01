# AzureManagedDiskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataDisksSkuType** | Pointer to **NullableInt32** | SKU type for data disks. | [optional] 
**OsDiskSkuType** | Pointer to **NullableInt32** | SKU type for OS disk. | [optional] 

## Methods

### NewAzureManagedDiskParams

`func NewAzureManagedDiskParams() *AzureManagedDiskParams`

NewAzureManagedDiskParams instantiates a new AzureManagedDiskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureManagedDiskParamsWithDefaults

`func NewAzureManagedDiskParamsWithDefaults() *AzureManagedDiskParams`

NewAzureManagedDiskParamsWithDefaults instantiates a new AzureManagedDiskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataDisksSkuType

`func (o *AzureManagedDiskParams) GetDataDisksSkuType() int32`

GetDataDisksSkuType returns the DataDisksSkuType field if non-nil, zero value otherwise.

### GetDataDisksSkuTypeOk

`func (o *AzureManagedDiskParams) GetDataDisksSkuTypeOk() (*int32, bool)`

GetDataDisksSkuTypeOk returns a tuple with the DataDisksSkuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDisksSkuType

`func (o *AzureManagedDiskParams) SetDataDisksSkuType(v int32)`

SetDataDisksSkuType sets DataDisksSkuType field to given value.

### HasDataDisksSkuType

`func (o *AzureManagedDiskParams) HasDataDisksSkuType() bool`

HasDataDisksSkuType returns a boolean if a field has been set.

### SetDataDisksSkuTypeNil

`func (o *AzureManagedDiskParams) SetDataDisksSkuTypeNil(b bool)`

 SetDataDisksSkuTypeNil sets the value for DataDisksSkuType to be an explicit nil

### UnsetDataDisksSkuType
`func (o *AzureManagedDiskParams) UnsetDataDisksSkuType()`

UnsetDataDisksSkuType ensures that no value is present for DataDisksSkuType, not even an explicit nil
### GetOsDiskSkuType

`func (o *AzureManagedDiskParams) GetOsDiskSkuType() int32`

GetOsDiskSkuType returns the OsDiskSkuType field if non-nil, zero value otherwise.

### GetOsDiskSkuTypeOk

`func (o *AzureManagedDiskParams) GetOsDiskSkuTypeOk() (*int32, bool)`

GetOsDiskSkuTypeOk returns a tuple with the OsDiskSkuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDiskSkuType

`func (o *AzureManagedDiskParams) SetOsDiskSkuType(v int32)`

SetOsDiskSkuType sets OsDiskSkuType field to given value.

### HasOsDiskSkuType

`func (o *AzureManagedDiskParams) HasOsDiskSkuType() bool`

HasOsDiskSkuType returns a boolean if a field has been set.

### SetOsDiskSkuTypeNil

`func (o *AzureManagedDiskParams) SetOsDiskSkuTypeNil(b bool)`

 SetOsDiskSkuTypeNil sets the value for OsDiskSkuType to be an explicit nil

### UnsetOsDiskSkuType
`func (o *AzureManagedDiskParams) UnsetOsDiskSkuType()`

UnsetOsDiskSkuType ensures that no value is present for OsDiskSkuType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


