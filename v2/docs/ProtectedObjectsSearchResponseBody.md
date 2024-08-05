# ProtectedObjectsSearchResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ProtectedObjectsSearchResponseBodyMetadata**](ProtectedObjectsSearchResponseBodyMetadata.md) |  | [optional] 
**NumResults** | Pointer to **NullableInt64** | Specifies the total number of search results which matches the search criteria. | [optional] 
**Objects** | Pointer to [**[]ProtectedObject**](ProtectedObject.md) | Specifies the list of Protected Objects. | [optional] 

## Methods

### NewProtectedObjectsSearchResponseBody

`func NewProtectedObjectsSearchResponseBody() *ProtectedObjectsSearchResponseBody`

NewProtectedObjectsSearchResponseBody instantiates a new ProtectedObjectsSearchResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedObjectsSearchResponseBodyWithDefaults

`func NewProtectedObjectsSearchResponseBodyWithDefaults() *ProtectedObjectsSearchResponseBody`

NewProtectedObjectsSearchResponseBodyWithDefaults instantiates a new ProtectedObjectsSearchResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ProtectedObjectsSearchResponseBody) GetMetadata() ProtectedObjectsSearchResponseBodyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProtectedObjectsSearchResponseBody) GetMetadataOk() (*ProtectedObjectsSearchResponseBodyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProtectedObjectsSearchResponseBody) SetMetadata(v ProtectedObjectsSearchResponseBodyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProtectedObjectsSearchResponseBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNumResults

`func (o *ProtectedObjectsSearchResponseBody) GetNumResults() int64`

GetNumResults returns the NumResults field if non-nil, zero value otherwise.

### GetNumResultsOk

`func (o *ProtectedObjectsSearchResponseBody) GetNumResultsOk() (*int64, bool)`

GetNumResultsOk returns a tuple with the NumResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumResults

`func (o *ProtectedObjectsSearchResponseBody) SetNumResults(v int64)`

SetNumResults sets NumResults field to given value.

### HasNumResults

`func (o *ProtectedObjectsSearchResponseBody) HasNumResults() bool`

HasNumResults returns a boolean if a field has been set.

### SetNumResultsNil

`func (o *ProtectedObjectsSearchResponseBody) SetNumResultsNil(b bool)`

 SetNumResultsNil sets the value for NumResults to be an explicit nil

### UnsetNumResults
`func (o *ProtectedObjectsSearchResponseBody) UnsetNumResults()`

UnsetNumResults ensures that no value is present for NumResults, not even an explicit nil
### GetObjects

`func (o *ProtectedObjectsSearchResponseBody) GetObjects() []ProtectedObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ProtectedObjectsSearchResponseBody) GetObjectsOk() (*[]ProtectedObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ProtectedObjectsSearchResponseBody) SetObjects(v []ProtectedObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ProtectedObjectsSearchResponseBody) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ProtectedObjectsSearchResponseBody) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ProtectedObjectsSearchResponseBody) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


