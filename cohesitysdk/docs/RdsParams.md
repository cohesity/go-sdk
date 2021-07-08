# RdsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneId** | Pointer to **NullableInt64** | Entity representing the availability zone to use while restoring the DB. | [optional] 
**DbInstanceId** | **NullableString** | The DB instance identifier to use for the restored DB. This field is required. | 
**DbOptionGroupId** | Pointer to **NullableInt64** | Entity representing the RDS option group to use while restoring the DB. | [optional] 
**DbParameterGroupId** | Pointer to **NullableInt64** | Entity representing the RDS parameter group to use while restoring the DB. | [optional] 
**DbPort** | Pointer to **NullableInt32** | Port to use for the DB in the restored RDS instance. | [optional] 
**EnableAutoMinorVersionUpgrade** | Pointer to **NullableBool** | Whether to enable auto minor version upgrade in the restored DB. | [optional] 
**EnableCopyTagsToSnapshots** | Pointer to **NullableBool** | Whether to enable copying of tags to snapshots of the DB. | [optional] 
**EnableDbAuthentication** | Pointer to **NullableBool** | Whether to enable IAM authentication for the DB. | [optional] 
**EnablePublicAccessibility** | Pointer to **NullableBool** | Whether this DB will be publicly accessible or not. | [optional] 
**IsMultiAzDeployment** | Pointer to **NullableBool** | Whether this is a multi-az deployment or not. | [optional] 

## Methods

### NewRdsParams

`func NewRdsParams(dbInstanceId NullableString, ) *RdsParams`

NewRdsParams instantiates a new RdsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsParamsWithDefaults

`func NewRdsParamsWithDefaults() *RdsParams`

NewRdsParamsWithDefaults instantiates a new RdsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneId

`func (o *RdsParams) GetAvailabilityZoneId() int64`

GetAvailabilityZoneId returns the AvailabilityZoneId field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdOk

`func (o *RdsParams) GetAvailabilityZoneIdOk() (*int64, bool)`

GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneId

`func (o *RdsParams) SetAvailabilityZoneId(v int64)`

SetAvailabilityZoneId sets AvailabilityZoneId field to given value.

### HasAvailabilityZoneId

`func (o *RdsParams) HasAvailabilityZoneId() bool`

HasAvailabilityZoneId returns a boolean if a field has been set.

### SetAvailabilityZoneIdNil

`func (o *RdsParams) SetAvailabilityZoneIdNil(b bool)`

 SetAvailabilityZoneIdNil sets the value for AvailabilityZoneId to be an explicit nil

### UnsetAvailabilityZoneId
`func (o *RdsParams) UnsetAvailabilityZoneId()`

UnsetAvailabilityZoneId ensures that no value is present for AvailabilityZoneId, not even an explicit nil
### GetDbInstanceId

`func (o *RdsParams) GetDbInstanceId() string`

GetDbInstanceId returns the DbInstanceId field if non-nil, zero value otherwise.

### GetDbInstanceIdOk

`func (o *RdsParams) GetDbInstanceIdOk() (*string, bool)`

GetDbInstanceIdOk returns a tuple with the DbInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInstanceId

`func (o *RdsParams) SetDbInstanceId(v string)`

SetDbInstanceId sets DbInstanceId field to given value.


### SetDbInstanceIdNil

`func (o *RdsParams) SetDbInstanceIdNil(b bool)`

 SetDbInstanceIdNil sets the value for DbInstanceId to be an explicit nil

### UnsetDbInstanceId
`func (o *RdsParams) UnsetDbInstanceId()`

UnsetDbInstanceId ensures that no value is present for DbInstanceId, not even an explicit nil
### GetDbOptionGroupId

`func (o *RdsParams) GetDbOptionGroupId() int64`

GetDbOptionGroupId returns the DbOptionGroupId field if non-nil, zero value otherwise.

### GetDbOptionGroupIdOk

`func (o *RdsParams) GetDbOptionGroupIdOk() (*int64, bool)`

GetDbOptionGroupIdOk returns a tuple with the DbOptionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOptionGroupId

`func (o *RdsParams) SetDbOptionGroupId(v int64)`

SetDbOptionGroupId sets DbOptionGroupId field to given value.

### HasDbOptionGroupId

`func (o *RdsParams) HasDbOptionGroupId() bool`

HasDbOptionGroupId returns a boolean if a field has been set.

### SetDbOptionGroupIdNil

