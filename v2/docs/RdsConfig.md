# RdsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbInstanceId** | **NullableString** | Specifies the DB instance identifier to use for the restored DB. | 
**DbOptionGroup** | Pointer to [**NullableRdsConfigDbOptionGroup**](RdsConfigDbOptionGroup.md) |  | [optional] 
**DbParameterGroup** | Pointer to [**NullableRdsConfigDbParameterGroup**](RdsConfigDbParameterGroup.md) |  | [optional] 
**DbPort** | **NullableInt32** | Specifies the port to use for the DB in the restored RDS instance. | 
**EnableAutoMinorVersionUpgrade** | **NullableBool** | Specifies whether to enable auto minor version upgrade in the restored DB. | 
**EnableCopyTagsToSnapshots** | **NullableBool** | Specifies whether to enable copying of tags to snapshots of the DB. | 
**EnableIamDbAuthentication** | **NullableBool** | Specifies whether to enable IAM authentication for the DB. | 
**EnablePublicAccessibility** | Pointer to **NullableBool** | Specifies whether this DB will be publicly accessible or not. | [optional] 
**IsMultiAzDeployment** | **NullableBool** | Specifies whether this is a multi-az deployment or not. | 
**PointInTimeUsecs** | Pointer to **NullableInt64** | Specifies a point in time for recovery in microseconds. | [optional] 

## Methods

### NewRdsConfig

`func NewRdsConfig(dbInstanceId NullableString, dbPort NullableInt32, enableAutoMinorVersionUpgrade NullableBool, enableCopyTagsToSnapshots NullableBool, enableIamDbAuthentication NullableBool, isMultiAzDeployment NullableBool, ) *RdsConfig`

NewRdsConfig instantiates a new RdsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsConfigWithDefaults

`func NewRdsConfigWithDefaults() *RdsConfig`

NewRdsConfigWithDefaults instantiates a new RdsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbInstanceId

`func (o *RdsConfig) GetDbInstanceId() string`

GetDbInstanceId returns the DbInstanceId field if non-nil, zero value otherwise.

### GetDbInstanceIdOk

`func (o *RdsConfig) GetDbInstanceIdOk() (*string, bool)`

GetDbInstanceIdOk returns a tuple with the DbInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInstanceId

`func (o *RdsConfig) SetDbInstanceId(v string)`

SetDbInstanceId sets DbInstanceId field to given value.


### SetDbInstanceIdNil

`func (o *RdsConfig) SetDbInstanceIdNil(b bool)`

 SetDbInstanceIdNil sets the value for DbInstanceId to be an explicit nil

### UnsetDbInstanceId
`func (o *RdsConfig) UnsetDbInstanceId()`

UnsetDbInstanceId ensures that no value is present for DbInstanceId, not even an explicit nil
### GetDbOptionGroup

`func (o *RdsConfig) GetDbOptionGroup() RdsConfigDbOptionGroup`

GetDbOptionGroup returns the DbOptionGroup field if non-nil, zero value otherwise.

### GetDbOptionGroupOk

`func (o *RdsConfig) GetDbOptionGroupOk() (*RdsConfigDbOptionGroup, bool)`

GetDbOptionGroupOk returns a tuple with the DbOptionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOptionGroup

`func (o *RdsConfig) SetDbOptionGroup(v RdsConfigDbOptionGroup)`

SetDbOptionGroup sets DbOptionGroup field to given value.

### HasDbOptionGroup

`func (o *RdsConfig) HasDbOptionGroup() bool`

HasDbOptionGroup returns a boolean if a field has been set.

### SetDbOptionGroupNil

`func (o *RdsConfig) SetDbOptionGroupNil(b bool)`

 SetDbOptionGroupNil sets the value for DbOptionGroup to be an explicit nil

### UnsetDbOptionGroup
`func (o *RdsConfig) UnsetDbOptionGroup()`

UnsetDbOptionGroup ensures that no value is present for DbOptionGroup, not even an explicit nil
### GetDbParameterGroup

`func (o *RdsConfig) GetDbParameterGroup() RdsConfigDbParameterGroup`

GetDbParameterGroup returns the DbParameterGroup field if non-nil, zero value otherwise.

### GetDbParameterGroupOk

`func (o *RdsConfig) GetDbParameterGroupOk() (*RdsConfigDbParameterGroup, bool)`

