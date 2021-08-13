# GetProtectedObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**ProtectedObjectInfo**](ProtectedObjectInfo.md) |  | [optional] 

## Methods

### NewGetProtectedObjectResponse

`func NewGetProtectedObjectResponse() *GetProtectedObjectResponse`

NewGetProtectedObjectResponse instantiates a new GetProtectedObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProtectedObjectResponseWithDefaults

`func NewGetProtectedObjectResponseWithDefaults() *GetProtectedObjectResponse`

NewGetProtectedObjectResponseWithDefaults instantiates a new GetProtectedObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *GetProtectedObjectResponse) GetObject() ProtectedObjectInfo`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetProtectedObjectResponse) GetObjectOk() (*ProtectedObjectInfo, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetProtectedObjectResponse) SetObject(v ProtectedObjectInfo)`

SetObject sets Object field to given value.

### HasObject

`func (o *GetProtectedObjectResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