`func (o *RdsParams) SetDbOptionGroupIdNil(b bool)`

 SetDbOptionGroupIdNil sets the value for DbOptionGroupId to be an explicit nil

### UnsetDbOptionGroupId
`func (o *RdsParams) UnsetDbOptionGroupId()`

UnsetDbOptionGroupId ensures that no value is present for DbOptionGroupId, not even an explicit nil
### GetDbParameterGroupId

`func (o *RdsParams) GetDbParameterGroupId() int64`

GetDbParameterGroupId returns the DbParameterGroupId field if non-nil, zero value otherwise.

### GetDbParameterGroupIdOk

`func (o *RdsParams) GetDbParameterGroupIdOk() (*int64, bool)`

GetDbParameterGroupIdOk returns a tuple with the DbParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbParameterGroupId

`func (o *RdsParams) SetDbParameterGroupId(v int64)`

SetDbParameterGroupId sets DbParameterGroupId field to given value.

### HasDbParameterGroupId

`func (o *RdsParams) HasDbParameterGroupId() bool`

HasDbParameterGroupId returns a boolean if a field has been set.

### SetDbParameterGroupIdNil

`func (o *RdsParams) SetDbParameterGroupIdNil(b bool)`

 SetDbParameterGroupIdNil sets the value for DbParameterGroupId to be an explicit nil

### UnsetDbParameterGroupId
`func (o *RdsParams) UnsetDbParameterGroupId()`

UnsetDbParameterGroupId ensures that no value is present for DbParameterGroupId, not even an explicit nil
### GetDbPort

`func (o *RdsParams) GetDbPort() int32`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *RdsParams) GetDbPortOk() (*int32, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *RdsParams) SetDbPort(v int32)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *RdsParams) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### SetDbPortNil

`func (o *RdsParams) SetDbPortNil(b bool)`

 SetDbPortNil sets the value for DbPort to be an explicit nil

### UnsetDbPort
`func (o *RdsParams) UnsetDbPort()`

UnsetDbPort ensures that no value is present for DbPort, not even an explicit nil
### GetEnableAutoMinorVersionUpgrade

`func (o *RdsParams) GetEnableAutoMinorVersionUpgrade() bool`

GetEnableAutoMinorVersionUpgrade returns the EnableAutoMinorVersionUpgrade field if non-nil, zero value otherwise.

### GetEnableAutoMinorVersionUpgradeOk

`func (o *RdsParams) GetEnableAutoMinorVersionUpgradeOk() (*bool, bool)`

GetEnableAutoMinorVersionUpgradeOk returns a tuple with the EnableAutoMinorVersionUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoMinorVersionUpgrade

`func (o *RdsParams) SetEnableAutoMinorVersionUpgrade(v bool)`

SetEnableAutoMinorVersionUpgrade sets EnableAutoMinorVersionUpgrade field to given value.

### HasEnableAutoMinorVersionUpgrade

`func (o *RdsParams) HasEnableAutoMinorVersionUpgrade() bool`

HasEnableAutoMinorVersionUpgrade returns a boolean if a field has been set.

### SetEnableAutoMinorVersionUpgradeNil

`func (o *RdsParams) SetEnableAutoMinorVersionUpgradeNil(b bool)`

 SetEnableAutoMinorVersionUpgradeNil sets the value for EnableAutoMinorVersionUpgrade to be an explicit nil

### UnsetEnableAutoMinorVersionUpgrade
`func (o *RdsParams) UnsetEnableAutoMinorVersionUpgrade()`

UnsetEnableAutoMinorVersionUpgrade ensures that no value is present for EnableAutoMinorVersionUpgrade, not even an explicit nil
### GetEnableCopyTagsToSnapshots

`func (o *RdsParams) GetEnableCopyTagsToSnapshots() bool`

GetEnableCopyTagsToSnapshots returns the EnableCopyTagsToSnapshots field if non-nil, zero value otherwise.

### GetEnableCopyTagsToSnapshotsOk

`func (o *RdsParams) GetEnableCopyTagsToSnapshotsOk() (*bool, bool)`

GetEnableCopyTagsToSnapshotsOk returns a tuple with the EnableCopyTagsToSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyTagsToSnapshots

`func (o *RdsParams) SetEnableCopyTagsToSnapshots(v bool)`

SetEnableCopyTagsToSnapshots sets EnableCopyTagsToSnapshots field to given value.

### HasEnableCopyTagsToSnapshots