GetDbParameterGroupOk returns a tuple with the DbParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbParameterGroup

`func (o *RdsConfig) SetDbParameterGroup(v RdsConfigDbParameterGroup)`

SetDbParameterGroup sets DbParameterGroup field to given value.

### HasDbParameterGroup

`func (o *RdsConfig) HasDbParameterGroup() bool`

HasDbParameterGroup returns a boolean if a field has been set.

### SetDbParameterGroupNil

`func (o *RdsConfig) SetDbParameterGroupNil(b bool)`

 SetDbParameterGroupNil sets the value for DbParameterGroup to be an explicit nil

### UnsetDbParameterGroup
`func (o *RdsConfig) UnsetDbParameterGroup()`

UnsetDbParameterGroup ensures that no value is present for DbParameterGroup, not even an explicit nil
### GetDbPort

`func (o *RdsConfig) GetDbPort() int32`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *RdsConfig) GetDbPortOk() (*int32, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *RdsConfig) SetDbPort(v int32)`

SetDbPort sets DbPort field to given value.


### SetDbPortNil

`func (o *RdsConfig) SetDbPortNil(b bool)`

 SetDbPortNil sets the value for DbPort to be an explicit nil

### UnsetDbPort
`func (o *RdsConfig) UnsetDbPort()`

UnsetDbPort ensures that no value is present for DbPort, not even an explicit nil
### GetEnableAutoMinorVersionUpgrade

`func (o *RdsConfig) GetEnableAutoMinorVersionUpgrade() bool`

GetEnableAutoMinorVersionUpgrade returns the EnableAutoMinorVersionUpgrade field if non-nil, zero value otherwise.

### GetEnableAutoMinorVersionUpgradeOk

`func (o *RdsConfig) GetEnableAutoMinorVersionUpgradeOk() (*bool, bool)`

GetEnableAutoMinorVersionUpgradeOk returns a tuple with the EnableAutoMinorVersionUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoMinorVersionUpgrade

`func (o *RdsConfig) SetEnableAutoMinorVersionUpgrade(v bool)`

SetEnableAutoMinorVersionUpgrade sets EnableAutoMinorVersionUpgrade field to given value.


### SetEnableAutoMinorVersionUpgradeNil

`func (o *RdsConfig) SetEnableAutoMinorVersionUpgradeNil(b bool)`

 SetEnableAutoMinorVersionUpgradeNil sets the value for EnableAutoMinorVersionUpgrade to be an explicit nil

### UnsetEnableAutoMinorVersionUpgrade
`func (o *RdsConfig) UnsetEnableAutoMinorVersionUpgrade()`

UnsetEnableAutoMinorVersionUpgrade ensures that no value is present for EnableAutoMinorVersionUpgrade, not even an explicit nil
### GetEnableCopyTagsToSnapshots

`func (o *RdsConfig) GetEnableCopyTagsToSnapshots() bool`

GetEnableCopyTagsToSnapshots returns the EnableCopyTagsToSnapshots field if non-nil, zero value otherwise.

### GetEnableCopyTagsToSnapshotsOk

`func (o *RdsConfig) GetEnableCopyTagsToSnapshotsOk() (*bool, bool)`

GetEnableCopyTagsToSnapshotsOk returns a tuple with the EnableCopyTagsToSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyTagsToSnapshots

`func (o *RdsConfig) SetEnableCopyTagsToSnapshots(v bool)`

SetEnableCopyTagsToSnapshots sets EnableCopyTagsToSnapshots field to given value.


### SetEnableCopyTagsToSnapshotsNil

`func (o *RdsConfig) SetEnableCopyTagsToSnapshotsNil(b bool)`

 SetEnableCopyTagsToSnapshotsNil sets the value for EnableCopyTagsToSnapshots to be an explicit nil

### UnsetEnableCopyTagsToSnapshots
`func (o *RdsConfig) UnsetEnableCopyTagsToSnapshots()`

UnsetEnableCopyTagsToSnapshots ensures that no value is present for EnableCopyTagsToSnapshots, not even an explicit nil
### GetEnableIamDbAuthentication

`func (o *RdsConfig) GetEnableIamDbAuthentication() bool`

GetEnableIamDbAuthentication returns the EnableIamDbAuthentication field if non-nil, zero value otherwise.

### GetEnableIamDbAuthenticationOk

`func (o *RdsConfig) GetEnableIamDbAuthenticationOk() (*bool, bool)`

GetEnableIamDbAuthenticationOk returns a tuple with the EnableIamDbAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIamDbAuthentication

`func (o *RdsConfig) SetEnableIamDbAuthentication(v bool)`

SetEnableIamDbAuthentication sets EnableIamDbAuthentication field to given value.


### SetEnableIamDbAuthenticationNil

`func (o *RdsConfig) SetEnableIamDbAuthenticationNil(b bool)`

 SetEnableIamDbAuthenticationNil sets the value for EnableIamDbAuthentication to be an explicit nil

### UnsetEnableIamDbAuthentication
`func (o *RdsConfig) UnsetEnableIamDbAuthentication()`

UnsetEnableIamDbAuthentication ensures that no value is present for EnableIamDbAuthentication, not even an explicit nil
### GetEnablePublicAccessibility

`func (o *RdsConfig) GetEnablePublicAccessibility() bool`

GetEnablePublicAccessibility returns the EnablePublicAccessibility field if non-nil, zero value otherwise.

### GetEnablePublicAccessibilityOk

`func (o *RdsConfig) GetEnablePublicAccessibilityOk() (*bool, bool)`

GetEnablePublicAccessibilityOk returns a tuple with the EnablePublicAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicAccessibility

`func (o *RdsConfig) SetEnablePublicAccessibility(v bool)`

SetEnablePublicAccessibility sets EnablePublicAccessibility field to given value.

### HasEnablePublicAccessibility

`func (o *RdsConfig) HasEnablePublicAccessibility() bool`

HasEnablePublicAccessibility returns a boolean if a field has been set.

### SetEnablePublicAccessibilityNil

`func (o *RdsConfig) SetEnablePublicAccessibilityNil(b bool)`

 SetEnablePublicAccessibilityNil sets the value for EnablePublicAccessibility to be an explicit nil

### UnsetEnablePublicAccessibility
`func (o *RdsConfig) UnsetEnablePublicAccessibility()`

UnsetEnablePublicAccessibility ensures that no value is present for EnablePublicAccessibility, not even an explicit nil
### GetIsMultiAzDeployment

`func (o *RdsConfig) GetIsMultiAzDeployment() bool`

GetIsMultiAzDeployment returns the IsMultiAzDeployment field if non-nil, zero value otherwise.

### GetIsMultiAzDeploymentOk

`func (o *RdsConfig) GetIsMultiAzDeploymentOk() (*bool, bool)`

GetIsMultiAzDeploymentOk returns a tuple with the IsMultiAzDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAzDeployment

`func (o *RdsConfig) SetIsMultiAzDeployment(v bool)`

SetIsMultiAzDeployment sets IsMultiAzDeployment field to given value.


### SetIsMultiAzDeploymentNil

`func (o *RdsConfig) SetIsMultiAzDeploymentNil(b bool)`

 SetIsMultiAzDeploymentNil sets the value for IsMultiAzDeployment to be an explicit nil

### UnsetIsMultiAzDeployment
`func (o *RdsConfig) UnsetIsMultiAzDeployment()`

UnsetIsMultiAzDeployment ensures that no value is present for IsMultiAzDeployment, not even an explicit nil
### GetPointInTimeUsecs

`func (o *RdsConfig) GetPointInTimeUsecs() int64`

GetPointInTimeUsecs returns the PointInTimeUsecs field if non-nil, zero value otherwise.

### GetPointInTimeUsecsOk

`func (o *RdsConfig) GetPointInTimeUsecsOk() (*int64, bool)`

GetPointInTimeUsecsOk returns a tuple with the PointInTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUsecs

`func (o *RdsConfig) SetPointInTimeUsecs(v int64)`

SetPointInTimeUsecs sets PointInTimeUsecs field to given value.

### HasPointInTimeUsecs

`func (o *RdsConfig) HasPointInTimeUsecs() bool`

HasPointInTimeUsecs returns a boolean if a field has been set.

### SetPointInTimeUsecsNil

`func (o *RdsConfig) SetPointInTimeUsecsNil(b bool)`

 SetPointInTimeUsecsNil sets the value for PointInTimeUsecs to be an explicit nil

### UnsetPointInTimeUsecs
`func (o *RdsConfig) UnsetPointInTimeUsecs()`

UnsetPointInTimeUsecs ensures that no value is present for PointInTimeUsecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


