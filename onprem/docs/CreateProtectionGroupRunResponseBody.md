# CreateProtectionGroupRunResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionGroupId** | Pointer to **NullableString** | Specifies id of the Protection Group which must be polled for seeing the new run. | [optional] 

## Methods

### NewCreateProtectionGroupRunResponseBody

`func NewCreateProtectionGroupRunResponseBody() *CreateProtectionGroupRunResponseBody`

NewCreateProtectionGroupRunResponseBody instantiates a new CreateProtectionGroupRunResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProtectionGroupRunResponseBodyWithDefaults

`func NewCreateProtectionGroupRunResponseBodyWithDefaults() *CreateProtectionGroupRunResponseBody`

NewCreateProtectionGroupRunResponseBodyWithDefaults instantiates a new CreateProtectionGroupRunResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionGroupId

`func (o *CreateProtectionGroupRunResponseBody) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CreateProtectionGroupRunResponseBody) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CreateProtectionGroupRunResponseBody) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CreateProtectionGroupRunResponseBody) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CreateProtectionGroupRunResponseBody) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CreateProtectionGroupRunResponseBody) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


