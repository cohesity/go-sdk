# SnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to **NullableString** | Specifies the environment type (such as kVMware or kSQL) that contains the source to backup. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**RelativeSnapshotDirectory** | Pointer to **NullableString** | Specifies the relative directory path from root path where the snapshot is stored. | [optional] 
**RootPath** | Pointer to **NullableString** | Specifies the root path where the snapshot is stored, using the following format: \&quot;/ViewBox/ViewName/fs\&quot;. | [optional] 
**SourceSnapshotCreateTimeUsecs** | Pointer to **NullableInt64** | Specifies the snapshot create time of the already created snapshot on the source | [optional] 
**SourceSnapshotName** | Pointer to **NullableString** | Specifies the name of the snapshot backed up in a Netapp Data-Protect Volume where we use already created snapshot on the source | [optional] 
**ViewName** | Pointer to **NullableString** | Specifies the name of the View that is cloned. NOTE: This field is only populated for View cloning. | [optional] 

## Methods

### NewSnapshotInfo

`func NewSnapshotInfo() *SnapshotInfo`

NewSnapshotInfo instantiates a new SnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotInfoWithDefaults

`func NewSnapshotInfoWithDefaults() *SnapshotInfo`

NewSnapshotInfoWithDefaults instantiates a new SnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *SnapshotInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SnapshotInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SnapshotInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SnapshotInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *SnapshotInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *SnapshotInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetRelativeSnapshotDirectory

`func (o *SnapshotInfo) GetRelativeSnapshotDirectory() string`

GetRelativeSnapshotDirectory returns the RelativeSnapshotDirectory field if non-nil, zero value otherwise.

### GetRelativeSnapshotDirectoryOk

`func (o *SnapshotInfo) GetRelativeSnapshotDirectoryOk() (*string, bool)`

GetRelativeSnapshotDirectoryOk returns a tuple with the RelativeSnapshotDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeSnapshotDirectory

`func (o *SnapshotInfo) SetRelativeSnapshotDirectory(v string)`

SetRelativeSnapshotDirectory sets RelativeSnapshotDirectory field to given value.

### HasRelativeSnapshotDirectory

`func (o *SnapshotInfo) HasRelativeSnapshotDirectory() bool`

HasRelativeSnapshotDirectory returns a boolean if a field has been set.

### SetRelativeSnapshotDirectoryNil

`func (o *SnapshotInfo) SetRelativeSnapshotDirectoryNil(b bool)`

 SetRelativeSnapshotDirectoryNil sets the value for RelativeSnapshotDirectory to be an explicit nil

### UnsetRelativeSnapshotDirectory
`func (o *SnapshotInfo) UnsetRelativeSnapshotDirectory()`

UnsetRelativeSnapshotDirectory ensures that no value is present for RelativeSnapshotDirectory, not even an explicit nil
### GetRootPath

`func (o *SnapshotInfo) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *SnapshotInfo) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *SnapshotInfo) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *SnapshotInfo) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### SetRootPathNil

`func (o *SnapshotInfo) SetRootPathNil(b bool)`

 SetRootPathNil sets the value for RootPath to be an explicit nil

### UnsetRootPath
`func (o *SnapshotInfo) UnsetRootPath()`

UnsetRootPath ensures that no value is present for RootPath, not even an explicit nil
### GetSourceSnapshotCreateTimeUsecs

`func (o *SnapshotInfo) GetSourceSnapshotCreateTimeUsecs() int64`

GetSourceSnapshotCreateTimeUsecs returns the SourceSnapshotCreateTimeUsecs field if non-nil, zero value otherwise.

### GetSourceSnapshotCreateTimeUsecsOk

`func (o *SnapshotInfo) GetSourceSnapshotCreateTimeUsecsOk() (*int64, bool)`

GetSourceSnapshotCreateTimeUsecsOk returns a tuple with the SourceSnapshotCreateTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotCreateTimeUsecs

`func (o *SnapshotInfo) SetSourceSnapshotCreateTimeUsecs(v int64)`

SetSourceSnapshotCreateTimeUsecs sets SourceSnapshotCreateTimeUsecs field to given value.

### HasSourceSnapshotCreateTimeUsecs

`func (o *SnapshotInfo) HasSourceSnapshotCreateTimeUsecs() bool`

HasSourceSnapshotCreateTimeUsecs returns a boolean if a field has been set.

### SetSourceSnapshotCreateTimeUsecsNil

`func (o *SnapshotInfo) SetSourceSnapshotCreateTimeUsecsNil(b bool)`

 SetSourceSnapshotCreateTimeUsecsNil sets the value for SourceSnapshotCreateTimeUsecs to be an explicit nil

### UnsetSourceSnapshotCreateTimeUsecs
`func (o *SnapshotInfo) UnsetSourceSnapshotCreateTimeUsecs()`

UnsetSourceSnapshotCreateTimeUsecs ensures that no value is present for SourceSnapshotCreateTimeUsecs, not even an explicit nil
### GetSourceSnapshotName

`func (o *SnapshotInfo) GetSourceSnapshotName() string`

GetSourceSnapshotName returns the SourceSnapshotName field if non-nil, zero value otherwise.

### GetSourceSnapshotNameOk

`func (o *SnapshotInfo) GetSourceSnapshotNameOk() (*string, bool)`

GetSourceSnapshotNameOk returns a tuple with the SourceSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotName

`func (o *SnapshotInfo) SetSourceSnapshotName(v string)`

SetSourceSnapshotName sets SourceSnapshotName field to given value.

### HasSourceSnapshotName

`func (o *SnapshotInfo) HasSourceSnapshotName() bool`

HasSourceSnapshotName returns a boolean if a field has been set.

### SetSourceSnapshotNameNil

`func (o *SnapshotInfo) SetSourceSnapshotNameNil(b bool)`

 SetSourceSnapshotNameNil sets the value for SourceSnapshotName to be an explicit nil

### UnsetSourceSnapshotName
`func (o *SnapshotInfo) UnsetSourceSnapshotName()`

UnsetSourceSnapshotName ensures that no value is present for SourceSnapshotName, not even an explicit nil
### GetViewName

`func (o *SnapshotInfo) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SnapshotInfo) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SnapshotInfo) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SnapshotInfo) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *SnapshotInfo) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *SnapshotInfo) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


