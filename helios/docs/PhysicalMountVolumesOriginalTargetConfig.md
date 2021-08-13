# PhysicalMountVolumesOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerCredentials** | Pointer to [**NullableCredentials**](Credentials.md) | Specifies credentials to access the target server. This is required if the server is of Linux OS. | [optional] 

## Methods

### NewPhysicalMountVolumesOriginalTargetConfig

`func NewPhysicalMountVolumesOriginalTargetConfig() *PhysicalMountVolumesOriginalTargetConfig`

NewPhysicalMountVolumesOriginalTargetConfig instantiates a new PhysicalMountVolumesOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalMountVolumesOriginalTargetConfigWithDefaults

`func NewPhysicalMountVolumesOriginalTargetConfigWithDefaults() *PhysicalMountVolumesOriginalTargetConfig`

NewPhysicalMountVolumesOriginalTargetConfigWithDefaults instantiates a new PhysicalMountVolumesOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerCredentials

`func (o *PhysicalMountVolumesOriginalTargetConfig) GetServerCredentials() Credentials`

GetServerCredentials returns the ServerCredentials field if non-nil, zero value otherwise.

### GetServerCredentialsOk

`func (o *PhysicalMountVolumesOriginalTargetConfig) GetServerCredentialsOk() (*Credentials, bool)`

GetServerCredentialsOk returns a tuple with the ServerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCredentials

`func (o *PhysicalMountVolumesOriginalTargetConfig) SetServerCredentials(v Credentials)`

SetServerCredentials sets ServerCredentials field to given value.

### HasServerCredentials

`func (o *PhysicalMountVolumesOriginalTargetConfig) HasServerCredentials() bool`

HasServerCredentials returns a boolean if a field has been set.

### SetServerCredentialsNil

`func (o *PhysicalMountVolumesOriginalTargetConfig) SetServerCredentialsNil(b bool)`

 SetServerCredentialsNil sets the value for ServerCredentials to be an explicit nil

### UnsetServerCredentials
`func (o *PhysicalMountVolumesOriginalTargetConfig) UnsetServerCredentials()`

UnsetServerCredentials ensures that no value is present for ServerCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


