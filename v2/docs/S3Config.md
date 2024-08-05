# S3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclConfig** | Pointer to [**S3ConfigAclConfig**](S3ConfigAclConfig.md) |  | [optional] 
**BucketPolicy** | Pointer to [**S3ConfigBucketPolicy**](S3ConfigBucketPolicy.md) |  | [optional] 
**EnableAbac** | Pointer to **NullableBool** | Specifies if this View has S3 ABAC enabled. This can only be set while creating a view. The ABAC server corresponding the tenant will be used for authentication and authorization checks.  | [optional] 
**LifecycleManagement** | Pointer to [**S3ConfigLifecycleManagement**](S3ConfigLifecycleManagement.md) |  | [optional] 
**OwnerInfo** | Pointer to [**S3ConfigOwnerInfo**](S3ConfigOwnerInfo.md) |  | [optional] 
**S3AccessPath** | Pointer to **NullableString** | Specifies the path to access this View as an S3 share. | [optional] [readonly] 
**S3EfficientMpuMaxSubfiles** | Pointer to **NullableInt32** | Specifies if this View has S3 MPU 2.0 enabled. This can set while editing a view.  | [optional] 
**S3EnableEfficientMpu** | Pointer to **NullableBool** | Specifies if this View has S3 MPU 2.0 enabled. This can set while editing a view.  | [optional] 
**S3MigrationAction** | Pointer to **NullableString** | Specifies the S3 migration action to be performed on this View. Supported migration actions are: [Enable, Cancel, Pause, Resume]. | [optional] 
**S3MigrationState** | Pointer to **NullableString** | Specifies the current S3 migration state for this View. A View can be under following migration states: [Eligible, Enable, Pause, Complete, UnderMigration]. | [optional] 
**Versioning** | Pointer to **NullableString** | Specifies the versioning state of S3 bucket. Buckets can be in one of three states: UnVersioned (default), VersioningEnabled, or VersioningSuspended. Once versioning is enabled for a bucket, it can never return to an UnVersioned state. However, versioning on the bucket can be suspended. | [optional] 

## Methods

### NewS3Config

`func NewS3Config() *S3Config`

NewS3Config instantiates a new S3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigWithDefaults

`func NewS3ConfigWithDefaults() *S3Config`

NewS3ConfigWithDefaults instantiates a new S3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclConfig

`func (o *S3Config) GetAclConfig() S3ConfigAclConfig`

GetAclConfig returns the AclConfig field if non-nil, zero value otherwise.

### GetAclConfigOk

`func (o *S3Config) GetAclConfigOk() (*S3ConfigAclConfig, bool)`

GetAclConfigOk returns a tuple with the AclConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclConfig

`func (o *S3Config) SetAclConfig(v S3ConfigAclConfig)`

SetAclConfig sets AclConfig field to given value.

### HasAclConfig

`func (o *S3Config) HasAclConfig() bool`

HasAclConfig returns a boolean if a field has been set.

### GetBucketPolicy

`func (o *S3Config) GetBucketPolicy() S3ConfigBucketPolicy`

GetBucketPolicy returns the BucketPolicy field if non-nil, zero value otherwise.

### GetBucketPolicyOk

`func (o *S3Config) GetBucketPolicyOk() (*S3ConfigBucketPolicy, bool)`

GetBucketPolicyOk returns a tuple with the BucketPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketPolicy

`func (o *S3Config) SetBucketPolicy(v S3ConfigBucketPolicy)`

SetBucketPolicy sets BucketPolicy field to given value.

### HasBucketPolicy

`func (o *S3Config) HasBucketPolicy() bool`

HasBucketPolicy returns a boolean if a field has been set.

### GetEnableAbac

`func (o *S3Config) GetEnableAbac() bool`

GetEnableAbac returns the EnableAbac field if non-nil, zero value otherwise.

### GetEnableAbacOk

`func (o *S3Config) GetEnableAbacOk() (*bool, bool)`

GetEnableAbacOk returns a tuple with the EnableAbac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAbac

`func (o *S3Config) SetEnableAbac(v bool)`

SetEnableAbac sets EnableAbac field to given value.

### HasEnableAbac

`func (o *S3Config) HasEnableAbac() bool`

HasEnableAbac returns a boolean if a field has been set.

### SetEnableAbacNil

