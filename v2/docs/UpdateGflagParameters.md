# UpdateGflagParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveNow** | Pointer to **NullableBool** | Specifies whether to apply the change immediately. If set to true, the gflag change will work without restarting the service. | [optional] 
**Reason** | Pointer to **NullableString** | Specifies the reason for clearing gflags. | [optional] 
**ServiceFlags** | Pointer to [**ServiceGflags**](ServiceGflags.md) |  | [optional] 

## Methods

### NewUpdateGflagParameters

`func NewUpdateGflagParameters() *UpdateGflagParameters`

NewUpdateGflagParameters instantiates a new UpdateGflagParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGflagParametersWithDefaults

`func NewUpdateGflagParametersWithDefaults() *UpdateGflagParameters`

NewUpdateGflagParametersWithDefaults instantiates a new UpdateGflagParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveNow

`func (o *UpdateGflagParameters) GetEffectiveNow() bool`

GetEffectiveNow returns the EffectiveNow field if non-nil, zero value otherwise.

### GetEffectiveNowOk

`func (o *UpdateGflagParameters) GetEffectiveNowOk() (*bool, bool)`

GetEffectiveNowOk returns a tuple with the EffectiveNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveNow

`func (o *UpdateGflagParameters) SetEffectiveNow(v bool)`

SetEffectiveNow sets EffectiveNow field to given value.

### HasEffectiveNow

`func (o *UpdateGflagParameters) HasEffectiveNow() bool`

HasEffectiveNow returns a boolean if a field has been set.

### SetEffectiveNowNil

`func (o *UpdateGflagParameters) SetEffectiveNowNil(b bool)`

 SetEffectiveNowNil sets the value for EffectiveNow to be an explicit nil

### UnsetEffectiveNow
`func (o *UpdateGflagParameters) UnsetEffectiveNow()`

UnsetEffectiveNow ensures that no value is present for EffectiveNow, not even an explicit nil
### GetReason

`func (o *UpdateGflagParameters) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateGflagParameters) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateGflagParameters) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateGflagParameters) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *UpdateGflagParameters) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UpdateGflagParameters) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetServiceFlags

`func (o *UpdateGflagParameters) GetServiceFlags() ServiceGflags`

GetServiceFlags returns the ServiceFlags field if non-nil, zero value otherwise.

### GetServiceFlagsOk

`func (o *UpdateGflagParameters) GetServiceFlagsOk() (*ServiceGflags, bool)`

GetServiceFlagsOk returns a tuple with the ServiceFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFlags

`func (o *UpdateGflagParameters) SetServiceFlags(v ServiceGflags)`

SetServiceFlags sets ServiceFlags field to given value.

### HasServiceFlags

`func (o *UpdateGflagParameters) HasServiceFlags() bool`

HasServiceFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


