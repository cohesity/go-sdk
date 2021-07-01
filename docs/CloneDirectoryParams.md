# CloneDirectoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDirectoryName** | Pointer to **NullableString** | Name of the new directory which will contain the clone contents. | [optional] 
**DestinationParentDirectoryPath** | Pointer to **NullableString** | Specifies the path of the destination parent directory. The source dir would be cloned as a child of destination parent dir. | [optional] 
**SourceDirectoryPath** | Pointer to **NullableString** | Specifies the path of the source directory | [optional] 

## Methods

### NewCloneDirectoryParams

`func NewCloneDirectoryParams() *CloneDirectoryParams`

NewCloneDirectoryParams instantiates a new CloneDirectoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneDirectoryParamsWithDefaults

`func NewCloneDirectoryParamsWithDefaults() *CloneDirectoryParams`

NewCloneDirectoryParamsWithDefaults instantiates a new CloneDirectoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDirectoryName

`func (o *CloneDirectoryParams) GetDestinationDirectoryName() string`

GetDestinationDirectoryName returns the DestinationDirectoryName field if non-nil, zero value otherwise.

### GetDestinationDirectoryNameOk

`func (o *CloneDirectoryParams) GetDestinationDirectoryNameOk() (*string, bool)`

GetDestinationDirectoryNameOk returns a tuple with the DestinationDirectoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDirectoryName

`func (o *CloneDirectoryParams) SetDestinationDirectoryName(v string)`

SetDestinationDirectoryName sets DestinationDirectoryName field to given value.

### HasDestinationDirectoryName

`func (o *CloneDirectoryParams) HasDestinationDirectoryName() bool`

HasDestinationDirectoryName returns a boolean if a field has been set.

### SetDestinationDirectoryNameNil

`func (o *CloneDirectoryParams) SetDestinationDirectoryNameNil(b bool)`

 SetDestinationDirectoryNameNil sets the value for DestinationDirectoryName to be an explicit nil

### UnsetDestinationDirectoryName
`func (o *CloneDirectoryParams) UnsetDestinationDirectoryName()`

UnsetDestinationDirectoryName ensures that no value is present for DestinationDirectoryName, not even an explicit nil
### GetDestinationParentDirectoryPath

`func (o *CloneDirectoryParams) GetDestinationParentDirectoryPath() string`

GetDestinationParentDirectoryPath returns the DestinationParentDirectoryPath field if non-nil, zero value otherwise.

### GetDestinationParentDirectoryPathOk

`func (o *CloneDirectoryParams) GetDestinationParentDirectoryPathOk() (*string, bool)`

GetDestinationParentDirectoryPathOk returns a tuple with the DestinationParentDirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationParentDirectoryPath

`func (o *CloneDirectoryParams) SetDestinationParentDirectoryPath(v string)`

SetDestinationParentDirectoryPath sets DestinationParentDirectoryPath field to given value.

### HasDestinationParentDirectoryPath

`func (o *CloneDirectoryParams) HasDestinationParentDirectoryPath() bool`

HasDestinationParentDirectoryPath returns a boolean if a field has been set.

### SetDestinationParentDirectoryPathNil

`func (o *CloneDirectoryParams) SetDestinationParentDirectoryPathNil(b bool)`

 SetDestinationParentDirectoryPathNil sets the value for DestinationParentDirectoryPath to be an explicit nil

### UnsetDestinationParentDirectoryPath
`func (o *CloneDirectoryParams) UnsetDestinationParentDirectoryPath()`

UnsetDestinationParentDirectoryPath ensures that no value is present for DestinationParentDirectoryPath, not even an explicit nil
### GetSourceDirectoryPath

`func (o *CloneDirectoryParams) GetSourceDirectoryPath() string`

GetSourceDirectoryPath returns the SourceDirectoryPath field if non-nil, zero value otherwise.

### GetSourceDirectoryPathOk

`func (o *CloneDirectoryParams) GetSourceDirectoryPathOk() (*string, bool)`

GetSourceDirectoryPathOk returns a tuple with the SourceDirectoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirectoryPath

`func (o *CloneDirectoryParams) SetSourceDirectoryPath(v string)`

SetSourceDirectoryPath sets SourceDirectoryPath field to given value.

### HasSourceDirectoryPath

`func (o *CloneDirectoryParams) HasSourceDirectoryPath() bool`

HasSourceDirectoryPath returns a boolean if a field has been set.

### SetSourceDirectoryPathNil

`func (o *CloneDirectoryParams) SetSourceDirectoryPathNil(b bool)`

 SetSourceDirectoryPathNil sets the value for SourceDirectoryPath to be an explicit nil

### UnsetSourceDirectoryPath
`func (o *CloneDirectoryParams) UnsetSourceDirectoryPath()`

UnsetSourceDirectoryPath ensures that no value is present for SourceDirectoryPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


