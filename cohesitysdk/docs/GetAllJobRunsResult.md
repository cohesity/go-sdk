# GetAllJobRunsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimeMsecs** | Pointer to **NullableInt64** | Specifies the end time of the run. | [optional] 
**EnvType** | Pointer to **NullableString** | Specifies the environment type of the job. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**JobId** | Pointer to **NullableString** | Specifies the job id. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the job name. | [optional] 
**JobRunId** | Pointer to **NullableString** | Specifies the job run id. | [optional] 
**JobType** | Pointer to **NullableString** | Specifies the job type, protection, replication, archival, apollo, indexing etc. | [optional] 
**StartTimeMsecs** | Pointer to **NullableInt64** | Specifies the start time of the run. | [optional] 
**ViewBoxId** | Pointer to **NullableInt64** | Specifies the view box id. | [optional] 

## Methods

### NewGetAllJobRunsResult

`func NewGetAllJobRunsResult() *GetAllJobRunsResult`

NewGetAllJobRunsResult instantiates a new GetAllJobRunsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllJobRunsResultWithDefaults

`func NewGetAllJobRunsResultWithDefaults() *GetAllJobRunsResult`

NewGetAllJobRunsResultWithDefaults instantiates a new GetAllJobRunsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimeMsecs

`func (o *GetAllJobRunsResult) GetEndTimeMsecs() int64`

GetEndTimeMsecs returns the EndTimeMsecs field if non-nil, zero value otherwise.

### GetEndTimeMsecsOk

`func (o *GetAllJobRunsResult) GetEndTimeMsecsOk() (*int64, bool)`

GetEndTimeMsecsOk returns a tuple with the EndTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeMsecs

`func (o *GetAllJobRunsResult) SetEndTimeMsecs(v int64)`

SetEndTimeMsecs sets EndTimeMsecs field to given value.

### HasEndTimeMsecs

`func (o *GetAllJobRunsResult) HasEndTimeMsecs() bool`

HasEndTimeMsecs returns a boolean if a field has been set.

### SetEndTimeMsecsNil

`func (o *GetAllJobRunsResult) SetEndTimeMsecsNil(b bool)`

 SetEndTimeMsecsNil sets the value for EndTimeMsecs to be an explicit nil

### UnsetEndTimeMsecs
`func (o *GetAllJobRunsResult) UnsetEndTimeMsecs()`

UnsetEndTimeMsecs ensures that no value is present for EndTimeMsecs, not even an explicit nil
### GetEnvType

`func (o *GetAllJobRunsResult) GetEnvType() string`

GetEnvType returns the EnvType field if non-nil, zero value otherwise.

### GetEnvTypeOk

`func (o *GetAllJobRunsResult) GetEnvTypeOk() (*string, bool)`

GetEnvTypeOk returns a tuple with the EnvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvType

`func (o *GetAllJobRunsResult) SetEnvType(v string)`

SetEnvType sets EnvType field to given value.

### HasEnvType

`func (o *GetAllJobRunsResult) HasEnvType() bool`

HasEnvType returns a boolean if a field has been set.

### SetEnvTypeNil

`func (o *GetAllJobRunsResult) SetEnvTypeNil(b bool)`

 SetEnvTypeNil sets the value for EnvType to be an explicit nil

### UnsetEnvType
`func (o *GetAllJobRunsResult) UnsetEnvType()`

UnsetEnvType ensures that no value is present for EnvType, not even an explicit nil
### GetJobId

`func (o *GetAllJobRunsResult) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GetAllJobRunsResult) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GetAllJobRunsResult) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *GetAllJobRunsResult) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *GetAllJobRunsResult) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *GetAllJobRunsResult) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetJobName

`func (o *GetAllJobRunsResult) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *GetAllJobRunsResult) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *GetAllJobRunsResult) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *GetAllJobRunsResult) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *GetAllJobRunsResult) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *GetAllJobRunsResult) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetJobRunId

`func (o *GetAllJobRunsResult) GetJobRunId() string`

GetJobRunId returns the JobRunId field if non-nil, zero value otherwise.

### GetJobRunIdOk

`func (o *GetAllJobRunsResult) GetJobRunIdOk() (*string, bool)`

GetJobRunIdOk returns a tuple with the JobRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunId

`func (o *GetAllJobRunsResult) SetJobRunId(v string)`

SetJobRunId sets JobRunId field to given value.

### HasJobRunId

`func (o *GetAllJobRunsResult) HasJobRunId() bool`

HasJobRunId returns a boolean if a field has been set.

### SetJobRunIdNil

`func (o *GetAllJobRunsResult) SetJobRunIdNil(b bool)`

 SetJobRunIdNil sets the value for JobRunId to be an explicit nil

### UnsetJobRunId
`func (o *GetAllJobRunsResult) UnsetJobRunId()`

UnsetJobRunId ensures that no value is present for JobRunId, not even an explicit nil
### GetJobType

`func (o *GetAllJobRunsResult) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *GetAllJobRunsResult) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *GetAllJobRunsResult) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *GetAllJobRunsResult) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### SetJobTypeNil

`func (o *GetAllJobRunsResult) SetJobTypeNil(b bool)`

 SetJobTypeNil sets the value for JobType to be an explicit nil

### UnsetJobType
`func (o *GetAllJobRunsResult) UnsetJobType()`

UnsetJobType ensures that no value is present for JobType, not even an explicit nil
### GetStartTimeMsecs

`func (o *GetAllJobRunsResult) GetStartTimeMsecs() int64`

GetStartTimeMsecs returns the StartTimeMsecs field if non-nil, zero value otherwise.

### GetStartTimeMsecsOk

`func (o *GetAllJobRunsResult) GetStartTimeMsecsOk() (*int64, bool)`

GetStartTimeMsecsOk returns a tuple with the StartTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeMsecs

`func (o *GetAllJobRunsResult) SetStartTimeMsecs(v int64)`

SetStartTimeMsecs sets StartTimeMsecs field to given value.

### HasStartTimeMsecs

`func (o *GetAllJobRunsResult) HasStartTimeMsecs() bool`

HasStartTimeMsecs returns a boolean if a field has been set.

### SetStartTimeMsecsNil

`func (o *GetAllJobRunsResult) SetStartTimeMsecsNil(b bool)`

 SetStartTimeMsecsNil sets the value for StartTimeMsecs to be an explicit nil

### UnsetStartTimeMsecs
`func (o *GetAllJobRunsResult) UnsetStartTimeMsecs()`

UnsetStartTimeMsecs ensures that no value is present for StartTimeMsecs, not even an explicit nil
### GetViewBoxId

`func (o *GetAllJobRunsResult) GetViewBoxId() int64`

GetViewBoxId returns the ViewBoxId field if non-nil, zero value otherwise.

### GetViewBoxIdOk

`func (o *GetAllJobRunsResult) GetViewBoxIdOk() (*int64, bool)`

GetViewBoxIdOk returns a tuple with the ViewBoxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewBoxId

`func (o *GetAllJobRunsResult) SetViewBoxId(v int64)`

SetViewBoxId sets ViewBoxId field to given value.

### HasViewBoxId

`func (o *GetAllJobRunsResult) HasViewBoxId() bool`

HasViewBoxId returns a boolean if a field has been set.

### SetViewBoxIdNil

`func (o *GetAllJobRunsResult) SetViewBoxIdNil(b bool)`

 SetViewBoxIdNil sets the value for ViewBoxId to be an explicit nil

### UnsetViewBoxId
`func (o *GetAllJobRunsResult) UnsetViewBoxId()`

UnsetViewBoxId ensures that no value is present for ViewBoxId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


