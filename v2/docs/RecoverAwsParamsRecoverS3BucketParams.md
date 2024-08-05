# RecoverAwsParamsRecoverS3BucketParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsS3BucketRestoreFilterPolicy** | Pointer to [**NullableRecoverAwsS3BucketParamsAwsS3BucketRestoreFilterPolicy**](RecoverAwsS3BucketParamsAwsS3BucketRestoreFilterPolicy.md) |  | [optional] 
**AwsTargetParams** | Pointer to [**NullableRecoverAwsS3BucketParamsAwsTargetParams**](RecoverAwsS3BucketParamsAwsTargetParams.md) |  | [optional] 
**RecoverProtectionGroupRunsParams** | Pointer to [**[]RecoverProtectionGroupRunParams**](RecoverProtectionGroupRunParams.md) | Specifies the Protection Group Runs params to recover. All the VM&#39;s that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object&#39;s snapshot. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 

## Methods

### NewRecoverAwsParamsRecoverS3BucketParams

`func NewRecoverAwsParamsRecoverS3BucketParams(targetEnvironment string, ) *RecoverAwsParamsRecoverS3BucketParams`

NewRecoverAwsParamsRecoverS3BucketParams instantiates a new RecoverAwsParamsRecoverS3BucketParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsParamsRecoverS3BucketParamsWithDefaults

`func NewRecoverAwsParamsRecoverS3BucketParamsWithDefaults() *RecoverAwsParamsRecoverS3BucketParams`

NewRecoverAwsParamsRecoverS3BucketParamsWithDefaults instantiates a new RecoverAwsParamsRecoverS3BucketParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsS3BucketRestoreFilterPolicy

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetAwsS3BucketRestoreFilterPolicy() RecoverAwsS3BucketParamsAwsS3BucketRestoreFilterPolicy`

GetAwsS3BucketRestoreFilterPolicy returns the AwsS3BucketRestoreFilterPolicy field if non-nil, zero value otherwise.

### GetAwsS3BucketRestoreFilterPolicyOk

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetAwsS3BucketRestoreFilterPolicyOk() (*RecoverAwsS3BucketParamsAwsS3BucketRestoreFilterPolicy, bool)`

GetAwsS3BucketRestoreFilterPolicyOk returns a tuple with the AwsS3BucketRestoreFilterPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsS3BucketRestoreFilterPolicy

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetAwsS3BucketRestoreFilterPolicy(v RecoverAwsS3BucketParamsAwsS3BucketRestoreFilterPolicy)`

SetAwsS3BucketRestoreFilterPolicy sets AwsS3BucketRestoreFilterPolicy field to given value.

### HasAwsS3BucketRestoreFilterPolicy

`func (o *RecoverAwsParamsRecoverS3BucketParams) HasAwsS3BucketRestoreFilterPolicy() bool`

HasAwsS3BucketRestoreFilterPolicy returns a boolean if a field has been set.

### SetAwsS3BucketRestoreFilterPolicyNil

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetAwsS3BucketRestoreFilterPolicyNil(b bool)`

 SetAwsS3BucketRestoreFilterPolicyNil sets the value for AwsS3BucketRestoreFilterPolicy to be an explicit nil

### UnsetAwsS3BucketRestoreFilterPolicy
`func (o *RecoverAwsParamsRecoverS3BucketParams) UnsetAwsS3BucketRestoreFilterPolicy()`

UnsetAwsS3BucketRestoreFilterPolicy ensures that no value is present for AwsS3BucketRestoreFilterPolicy, not even an explicit nil
### GetAwsTargetParams

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetAwsTargetParams() RecoverAwsS3BucketParamsAwsTargetParams`

GetAwsTargetParams returns the AwsTargetParams field if non-nil, zero value otherwise.

### GetAwsTargetParamsOk

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetAwsTargetParamsOk() (*RecoverAwsS3BucketParamsAwsTargetParams, bool)`

GetAwsTargetParamsOk returns a tuple with the AwsTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetParams

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetAwsTargetParams(v RecoverAwsS3BucketParamsAwsTargetParams)`

SetAwsTargetParams sets AwsTargetParams field to given value.

### HasAwsTargetParams

`func (o *RecoverAwsParamsRecoverS3BucketParams) HasAwsTargetParams() bool`

HasAwsTargetParams returns a boolean if a field has been set.

### SetAwsTargetParamsNil

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetAwsTargetParamsNil(b bool)`

 SetAwsTargetParamsNil sets the value for AwsTargetParams to be an explicit nil

### UnsetAwsTargetParams
`func (o *RecoverAwsParamsRecoverS3BucketParams) UnsetAwsTargetParams()`

UnsetAwsTargetParams ensures that no value is present for AwsTargetParams, not even an explicit nil
### GetRecoverProtectionGroupRunsParams

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetRecoverProtectionGroupRunsParams() []RecoverProtectionGroupRunParams`

GetRecoverProtectionGroupRunsParams returns the RecoverProtectionGroupRunsParams field if non-nil, zero value otherwise.

### GetRecoverProtectionGroupRunsParamsOk

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetRecoverProtectionGroupRunsParamsOk() (*[]RecoverProtectionGroupRunParams, bool)`

GetRecoverProtectionGroupRunsParamsOk returns a tuple with the RecoverProtectionGroupRunsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverProtectionGroupRunsParams

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetRecoverProtectionGroupRunsParams(v []RecoverProtectionGroupRunParams)`

SetRecoverProtectionGroupRunsParams sets RecoverProtectionGroupRunsParams field to given value.

### HasRecoverProtectionGroupRunsParams

`func (o *RecoverAwsParamsRecoverS3BucketParams) HasRecoverProtectionGroupRunsParams() bool`

HasRecoverProtectionGroupRunsParams returns a boolean if a field has been set.

### SetRecoverProtectionGroupRunsParamsNil

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetRecoverProtectionGroupRunsParamsNil(b bool)`

 SetRecoverProtectionGroupRunsParamsNil sets the value for RecoverProtectionGroupRunsParams to be an explicit nil

### UnsetRecoverProtectionGroupRunsParams
`func (o *RecoverAwsParamsRecoverS3BucketParams) UnsetRecoverProtectionGroupRunsParams()`

UnsetRecoverProtectionGroupRunsParams ensures that no value is present for RecoverProtectionGroupRunsParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverAwsParamsRecoverS3BucketParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverAwsParamsRecoverS3BucketParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


