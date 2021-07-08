# ApplicationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationTreeInfo** | Pointer to [**[]ProtectionSourceNode**](ProtectionSourceNode.md) | Application Server and the subtrees below them.  Specifies the application subtree used to store additional application level Objects. Different environments use the subtree to store application level information. For example for SQL Server, this subtree stores the SQL Server instances running on a VM or a Physical Server. overrideDescription: true | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment type of the application such as &#39;kSQL&#39;, &#39;kExchange&#39; registered on the Protection Source. overrideDescription: true Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 

## Methods

### NewApplicationInfo

`func NewApplicationInfo() *ApplicationInfo`

NewApplicationInfo instantiates a new ApplicationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInfoWithDefaults

`func NewApplicationInfoWithDefaults() *ApplicationInfo`

NewApplicationInfoWithDefaults instantiates a new ApplicationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationTreeInfo

`func (o *ApplicationInfo) GetApplicationTreeInfo() []ProtectionSourceNode`

GetApplicationTreeInfo returns the ApplicationTreeInfo field if non-nil, zero value otherwise.

### GetApplicationTreeInfoOk

`func (o *ApplicationInfo) GetApplicationTreeInfoOk() (*[]ProtectionSourceNode, bool)`

GetApplicationTreeInfoOk returns a tuple with the ApplicationTreeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTreeInfo

`func (o *ApplicationInfo) SetApplicationTreeInfo(v []ProtectionSourceNode)`

SetApplicationTreeInfo sets ApplicationTreeInfo field to given value.

### HasApplicationTreeInfo

`func (o *ApplicationInfo) HasApplicationTreeInfo() bool`

HasApplicationTreeInfo returns a boolean if a field has been set.

### SetApplicationTreeInfoNil

`func (o *ApplicationInfo) SetApplicationTreeInfoNil(b bool)`

 SetApplicationTreeInfoNil sets the value for ApplicationTreeInfo to be an explicit nil

### UnsetApplicationTreeInfo
`func (o *ApplicationInfo) UnsetApplicationTreeInfo()`

UnsetApplicationTreeInfo ensures that no value is present for ApplicationTreeInfo, not even an explicit nil
### GetEnvironment

`func (o *ApplicationInfo) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationInfo) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationInfo) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationInfo) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ApplicationInfo) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ApplicationInfo) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


