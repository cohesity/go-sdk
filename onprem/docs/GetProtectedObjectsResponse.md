# GetProtectedObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]ProtectedObjectInfo**](ProtectedObjectInfo.md) | Specifies the protected object backup configuration and lastRun details if it has happned. | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewGetProtectedObjectsResponse

`func NewGetProtectedObjectsResponse() *GetProtectedObjectsResponse`

NewGetProtectedObjectsResponse instantiates a new GetProtectedObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProtectedObjectsResponseWithDefaults

`func NewGetProtectedObjectsResponseWithDefaults() *GetProtectedObjectsResponse`

NewGetProtectedObjectsResponseWithDefaults instantiates a new GetProtectedObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *GetProtectedObjectsResponse) GetObjects() []ProtectedObjectInfo`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetProtectedObjectsResponse) GetObjectsOk() (*[]ProtectedObjectInfo, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetProtectedObjectsResponse) SetObjects(v []ProtectedObjectInfo)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *GetProtectedObjectsResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *GetProtectedObjectsResponse) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *GetProtectedObjectsResponse) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPaginationInfo

`func (o *GetProtectedObjectsResponse) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *GetProtectedObjectsResponse) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *GetProtectedObjectsResponse) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *GetProtectedObjectsResponse) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