`func (o *S3Config) SetEnableAbacNil(b bool)`

 SetEnableAbacNil sets the value for EnableAbac to be an explicit nil

### UnsetEnableAbac
`func (o *S3Config) UnsetEnableAbac()`

UnsetEnableAbac ensures that no value is present for EnableAbac, not even an explicit nil
### GetLifecycleManagement

`func (o *S3Config) GetLifecycleManagement() S3ConfigLifecycleManagement`

GetLifecycleManagement returns the LifecycleManagement field if non-nil, zero value otherwise.

### GetLifecycleManagementOk

`func (o *S3Config) GetLifecycleManagementOk() (*S3ConfigLifecycleManagement, bool)`

GetLifecycleManagementOk returns a tuple with the LifecycleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleManagement

`func (o *S3Config) SetLifecycleManagement(v S3ConfigLifecycleManagement)`

SetLifecycleManagement sets LifecycleManagement field to given value.

### HasLifecycleManagement

`func (o *S3Config) HasLifecycleManagement() bool`

HasLifecycleManagement returns a boolean if a field has been set.

### GetOwnerInfo

`func (o *S3Config) GetOwnerInfo() S3ConfigOwnerInfo`

GetOwnerInfo returns the OwnerInfo field if non-nil, zero value otherwise.

### GetOwnerInfoOk

`func (o *S3Config) GetOwnerInfoOk() (*S3ConfigOwnerInfo, bool)`

GetOwnerInfoOk returns a tuple with the OwnerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInfo

`func (o *S3Config) SetOwnerInfo(v S3ConfigOwnerInfo)`

SetOwnerInfo sets OwnerInfo field to given value.

### HasOwnerInfo

`func (o *S3Config) HasOwnerInfo() bool`

HasOwnerInfo returns a boolean if a field has been set.

### GetS3AccessPath

`func (o *S3Config) GetS3AccessPath() string`

GetS3AccessPath returns the S3AccessPath field if non-nil, zero value otherwise.

### GetS3AccessPathOk

`func (o *S3Config) GetS3AccessPathOk() (*string, bool)`

GetS3AccessPathOk returns a tuple with the S3AccessPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessPath

`func (o *S3Config) SetS3AccessPath(v string)`

SetS3AccessPath sets S3AccessPath field to given value.

### HasS3AccessPath

`func (o *S3Config) HasS3AccessPath() bool`

HasS3AccessPath returns a boolean if a field has been set.

### SetS3AccessPathNil

`func (o *S3Config) SetS3AccessPathNil(b bool)`

 SetS3AccessPathNil sets the value for S3AccessPath to be an explicit nil

### UnsetS3AccessPath
`func (o *S3Config) UnsetS3AccessPath()`

UnsetS3AccessPath ensures that no value is present for S3AccessPath, not even an explicit nil
### GetS3EfficientMpuMaxSubfiles

`func (o *S3Config) GetS3EfficientMpuMaxSubfiles() int32`

GetS3EfficientMpuMaxSubfiles returns the S3EfficientMpuMaxSubfiles field if non-nil, zero value otherwise.

### GetS3EfficientMpuMaxSubfilesOk

`func (o *S3Config) GetS3EfficientMpuMaxSubfilesOk() (*int32, bool)`

GetS3EfficientMpuMaxSubfilesOk returns a tuple with the S3EfficientMpuMaxSubfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3EfficientMpuMaxSubfiles

`func (o *S3Config) SetS3EfficientMpuMaxSubfiles(v int32)`

SetS3EfficientMpuMaxSubfiles sets S3EfficientMpuMaxSubfiles field to given value.

### HasS3EfficientMpuMaxSubfiles

`func (o *S3Config) HasS3EfficientMpuMaxSubfiles() bool`

HasS3EfficientMpuMaxSubfiles returns a boolean if a field has been set.

### SetS3EfficientMpuMaxSubfilesNil

`func (o *S3Config) SetS3EfficientMpuMaxSubfilesNil(b bool)`

 SetS3EfficientMpuMaxSubfilesNil sets the value for S3EfficientMpuMaxSubfiles to be an explicit nil

### UnsetS3EfficientMpuMaxSubfiles
`func (o *S3Config) UnsetS3EfficientMpuMaxSubfiles()`

