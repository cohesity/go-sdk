# DataMigrationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysToKeep** | Pointer to **NullableInt64** | Specifies how many days to retain Snapshots on the Cohesity Cluster. | [optional] 
**SchedulingPolicy** | Pointer to [**SchedulingPolicy**](SchedulingPolicy.md) |  | [optional] 
**WormRetentionType** | Pointer to **NullableString** | Specifies WORM retention type for the files. When a WORM retention type is specified, the files will be kept until the maximum of the retention time. During that time, the files cannot be deleted. &#39;kNone&#39; implies there is no WORM retention set. &#39;kCompliance&#39; implies WORM retention is set for compliance reason. &#39;kAdministrative&#39; implies WORM retention is set for administrative purposes. | [optional] 

## Methods

### NewDataMigrationPolicy

`func NewDataMigrationPolicy() *DataMigrationPolicy`

NewDataMigrationPolicy instantiates a new DataMigrationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMigrationPolicyWithDefaults

`func NewDataMigrationPolicyWithDefaults() *DataMigrationPolicy`

NewDataMigrationPolicyWithDefaults instantiates a new DataMigrationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysToKeep

`func (o *DataMigrationPolicy) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *DataMigrationPolicy) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *DataMigrationPolicy) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *DataMigrationPolicy) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *DataMigrationPolicy) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *DataMigrationPolicy) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetSchedulingPolicy

`func (o *DataMigrationPolicy) GetSchedulingPolicy() SchedulingPolicy`

GetSchedulingPolicy returns the SchedulingPolicy field if non-nil, zero value otherwise.

### GetSchedulingPolicyOk

`func (o *DataMigrationPolicy) GetSchedulingPolicyOk() (*SchedulingPolicy, bool)`

GetSchedulingPolicyOk returns a tuple with the SchedulingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingPolicy

`func (o *DataMigrationPolicy) SetSchedulingPolicy(v SchedulingPolicy)`

SetSchedulingPolicy sets SchedulingPolicy field to given value.

### HasSchedulingPolicy

`func (o *DataMigrationPolicy) HasSchedulingPolicy() bool`

HasSchedulingPolicy returns a boolean if a field has been set.

### GetWormRetentionType

`func (o *DataMigrationPolicy) GetWormRetentionType() string`

GetWormRetentionType returns the WormRetentionType field if non-nil, zero value otherwise.

### GetWormRetentionTypeOk

`func (o *DataMigrationPolicy) GetWormRetentionTypeOk() (*string, bool)`

GetWormRetentionTypeOk returns a tuple with the WormRetentionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormRetentionType

`func (o *DataMigrationPolicy) SetWormRetentionType(v string)`

SetWormRetentionType sets WormRetentionType field to given value.

### HasWormRetentionType

`func (o *DataMigrationPolicy) HasWormRetentionType() bool`

HasWormRetentionType returns a boolean if a field has been set.

### SetWormRetentionTypeNil

`func (o *DataMigrationPolicy) SetWormRetentionTypeNil(b bool)`

 SetWormRetentionTypeNil sets the value for WormRetentionType to be an explicit nil

### UnsetWormRetentionType
`func (o *DataMigrationPolicy) UnsetWormRetentionType()`

UnsetWormRetentionType ensures that no value is present for WormRetentionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


