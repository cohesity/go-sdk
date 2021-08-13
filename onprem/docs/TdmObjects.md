# TdmObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | Pointer to [**[]TdmObject**](TdmObject.md) | Specifies the collection of TDM objects, filtered by the specified criteria. | [optional] 
**PaginationInfo** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewTdmObjects

`func NewTdmObjects() *TdmObjects`

NewTdmObjects instantiates a new TdmObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTdmObjectsWithDefaults

`func NewTdmObjectsWithDefaults() *TdmObjects`

NewTdmObjectsWithDefaults instantiates a new TdmObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *TdmObjects) GetObjects() []TdmObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *TdmObjects) GetObjectsOk() (*[]TdmObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *TdmObjects) SetObjects(v []TdmObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *TdmObjects) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetPaginationInfo

`func (o *TdmObjects) GetPaginationInfo() PaginationInfo`

GetPaginationInfo returns the PaginationInfo field if non-nil, zero value otherwise.

### GetPaginationInfoOk

`func (o *TdmObjects) GetPaginationInfoOk() (*PaginationInfo, bool)`

GetPaginationInfoOk returns a tuple with the PaginationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationInfo

`func (o *TdmObjects) SetPaginationInfo(v PaginationInfo)`

SetPaginationInfo sets PaginationInfo field to given value.

### HasPaginationInfo

`func (o *TdmObjects) HasPaginationInfo() bool`

HasPaginationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


