# GpfsObjectProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **NullableString** | Specifies the protocol of the NAS device being backed up. | [optional] 

## Methods

### NewGpfsObjectProtectionParams

`func NewGpfsObjectProtectionParams() *GpfsObjectProtectionParams`

NewGpfsObjectProtectionParams instantiates a new GpfsObjectProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpfsObjectProtectionParamsWithDefaults

`func NewGpfsObjectProtectionParamsWithDefaults() *GpfsObjectProtectionParams`

NewGpfsObjectProtectionParamsWithDefaults instantiates a new GpfsObjectProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *GpfsObjectProtectionParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GpfsObjectProtectionParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GpfsObjectProtectionParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GpfsObjectProtectionParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GpfsObjectProtectionParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GpfsObjectProtectionParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


