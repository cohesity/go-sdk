# EndpointOnlyRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 

## Methods

### NewEndpointOnlyRegistrationParams

`func NewEndpointOnlyRegistrationParams(endpoint string, ) *EndpointOnlyRegistrationParams`

NewEndpointOnlyRegistrationParams instantiates a new EndpointOnlyRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointOnlyRegistrationParamsWithDefaults

`func NewEndpointOnlyRegistrationParamsWithDefaults() *EndpointOnlyRegistrationParams`

NewEndpointOnlyRegistrationParamsWithDefaults instantiates a new EndpointOnlyRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *EndpointOnlyRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EndpointOnlyRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EndpointOnlyRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDescription

`func (o *EndpointOnlyRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EndpointOnlyRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EndpointOnlyRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EndpointOnlyRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EndpointOnlyRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EndpointOnlyRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


