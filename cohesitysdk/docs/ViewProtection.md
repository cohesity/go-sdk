# ViewProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inactive** | Pointer to **NullableBool** | Specifies if this View is an inactive View that was created on this Remote Cluster to store the Snapshots created by replication. This inactive View cannot be NFS or SMB mounted. | [optional] 
**MagnetoEntityId** | Pointer to **NullableInt64** | Specifies the id of the Protection Source that is using this View. | [optional] 
**ProtectionJobs** | Pointer to [**[]ProtectionJobInfo**](ProtectionJobInfo.md) | Array of Protection Jobs.  Specifies the Protection Jobs that are protecting the View. | [optional] 

## Methods

### NewViewProtection

`func NewViewProtection() *ViewProtection`

NewViewProtection instantiates a new ViewProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewProtectionWithDefaults

`func NewViewProtectionWithDefaults() *ViewProtection`

NewViewProtectionWithDefaults instantiates a new ViewProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInactive

`func (o *ViewProtection) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *ViewProtection) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *ViewProtection) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *ViewProtection) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### SetInactiveNil

`func (o *ViewProtection) SetInactiveNil(b bool)`

 SetInactiveNil sets the value for Inactive to be an explicit nil

### UnsetInactive
`func (o *ViewProtection) UnsetInactive()`

UnsetInactive ensures that no value is present for Inactive, not even an explicit nil
### GetMagnetoEntityId

`func (o *ViewProtection) GetMagnetoEntityId() int64`

GetMagnetoEntityId returns the MagnetoEntityId field if non-nil, zero value otherwise.

### GetMagnetoEntityIdOk

`func (o *ViewProtection) GetMagnetoEntityIdOk() (*int64, bool)`

GetMagnetoEntityIdOk returns a tuple with the MagnetoEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetoEntityId

`func (o *ViewProtection) SetMagnetoEntityId(v int64)`

SetMagnetoEntityId sets MagnetoEntityId field to given value.

### HasMagnetoEntityId

`func (o *ViewProtection) HasMagnetoEntityId() bool`

HasMagnetoEntityId returns a boolean if a field has been set.

### SetMagnetoEntityIdNil

`func (o *ViewProtection) SetMagnetoEntityIdNil(b bool)`

 SetMagnetoEntityIdNil sets the value for MagnetoEntityId to be an explicit nil

### UnsetMagnetoEntityId
`func (o *ViewProtection) UnsetMagnetoEntityId()`

UnsetMagnetoEntityId ensures that no value is present for MagnetoEntityId, not even an explicit nil
### GetProtectionJobs

`func (o *ViewProtection) GetProtectionJobs() []ProtectionJobInfo`

GetProtectionJobs returns the ProtectionJobs field if non-nil, zero value otherwise.

### GetProtectionJobsOk

`func (o *ViewProtection) GetProtectionJobsOk() (*[]ProtectionJobInfo, bool)`

GetProtectionJobsOk returns a tuple with the ProtectionJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobs

`func (o *ViewProtection) SetProtectionJobs(v []ProtectionJobInfo)`

SetProtectionJobs sets ProtectionJobs field to given value.

### HasProtectionJobs

`func (o *ViewProtection) HasProtectionJobs() bool`

HasProtectionJobs returns a boolean if a field has been set.

### SetProtectionJobsNil

`func (o *ViewProtection) SetProtectionJobsNil(b bool)`

 SetProtectionJobsNil sets the value for ProtectionJobs to be an explicit nil

### UnsetProtectionJobs
`func (o *ViewProtection) UnsetProtectionJobs()`

UnsetProtectionJobs ensures that no value is present for ProtectionJobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


