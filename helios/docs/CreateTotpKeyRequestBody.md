# CreateTotpKeyRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpKeyName** | Pointer to **NullableString** | Specifies the key name. | [optional] 

## Methods

### NewCreateTotpKeyRequestBody

`func NewCreateTotpKeyRequestBody() *CreateTotpKeyRequestBody`

NewCreateTotpKeyRequestBody instantiates a new CreateTotpKeyRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTotpKeyRequestBodyWithDefaults

`func NewCreateTotpKeyRequestBodyWithDefaults() *CreateTotpKeyRequestBody`

NewCreateTotpKeyRequestBodyWithDefaults instantiates a new CreateTotpKeyRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpKeyName

`func (o *CreateTotpKeyRequestBody) GetTotpKeyName() string`

GetTotpKeyName returns the TotpKeyName field if non-nil, zero value otherwise.

### GetTotpKeyNameOk

`func (o *CreateTotpKeyRequestBody) GetTotpKeyNameOk() (*string, bool)`

GetTotpKeyNameOk returns a tuple with the TotpKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpKeyName

`func (o *CreateTotpKeyRequestBody) SetTotpKeyName(v string)`

SetTotpKeyName sets TotpKeyName field to given value.

### HasTotpKeyName

`func (o *CreateTotpKeyRequestBody) HasTotpKeyName() bool`

HasTotpKeyName returns a boolean if a field has been set.

### SetTotpKeyNameNil

`func (o *CreateTotpKeyRequestBody) SetTotpKeyNameNil(b bool)`

 SetTotpKeyNameNil sets the value for TotpKeyName to be an explicit nil

### UnsetTotpKeyName
`func (o *CreateTotpKeyRequestBody) UnsetTotpKeyName()`

UnsetTotpKeyName ensures that no value is present for TotpKeyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


