# DeployDBInstancesToRDSParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoMinorVersionUpgrade** | Pointer to **NullableBool** | Whether to enable auto minor version upgrade in the restored DB. | [optional] 
**AvailabilityZone** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**CopyTagsToSnapshots** | Pointer to **NullableBool** | Whether to enable copying of tags to snapshots of the DB. | [optional] 
**DbInstanceId** | Pointer to **NullableString** | The DB instance identifier to use for the restored DB. This field is required. | [optional] 
**DbOptionGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**DbParameterGroup** | Pointer to [**EntityProto**](EntityProto.md) |  | [optional] 
**DbPort** | Pointer to **NullableInt32** | Port to use for the DB in the restored RDS instance. | [optional] 
**IamDbAuthentication** | Pointer to **NullableBool** | Whether to enable IAM authentication for the DB. | [optional] 
**MultiAzDeployment** | Pointer to **NullableBool** | Whether this is a multi-az deployment or not. | [optional] 
**PointInTimeParams** | Pointer to [**DeployDBInstancesToRDSParamsPointInTimeRestoreParams**](DeployDBInstancesToRDSParamsPointInTimeRestoreParams.md) |  | [optional] 
**PublicAccessibility** | Pointer to **NullableBool** | Whether this DB will be publicly accessible or not. | [optional] 

## Methods

### NewDeployDBInstancesToRDSParams

`func NewDeployDBInstancesToRDSParams() *DeployDBInstancesToRDSParams`

NewDeployDBInstancesToRDSParams instantiates a new DeployDBInstancesToRDSParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployDBInstancesToRDSParamsWithDefaults

`func NewDeployDBInstancesToRDSParamsWithDefaults() *DeployDBInstancesToRDSParams`

NewDeployDBInstancesToRDSParamsWithDefaults instantiates a new DeployDBInstancesToRDSParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoMinorVersionUpgrade

`func (o *DeployDBInstancesToRDSParams) GetAutoMinorVersionUpgrade() bool`

GetAutoMinorVersionUpgrade returns the AutoMinorVersionUpgrade field if non-nil, zero value otherwise.

### GetAutoMinorVersionUpgradeOk

`func (o *DeployDBInstancesToRDSParams) GetAutoMinorVersionUpgradeOk() (*bool, bool)`

GetAutoMinorVersionUpgradeOk returns a tuple with the AutoMinorVersionUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMinorVersionUpgrade

`func (o *DeployDBInstancesToRDSParams) SetAutoMinorVersionUpgrade(v bool)`

SetAutoMinorVersionUpgrade sets AutoMinorVersionUpgrade field to given value.

### HasAutoMinorVersionUpgrade

`func (o *DeployDBInstancesToRDSParams) HasAutoMinorVersionUpgrade() bool`

HasAutoMinorVersionUpgrade returns a boolean if a field has been set.

### SetAutoMinorVersionUpgradeNil

`func (o *DeployDBInstancesToRDSParams) SetAutoMinorVersionUpgradeNil(b bool)`

 SetAutoMinorVersionUpgradeNil sets the value for AutoMinorVersionUpgrade to be an explicit nil

### UnsetAutoMinorVersionUpgrade
`func (o *DeployDBInstancesToRDSParams) UnsetAutoMinorVersionUpgrade()`

UnsetAutoMinorVersionUpgrade ensures that no value is present for AutoMinorVersionUpgrade, not even an explicit nil
### GetAvailabilityZone

`func (o *DeployDBInstancesToRDSParams) GetAvailabilityZone() EntityProto`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *DeployDBInstancesToRDSParams) GetAvailabilityZoneOk() (*EntityProto, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *DeployDBInstancesToRDSParams) SetAvailabilityZone(v EntityProto)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *DeployDBInstancesToRDSParams) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCopyTagsToSnapshots

`func (o *DeployDBInstancesToRDSParams) GetCopyTagsToSnapshots() bool`

GetCopyTagsToSnapshots returns the CopyTagsToSnapshots field if non-nil, zero value otherwise.

### GetCopyTagsToSnapshotsOk

`func (o *DeployDBInstancesToRDSParams) GetCopyTagsToSnapshotsOk() (*bool, bool)`

GetCopyTagsToSnapshotsOk returns a tuple with the CopyTagsToSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTagsToSnapshots

`func (o *DeployDBInstancesToRDSParams) SetCopyTagsToSnapshots(v bool)`

SetCopyTagsToSnapshots sets CopyTagsToSnapshots field to given value.

### HasCopyTagsToSnapshots

`func (o *DeployDBInstancesToRDSParams) HasCopyTagsToSnapshots() bool`

