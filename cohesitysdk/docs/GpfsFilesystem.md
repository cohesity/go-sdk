# GpfsFilesystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Specifies the id of the file system. | [optional] 
**Path** | Pointer to **NullableString** | Specifies the path of the file system. | [optional] 

## Methods

### NewGpfsFilesystem

`func NewGpfsFilesystem() *GpfsFilesystem`

NewGpfsFilesystem instantiates a new GpfsFilesystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsFilesystemWithDefaults

`func NewGpfsFilesystemWithDefaults() *GpfsFilesystem`

NewGpfsFilesystemWithDefaults instantiates a new GpfsFilesystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GpfsFilesystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpfsFilesystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpfsFilesystem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GpfsFilesystem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GpfsFilesystem) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GpfsFilesystem) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *GpfsFilesystem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GpfsFilesystem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GpfsFilesystem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GpfsFilesystem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GpfsFilesystem) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GpfsFilesystem) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


