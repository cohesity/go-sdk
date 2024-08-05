# ObjectSnapshotNetappParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedNasMountProtocols** | Pointer to **[]string** | Specifies a list of NAS mount protocols supported by this object. | [optional] 
**VolumeExtendedStyle** | Pointer to **NullableString** | Specifies the extended style of a NetApp volume. | [optional] 
**VolumeType** | Pointer to **NullableString** | Specifies the Netapp volume type. | [optional] 

## Methods

### NewObjectSnapshotNetappParams

`func NewObjectSnapshotNetappParams() *ObjectSnapshotNetappParams`

NewObjectSnapshotNetappParams instantiates a new ObjectSnapshotNetappParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotNetappParamsWithDefaults

`func NewObjectSnapshotNetappParamsWithDefaults() *ObjectSnapshotNetappParams`

NewObjectSnapshotNetappParamsWithDefaults instantiates a new ObjectSnapshotNetappParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedNasMountProtocols

`func (o *ObjectSnapshotNetappParams) GetSupportedNasMountProtocols() []string`

GetSupportedNasMountProtocols returns the SupportedNasMountProtocols field if non-nil, zero value otherwise.

### GetSupportedNasMountProtocolsOk

`func (o *ObjectSnapshotNetappParams) GetSupportedNasMountProtocolsOk() (*[]string, bool)`

GetSupportedNasMountProtocolsOk returns a tuple with the SupportedNasMountProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedNasMountProtocols

`func (o *ObjectSnapshotNetappParams) SetSupportedNasMountProtocols(v []string)`

SetSupportedNasMountProtocols sets SupportedNasMountProtocols field to given value.

### HasSupportedNasMountProtocols

`func (o *ObjectSnapshotNetappParams) HasSupportedNasMountProtocols() bool`

HasSupportedNasMountProtocols returns a boolean if a field has been set.

### SetSupportedNasMountProtocolsNil

`func (o *ObjectSnapshotNetappParams) SetSupportedNasMountProtocolsNil(b bool)`

 SetSupportedNasMountProtocolsNil sets the value for SupportedNasMountProtocols to be an explicit nil

### UnsetSupportedNasMountProtocols
`func (o *ObjectSnapshotNetappParams) UnsetSupportedNasMountProtocols()`

UnsetSupportedNasMountProtocols ensures that no value is present for SupportedNasMountProtocols, not even an explicit nil
### GetVolumeExtendedStyle

`func (o *ObjectSnapshotNetappParams) GetVolumeExtendedStyle() string`

GetVolumeExtendedStyle returns the VolumeExtendedStyle field if non-nil, zero value otherwise.

### GetVolumeExtendedStyleOk

`func (o *ObjectSnapshotNetappParams) GetVolumeExtendedStyleOk() (*string, bool)`

GetVolumeExtendedStyleOk returns a tuple with the VolumeExtendedStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeExtendedStyle

`func (o *ObjectSnapshotNetappParams) SetVolumeExtendedStyle(v string)`

SetVolumeExtendedStyle sets VolumeExtendedStyle field to given value.

### HasVolumeExtendedStyle

`func (o *ObjectSnapshotNetappParams) HasVolumeExtendedStyle() bool`

HasVolumeExtendedStyle returns a boolean if a field has been set.

### SetVolumeExtendedStyleNil

`func (o *ObjectSnapshotNetappParams) SetVolumeExtendedStyleNil(b bool)`

 SetVolumeExtendedStyleNil sets the value for VolumeExtendedStyle to be an explicit nil

### UnsetVolumeExtendedStyle
`func (o *ObjectSnapshotNetappParams) UnsetVolumeExtendedStyle()`

UnsetVolumeExtendedStyle ensures that no value is present for VolumeExtendedStyle, not even an explicit nil
### GetVolumeType

`func (o *ObjectSnapshotNetappParams) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ObjectSnapshotNetappParams) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ObjectSnapshotNetappParams) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ObjectSnapshotNetappParams) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *ObjectSnapshotNetappParams) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *ObjectSnapshotNetappParams) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


