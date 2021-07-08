# CloneAppViewParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountPathIdentifier** | Pointer to **NullableString** | Mount path identifier, which identifies the sub-dir where the cohesity view for App recovery will be mounted. | [optional] 

## Methods

### NewCloneAppViewParams

`func NewCloneAppViewParams() *CloneAppViewParams`

NewCloneAppViewParams instantiates a new CloneAppViewParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneAppViewParamsWithDefaults

`func NewCloneAppViewParamsWithDefaults() *CloneAppViewParams`

NewCloneAppViewParamsWithDefaults instantiates a new CloneAppViewParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountPathIdentifier

`func (o *CloneAppViewParams) GetMountPathIdentifier() string`

GetMountPathIdentifier returns the MountPathIdentifier field if non-nil, zero value otherwise.

### GetMountPathIdentifierOk

`func (o *CloneAppViewParams) GetMountPathIdentifierOk() (*string, bool)`

GetMountPathIdentifierOk returns a tuple with the MountPathIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPathIdentifier

`func (o *CloneAppViewParams) SetMountPathIdentifier(v string)`

SetMountPathIdentifier sets MountPathIdentifier field to given value.

### HasMountPathIdentifier

`func (o *CloneAppViewParams) HasMountPathIdentifier() bool`

HasMountPathIdentifier returns a boolean if a field has been set.

### SetMountPathIdentifierNil

`func (o *CloneAppViewParams) SetMountPathIdentifierNil(b bool)`

 SetMountPathIdentifierNil sets the value for MountPathIdentifier to be an explicit nil

### UnsetMountPathIdentifier
`func (o *CloneAppViewParams) UnsetMountPathIdentifier()`

UnsetMountPathIdentifier ensures that no value is present for MountPathIdentifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


