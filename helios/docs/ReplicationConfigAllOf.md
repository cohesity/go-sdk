# ReplicationConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetType** | **string** | Specifies the type of target to which replication need to be performed. | 
**RemoteTargetConfig** | Pointer to [**RemoteTargetConfig**](RemoteTargetConfig.md) |  | [optional] 
**AwsTargetConfig** | Pointer to [**AWSTargetConfig**](AWSTargetConfig.md) |  | [optional] 
**AzureTargetConfig** | Pointer to [**AzureTargetConfig**](AzureTargetConfig.md) |  | [optional] 

## Methods

### NewReplicationConfigAllOf

`func NewReplicationConfigAllOf(targetType string, ) *ReplicationConfigAllOf`

NewReplicationConfigAllOf instantiates a new ReplicationConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationConfigAllOfWithDefaults

`func NewReplicationConfigAllOfWithDefaults() *ReplicationConfigAllOf`

NewReplicationConfigAllOfWithDefaults instantiates a new ReplicationConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetType

`func (o *ReplicationConfigAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ReplicationConfigAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ReplicationConfigAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.


### GetRemoteTargetConfig

`func (o *ReplicationConfigAllOf) GetRemoteTargetConfig() RemoteTargetConfig`

GetRemoteTargetConfig returns the RemoteTargetConfig field if non-nil, zero value otherwise.

### GetRemoteTargetConfigOk

`func (o *ReplicationConfigAllOf) GetRemoteTargetConfigOk() (*RemoteTargetConfig, bool)`

GetRemoteTargetConfigOk returns a tuple with the RemoteTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargetConfig

`func (o *ReplicationConfigAllOf) SetRemoteTargetConfig(v RemoteTargetConfig)`

SetRemoteTargetConfig sets RemoteTargetConfig field to given value.

### HasRemoteTargetConfig

`func (o *ReplicationConfigAllOf) HasRemoteTargetConfig() bool`

HasRemoteTargetConfig returns a boolean if a field has been set.

### GetAwsTargetConfig

`func (o *ReplicationConfigAllOf) GetAwsTargetConfig() AWSTargetConfig`

GetAwsTargetConfig returns the AwsTargetConfig field if non-nil, zero value otherwise.

### GetAwsTargetConfigOk

`func (o *ReplicationConfigAllOf) GetAwsTargetConfigOk() (*AWSTargetConfig, bool)`

GetAwsTargetConfigOk returns a tuple with the AwsTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsTargetConfig

`func (o *ReplicationConfigAllOf) SetAwsTargetConfig(v AWSTargetConfig)`

SetAwsTargetConfig sets AwsTargetConfig field to given value.

### HasAwsTargetConfig

`func (o *ReplicationConfigAllOf) HasAwsTargetConfig() bool`

HasAwsTargetConfig returns a boolean if a field has been set.

### GetAzureTargetConfig

`func (o *ReplicationConfigAllOf) GetAzureTargetConfig() AzureTargetConfig`

GetAzureTargetConfig returns the AzureTargetConfig field if non-nil, zero value otherwise.

### GetAzureTargetConfigOk

`func (o *ReplicationConfigAllOf) GetAzureTargetConfigOk() (*AzureTargetConfig, bool)`

GetAzureTargetConfigOk returns a tuple with the AzureTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTargetConfig

`func (o *ReplicationConfigAllOf) SetAzureTargetConfig(v AzureTargetConfig)`

SetAzureTargetConfig sets AzureTargetConfig field to given value.

### HasAzureTargetConfig

`func (o *ReplicationConfigAllOf) HasAzureTargetConfig() bool`

HasAzureTargetConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