UnsetS3EfficientMpuMaxSubfiles ensures that no value is present for S3EfficientMpuMaxSubfiles, not even an explicit nil
### GetS3EnableEfficientMpu

`func (o *S3Config) GetS3EnableEfficientMpu() bool`

GetS3EnableEfficientMpu returns the S3EnableEfficientMpu field if non-nil, zero value otherwise.

### GetS3EnableEfficientMpuOk

`func (o *S3Config) GetS3EnableEfficientMpuOk() (*bool, bool)`

GetS3EnableEfficientMpuOk returns a tuple with the S3EnableEfficientMpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3EnableEfficientMpu

`func (o *S3Config) SetS3EnableEfficientMpu(v bool)`

SetS3EnableEfficientMpu sets S3EnableEfficientMpu field to given value.

### HasS3EnableEfficientMpu

`func (o *S3Config) HasS3EnableEfficientMpu() bool`

HasS3EnableEfficientMpu returns a boolean if a field has been set.

### SetS3EnableEfficientMpuNil

`func (o *S3Config) SetS3EnableEfficientMpuNil(b bool)`

 SetS3EnableEfficientMpuNil sets the value for S3EnableEfficientMpu to be an explicit nil

### UnsetS3EnableEfficientMpu
`func (o *S3Config) UnsetS3EnableEfficientMpu()`

UnsetS3EnableEfficientMpu ensures that no value is present for S3EnableEfficientMpu, not even an explicit nil
### GetS3MigrationAction

`func (o *S3Config) GetS3MigrationAction() string`

GetS3MigrationAction returns the S3MigrationAction field if non-nil, zero value otherwise.

### GetS3MigrationActionOk

`func (o *S3Config) GetS3MigrationActionOk() (*string, bool)`

GetS3MigrationActionOk returns a tuple with the S3MigrationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationAction

`func (o *S3Config) SetS3MigrationAction(v string)`

SetS3MigrationAction sets S3MigrationAction field to given value.

### HasS3MigrationAction

`func (o *S3Config) HasS3MigrationAction() bool`

HasS3MigrationAction returns a boolean if a field has been set.

### SetS3MigrationActionNil

`func (o *S3Config) SetS3MigrationActionNil(b bool)`

 SetS3MigrationActionNil sets the value for S3MigrationAction to be an explicit nil

### UnsetS3MigrationAction
`func (o *S3Config) UnsetS3MigrationAction()`

UnsetS3MigrationAction ensures that no value is present for S3MigrationAction, not even an explicit nil
### GetS3MigrationState

`func (o *S3Config) GetS3MigrationState() string`

GetS3MigrationState returns the S3MigrationState field if non-nil, zero value otherwise.

### GetS3MigrationStateOk

`func (o *S3Config) GetS3MigrationStateOk() (*string, bool)`

GetS3MigrationStateOk returns a tuple with the S3MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3MigrationState

`func (o *S3Config) SetS3MigrationState(v string)`

SetS3MigrationState sets S3MigrationState field to given value.

### HasS3MigrationState

`func (o *S3Config) HasS3MigrationState() bool`

HasS3MigrationState returns a boolean if a field has been set.

### SetS3MigrationStateNil

`func (o *S3Config) SetS3MigrationStateNil(b bool)`

 SetS3MigrationStateNil sets the value for S3MigrationState to be an explicit nil

### UnsetS3MigrationState
`func (o *S3Config) UnsetS3MigrationState()`

UnsetS3MigrationState ensures that no value is present for S3MigrationState, not even an explicit nil
### GetVersioning

`func (o *S3Config) GetVersioning() string`

GetVersioning returns the Versioning field if non-nil, zero value otherwise.

### GetVersioningOk

`func (o *S3Config) GetVersioningOk() (*string, bool)`

GetVersioningOk returns a tuple with the Versioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersioning

`func (o *S3Config) SetVersioning(v string)`

SetVersioning sets Versioning field to given value.

### HasVersioning

`func (o *S3Config) HasVersioning() bool`

HasVersioning returns a boolean if a field has been set.

### SetVersioningNil

`func (o *S3Config) SetVersioningNil(b bool)`

 SetVersioningNil sets the value for Versioning to be an explicit nil

### UnsetVersioning
`func (o *S3Config) UnsetVersioning()`

UnsetVersioning ensures that no value is present for Versioning, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


