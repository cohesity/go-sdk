# DowntieredDataLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the ID of the object. | 
**ViewName** | **NullableString** | Specifies the view name. | 
**MountPath** | Pointer to **NullableString** | Specifies the mount path inside the view. | [optional] 

## Methods

### NewDowntieredDataLocation

`func NewDowntieredDataLocation(id NullableInt64, viewName NullableString, ) *DowntieredDataLocation`

NewDowntieredDataLocation instantiates a new DowntieredDataLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDowntieredDataLocationWithDefaults

`func NewDowntieredDataLocationWithDefaults() *DowntieredDataLocation`

NewDowntieredDataLocationWithDefaults instantiates a new DowntieredDataLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DowntieredDataLocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DowntieredDataLocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DowntieredDataLocation) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *DowntieredDataLocation) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DowntieredDataLocation) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetViewName

`func (o *DowntieredDataLocation) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DowntieredDataLocation) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DowntieredDataLocation) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### SetViewNameNil

`func (o *DowntieredDataLocation) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *DowntieredDataLocation) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetMountPath

`func (o *DowntieredDataLocation) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *DowntieredDataLocation) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *DowntieredDataLocation) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *DowntieredDataLocation) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### SetMountPathNil

`func (o *DowntieredDataLocation) SetMountPathNil(b bool)`

 SetMountPathNil sets the value for MountPath to be an explicit nil

### UnsetMountPath
`func (o *DowntieredDataLocation) UnsetMountPath()`

UnsetMountPath ensures that no value is present for MountPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


