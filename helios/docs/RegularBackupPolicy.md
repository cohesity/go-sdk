# RegularBackupPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incremental** | Pointer to [**IncrementalBackupPolicy**](IncrementalBackupPolicy.md) |  | [optional] 
**Full** | Pointer to [**FullBackupPolicy**](FullBackupPolicy.md) |  | [optional] 
**Retention** | [**Retention**](Retention.md) |  | 
**PrimaryBackupTarget** | Pointer to [**PrimaryBackupTarget**](PrimaryBackupTarget.md) |  | [optional] 

## Methods

### NewRegularBackupPolicy

`func NewRegularBackupPolicy(retention Retention, ) *RegularBackupPolicy`

NewRegularBackupPolicy instantiates a new RegularBackupPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegularBackupPolicyWithDefaults

`func NewRegularBackupPolicyWithDefaults() *RegularBackupPolicy`

NewRegularBackupPolicyWithDefaults instantiates a new RegularBackupPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncremental

`func (o *RegularBackupPolicy) GetIncremental() IncrementalBackupPolicy`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *RegularBackupPolicy) GetIncrementalOk() (*IncrementalBackupPolicy, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *RegularBackupPolicy) SetIncremental(v IncrementalBackupPolicy)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *RegularBackupPolicy) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.

### GetFull

`func (o *RegularBackupPolicy) GetFull() FullBackupPolicy`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *RegularBackupPolicy) GetFullOk() (*FullBackupPolicy, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *RegularBackupPolicy) SetFull(v FullBackupPolicy)`

SetFull sets Full field to given value.

### HasFull

`func (o *RegularBackupPolicy) HasFull() bool`

HasFull returns a boolean if a field has been set.

### GetRetention

`func (o *RegularBackupPolicy) GetRetention() Retention`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *RegularBackupPolicy) GetRetentionOk() (*Retention, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *RegularBackupPolicy) SetRetention(v Retention)`

SetRetention sets Retention field to given value.


### GetPrimaryBackupTarget

`func (o *RegularBackupPolicy) GetPrimaryBackupTarget() PrimaryBackupTarget`

GetPrimaryBackupTarget returns the PrimaryBackupTarget field if non-nil, zero value otherwise.

### GetPrimaryBackupTargetOk

`func (o *RegularBackupPolicy) GetPrimaryBackupTargetOk() (*PrimaryBackupTarget, bool)`

GetPrimaryBackupTargetOk returns a tuple with the PrimaryBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBackupTarget

`func (o *RegularBackupPolicy) SetPrimaryBackupTarget(v PrimaryBackupTarget)`

SetPrimaryBackupTarget sets PrimaryBackupTarget field to given value.

### HasPrimaryBackupTarget

`func (o *RegularBackupPolicy) HasPrimaryBackupTarget() bool`

HasPrimaryBackupTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


