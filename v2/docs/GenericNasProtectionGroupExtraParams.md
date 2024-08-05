# GenericNasProtectionGroupExtraParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectCloudArchive** | Pointer to **NullableBool** | Specifies whether or not to store the snapshots in this run directly in an Archive Target instead of on the Cluster. If this is set to true, the associated policy must have exactly one Archive Target associated with it and the policy must be set up to archive after every run. Also, a Storage Domain cannot be specified. Default behavior is &#39;false&#39;. | [optional] 
**NativeFormat** | Pointer to **NullableBool** | Specifies whether or not to enable native format for direct archive job. This field is set to true if native format should be used for archiving. | [optional] 
**Protocol** | Pointer to **NullableString** | Specifies the preferred protocol to use if this device supports multiple protocols. | [optional] 

## Methods

### NewGenericNasProtectionGroupExtraParams

`func NewGenericNasProtectionGroupExtraParams() *GenericNasProtectionGroupExtraParams`

NewGenericNasProtectionGroupExtraParams instantiates a new GenericNasProtectionGroupExtraParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericNasProtectionGroupExtraParamsWithDefaults

`func NewGenericNasProtectionGroupExtraParamsWithDefaults() *GenericNasProtectionGroupExtraParams`

NewGenericNasProtectionGroupExtraParamsWithDefaults instantiates a new GenericNasProtectionGroupExtraParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectCloudArchive

`func (o *GenericNasProtectionGroupExtraParams) GetDirectCloudArchive() bool`

GetDirectCloudArchive returns the DirectCloudArchive field if non-nil, zero value otherwise.

### GetDirectCloudArchiveOk

`func (o *GenericNasProtectionGroupExtraParams) GetDirectCloudArchiveOk() (*bool, bool)`

GetDirectCloudArchiveOk returns a tuple with the DirectCloudArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectCloudArchive

`func (o *GenericNasProtectionGroupExtraParams) SetDirectCloudArchive(v bool)`

SetDirectCloudArchive sets DirectCloudArchive field to given value.

### HasDirectCloudArchive

`func (o *GenericNasProtectionGroupExtraParams) HasDirectCloudArchive() bool`

HasDirectCloudArchive returns a boolean if a field has been set.

### SetDirectCloudArchiveNil

`func (o *GenericNasProtectionGroupExtraParams) SetDirectCloudArchiveNil(b bool)`

 SetDirectCloudArchiveNil sets the value for DirectCloudArchive to be an explicit nil

### UnsetDirectCloudArchive
`func (o *GenericNasProtectionGroupExtraParams) UnsetDirectCloudArchive()`

UnsetDirectCloudArchive ensures that no value is present for DirectCloudArchive, not even an explicit nil
### GetNativeFormat

`func (o *GenericNasProtectionGroupExtraParams) GetNativeFormat() bool`

GetNativeFormat returns the NativeFormat field if non-nil, zero value otherwise.

### GetNativeFormatOk

`func (o *GenericNasProtectionGroupExtraParams) GetNativeFormatOk() (*bool, bool)`

GetNativeFormatOk returns a tuple with the NativeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeFormat

`func (o *GenericNasProtectionGroupExtraParams) SetNativeFormat(v bool)`

SetNativeFormat sets NativeFormat field to given value.

### HasNativeFormat

`func (o *GenericNasProtectionGroupExtraParams) HasNativeFormat() bool`

HasNativeFormat returns a boolean if a field has been set.

### SetNativeFormatNil

`func (o *GenericNasProtectionGroupExtraParams) SetNativeFormatNil(b bool)`

 SetNativeFormatNil sets the value for NativeFormat to be an explicit nil

### UnsetNativeFormat
`func (o *GenericNasProtectionGroupExtraParams) UnsetNativeFormat()`

UnsetNativeFormat ensures that no value is present for NativeFormat, not even an explicit nil
### GetProtocol

`func (o *GenericNasProtectionGroupExtraParams) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GenericNasProtectionGroupExtraParams) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GenericNasProtectionGroupExtraParams) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GenericNasProtectionGroupExtraParams) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GenericNasProtectionGroupExtraParams) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GenericNasProtectionGroupExtraParams) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


