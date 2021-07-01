# ListStorageProfilesResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageProfiles** | Pointer to [**[]VcdStorageProfile**](VcdStorageProfile.md) | Specifies a list of storage profiles. | [optional] 

## Methods

### NewListStorageProfilesResponseBody

`func NewListStorageProfilesResponseBody() *ListStorageProfilesResponseBody`

NewListStorageProfilesResponseBody instantiates a new ListStorageProfilesResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageProfilesResponseBodyWithDefaults

`func NewListStorageProfilesResponseBodyWithDefaults() *ListStorageProfilesResponseBody`

NewListStorageProfilesResponseBodyWithDefaults instantiates a new ListStorageProfilesResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageProfiles

`func (o *ListStorageProfilesResponseBody) GetStorageProfiles() []VcdStorageProfile`

GetStorageProfiles returns the StorageProfiles field if non-nil, zero value otherwise.

### GetStorageProfilesOk

`func (o *ListStorageProfilesResponseBody) GetStorageProfilesOk() (*[]VcdStorageProfile, bool)`

GetStorageProfilesOk returns a tuple with the StorageProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfiles

`func (o *ListStorageProfilesResponseBody) SetStorageProfiles(v []VcdStorageProfile)`

SetStorageProfiles sets StorageProfiles field to given value.

### HasStorageProfiles

`func (o *ListStorageProfilesResponseBody) HasStorageProfiles() bool`

HasStorageProfiles returns a boolean if a field has been set.

### SetStorageProfilesNil

`func (o *ListStorageProfilesResponseBody) SetStorageProfilesNil(b bool)`

 SetStorageProfilesNil sets the value for StorageProfiles to be an explicit nil

### UnsetStorageProfiles
`func (o *ListStorageProfilesResponseBody) UnsetStorageProfiles()`

UnsetStorageProfiles ensures that no value is present for StorageProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


