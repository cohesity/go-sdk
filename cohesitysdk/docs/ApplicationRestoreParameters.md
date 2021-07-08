# ApplicationRestoreParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationEnvironment** | Pointer to **NullableString** | Specifies the Environment of the Application server to restore like &#39;kSQL&#39;,&#39;kAD&#39;,or &#39;kExchange&#39;. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**ApplicationRestoreObjects** | Pointer to [**[]ApplicationRestoreObject**](ApplicationRestoreObject.md) | Specifies the Application Server objects whose data should be restored. | [optional] 
**HostingProtectionSource** | Pointer to [**RestoreObjectDetails**](RestoreObjectDetails.md) |  | [optional] 
**ProtectionSourceAndApplicationObjects** | Pointer to [**[]ProtectionSourceAndApplicationRestoreObjects**](ProtectionSourceAndApplicationRestoreObjects.md) | Specifies the list of hosting protection source and Application restore objects tuple. | [optional] 

## Methods

### NewApplicationRestoreParameters

`func NewApplicationRestoreParameters() *ApplicationRestoreParameters`

NewApplicationRestoreParameters instantiates a new ApplicationRestoreParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRestoreParametersWithDefaults

`func NewApplicationRestoreParametersWithDefaults() *ApplicationRestoreParameters`

NewApplicationRestoreParametersWithDefaults instantiates a new ApplicationRestoreParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationEnvironment

`func (o *ApplicationRestoreParameters) GetApplicationEnvironment() string`

GetApplicationEnvironment returns the ApplicationEnvironment field if non-nil, zero value otherwise.

### GetApplicationEnvironmentOk

`func (o *ApplicationRestoreParameters) GetApplicationEnvironmentOk() (*string, bool)`

GetApplicationEnvironmentOk returns a tuple with the ApplicationEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEnvironment

`func (o *ApplicationRestoreParameters) SetApplicationEnvironment(v string)`

SetApplicationEnvironment sets ApplicationEnvironment field to given value.

### HasApplicationEnvironment

`func (o *ApplicationRestoreParameters) HasApplicationEnvironment() bool`

HasApplicationEnvironment returns a boolean if a field has been set.

### SetApplicationEnvironmentNil

`func (o *ApplicationRestoreParameters) SetApplicationEnvironmentNil(b bool)`

 SetApplicationEnvironmentNil sets the value for ApplicationEnvironment to be an explicit nil

### UnsetApplicationEnvironment
`func (o *ApplicationRestoreParameters) UnsetApplicationEnvironment()`

UnsetApplicationEnvironment ensures that no value is present for ApplicationEnvironment, not even an explicit nil
### GetApplicationRestoreObjects

`func (o *ApplicationRestoreParameters) GetApplicationRestoreObjects() []ApplicationRestoreObject`

GetApplicationRestoreObjects returns the ApplicationRestoreObjects field if non-nil, zero value otherwise.

### GetApplicationRestoreObjectsOk

`func (o *ApplicationRestoreParameters) GetApplicationRestoreObjectsOk() (*[]ApplicationRestoreObject, bool)`

GetApplicationRestoreObjectsOk returns a tuple with the ApplicationRestoreObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationRestoreObjects

`func (o *ApplicationRestoreParameters) SetApplicationRestoreObjects(v []ApplicationRestoreObject)`

SetApplicationRestoreObjects sets ApplicationRestoreObjects field to given value.

### HasApplicationRestoreObjects

`func (o *ApplicationRestoreParameters) HasApplicationRestoreObjects() bool`

HasApplicationRestoreObjects returns a boolean if a field has been set.

### SetApplicationRestoreObjectsNil

`func (o *ApplicationRestoreParameters) SetApplicationRestoreObjectsNil(b bool)`

 SetApplicationRestoreObjectsNil sets the value for ApplicationRestoreObjects to be an explicit nil

### UnsetApplicationRestoreObjects
`func (o *ApplicationRestoreParameters) UnsetApplicationRestoreObjects()`

UnsetApplicationRestoreObjects ensures that no value is present for ApplicationRestoreObjects, not even an explicit nil
### GetHostingProtectionSource

`func (o *ApplicationRestoreParameters) GetHostingProtectionSource() RestoreObjectDetails`

GetHostingProtectionSource returns the HostingProtectionSource field if non-nil, zero value otherwise.

### GetHostingProtectionSourceOk

`func (o *ApplicationRestoreParameters) GetHostingProtectionSourceOk() (*RestoreObjectDetails, bool)`

GetHostingProtectionSourceOk returns a tuple with the HostingProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingProtectionSource

`func (o *ApplicationRestoreParameters) SetHostingProtectionSource(v RestoreObjectDetails)`

SetHostingProtectionSource sets HostingProtectionSource field to given value.

### HasHostingProtectionSource

`func (o *ApplicationRestoreParameters) HasHostingProtectionSource() bool`

HasHostingProtectionSource returns a boolean if a field has been set.

### GetProtectionSourceAndApplicationObjects

`func (o *ApplicationRestoreParameters) GetProtectionSourceAndApplicationObjects() []ProtectionSourceAndApplicationRestoreObjects`

GetProtectionSourceAndApplicationObjects returns the ProtectionSourceAndApplicationObjects field if non-nil, zero value otherwise.

### GetProtectionSourceAndApplicationObjectsOk

`func (o *ApplicationRestoreParameters) GetProtectionSourceAndApplicationObjectsOk() (*[]ProtectionSourceAndApplicationRestoreObjects, bool)`

GetProtectionSourceAndApplicationObjectsOk returns a tuple with the ProtectionSourceAndApplicationObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionSourceAndApplicationObjects

`func (o *ApplicationRestoreParameters) SetProtectionSourceAndApplicationObjects(v []ProtectionSourceAndApplicationRestoreObjects)`

SetProtectionSourceAndApplicationObjects sets ProtectionSourceAndApplicationObjects field to given value.

### HasProtectionSourceAndApplicationObjects

`func (o *ApplicationRestoreParameters) HasProtectionSourceAndApplicationObjects() bool`

HasProtectionSourceAndApplicationObjects returns a boolean if a field has been set.

### SetProtectionSourceAndApplicationObjectsNil

`func (o *ApplicationRestoreParameters) SetProtectionSourceAndApplicationObjectsNil(b bool)`

 SetProtectionSourceAndApplicationObjectsNil sets the value for ProtectionSourceAndApplicationObjects to be an explicit nil

### UnsetProtectionSourceAndApplicationObjects
`func (o *ApplicationRestoreParameters) UnsetProtectionSourceAndApplicationObjects()`

UnsetProtectionSourceAndApplicationObjects ensures that no value is present for ProtectionSourceAndApplicationObjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


