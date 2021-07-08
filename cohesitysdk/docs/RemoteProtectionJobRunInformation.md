# RemoteProtectionJobRunInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **NullableString** | Specifies the name of the original Cluster that archived the data to the Vault. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type (such as kVMware or kSQL) of the original archived Protection Job. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**JobName** | Pointer to **NullableString** | Specifies the name of the Protection Job on the original Cluster. | [optional] 
**JobUid** | Pointer to [**NullableUniversalId**](UniversalId.md) | Specifies the globally unique id of the original Protection Job that archived the data to the Vault. This id is assigned by the original Cluster that archived the data. | [optional] 
**ProtectionJobRuns** | Pointer to [**[]RemoteProtectionJobRunInstance**](RemoteProtectionJobRunInstance.md) | Array of Protection Job Run Details.  Specifies the list of Protection Job Runs (Snapshot) details for a Protection Job archived to a Vault. | [optional] 

## Methods

### NewRemoteProtectionJobRunInformation

`func NewRemoteProtectionJobRunInformation() *RemoteProtectionJobRunInformation`

NewRemoteProtectionJobRunInformation instantiates a new RemoteProtectionJobRunInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProtectionJobRunInformationWithDefaults

`func NewRemoteProtectionJobRunInformationWithDefaults() *RemoteProtectionJobRunInformation`

NewRemoteProtectionJobRunInformationWithDefaults instantiates a new RemoteProtectionJobRunInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *RemoteProtectionJobRunInformation) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RemoteProtectionJobRunInformation) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RemoteProtectionJobRunInformation) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RemoteProtectionJobRunInformation) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *RemoteProtectionJobRunInformation) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *RemoteProtectionJobRunInformation) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetEnvironment

`func (o *RemoteProtectionJobRunInformation) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RemoteProtectionJobRunInformation) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RemoteProtectionJobRunInformation) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RemoteProtectionJobRunInformation) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RemoteProtectionJobRunInformation) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RemoteProtectionJobRunInformation) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetJobName

`func (o *RemoteProtectionJobRunInformation) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *RemoteProtectionJobRunInformation) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *RemoteProtectionJobRunInformation) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *RemoteProtectionJobRunInformation) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### SetJobNameNil

`func (o *RemoteProtectionJobRunInformation) SetJobNameNil(b bool)`

 SetJobNameNil sets the value for JobName to be an explicit nil

### UnsetJobName
`func (o *RemoteProtectionJobRunInformation) UnsetJobName()`

UnsetJobName ensures that no value is present for JobName, not even an explicit nil
### GetJobUid

`func (o *RemoteProtectionJobRunInformation) GetJobUid() UniversalId`

GetJobUid returns the JobUid field if non-nil, zero value otherwise.

### GetJobUidOk

`func (o *RemoteProtectionJobRunInformation) GetJobUidOk() (*UniversalId, bool)`

GetJobUidOk returns a tuple with the JobUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobUid

`func (o *RemoteProtectionJobRunInformation) SetJobUid(v UniversalId)`

SetJobUid sets JobUid field to given value.

### HasJobUid

`func (o *RemoteProtectionJobRunInformation) HasJobUid() bool`

HasJobUid returns a boolean if a field has been set.

### SetJobUidNil

`func (o *RemoteProtectionJobRunInformation) SetJobUidNil(b bool)`

 SetJobUidNil sets the value for JobUid to be an explicit nil

### UnsetJobUid
`func (o *RemoteProtectionJobRunInformation) UnsetJobUid()`

UnsetJobUid ensures that no value is present for JobUid, not even an explicit nil
### GetProtectionJobRuns

`func (o *RemoteProtectionJobRunInformation) GetProtectionJobRuns() []RemoteProtectionJobRunInstance`

GetProtectionJobRuns returns the ProtectionJobRuns field if non-nil, zero value otherwise.

### GetProtectionJobRunsOk

`func (o *RemoteProtectionJobRunInformation) GetProtectionJobRunsOk() (*[]RemoteProtectionJobRunInstance, bool)`

GetProtectionJobRunsOk returns a tuple with the ProtectionJobRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionJobRuns

`func (o *RemoteProtectionJobRunInformation) SetProtectionJobRuns(v []RemoteProtectionJobRunInstance)`

SetProtectionJobRuns sets ProtectionJobRuns field to given value.

### HasProtectionJobRuns

`func (o *RemoteProtectionJobRunInformation) HasProtectionJobRuns() bool`

HasProtectionJobRuns returns a boolean if a field has been set.

### SetProtectionJobRunsNil

`func (o *RemoteProtectionJobRunInformation) SetProtectionJobRunsNil(b bool)`

 SetProtectionJobRunsNil sets the value for ProtectionJobRuns to be an explicit nil

### UnsetProtectionJobRuns
`func (o *RemoteProtectionJobRunInformation) UnsetProtectionJobRuns()`

UnsetProtectionJobRuns ensures that no value is present for ProtectionJobRuns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


