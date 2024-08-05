# DataLocator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeDiskInfo** | Pointer to [**NodeDiskLocation**](NodeDiskLocation.md) |  | [optional] 
**TopLevelFolder** | Pointer to **string** | Denotes the top-level-directory in one of the above locations. | [optional] 
**Type** | Pointer to **string** | The type of the location. | [optional] 
**VaultInfo** | Pointer to [**VaultLocation**](VaultLocation.md) |  | [optional] 
**ViewInfo** | Pointer to [**ViewLocation**](ViewLocation.md) |  | [optional] 

## Methods

### NewDataLocator

`func NewDataLocator() *DataLocator`

NewDataLocator instantiates a new DataLocator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLocatorWithDefaults

`func NewDataLocatorWithDefaults() *DataLocator`

NewDataLocatorWithDefaults instantiates a new DataLocator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeDiskInfo

`func (o *DataLocator) GetNodeDiskInfo() NodeDiskLocation`

GetNodeDiskInfo returns the NodeDiskInfo field if non-nil, zero value otherwise.

### GetNodeDiskInfoOk

`func (o *DataLocator) GetNodeDiskInfoOk() (*NodeDiskLocation, bool)`

GetNodeDiskInfoOk returns a tuple with the NodeDiskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDiskInfo

`func (o *DataLocator) SetNodeDiskInfo(v NodeDiskLocation)`

SetNodeDiskInfo sets NodeDiskInfo field to given value.

### HasNodeDiskInfo

`func (o *DataLocator) HasNodeDiskInfo() bool`

HasNodeDiskInfo returns a boolean if a field has been set.

### GetTopLevelFolder

`func (o *DataLocator) GetTopLevelFolder() string`

GetTopLevelFolder returns the TopLevelFolder field if non-nil, zero value otherwise.

### GetTopLevelFolderOk

`func (o *DataLocator) GetTopLevelFolderOk() (*string, bool)`

GetTopLevelFolderOk returns a tuple with the TopLevelFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelFolder

`func (o *DataLocator) SetTopLevelFolder(v string)`

SetTopLevelFolder sets TopLevelFolder field to given value.

### HasTopLevelFolder

`func (o *DataLocator) HasTopLevelFolder() bool`

HasTopLevelFolder returns a boolean if a field has been set.

### GetType

`func (o *DataLocator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataLocator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataLocator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataLocator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVaultInfo

`func (o *DataLocator) GetVaultInfo() VaultLocation`

GetVaultInfo returns the VaultInfo field if non-nil, zero value otherwise.

### GetVaultInfoOk

`func (o *DataLocator) GetVaultInfoOk() (*VaultLocation, bool)`

GetVaultInfoOk returns a tuple with the VaultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultInfo

`func (o *DataLocator) SetVaultInfo(v VaultLocation)`

SetVaultInfo sets VaultInfo field to given value.

### HasVaultInfo

`func (o *DataLocator) HasVaultInfo() bool`

HasVaultInfo returns a boolean if a field has been set.

### GetViewInfo

`func (o *DataLocator) GetViewInfo() ViewLocation`

GetViewInfo returns the ViewInfo field if non-nil, zero value otherwise.

### GetViewInfoOk

`func (o *DataLocator) GetViewInfoOk() (*ViewLocation, bool)`

GetViewInfoOk returns a tuple with the ViewInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewInfo

`func (o *DataLocator) SetViewInfo(v ViewLocation)`

SetViewInfo sets ViewInfo field to given value.

### HasViewInfo

`func (o *DataLocator) HasViewInfo() bool`

HasViewInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


