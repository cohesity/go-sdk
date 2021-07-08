# MountVolumesTaskStateProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullNasPath** | Pointer to **NullableString** | Contains the SMB/NFS path of the share we expose to clients. The share contains the files pertinent to the original backup run type. | [optional] 
**HostEntity** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**MountInfo** | Pointer to [**MountVolumesInfoProto**](MountVolumesInfoProto.md) |  | [optional] 
**MountParams** | Pointer to [**MountVolumesParams**](MountVolumesParams.md) |  | [optional] 

## Methods

### NewMountVolumesTaskStateProto

`func NewMountVolumesTaskStateProto() *MountVolumesTaskStateProto`

NewMountVolumesTaskStateProto instantiates a new MountVolumesTaskStateProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountVolumesTaskStateProtoWithDefaults

`func NewMountVolumesTaskStateProtoWithDefaults() *MountVolumesTaskStateProto`

NewMountVolumesTaskStateProtoWithDefaults instantiates a new MountVolumesTaskStateProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullNasPath

`func (o *MountVolumesTaskStateProto) GetFullNasPath() string`

GetFullNasPath returns the FullNasPath field if non-nil, zero value otherwise.

### GetFullNasPathOk

`func (o *MountVolumesTaskStateProto) GetFullNasPathOk() (*string, bool)`

GetFullNasPathOk returns a tuple with the FullNasPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullNasPath

`func (o *MountVolumesTaskStateProto) SetFullNasPath(v string)`

SetFullNasPath sets FullNasPath field to given value.

### HasFullNasPath

`func (o *MountVolumesTaskStateProto) HasFullNasPath() bool`

HasFullNasPath returns a boolean if a field has been set.

### SetFullNasPathNil

`func (o *MountVolumesTaskStateProto) SetFullNasPathNil(b bool)`

 SetFullNasPathNil sets the value for FullNasPath to be an explicit nil

### UnsetFullNasPath
`func (o *MountVolumesTaskStateProto) UnsetFullNasPath()`

UnsetFullNasPath ensures that no value is present for FullNasPath, not even an explicit nil
### GetHostEntity

`func (o *MountVolumesTaskStateProto) GetHostEntity() EntityProto`

GetHostEntity returns the HostEntity field if non-nil, zero value otherwise.

### GetHostEntityOk

`func (o *MountVolumesTaskStateProto) GetHostEntityOk() (*EntityProto, bool)`

GetHostEntityOk returns a tuple with the HostEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEntity

`func (o *MountVolumesTaskStateProto) SetHostEntity(v EntityProto)`

SetHostEntity sets HostEntity field to given value.

### HasHostEntity

`func (o *MountVolumesTaskStateProto) HasHostEntity() bool`

HasHostEntity returns a boolean if a field has been set.

### GetMountInfo

`func (o *MountVolumesTaskStateProto) GetMountInfo() MountVolumesInfoProto`

GetMountInfo returns the MountInfo field if non-nil, zero value otherwise.

### GetMountInfoOk

`func (o *MountVolumesTaskStateProto) GetMountInfoOk() (*MountVolumesInfoProto, bool)`

GetMountInfoOk returns a tuple with the MountInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountInfo

`func (o *MountVolumesTaskStateProto) SetMountInfo(v MountVolumesInfoProto)`

SetMountInfo sets MountInfo field to given value.

### HasMountInfo

`func (o *MountVolumesTaskStateProto) HasMountInfo() bool`

HasMountInfo returns a boolean if a field has been set.

### GetMountParams

`func (o *MountVolumesTaskStateProto) GetMountParams() MountVolumesParams`

GetMountParams returns the MountParams field if non-nil, zero value otherwise.

### GetMountParamsOk

`func (o *MountVolumesTaskStateProto) GetMountParamsOk() (*MountVolumesParams, bool)`

GetMountParamsOk returns a tuple with the MountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountParams

`func (o *MountVolumesTaskStateProto) SetMountParams(v MountVolumesParams)`

SetMountParams sets MountParams field to given value.

### HasMountParams

`func (o *MountVolumesTaskStateProto) HasMountParams() bool`

HasMountParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


