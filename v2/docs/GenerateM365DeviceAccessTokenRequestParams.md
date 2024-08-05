# GenerateM365DeviceAccessTokenRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceCode** | Pointer to **NullableString** | Specifies the string used to verify the session between the client and the authorization server. The client uses this parameter to request the access token from the authorization server. | [optional] 
**Domain** | Pointer to **NullableString** | Specifies the Microsoft365 domain. | [optional] 

## Methods

### NewGenerateM365DeviceAccessTokenRequestParams

`func NewGenerateM365DeviceAccessTokenRequestParams() *GenerateM365DeviceAccessTokenRequestParams`

NewGenerateM365DeviceAccessTokenRequestParams instantiates a new GenerateM365DeviceAccessTokenRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateM365DeviceAccessTokenRequestParamsWithDefaults

`func NewGenerateM365DeviceAccessTokenRequestParamsWithDefaults() *GenerateM365DeviceAccessTokenRequestParams`

NewGenerateM365DeviceAccessTokenRequestParamsWithDefaults instantiates a new GenerateM365DeviceAccessTokenRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceCode

`func (o *GenerateM365DeviceAccessTokenRequestParams) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *GenerateM365DeviceAccessTokenRequestParams) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *GenerateM365DeviceAccessTokenRequestParams) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.

### HasDeviceCode

`func (o *GenerateM365DeviceAccessTokenRequestParams) HasDeviceCode() bool`

HasDeviceCode returns a boolean if a field has been set.

### SetDeviceCodeNil

`func (o *GenerateM365DeviceAccessTokenRequestParams) SetDeviceCodeNil(b bool)`

 SetDeviceCodeNil sets the value for DeviceCode to be an explicit nil

### UnsetDeviceCode
`func (o *GenerateM365DeviceAccessTokenRequestParams) UnsetDeviceCode()`

UnsetDeviceCode ensures that no value is present for DeviceCode, not even an explicit nil
### GetDomain

`func (o *GenerateM365DeviceAccessTokenRequestParams) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GenerateM365DeviceAccessTokenRequestParams) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GenerateM365DeviceAccessTokenRequestParams) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GenerateM365DeviceAccessTokenRequestParams) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GenerateM365DeviceAccessTokenRequestParams) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GenerateM365DeviceAccessTokenRequestParams) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


