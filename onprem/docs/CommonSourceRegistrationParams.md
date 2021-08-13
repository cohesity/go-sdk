# CommonSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Specifies the username to access target entity. | 
**Password** | **string** | Specifies the password to access target entity. | 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 

## Methods

### NewCommonSourceRegistrationParams

`func NewCommonSourceRegistrationParams(username string, password string, endpoint string, ) *CommonSourceRegistrationParams`

NewCommonSourceRegistrationParams instantiates a new CommonSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSourceRegistrationParamsWithDefaults

`func NewCommonSourceRegistrationParamsWithDefaults() *CommonSourceRegistrationParams`

NewCommonSourceRegistrationParamsWithDefaults instantiates a new CommonSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CommonSourceRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CommonSourceRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CommonSourceRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CommonSourceRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CommonSourceRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CommonSourceRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEndpoint

`func (o *CommonSourceRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CommonSourceRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CommonSourceRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDescription

`func (o *CommonSourceRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonSourceRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonSourceRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonSourceRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommonSourceRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommonSourceRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


