# CommonSearchIndexedObjectsResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** | Specifies the object type. | [optional] 
**Count** | Pointer to **NullableInt32** | Specifies the total number of indexed objects that match the filter and search criteria. Use this value to determine how many additional requests are required to get the full result. | [optional] 
**PaginationCookie** | Pointer to **NullableString** | Specifies cookie for resuming search if pagination is being used. | [optional] 

## Methods

### NewCommonSearchIndexedObjectsResponseParams

`func NewCommonSearchIndexedObjectsResponseParams() *CommonSearchIndexedObjectsResponseParams`

NewCommonSearchIndexedObjectsResponseParams instantiates a new CommonSearchIndexedObjectsResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSearchIndexedObjectsResponseParamsWithDefaults

`func NewCommonSearchIndexedObjectsResponseParamsWithDefaults() *CommonSearchIndexedObjectsResponseParams`

NewCommonSearchIndexedObjectsResponseParamsWithDefaults instantiates a new CommonSearchIndexedObjectsResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *CommonSearchIndexedObjectsResponseParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommonSearchIndexedObjectsResponseParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommonSearchIndexedObjectsResponseParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CommonSearchIndexedObjectsResponseParams) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetCount

`func (o *CommonSearchIndexedObjectsResponseParams) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CommonSearchIndexedObjectsResponseParams) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CommonSearchIndexedObjectsResponseParams) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CommonSearchIndexedObjectsResponseParams) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CommonSearchIndexedObjectsResponseParams) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CommonSearchIndexedObjectsResponseParams) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetPaginationCookie

`func (o *CommonSearchIndexedObjectsResponseParams) GetPaginationCookie() string`

GetPaginationCookie returns the PaginationCookie field if non-nil, zero value otherwise.

### GetPaginationCookieOk

`func (o *CommonSearchIndexedObjectsResponseParams) GetPaginationCookieOk() (*string, bool)`

GetPaginationCookieOk returns a tuple with the PaginationCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginationCookie

`func (o *CommonSearchIndexedObjectsResponseParams) SetPaginationCookie(v string)`

SetPaginationCookie sets PaginationCookie field to given value.

### HasPaginationCookie

`func (o *CommonSearchIndexedObjectsResponseParams) HasPaginationCookie() bool`

HasPaginationCookie returns a boolean if a field has been set.

### SetPaginationCookieNil

`func (o *CommonSearchIndexedObjectsResponseParams) SetPaginationCookieNil(b bool)`

 SetPaginationCookieNil sets the value for PaginationCookie to be an explicit nil

### UnsetPaginationCookie
`func (o *CommonSearchIndexedObjectsResponseParams) UnsetPaginationCookie()`

UnsetPaginationCookie ensures that no value is present for PaginationCookie, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


