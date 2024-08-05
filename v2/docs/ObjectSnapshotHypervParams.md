# ObjectSnapshotHypervParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **NullableString** | Specifies the protection type of HyperV snapshots. | [optional] 

## Methods

### NewObjectSnapshotHypervParams

`func NewObjectSnapshotHypervParams() *ObjectSnapshotHypervParams`

NewObjectSnapshotHypervParams instantiates a new ObjectSnapshotHypervParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotHypervParamsWithDefaults

`func NewObjectSnapshotHypervParamsWithDefaults() *ObjectSnapshotHypervParams`

NewObjectSnapshotHypervParamsWithDefaults instantiates a new ObjectSnapshotHypervParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionType

`func (o *ObjectSnapshotHypervParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ObjectSnapshotHypervParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ObjectSnapshotHypervParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *ObjectSnapshotHypervParams) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *ObjectSnapshotHypervParams) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ObjectSnapshotHypervParams) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


