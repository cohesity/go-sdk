# ObjectsSearchResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** | Specifies the number of objects to be fetched for the specified pagination cookie. | [optional] 
**Objects** | Pointer to [**[]SearchObject**](SearchObject.md) | Specifies the list of Objects. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | [optional] 

## Methods

### NewObjectsSearchResponseBody

`func NewObjectsSearchResponseBody() *ObjectsSearchResponseBody`

NewObjectsSearchResponseBody instantiates a new ObjectsSearchResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectsSearchResponseBodyWithDefaults

`func NewObjectsSearchResponseBodyWithDefaults() *ObjectsSearchResponseBody`

NewObjectsSearchResponseBodyWithDefaults instantiates a new ObjectsSearchResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ObjectsSearchResponseBody) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ObjectsSearchResponseBody) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ObjectsSearchResponseBody) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ObjectsSearchResponseBody) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ObjectsSearchResponseBody) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ObjectsSearchResponseBody) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetObjects

`func (o *ObjectsSearchResponseBody) GetObjects() []SearchObject`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ObjectsSearchResponseBody) GetObjectsOk() (*[]SearchObject, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ObjectsSearchResponseBody) SetObjects(v []SearchObject)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ObjectsSearchResponseBody) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### SetObjectsNil

`func (o *ObjectsSearchResponseBody) SetObjectsNil(b bool)`

 SetObjectsNil sets the value for Objects to be an explicit nil

### UnsetObjects
`func (o *ObjectsSearchResponseBody) UnsetObjects()`

UnsetObjects ensures that no value is present for Objects, not even an explicit nil
### GetPaginationCookie

`func (o *ObjectsSearchResponseBody) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *ObjectsSearchResponseBody) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *ObjectsSearchResponseBody) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *ObjectsSearchResponseBody) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *ObjectsSearchResponseBody) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *ObjectsSearchResponseBody) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


