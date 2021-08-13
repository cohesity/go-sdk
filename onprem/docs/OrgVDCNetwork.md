# OrgVDCNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcdUuid** | **NullableString** | Specifies the UUID of network associated with the Virtual Cloud director. | 
**Name** | **NullableString** | Specifies the name of the catalog. | 
**VCenterUuid** | Pointer to **NullableString** | Specifies the UUID of network associated with the Vcenter. | [optional] [readonly] 

## Methods

### NewOrgVDCNetwork

`func NewOrgVDCNetwork(vcdUuid NullableString, name NullableString, ) *OrgVDCNetwork`

NewOrgVDCNetwork instantiates a new OrgVDCNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgVDCNetworkWithDefaults

`func NewOrgVDCNetworkWithDefaults() *OrgVDCNetwork`

NewOrgVDCNetworkWithDefaults instantiates a new OrgVDCNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcdUuid

`func (o *OrgVDCNetwork) GetVcdUuid() string`

GetVcdUuid returns the VcdUuid field if non-nil, zero value otherwise.

### GetVcdUuidOk

`func (o *OrgVDCNetwork) GetVcdUuidOk() (*string, bool)`

GetVcdUuidOk returns a tuple with the VcdUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcdUuid

`func (o *OrgVDCNetwork) SetVcdUuid(v string)`

SetVcdUuid sets VcdUuid field to given value.


### SetVcdUuidNil

`func (o *OrgVDCNetwork) SetVcdUuidNil(b bool)`

 SetVcdUuidNil sets the value for VcdUuid to be an explicit nil

### UnsetVcdUuid
`func (o *OrgVDCNetwork) UnsetVcdUuid()`

UnsetVcdUuid ensures that no value is present for VcdUuid, not even an explicit nil
### GetName

`func (o *OrgVDCNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgVDCNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgVDCNetwork) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *OrgVDCNetwork) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrgVDCNetwork) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVCenterUuid

`func (o *OrgVDCNetwork) GetVCenterUuid() string`

GetVCenterUuid returns the VCenterUuid field if non-nil, zero value otherwise.

### GetVCenterUuidOk

`func (o *OrgVDCNetwork) GetVCenterUuidOk() (*string, bool)`

GetVCenterUuidOk returns a tuple with the VCenterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCenterUuid

`func (o *OrgVDCNetwork) SetVCenterUuid(v string)`

SetVCenterUuid sets VCenterUuid field to given value.

### HasVCenterUuid

`func (o *OrgVDCNetwork) HasVCenterUuid() bool`

HasVCenterUuid returns a boolean if a field has been set.

### SetVCenterUuidNil

`func (o *OrgVDCNetwork) SetVCenterUuidNil(b bool)`

 SetVCenterUuidNil sets the value for VCenterUuid to be an explicit nil

### UnsetVCenterUuid
`func (o *OrgVDCNetwork) UnsetVCenterUuid()`

UnsetVCenterUuid ensures that no value is present for VCenterUuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


