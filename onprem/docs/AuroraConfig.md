# AuroraConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbPort** | **NullableInt32** | Specifies the port to use for the DB in the restored Aurora instance. | 
**DbInstanceId** | **NullableString** | Specifies the DB instance identifier to use for the restored DB. | 
**IsMultiAzDeployment** | **NullableBool** | Specifies whether this is a multi-az deployment or not. | 
**EnablePublicAccessibility** | Pointer to **NullableBool** | Specifies whether this DB will be publicly accessible or not. | [optional] 
**EnableIamDbAuthentication** | **NullableBool** | Specifies whether to enable IAM authentication for the DB. | 
**EnableCopyTagsToSnapshots** | **NullableBool** | Specifies whether to enable copying of tags to snapshots of the DB. | 
**EnableAutoMinorVersionUpgrade** | **NullableBool** | Specifies whether to enable auto minor version upgrade in the restored DB. | 
**DbOptionGroup** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies entity representing the Aurora option group to use while restoring the DB. | [optional] 
**DbParameterGroup** | Pointer to [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the entity representing the Aurora parameter group to use while restoring the DB. | [optional] 

## Methods

### NewAuroraConfig

`func NewAuroraConfig(dbPort NullableInt32, dbInstanceId NullableString, isMultiAzDeployment NullableBool, enableIamDbAuthentication NullableBool, enableCopyTagsToSnapshots NullableBool, enableAutoMinorVersionUpgrade NullableBool, ) *AuroraConfig`

NewAuroraConfig instantiates a new AuroraConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuroraConfigWithDefaults

`func NewAuroraConfigWithDefaults() *AuroraConfig`

NewAuroraConfigWithDefaults instantiates a new AuroraConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbPort

`func (o *AuroraConfig) GetDbPort() int32`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *AuroraConfig) GetDbPortOk() (*int32, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *AuroraConfig) SetDbPort(v int32)`

SetDbPort sets DbPort field to given value.


### SetDbPortNil

`func (o *AuroraConfig) SetDbPortNil(b bool)`

 SetDbPortNil sets the value for DbPort to be an explicit nil

### UnsetDbPort
`func (o *AuroraConfig) UnsetDbPort()`

UnsetDbPort ensures that no value is present for DbPort, not even an explicit nil
### GetDbInstanceId

`func (o *AuroraConfig) GetDbInstanceId() string`

GetDbInstanceId returns the DbInstanceId field if non-nil, zero value otherwise.

### GetDbInstanceIdOk

`func (o *AuroraConfig) GetDbInstanceIdOk() (*string, bool)`

GetDbInstanceIdOk returns a tuple with the DbInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInstanceId

`func (o *AuroraConfig) SetDbInstanceId(v string)`

SetDbInstanceId sets DbInstanceId field to given value.


### SetDbInstanceIdNil

`func (o *AuroraConfig) SetDbInstanceIdNil(b bool)`

 SetDbInstanceIdNil sets the value for DbInstanceId to be an explicit nil

### UnsetDbInstanceId
`func (o *AuroraConfig) UnsetDbInstanceId()`

UnsetDbInstanceId ensures that no value is present for DbInstanceId, not even an explicit nil
### GetIsMultiAzDeployment

`func (o *AuroraConfig) GetIsMultiAzDeployment() bool`

GetIsMultiAzDeployment returns the IsMultiAzDeployment field if non-nil, zero value otherwise.

### GetIsMultiAzDeploymentOk

`func (o *AuroraConfig) GetIsMultiAzDeploymentOk() (*bool, bool)`

GetIsMultiAzDeploymentOk returns a tuple with the IsMultiAzDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAzDeployment

`func (o *AuroraConfig) SetIsMultiAzDeployment(v bool)`

SetIsMultiAzDeployment sets IsMultiAzDeployment field to given value.


### SetIsMultiAzDeploymentNil

`func (o *AuroraConfig) SetIsMultiAzDeploymentNil(b bool)`

 SetIsMultiAzDeploymentNil sets the value for IsMultiAzDeployment to be an explicit nil

### UnsetIsMultiAzDeployment
`func (o *AuroraConfig) UnsetIsMultiAzDeployment()`

UnsetIsMultiAzDeployment ensures that no value is present for IsMultiAzDeployment, not even an explicit nil
### GetEnablePublicAccessibility

`func (o *AuroraConfig) GetEnablePublicAccessibility() bool`

GetEnablePublicAccessibility returns the EnablePublicAccessibility field if non-nil, zero value otherwise.

### GetEnablePublicAccessibilityOk

`func (o *AuroraConfig) GetEnablePublicAccessibilityOk() (*bool, bool)`

GetEnablePublicAccessibilityOk returns a tuple with the EnablePublicAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicAccessibility

`func (o *AuroraConfig) SetEnablePublicAccessibility(v bool)`

SetEnablePublicAccessibility sets EnablePublicAccessibility field to given value.

### HasEnablePublicAccessibility

`func (o *AuroraConfig) HasEnablePublicAccessibility() bool`

HasEnablePublicAccessibility returns a boolean if a field has been set.

### SetEnablePublicAccessibilityNil

`func (o *AuroraConfig) SetEnablePublicAccessibilityNil(b bool)`

 SetEnablePublicAccessibilityNil sets the value for EnablePublicAccessibility to be an explicit nil

### UnsetEnablePublicAccessibility
`func (o *AuroraConfig) UnsetEnablePublicAccessibility()`

UnsetEnablePublicAccessibility ensures that no value is present for EnablePublicAccessibility, not even an explicit nil
### GetEnableIamDbAuthentication

`func (o *AuroraConfig) GetEnableIamDbAuthentication() bool`

GetEnableIamDbAuthentication returns the EnableIamDbAuthentication field if non-nil, zero value otherwise.

### GetEnableIamDbAuthenticationOk

`func (o *AuroraConfig) GetEnableIamDbAuthenticationOk() (*bool, bool)`

GetEnableIamDbAuthenticationOk returns a tuple with the EnableIamDbAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIamDbAuthentication

`func (o *AuroraConfig) SetEnableIamDbAuthentication(v bool)`

SetEnableIamDbAuthentication sets EnableIamDbAuthentication field to given value.


### SetEnableIamDbAuthenticationNil

`func (o *AuroraConfig) SetEnableIamDbAuthenticationNil(b bool)`

 SetEnableIamDbAuthenticationNil sets the value for EnableIamDbAuthentication to be an explicit nil

### UnsetEnableIamDbAuthentication
`func (o *AuroraConfig) UnsetEnableIamDbAuthentication()`

UnsetEnableIamDbAuthentication ensures that no value is present for EnableIamDbAuthentication, not even an explicit nil
### GetEnableCopyTagsToSnapshots

`func (o *AuroraConfig) GetEnableCopyTagsToSnapshots() bool`

GetEnableCopyTagsToSnapshots returns the EnableCopyTagsToSnapshots field if non-nil, zero value otherwise.

### GetEnableCopyTagsToSnapshotsOk

`func (o *AuroraConfig) GetEnableCopyTagsToSnapshotsOk() (*bool, bool)`

GetEnableCopyTagsToSnapshotsOk returns a tuple with the EnableCopyTagsToSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyTagsToSnapshots

`func (o *AuroraConfig) SetEnableCopyTagsToSnapshots(v bool)`

SetEnableCopyTagsToSnapshots sets EnableCopyTagsToSnapshots field to given value.


### SetEnableCopyTagsToSnapshotsNil

`func (o *AuroraConfig) SetEnableCopyTagsToSnapshotsNil(b bool)`

 SetEnableCopyTagsToSnapshotsNil sets the value for EnableCopyTagsToSnapshots to be an explicit nil

### UnsetEnableCopyTagsToSnapshots
`func (o *AuroraConfig) UnsetEnableCopyTagsToSnapshots()`

UnsetEnableCopyTagsToSnapshots ensures that no value is present for EnableCopyTagsToSnapshots, not even an explicit nil
### GetEnableAutoMinorVersionUpgrade

`func (o *AuroraConfig) GetEnableAutoMinorVersionUpgrade() bool`

GetEnableAutoMinorVersionUpgrade returns the EnableAutoMinorVersionUpgrade field if non-nil, zero value otherwise.

### GetEnableAutoMinorVersionUpgradeOk

`func (o *AuroraConfig) GetEnableAutoMinorVersionUpgradeOk() (*bool, bool)`

GetEnableAutoMinorVersionUpgradeOk returns a tuple with the EnableAutoMinorVersionUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoMinorVersionUpgrade

`func (o *AuroraConfig) SetEnableAutoMinorVersionUpgrade(v bool)`

SetEnableAutoMinorVersionUpgrade sets EnableAutoMinorVersionUpgrade field to given value.


### SetEnableAutoMinorVersionUpgradeNil

`func (o *AuroraConfig) SetEnableAutoMinorVersionUpgradeNil(b bool)`

 SetEnableAutoMinorVersionUpgradeNil sets the value for EnableAutoMinorVersionUpgrade to be an explicit nil

### UnsetEnableAutoMinorVersionUpgrade
`func (o *AuroraConfig) UnsetEnableAutoMinorVersionUpgrade()`

UnsetEnableAutoMinorVersionUpgrade ensures that no value is present for EnableAutoMinorVersionUpgrade, not even an explicit nil
### GetDbOptionGroup

`func (o *AuroraConfig) GetDbOptionGroup() RecoveryObjectIdentifier`

GetDbOptionGroup returns the DbOptionGroup field if non-nil, zero value otherwise.

### GetDbOptionGroupOk

`func (o *AuroraConfig) GetDbOptionGroupOk() (*RecoveryObjectIdentifier, bool)`

GetDbOptionGroupOk returns a tuple with the DbOptionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOptionGroup

`func (o *AuroraConfig) SetDbOptionGroup(v RecoveryObjectIdentifier)`

SetDbOptionGroup sets DbOptionGroup field to given value.

### HasDbOptionGroup

`func (o *AuroraConfig) HasDbOptionGroup() bool`

HasDbOptionGroup returns a boolean if a field has been set.

### SetDbOptionGroupNil

`func (o *AuroraConfig) SetDbOptionGroupNil(b bool)`

 SetDbOptionGroupNil sets the value for DbOptionGroup to be an explicit nil

### UnsetDbOptionGroup
`func (o *AuroraConfig) UnsetDbOptionGroup()`

UnsetDbOptionGroup ensures that no value is present for DbOptionGroup, not even an explicit nil
### GetDbParameterGroup

`func (o *AuroraConfig) GetDbParameterGroup() RecoveryObjectIdentifier`

GetDbParameterGroup returns the DbParameterGroup field if non-nil, zero value otherwise.

### GetDbParameterGroupOk

`func (o *AuroraConfig) GetDbParameterGroupOk() (*RecoveryObjectIdentifier, bool)`

GetDbParameterGroupOk returns a tuple with the DbParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbParameterGroup

`func (o *AuroraConfig) SetDbParameterGroup(v RecoveryObjectIdentifier)`

SetDbParameterGroup sets DbParameterGroup field to given value.

### HasDbParameterGroup

`func (o *AuroraConfig) HasDbParameterGroup() bool`

HasDbParameterGroup returns a boolean if a field has been set.

### SetDbParameterGroupNil

`func (o *AuroraConfig) SetDbParameterGroupNil(b bool)`

 SetDbParameterGroupNil sets the value for DbParameterGroup to be an explicit nil

### UnsetDbParameterGroup
`func (o *AuroraConfig) UnsetDbParameterGroup()`

UnsetDbParameterGroup ensures that no value is present for DbParameterGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


