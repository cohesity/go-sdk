# RecoverVolumeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationVolumeGuid** | **NullableString** | Specifies the guid of the destination volume. | 
**SourceVolumeGuid** | **NullableString** | Specifies the guid of the source volume. | 

## Methods

### NewRecoverVolumeMapping

`func NewRecoverVolumeMapping(destinationVolumeGuid NullableString, sourceVolumeGuid NullableString, ) *RecoverVolumeMapping`

NewRecoverVolumeMapping instantiates a new RecoverVolumeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverVolumeMappingWithDefaults

`func NewRecoverVolumeMappingWithDefaults() *RecoverVolumeMapping`

NewRecoverVolumeMappingWithDefaults instantiates a new RecoverVolumeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationVolumeGuid

`func (o *RecoverVolumeMapping) GetDestinationVolumeGuid() string`

GetDestinationVolumeGuid returns the DestinationVolumeGuid field if non-nil, zero value otherwise.

### GetDestinationVolumeGuidOk

`func (o *RecoverVolumeMapping) GetDestinationVolumeGuidOk() (*string, bool)`

GetDestinationVolumeGuidOk returns a tuple with the DestinationVolumeGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationVolumeGuid

`func (o *RecoverVolumeMapping) SetDestinationVolumeGuid(v string)`

SetDestinationVolumeGuid sets DestinationVolumeGuid field to given value.


### SetDestinationVolumeGuidNil

`func (o *RecoverVolumeMapping) SetDestinationVolumeGuidNil(b bool)`

 SetDestinationVolumeGuidNil sets the value for DestinationVolumeGuid to be an explicit nil

### UnsetDestinationVolumeGuid
`func (o *RecoverVolumeMapping) UnsetDestinationVolumeGuid()`

UnsetDestinationVolumeGuid ensures that no value is present for DestinationVolumeGuid, not even an explicit nil
### GetSourceVolumeGuid

`func (o *RecoverVolumeMapping) GetSourceVolumeGuid() string`

GetSourceVolumeGuid returns the SourceVolumeGuid field if non-nil, zero value otherwise.

### GetSourceVolumeGuidOk

`func (o *RecoverVolumeMapping) GetSourceVolumeGuidOk() (*string, bool)`

GetSourceVolumeGuidOk returns a tuple with the SourceVolumeGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeGuid

`func (o *RecoverVolumeMapping) SetSourceVolumeGuid(v string)`

SetSourceVolumeGuid sets SourceVolumeGuid field to given value.


### SetSourceVolumeGuidNil

`func (o *RecoverVolumeMapping) SetSourceVolumeGuidNil(b bool)`

 SetSourceVolumeGuidNil sets the value for SourceVolumeGuid to be an explicit nil

### UnsetSourceVolumeGuid
`func (o *RecoverVolumeMapping) UnsetSourceVolumeGuid()`

UnsetSourceVolumeGuid ensures that no value is present for SourceVolumeGuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


