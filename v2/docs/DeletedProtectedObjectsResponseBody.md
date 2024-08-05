# DeletedProtectedObjectsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]ProtectedObjectInfo**](ProtectedObjectInfo.md) | Specifies the protected deleted objects. | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewDeletedProtectedObjectsResponseBody

`func NewDeletedProtectedObjectsResponseBody() *DeletedProtectedObjectsResponseBody`

NewDeletedProtectedObjectsResponseBody instantiates a new DeletedProtectedObjectsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedProtectedObjectsResponseBodyWithDefaults

`func NewDeletedProtectedObjectsResponseBodyWithDefaults() *DeletedProtectedObjectsResponseBody`

NewDeletedProtectedObjectsResponseBodyWithDefaults instantiates a new DeletedProtectedObjectsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *DeletedProtectedObjectsResponseBody) GetObjects() []ProtectedObjectInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *DeletedProtectedObjectsResponseBody) GetObjectsOk() (*[]ProtectedObjectInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *DeletedProtectedObjectsResponseBody) SetObjects(v []ProtectedObjectInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *DeletedProtectedObjectsResponseBody) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *DeletedProtectedObjectsResponseBody) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *DeletedProtectedObjectsResponseBody) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPaginationInfo

`func (o *DeletedProtectedObjectsResponseBody) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *DeletedProtectedObjectsResponseBody) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *DeletedProtectedObjectsResponseBody) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *DeletedProtectedObjectsResponseBody) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


