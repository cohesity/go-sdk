# EmblemServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmblemHostAddress** | Pointer to **NullableString** | Address of the service to connect to. | [optional] 
**EmblemPort** | Pointer to **NullableInt32** | Port on which service is listening. | [optional] 

## Methods

### NewEmblemServiceInfo

`func NewEmblemServiceInfo() *EmblemServiceInfo`

NewEmblemServiceInfo instantiates a new EmblemServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmblemServiceInfoWithDefaults

`func NewEmblemServiceInfoWithDefaults() *EmblemServiceInfo`

NewEmblemServiceInfoWithDefaults instantiates a new EmblemServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmblemHostAddress

`func (o *EmblemServiceInfo) GetEmblemHostAddress() string`

GetEmblemHostAddress returns the EmblemHostAddress field if non-nil, zero value otherwise.

### GetEmblemHostAddressOk

`func (o *EmblemServiceInfo) GetEmblemHostAddressOk() (*string, bool)`

GetEmblemHostAddressOk returns a tuple with the EmblemHostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmblemHostAddress

`func (o *EmblemServiceInfo) SetEmblemHostAddress(v string)`

SetEmblemHostAddress sets EmblemHostAddress field to given value.

### HasEmblemHostAddress

`func (o *EmblemServiceInfo) HasEmblemHostAddress() bool`

HasEmblemHostAddress returns a boolean if a field has been set.

### SetEmblemHostAddressNil

`func (o *EmblemServiceInfo) SetEmblemHostAddressNil(b bool)`

 SetEmblemHostAddressNil sets the value for EmblemHostAddress to be an explicit nil

### UnsetEmblemHostAddress
`func (o *EmblemServiceInfo) UnsetEmblemHostAddress()`

UnsetEmblemHostAddress ensures that no value is present for EmblemHostAddress, not even an explicit nil
### GetEmblemPort

`func (o *EmblemServiceInfo) GetEmblemPort() int32`

GetEmblemPort returns the EmblemPort field if non-nil, zero value otherwise.

### GetEmblemPortOk

`func (o *EmblemServiceInfo) GetEmblemPortOk() (*int32, bool)`

GetEmblemPortOk returns a tuple with the EmblemPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmblemPort

`func (o *EmblemServiceInfo) SetEmblemPort(v int32)`

SetEmblemPort sets EmblemPort field to given value.

### HasEmblemPort

`func (o *EmblemServiceInfo) HasEmblemPort() bool`

HasEmblemPort returns a boolean if a field has been set.

### SetEmblemPortNil

`func (o *EmblemServiceInfo) SetEmblemPortNil(b bool)`

 SetEmblemPortNil sets the value for EmblemPort to be an explicit nil

### UnsetEmblemPort
`func (o *EmblemServiceInfo) UnsetEmblemPort()`

UnsetEmblemPort ensures that no value is present for EmblemPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