HasCopyTagsToSnapshots returns a boolean if a field has been set.

### SetCopyTagsToSnapshotsNil

`func (o *DeployDBInstancesToRDSParams) SetCopyTagsToSnapshotsNil(b bool)`

 SetCopyTagsToSnapshotsNil sets the value for CopyTagsToSnapshots to be an explicit nil

### UnsetCopyTagsToSnapshots
`func (o *DeployDBInstancesToRDSParams) UnsetCopyTagsToSnapshots()`

UnsetCopyTagsToSnapshots ensures that no value is present for CopyTagsToSnapshots, not even an explicit nil
### GetDbInstanceId

`func (o *DeployDBInstancesToRDSParams) GetDbInstanceId() string`

GetDbInstanceId returns the DbInstanceId field if non-nil, zero value otherwise.

### GetDbInstanceIdOk

`func (o *DeployDBInstancesToRDSParams) GetDbInstanceIdOk() (*string, bool)`

GetDbInstanceIdOk returns a tuple with the DbInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbInstanceId

`func (o *DeployDBInstancesToRDSParams) SetDbInstanceId(v string)`

SetDbInstanceId sets DbInstanceId field to given value.

### HasDbInstanceId

`func (o *DeployDBInstancesToRDSParams) HasDbInstanceId() bool`

HasDbInstanceId returns a boolean if a field has been set.

### SetDbInstanceIdNil

`func (o *DeployDBInstancesToRDSParams) SetDbInstanceIdNil(b bool)`

 SetDbInstanceIdNil sets the value for DbInstanceId to be an explicit nil

### UnsetDbInstanceId
`func (o *DeployDBInstancesToRDSParams) UnsetDbInstanceId()`

UnsetDbInstanceId ensures that no value is present for DbInstanceId, not even an explicit nil
### GetDbOptionGroup

`func (o *DeployDBInstancesToRDSParams) GetDbOptionGroup() EntityProto`

GetDbOptionGroup returns the DbOptionGroup field if non-nil, zero value otherwise.

### GetDbOptionGroupOk

`func (o *DeployDBInstancesToRDSParams) GetDbOptionGroupOk() (*EntityProto, bool)`

GetDbOptionGroupOk returns a tuple with the DbOptionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbOptionGroup

`func (o *DeployDBInstancesToRDSParams) SetDbOptionGroup(v EntityProto)`

SetDbOptionGroup sets DbOptionGroup field to given value.

### HasDbOptionGroup

`func (o *DeployDBInstancesToRDSParams) HasDbOptionGroup() bool`

HasDbOptionGroup returns a boolean if a field has been set.

### GetDbParameterGroup

`func (o *DeployDBInstancesToRDSParams) GetDbParameterGroup() EntityProto`

GetDbParameterGroup returns the DbParameterGroup field if non-nil, zero value otherwise.

### GetDbParameterGroupOk

`func (o *DeployDBInstancesToRDSParams) GetDbParameterGroupOk() (*EntityProto, bool)`

GetDbParameterGroupOk returns a tuple with the DbParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbParameterGroup

`func (o *DeployDBInstancesToRDSParams) SetDbParameterGroup(v EntityProto)`

SetDbParameterGroup sets DbParameterGroup field to given value.

### HasDbParameterGroup

`func (o *DeployDBInstancesToRDSParams) HasDbParameterGroup() bool`

HasDbParameterGroup returns a boolean if a field has been set.

### GetDbPort

`func (o *DeployDBInstancesToRDSParams) GetDbPort() int32`

GetDbPort returns the DbPort field if non-nil, zero value otherwise.

### GetDbPortOk

`func (o *DeployDBInstancesToRDSParams) GetDbPortOk() (*int32, bool)`

GetDbPortOk returns a tuple with the DbPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPort

`func (o *DeployDBInstancesToRDSParams) SetDbPort(v int32)`

SetDbPort sets DbPort field to given value.

### HasDbPort

`func (o *DeployDBInstancesToRDSParams) HasDbPort() bool`

HasDbPort returns a boolean if a field has been set.

### SetDbPortNil

`func (o *DeployDBInstancesToRDSParams) SetDbPortNil(b bool)`

 SetDbPortNil sets the value for DbPort to be an explicit nil

### UnsetDbPort
`func (o *DeployDBInstancesToRDSParams) UnsetDbPort()`

UnsetDbPort ensures that no value is present for DbPort, not even an explicit nil
### GetIamDbAuthentication

`func (o *DeployDBInstancesToRDSParams) GetIamDbAuthentication() bool`

GetIamDbAuthentication returns the IamDbAuthentication field if non-nil, zero value otherwise.

