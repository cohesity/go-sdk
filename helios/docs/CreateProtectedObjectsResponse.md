# CreateProtectedObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectedObjects** | Pointer to [**[]ObjectProtectionSummary**](ObjectProtectionSummary.md) | Specifies the list of protected objects. | [optional] 

## Methods

### NewCreateProtectedObjectsResponse

`func NewCreateProtectedObjectsResponse() *CreateProtectedObjectsResponse`

NewCreateProtectedObjectsResponse instantiates a new CreateProtectedObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProtectedObjectsResponseWithDefaults

`func NewCreateProtectedObjectsResponseWithDefaults() *CreateProtectedObjectsResponse`

NewCreateProtectedObjectsResponseWithDefaults instantiates a new CreateProtectedObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectedObjects

`func (o *CreateProtectedObjectsResponse) GetProtectedObjects() []ObjectProtectionSummary`

GetProtectedObjects returns the ProtectedObjects field if non-nil, zero value otherwise.

### GetProtectedObjectsOk

`func (o *CreateProtectedObjectsResponse) GetProtectedObjectsOk() (*[]ObjectProtectionSummary, bool)`

GetProtectedObjectsOk returns a tuple with the ProtectedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedObjects

`func (o *CreateProtectedObjectsResponse) SetProtectedObjects(v []ObjectProtectionSummary)`

SetProtectedObjects sets ProtectedObjects field to given value.

### HasProtectedObjects

`func (o *CreateProtectedObjectsResponse) HasProtectedObjects() bool`

HasProtectedObjects returns a boolean if a field has been set.

### SetProtectedObjectsNil

`func (o *CreateProtectedObjectsResponse) SetProtectedObjectsNil(b bool)`

 SetProtectedObjectsNil sets the value for ProtectedObjects to be an explicit nil

### UnsetProtectedObjects
`func (o *CreateProtectedObjectsResponse) UnsetProtectedObjects()`

UnsetProtectedObjects ensures that no value is present for ProtectedObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


