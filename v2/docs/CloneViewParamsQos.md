# CloneViewParamsQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Specifies the name of the QoS Policy. If not specified, the default is &#39;BackupTargetLow&#39;.  BackupTargetAuto: (Applicable only for C6K Platform) Use this policy for workloads such as backups, which keep many I/Os outstanding. This policy splits I/Os across SSDs and HDDs to achieve maximum performance based on the current usage. The priority for processing workload with this policy is the same as Backup Target High and Backup Target SSD.  JournaledSequentialDump: Use this policy for workloads that write large amounts of data sequentially to a very small number of files and do not keep many I/Os outstanding. By default data is written to the SSD and has the highest priority and low latency.  TestAndDevHigh: Use this policy for workloads that require lower I/O latency or do not keep many I/Os outstanding, as the I/Os are given higher priority compared to other QoS policies. Data is written to the SSD.  TestAndDevLow: The same as TestAndDev High, except that the I/Os with this QoS policy are given lower priority compared to I/Os with TestAndDev High when there is contention.  BackupTargetCommvault: Use this policy to intelligently detect and exclude application-specific markers to achieve better deduplication when CommVault backup application is writing to a Cohesity View. Data is written to SSD and has the same priority and latency as TestAndDev High.  BackupTargetSSD: Use this policy for workloads such as backups, which keep many I/Os outstanding, but in this case, DataPlatform sends both sequential and random I/Os to SSD. The latency is lower than other Backup Target policies. The priority for processing workload with this policy is the same as Backup Target Auto.  BackupTargetHigh: Use this policy for non-latency sensitive workloads such as backups, which keep many I/Os outstanding. Data is written to HDD and has higher latency compared to other QoS policies writing to a SSD The priority for processing workload with this policy is the same as Backup Target Auto.  BackupTargetLow: The same as Backup Target High, except that the priority for processing workloads with this policy is lower than workloads with Backup Target Auto / High /SSD when there is contention. | [optional] 

## Methods

### NewCloneViewParamsQos

`func NewCloneViewParamsQos() *CloneViewParamsQos`

NewCloneViewParamsQos instantiates a new CloneViewParamsQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneViewParamsQosWithDefaults

`func NewCloneViewParamsQosWithDefaults() *CloneViewParamsQos`

NewCloneViewParamsQosWithDefaults instantiates a new CloneViewParamsQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneViewParamsQos) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneViewParamsQos) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneViewParamsQos) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloneViewParamsQos) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CloneViewParamsQos) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CloneViewParamsQos) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


