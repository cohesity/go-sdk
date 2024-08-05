# AzureNativeObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTransferInfo** | Pointer to [**DataTransferInfo**](DataTransferInfo.md) |  | [optional] 
**DiskExclusionParams** | Pointer to [**AzureDiskExclusionParams**](AzureDiskExclusionParams.md) |  | [optional] 
**Objects** | Pointer to [**[]AzureObjectLevelParams**](AzureObjectLevelParams.md) | Specifies the objects to be protected. | [optional] 

## Methods

### NewAzureNativeObjectProtectionParams

`func NewAzureNativeObjectProtectionParams() *AzureNativeObjectProtectionParams`

NewAzureNativeObjectProtectionParams instantiates a new AzureNativeObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureNativeObjectProtectionParamsWithDefaults

`func NewAzureNativeObjectProtectionParamsWithDefaults() *AzureNativeObjectProtectionParams`

NewAzureNativeObjectProtectionParamsWithDefaults instantiates a new AzureNativeObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTransferInfo

`func (o *AzureNativeObjectProtectionParams) GetDataTransferInfo() DataTransferInfo`

GetDataTransferInfo returns the DataTransferInfo field if non-nil, zero value otherwise.

### GetDataTransferInfoOk

`func (o *AzureNativeObjectProtectionParams) GetDataTransferInfoOk() (*DataTransferInfo, bool)`

GetDataTransferInfoOk returns a tuple with the DataTransferInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferInfo

`func (o *AzureNativeObjectProtectionParams) SetDataTransferInfo(v DataTransferInfo)`

SetDataTransferInfo sets DataTransferInfo field to given value.

### HasDataTransferInfo

`func (o *AzureNativeObjectProtectionParams) HasDataTransferInfo() bool`

HasDataTransferInfo returns a boolean if a field has been set.

### GetDiskExclusionParams

`func (o *AzureNativeObjectProtectionParams) GetDiskExclusionParams() AzureDiskExclusionParams`

GetDiskExclusionParams returns the DiskExclusionParams field if non-nil, zero value otherwise.

### GetDiskExclusionParamsOk

`func (o *AzureNativeObjectProtectionParams) GetDiskExclusionParamsOk() (*AzureDiskExclusionParams, bool)`

GetDiskExclusionParamsOk returns a tuple with the DiskExclusionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskExclusionParams

`func (o *AzureNativeObjectProtectionParams) SetDiskExclusionParams(v AzureDiskExclusionParams)`

SetDiskExclusionParams sets DiskExclusionParams field to given value.

### HasDiskExclusionParams

`func (o *AzureNativeObjectProtectionParams) HasDiskExclusionParams() bool`

HasDiskExclusionParams returns a boolean if a field has been set.

### GetObjects

`func (o *AzureNativeObjectProtectionParams) GetObjects() []AzureObjectLevelParams`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *AzureNativeObjectProtectionParams) GetObjectsOk() (*[]AzureObjectLevelParams, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *AzureNativeObjectProtectionParams) SetObjects(v []AzureObjectLevelParams)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *AzureNativeObjectProtectionParams) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


