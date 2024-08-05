# CreateBifrostConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Specifies the name of the connection. | 

## Methods

### NewCreateBifrostConnectionRequest

`func NewCreateBifrostConnectionRequest(name NullableString, ) *CreateBifrostConnectionRequest`

NewCreateBifrostConnectionRequest instantiates a new CreateBifrostConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBifrostConnectionRequestWithDefaults

`func NewCreateBifrostConnectionRequestWithDefaults() *CreateBifrostConnectionRequest`

NewCreateBifrostConnectionRequestWithDefaults instantiates a new CreateBifrostConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBifrostConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBifrostConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBifrostConnectionRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateBifrostConnectionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateBifrostConnectionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


