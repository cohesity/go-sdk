# CdpBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retention** | [**CdpRetention**](CdpRetention.md) |  | 

## Methods

### NewCdpBackupPolicy

`func NewCdpBackupPolicy(retention CdpRetention, ) *CdpBackupPolicy`

NewCdpBackupPolicy instantiates a new CdpBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdpBackupPolicyWithDefaults

`func NewCdpBackupPolicyWithDefaults() *CdpBackupPolicy`

NewCdpBackupPolicyWithDefaults instantiates a new CdpBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetention

`func (o *CdpBackupPolicy) GetRetention() CdpRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *CdpBackupPolicy) GetRetentionOk() (*CdpRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *CdpBackupPolicy) SetRetention(v CdpRetention)`

SetRetention sets Retention field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


