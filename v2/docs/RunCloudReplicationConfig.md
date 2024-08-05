# RunCloudReplicationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsTarget** | Pointer to [**AWSTargetConfig**](AWSTargetConfig.md) |  | [optional] 
**AzureTarget** | Pointer to [**AzureTargetConfig**](AzureTargetConfig.md) |  | [optional] 
**OnLegalHold** | Pointer to **NullableBool** | Specifies if the Run is on legal hold. | [optional] 
**Retention** | Pointer to [**Retention**](Retention.md) |  | [optional] 
**TargetType** | **string** | Specifies the type of target to which replication need to be performed. | 

## Methods

### NewRunCloudReplicationConfig

`func NewRunCloudReplicationConfig(targetType string, ) *RunCloudReplicationConfig`

NewRunCloudReplicationConfig instantiates a new RunCloudReplicationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunCloudReplicationConfigWithDefaults

`func NewRunCloudReplicationConfigWithDefaults() *RunCloudReplicationConfig`

NewRunCloudReplicationConfigWithDefaults instantiates a new RunCloudReplicationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsTarget

`func (o *RunCloudReplicationConfig) GetAwsTarget() AWSTargetConfig`

GetAwsTarget returns the AwsTarget field if non-nil, zero value otherwise.

### GetAwsTargetOk

`func (o *RunCloudReplicationConfig) GetAwsTargetOk() (*AWSTargetConfig, bool)`

GetAwsTargetOk returns a tuple with the AwsTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTarget

`func (o *RunCloudReplicationConfig) SetAwsTarget(v AWSTargetConfig)`

SetAwsTarget sets AwsTarget field to given value.

### HasAwsTarget

`func (o *RunCloudReplicationConfig) HasAwsTarget() bool`

HasAwsTarget returns a boolean if a field has been set.

### GetAzureTarget

`func (o *RunCloudReplicationConfig) GetAzureTarget() AzureTargetConfig`

GetAzureTarget returns the AzureTarget field if non-nil, zero value otherwise.

### GetAzureTargetOk

`func (o *RunCloudReplicationConfig) GetAzureTargetOk() (*AzureTargetConfig, bool)`

GetAzureTargetOk returns a tuple with the AzureTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTarget

`func (o *RunCloudReplicationConfig) SetAzureTarget(v AzureTargetConfig)`

SetAzureTarget sets AzureTarget field to given value.

### HasAzureTarget

`func (o *RunCloudReplicationConfig) HasAzureTarget() bool`

HasAzureTarget returns a boolean if a field has been set.

### GetOnLegalHold

`func (o *RunCloudReplicationConfig) GetOnLegalHold() bool`

GetOnLegalHold returns the OnLegalHold field if non-nil, zero value otherwise.

### GetOnLegalHoldOk

`func (o *RunCloudReplicationConfig) GetOnLegalHoldOk() (*bool, bool)`

GetOnLegalHoldOk returns a tuple with the OnLegalHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnLegalHold

`func (o *RunCloudReplicationConfig) SetOnLegalHold(v bool)`

SetOnLegalHold sets OnLegalHold field to given value.

### HasOnLegalHold

`func (o *RunCloudReplicationConfig) HasOnLegalHold() bool`

HasOnLegalHold returns a boolean if a field has been set.

### SetOnLegalHoldNil

`func (o *RunCloudReplicationConfig) SetOnLegalHoldNil(b bool)`

 SetOnLegalHoldNil sets the value for OnLegalHold to be an explicit nil

### UnsetOnLegalHold
`func (o *RunCloudReplicationConfig) UnsetOnLegalHold()`

UnsetOnLegalHold ensures that no value is present for OnLegalHold, not even an explicit nil
### GetRetention

`func (o *RunCloudReplicationConfig) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RunCloudReplicationConfig) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RunCloudReplicationConfig) SetRetention(v Retention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *RunCloudReplicationConfig) HasRetention() bool`

HasRetention returns a boolean if a field has been set.

### GetTargetType

`func (o *RunCloudReplicationConfig) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *RunCloudReplicationConfig) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *RunCloudReplicationConfig) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


