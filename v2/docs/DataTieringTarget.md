# DataTieringTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPath** | Pointer to **NullableString** | Specifies the mount path inside the view. | [optional] 
**StorageDomainId** | **NullableInt64** | Specifies the Storage Domain ID where the view will be kept. | 
**ViewName** | **NullableString** | Specifies the view name. | 

## Methods

### NewDataTieringTarget

`func NewDataTieringTarget(storageDomainId NullableInt64, viewName NullableString, ) *DataTieringTarget`

NewDataTieringTarget instantiates a new DataTieringTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTieringTargetWithDefaults

`func NewDataTieringTargetWithDefaults() *DataTieringTarget`

NewDataTieringTargetWithDefaults instantiates a new DataTieringTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPath

`func (o *DataTieringTarget) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *DataTieringTarget) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *DataTieringTarget) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *DataTieringTarget) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *DataTieringTarget) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *DataTieringTarget) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil
### GetStorageDomainId

`func (o *DataTieringTarget) GetStorageDomainId() int64`

GetStorageDomainId returns the StorageDomainId field if non-nil, zero value otherwise.

### GetStorageDomainIdOk

`func (o *DataTieringTarget) GetStorageDomainIdOk() (*int64, bool)`

GetStorageDomainIdOk returns a tuple with the StorageDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDomainId

`func (o *DataTieringTarget) SetStorageDomainId(v int64)`

SetStorageDomainId sets StorageDomainId field to given value.


### SetStorageDomainIdNil

`func (o *DataTieringTarget) SetStorageDomainIdNil(b bool)`

 SetStorageDomainIdNil sets the value for StorageDomainId to be an explicit nil

### UnsetStorageDomainId
`func (o *DataTieringTarget) UnsetStorageDomainId()`

UnsetStorageDomainId ensures that no value is present for StorageDomainId, not even an explicit nil
### GetViewName

`func (o *DataTieringTarget) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DataTieringTarget) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DataTieringTarget) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### SetViewNameNil

`func (o *DataTieringTarget) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *DataTieringTarget) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


