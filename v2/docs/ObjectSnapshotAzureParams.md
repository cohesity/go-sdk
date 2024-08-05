# ObjectSnapshotAzureParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectionType** | Pointer to **NullableString** | Specifies the protection type of Azure snapshots. | [optional] 

## Methods

### NewObjectSnapshotAzureParams

`func NewObjectSnapshotAzureParams() *ObjectSnapshotAzureParams`

NewObjectSnapshotAzureParams instantiates a new ObjectSnapshotAzureParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotAzureParamsWithDefaults

`func NewObjectSnapshotAzureParamsWithDefaults() *ObjectSnapshotAzureParams`

NewObjectSnapshotAzureParamsWithDefaults instantiates a new ObjectSnapshotAzureParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectionType

`func (o *ObjectSnapshotAzureParams) GetProtectionType() string`

GetProtectionType returns the ProtectionType field if non-nil, zero value otherwise.

### GetProtectionTypeOk

`func (o *ObjectSnapshotAzureParams) GetProtectionTypeOk() (*string, bool)`

GetProtectionTypeOk returns a tuple with the ProtectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionType

`func (o *ObjectSnapshotAzureParams) SetProtectionType(v string)`

SetProtectionType sets ProtectionType field to given value.

### HasProtectionType

`func (o *ObjectSnapshotAzureParams) HasProtectionType() bool`

HasProtectionType returns a boolean if a field has been set.

### SetProtectionTypeNil

`func (o *ObjectSnapshotAzureParams) SetProtectionTypeNil(b bool)`

 SetProtectionTypeNil sets the value for ProtectionType to be an explicit nil

### UnsetProtectionType
`func (o *ObjectSnapshotAzureParams) UnsetProtectionType()`

UnsetProtectionType ensures that no value is present for ProtectionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