### GetIamDbAuthenticationOk

`func (o *DeployDBInstancesToRDSParams) GetIamDbAuthenticationOk() (*bool, bool)`

GetIamDbAuthenticationOk returns a tuple with the IamDbAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamDbAuthentication

`func (o *DeployDBInstancesToRDSParams) SetIamDbAuthentication(v bool)`

SetIamDbAuthentication sets IamDbAuthentication field to given value.

### HasIamDbAuthentication

`func (o *DeployDBInstancesToRDSParams) HasIamDbAuthentication() bool`

HasIamDbAuthentication returns a boolean if a field has been set.

### SetIamDbAuthenticationNil

`func (o *DeployDBInstancesToRDSParams) SetIamDbAuthenticationNil(b bool)`

 SetIamDbAuthenticationNil sets the value for IamDbAuthentication to be an explicit nil

### UnsetIamDbAuthentication
`func (o *DeployDBInstancesToRDSParams) UnsetIamDbAuthentication()`

UnsetIamDbAuthentication ensures that no value is present for IamDbAuthentication, not even an explicit nil
### GetMultiAzDeployment

`func (o *DeployDBInstancesToRDSParams) GetMultiAzDeployment() bool`

GetMultiAzDeployment returns the MultiAzDeployment field if non-nil, zero value otherwise.

### GetMultiAzDeploymentOk

`func (o *DeployDBInstancesToRDSParams) GetMultiAzDeploymentOk() (*bool, bool)`

GetMultiAzDeploymentOk returns a tuple with the MultiAzDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAzDeployment

`func (o *DeployDBInstancesToRDSParams) SetMultiAzDeployment(v bool)`

SetMultiAzDeployment sets MultiAzDeployment field to given value.

### HasMultiAzDeployment

`func (o *DeployDBInstancesToRDSParams) HasMultiAzDeployment() bool`

HasMultiAzDeployment returns a boolean if a field has been set.

### SetMultiAzDeploymentNil

`func (o *DeployDBInstancesToRDSParams) SetMultiAzDeploymentNil(b bool)`

 SetMultiAzDeploymentNil sets the value for MultiAzDeployment to be an explicit nil

### UnsetMultiAzDeployment
`func (o *DeployDBInstancesToRDSParams) UnsetMultiAzDeployment()`

UnsetMultiAzDeployment ensures that no value is present for MultiAzDeployment, not even an explicit nil
### GetPointInTimeParams

`func (o *DeployDBInstancesToRDSParams) GetPointInTimeParams() DeployDBInstancesToRDSParamsPointInTimeRestoreParams`

GetPointInTimeParams returns the PointInTimeParams field if non-nil, zero value otherwise.

### GetPointInTimeParamsOk

`func (o *DeployDBInstancesToRDSParams) GetPointInTimeParamsOk() (*DeployDBInstancesToRDSParamsPointInTimeRestoreParams, bool)`

GetPointInTimeParamsOk returns a tuple with the PointInTimeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeParams

`func (o *DeployDBInstancesToRDSParams) SetPointInTimeParams(v DeployDBInstancesToRDSParamsPointInTimeRestoreParams)`

SetPointInTimeParams sets PointInTimeParams field to given value.

### HasPointInTimeParams

`func (o *DeployDBInstancesToRDSParams) HasPointInTimeParams() bool`

HasPointInTimeParams returns a boolean if a field has been set.

### GetPublicAccessibility

`func (o *DeployDBInstancesToRDSParams) GetPublicAccessibility() bool`

GetPublicAccessibility returns the PublicAccessibility field if non-nil, zero value otherwise.

### GetPublicAccessibilityOk

`func (o *DeployDBInstancesToRDSParams) GetPublicAccessibilityOk() (*bool, bool)`

GetPublicAccessibilityOk returns a tuple with the PublicAccessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessibility

`func (o *DeployDBInstancesToRDSParams) SetPublicAccessibility(v bool)`

SetPublicAccessibility sets PublicAccessibility field to given value.

### HasPublicAccessibility

`func (o *DeployDBInstancesToRDSParams) HasPublicAccessibility() bool`

HasPublicAccessibility returns a boolean if a field has been set.

### SetPublicAccessibilityNil

`func (o *DeployDBInstancesToRDSParams) SetPublicAccessibilityNil(b bool)`

 SetPublicAccessibilityNil sets the value for PublicAccessibility to be an explicit nil

### UnsetPublicAccessibility
`func (o *DeployDBInstancesToRDSParams) UnsetPublicAccessibility()`

UnsetPublicAccessibility ensures that no value is present for PublicAccessibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