`func (o *RdsParams) HasEnableCopyTagsToSnapshots() bool`

HasEnableCopyTagsToSnapshots returns a boolean if a field has been set.

### SetEnableCopyTagsToSnapshotsNil

`func (o *RdsParams) SetEnableCopyTagsToSnapshotsNil(b bool)`

 SetEnableCopyTagsToSnapshotsNil sets the value for EnableCopyTagsToSnapshots to be an explicit nil

### UnsetEnableCopyTagsToSnapshots
`func (o *RdsParams) UnsetEnableCopyTagsToSnapshots()`

UnsetEnableCopyTagsToSnapshots ensures that no value is present for EnableCopyTagsToSnapshots, not even an explicit nil
### GetEnableDbAuthentication

`func (o *RdsParams) GetEnableDbAuthentication() bool`

GetEnableDbAuthentication returns the EnableDbAuthentication field if non-nil, zero value otherwise.

### GetEnableDbAuthenticationOk

`func (o *RdsParams) GetEnableDbAuthenticationOk() (*bool, bool)`

GetEnableDbAuthenticationOk returns a tuple with the EnableDbAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDbAuthentication

`func (o *RdsParams) SetEnableDbAuthentication(v bool)`

SetEnableDbAuthentication sets EnableDbAuthentication field to given value.

### HasEnableDbAuthentication

`func (o *RdsParams) HasEnableDbAuthentication() bool`

HasEnableDbAuthentication returns a boolean if a field has been set.

### SetEnableDbAuthenticationNil

`func (o *RdsParams) SetEnableDbAuthenticationNil(b bool)`

 SetEnableDbAuthenticationNil sets the value for EnableDbAuthentication to be an explicit nil

### UnsetEnableDbAuthentication
`func (o *RdsParams) UnsetEnableDbAuthentication()`

UnsetEnableDbAuthentication ensures that no value is present for EnableDbAuthentication, not even an explicit nil
### GetEnablePublicAccessibility

`func (o *RdsParams) GetEnablePublicAccessibility() bool`

GetEnablePublicAccessibility returns the EnablePublicAccessibility field if non-nil, zero value otherwise.

### GetEnablePublicAccessibilityOk

`func (o *RdsParams) GetEnablePublicAccessibilityOk() (*bool, bool)`

GetEnablePublicAccessibilityOk returns a tuple with the EnablePublicAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicAccessibility

`func (o *RdsParams) SetEnablePublicAccessibility(v bool)`

SetEnablePublicAccessibility sets EnablePublicAccessibility field to given value.

### HasEnablePublicAccessibility

`func (o *RdsParams) HasEnablePublicAccessibility() bool`

HasEnablePublicAccessibility returns a boolean if a field has been set.

### SetEnablePublicAccessibilityNil

`func (o *RdsParams) SetEnablePublicAccessibilityNil(b bool)`

 SetEnablePublicAccessibilityNil sets the value for EnablePublicAccessibility to be an explicit nil

### UnsetEnablePublicAccessibility
`func (o *RdsParams) UnsetEnablePublicAccessibility()`

UnsetEnablePublicAccessibility ensures that no value is present for EnablePublicAccessibility, not even an explicit nil
### GetIsMultiAzDeployment

`func (o *RdsParams) GetIsMultiAzDeployment() bool`

GetIsMultiAzDeployment returns the IsMultiAzDeployment field if non-nil, zero value otherwise.

### GetIsMultiAzDeploymentOk

`func (o *RdsParams) GetIsMultiAzDeploymentOk() (*bool, bool)`

GetIsMultiAzDeploymentOk returns a tuple with the IsMultiAzDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAzDeployment

`func (o *RdsParams) SetIsMultiAzDeployment(v bool)`

SetIsMultiAzDeployment sets IsMultiAzDeployment field to given value.

### HasIsMultiAzDeployment

`func (o *RdsParams) HasIsMultiAzDeployment() bool`

HasIsMultiAzDeployment returns a boolean if a field has been set.

### SetIsMultiAzDeploymentNil

`func (o *RdsParams) SetIsMultiAzDeploymentNil(b bool)`

 SetIsMultiAzDeploymentNil sets the value for IsMultiAzDeployment to be an explicit nil

### UnsetIsMultiAzDeployment
`func (o *RdsParams) UnsetIsMultiAzDeployment()`

UnsetIsMultiAzDeployment ensures that no value is present for IsMultiAzDeployment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


