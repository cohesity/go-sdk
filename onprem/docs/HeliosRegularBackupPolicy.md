# HeliosRegularBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incremental** | Pointer to [**HeliosIncrementalBackupPolicy**](HeliosIncrementalBackupPolicy.md) |  | [optional] 
**Full** | Pointer to [**HeliosFullBackupPolicy**](HeliosFullBackupPolicy.md) |  | [optional] 
**Retention** | Pointer to [**HeliosRetention**](HeliosRetention.md) |  | [optional] 

## Methods

### NewHeliosRegularBackupPolicy

`func NewHeliosRegularBackupPolicy() *HeliosRegularBackupPolicy`

NewHeliosRegularBackupPolicy instantiates a new HeliosRegularBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosRegularBackupPolicyWithDefaults

`func NewHeliosRegularBackupPolicyWithDefaults() *HeliosRegularBackupPolicy`

NewHeliosRegularBackupPolicyWithDefaults instantiates a new HeliosRegularBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncremental

`func (o *HeliosRegularBackupPolicy) GetIncremental() HeliosIncrementalBackupPolicy`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *HeliosRegularBackupPolicy) GetIncrementalOk() (*HeliosIncrementalBackupPolicy, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *HeliosRegularBackupPolicy) SetIncremental(v HeliosIncrementalBackupPolicy)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *HeliosRegularBackupPolicy) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.

### GetFull

`func (o *HeliosRegularBackupPolicy) GetFull() HeliosFullBackupPolicy`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *HeliosRegularBackupPolicy) GetFullOk() (*HeliosFullBackupPolicy, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *HeliosRegularBackupPolicy) SetFull(v HeliosFullBackupPolicy)`

SetFull sets Full field to given value.

### HasFull

`func (o *HeliosRegularBackupPolicy) HasFull() bool`

HasFull returns a boolean if a field has been set.

### GetRetention

`func (o *HeliosRegularBackupPolicy) GetRetention() HeliosRetention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *HeliosRegularBackupPolicy) GetRetentionOk() (*HeliosRetention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *HeliosRegularBackupPolicy) SetRetention(v HeliosRetention)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *HeliosRegularBackupPolicy) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


