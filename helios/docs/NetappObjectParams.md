# NetappObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedNasMountProtocols** | Pointer to **[]string** | Specifies a list of NAS mount protocols supported by this object. | [optional] 
**VolumeType** | Pointer to **NullableString** | Specifies the Netapp volume type. | [optional] 

## Methods

### NewNetappObjectParams

`func NewNetappObjectParams() *NetappObjectParams`

NewNetappObjectParams instantiates a new NetappObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappObjectParamsWithDefaults

`func NewNetappObjectParamsWithDefaults() *NetappObjectParams`

NewNetappObjectParamsWithDefaults instantiates a new NetappObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedNasMountProtocols

`func (o *NetappObjectParams) GetSupportedNasMountProtocols() []string`

GetSupportedNasMountProtocols returns the SupportedNasMountProtocols field if non-nil, zero value otherwise.

### GetSupportedNasMountProtocolsOk

`func (o *NetappObjectParams) GetSupportedNasMountProtocolsOk() (*[]string, bool)`

GetSupportedNasMountProtocolsOk returns a tuple with the SupportedNasMountProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedNasMountProtocols

`func (o *NetappObjectParams) SetSupportedNasMountProtocols(v []string)`

SetSupportedNasMountProtocols sets SupportedNasMountProtocols field to given value.

### HasSupportedNasMountProtocols

`func (o *NetappObjectParams) HasSupportedNasMountProtocols() bool`

HasSupportedNasMountProtocols returns a boolean if a field has been set.

### SetSupportedNasMountProtocolsNil

`func (o *NetappObjectParams) SetSupportedNasMountProtocolsNil(b bool)`

 SetSupportedNasMountProtocolsNil sets the value for SupportedNasMountProtocols to be an explicit nil

### UnsetSupportedNasMountProtocols
`func (o *NetappObjectParams) UnsetSupportedNasMountProtocols()`

UnsetSupportedNasMountProtocols ensures that no value is present for SupportedNasMountProtocols, not even an explicit nil
### GetVolumeType

`func (o *NetappObjectParams) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *NetappObjectParams) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *NetappObjectParams) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *NetappObjectParams) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *NetappObjectParams) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *NetappObjectParams) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


