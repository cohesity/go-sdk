# ErrorClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewErrorClass

`func NewErrorClass() *ErrorClass`

NewErrorClass instantiates a new ErrorClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorClassWithDefaults

`func NewErrorClassWithDefaults() *ErrorClass`

NewErrorClassWithDefaults instantiates a new ErrorClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *ErrorClass) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *ErrorClass) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *ErrorClass) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *ErrorClass) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetCount

`func (o *ErrorClass) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ErrorClass) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ErrorClass) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ErrorClass) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ErrorClass) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ErrorClass) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


