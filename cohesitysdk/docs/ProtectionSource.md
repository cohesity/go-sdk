# ProtectionSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcropolisProtectionSource** | Pointer to [**NullableAcropolisProtectionSource**](AcropolisProtectionSource.md) | Specifies details about an Acropolis Protection Source when the environment is set to &#39;kAcropolis&#39;. | [optional] 
**AdProtectionSource** | Pointer to [**NullableAdProtectionSource**](AdProtectionSource.md) | Specifies details about an AD Protection Source when the environment is set to &#39;kAD&#39;. | [optional] 
**AwsProtectionSource** | Pointer to [**NullableAwsProtectionSource**](AwsProtectionSource.md) | Specifies details about an AWS Protection Source when the environment is set to &#39;kAWS&#39;. | [optional] 
**AzureProtectionSource** | Pointer to [**NullableAzureProtectionSource**](AzureProtectionSource.md) | Specifies details about an Azure Protection Source when the environment is set to &#39;kAzure&#39;. | [optional] 
**CassandraProtectionSource** | Pointer to [**NullableCassandraProtectionSource**](CassandraProtectionSource.md) | Specifies details about a Cassandra Protection Source when the environment is set to &#39;kCassandra&#39;. | [optional] 
**CouchbaseProtectionSource** | Pointer to [**NullableCouchbaseProtectionSource**](CouchbaseProtectionSource.md) | Specifies details about a Couchbase Protection Source when the environment is set to &#39;kCouchbase&#39;. | [optional] 
**ElastifileProtectionSource** | Pointer to [**NullableElastifileProtectionSource**](ElastifileProtectionSource.md) | Specifies details about a Elastifile Protection Source when the environment is set to &#39;kElastifile&#39;. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment (such as &#39;kVMware&#39; or &#39;kSQL&#39;) where the Protection Source exists. Depending on the environment, one of the following Protection Sources are initialized.  NOTE: kPuppeteer refers to Cohesity&#39;s Remote Adapter. Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**ExchangeProtectionSource** | Pointer to [**NullableExchangeProtectionSource**](ExchangeProtectionSource.md) | Specifies details about an Exchange Protection Source when the environment is set to &#39;kExchange&#39;. | [optional] 
**FlashBladeProtectionSource** | Pointer to [**NullableFlashBladeProtectionSource**](FlashBladeProtectionSource.md) | Specifies details about a Pure Storage FlashBlade Protection Source when the environment is set to &#39;kFlashBlade&#39;. | [optional] 
**GcpProtectionSource** | Pointer to [**NullableGcpProtectionSource**](GcpProtectionSource.md) | Specifies details about an GCP Protection Source when the environment is set to &#39;kGCP&#39;. | [optional] 
**GpfsProtectionSource** | Pointer to [**NullableGpfsProtectionSource**](GpfsProtectionSource.md) | Specifies details about an GPFS Protection Source when the environment is set to &#39;kGPFS&#39;. | [optional] 
**HbaseProtectionSource** | Pointer to [**NullableHBaseProtectionSource**](HBaseProtectionSource.md) | Specifies details about a HBase Protection Source when the environment is set to &#39;kHBase&#39;. | [optional] 
**HdfsProtectionSource** | Pointer to [**NullableHdfsProtectionSource**](HdfsProtectionSource.md) | Specifies details about a Hdfs Protection Source when the environment is set to &#39;kHdfs&#39;. | [optional] 
**HiveProtectionSource** | Pointer to [**NullableHiveProtectionSource**](HiveProtectionSource.md) | Specifies details about a Hive Protection Source when the environment is set to &#39;kHive&#39;. | [optional] 
**HyperFlexProtectionSource** | Pointer to [**NullableHyperFlexProtectionSource**](HyperFlexProtectionSource.md) | Specifies details about a HyperFlex Storage Snapshot source when the environment is set to &#39;kHyperFlex&#39; | [optional] 
**HypervProtectionSource** | Pointer to [**NullableHypervProtectionSource**](HypervProtectionSource.md) | Specifies details about a HyperV Protection Source when the environment is set to &#39;kHyperV&#39;. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies an id of the Protection Source. | [optional] 
**IsilonProtectionSource** | Pointer to [**NullableIsilonProtectionSource**](IsilonProtectionSource.md) | Specifies details about an Isilon OneFs Protection Source when the environment is set to &#39;kIsilon&#39;. | [optional] 
**KubernetesProtectionSource** | Pointer to [**NullableKubernetesProtectionSource**](KubernetesProtectionSource.md) | Specifies details about a Kubernetes Protection Source when the environment is set to &#39;kKubernetes&#39;. | [optional] 
**KvmProtectionSource** | Pointer to [**NullableKvmProtectionSource**](KvmProtectionSource.md) | Specifies details about a KVM Protection Source when the environment is set to &#39;kKVM&#39;. | [optional] 
**MongodbProtectionSource** | Pointer to [**NullableMongoDBProtectionSource**](MongoDBProtectionSource.md) | Specifies details about a MongoDB Protection Source when the environment is set to &#39;kMongoDB&#39;. | [optional] 
**Name** | Pointer to **NullableString** | Specifies a name of the Protection Source. | [optional] 
**NasProtectionSource** | Pointer to [**NullableNasProtectionSource**](NasProtectionSource.md) | Specifies details about a Generic NAS Protection Source when the environment is set to &#39;kGenericNas&#39;. | [optional] 
**NetappProtectionSource** | Pointer to [**NullableNetappProtectionSource**](NetappProtectionSource.md) | Specifies details about a NetApp Protection Source when the environment is set to &#39;kNetapp&#39;. | [optional] 
**NimbleProtectionSource** | Pointer to [**NullableNimbleProtectionSource**](NimbleProtectionSource.md) | Specifies details about a SAN Protection Source when the environment is set to &#39;kNimble&#39;. | [optional] 
**Office365ProtectionSource** | Pointer to [**NullableOffice365ProtectionSource**](Office365ProtectionSource.md) | Specifies details about an Office 365 Protection Source when the environment is set to &#39;kO365&#39;. | [optional] 
**OracleProtectionSource** | Pointer to [**NullableOracleProtectionSource**](OracleProtectionSource.md) | Specifies details about an Oracle Protection Source when the environment is set to &#39;kOracle&#39;. | [optional] 
**ParentId** | Pointer to **NullableInt64** | Specifies an id of the parent of the Protection Source. | [optional] 
**PhysicalProtectionSource** | Pointer to [**NullablePhysicalProtectionSource**](PhysicalProtectionSource.md) | Specifies details about a Physical Protection Source when the environment is set to &#39;kPhysical&#39;. | [optional] 
**PureProtectionSource** | Pointer to [**NullablePureProtectionSource**](PureProtectionSource.md) | Specifies details about a Pure Protection Source when the environment is set to &#39;kPure&#39;. | [optional] 
**SqlProtectionSource** | Pointer to [**NullableSqlProtectionSource**](SqlProtectionSource.md) | Specifies details about a SQL Protection Source when the environment is set to &#39;kSQL&#39;. | [optional] 
**ViewProtectionSource** | Pointer to [**NullableViewProtectionSource**](ViewProtectionSource.md) | Specifies details about a View Protection Source when the environment is set to &#39;kView&#39;. | [optional] 
**VmWareProtectionSource** | Pointer to [**NullableVMwareProtectionSource**](VMwareProtectionSource.md) | Specifies details about a VMware Protection Source when the environment is set to &#39;kVMware&#39;. | [optional] 

## Methods

### NewProtectionSource

`func NewProtectionSource() *ProtectionSource`

NewProtectionSource instantiates a new ProtectionSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectionSourceWithDefaults

`func NewProtectionSourceWithDefaults() *ProtectionSource`

NewProtectionSourceWithDefaults instantiates a new ProtectionSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcropolisProtectionSource

`func (o *ProtectionSource) GetAcropolisProtectionSource() AcropolisProtectionSource`

GetAcropolisProtectionSource returns the AcropolisProtectionSource field if non-nil, zero value otherwise.

### GetAcropolisProtectionSourceOk

`func (o *ProtectionSource) GetAcropolisProtectionSourceOk() (*AcropolisProtectionSource, bool)`

GetAcropolisProtectionSourceOk returns a tuple with the AcropolisProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisProtectionSource

`func (o *ProtectionSource) SetAcropolisProtectionSource(v AcropolisProtectionSource)`

SetAcropolisProtectionSource sets AcropolisProtectionSource field to given value.

### HasAcropolisProtectionSource

`func (o *ProtectionSource) HasAcropolisProtectionSource() bool`

HasAcropolisProtectionSource returns a boolean if a field has been set.

### SetAcropolisProtectionSourceNil

`func (o *ProtectionSource) SetAcropolisProtectionSourceNil(b bool)`

 SetAcropolisProtectionSourceNil sets the value for AcropolisProtectionSource to be an explicit nil

### UnsetAcropolisProtectionSource
`func (o *ProtectionSource) UnsetAcropolisProtectionSource()`

UnsetAcropolisProtectionSource ensures that no value is present for AcropolisProtectionSource, not even an explicit nil
### GetAdProtectionSource

`func (o *ProtectionSource) GetAdProtectionSource() AdProtectionSource`

GetAdProtectionSource returns the AdProtectionSource field if non-nil, zero value otherwise.

### GetAdProtectionSourceOk

`func (o *ProtectionSource) GetAdProtectionSourceOk() (*AdProtectionSource, bool)`

GetAdProtectionSourceOk returns a tuple with the AdProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdProtectionSource

`func (o *ProtectionSource) SetAdProtectionSource(v AdProtectionSource)`

SetAdProtectionSource sets AdProtectionSource field to given value.

### HasAdProtectionSource

`func (o *ProtectionSource) HasAdProtectionSource() bool`

HasAdProtectionSource returns a boolean if a field has been set.

### SetAdProtectionSourceNil

`func (o *ProtectionSource) SetAdProtectionSourceNil(b bool)`

 SetAdProtectionSourceNil sets the value for AdProtectionSource to be an explicit nil

### UnsetAdProtectionSource
`func (o *ProtectionSource) UnsetAdProtectionSource()`

UnsetAdProtectionSource ensures that no value is present for AdProtectionSource, not even an explicit nil
### GetAwsProtectionSource

`func (o *ProtectionSource) GetAwsProtectionSource() AwsProtectionSource`

GetAwsProtectionSource returns the AwsProtectionSource field if non-nil, zero value otherwise.

### GetAwsProtectionSourceOk

`func (o *ProtectionSource) GetAwsProtectionSourceOk() (*AwsProtectionSource, bool)`

GetAwsProtectionSourceOk returns a tuple with the AwsProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsProtectionSource

`func (o *ProtectionSource) SetAwsProtectionSource(v AwsProtectionSource)`

SetAwsProtectionSource sets AwsProtectionSource field to given value.

### HasAwsProtectionSource

`func (o *ProtectionSource) HasAwsProtectionSource() bool`

HasAwsProtectionSource returns a boolean if a field has been set.

### SetAwsProtectionSourceNil

`func (o *ProtectionSource) SetAwsProtectionSourceNil(b bool)`

 SetAwsProtectionSourceNil sets the value for AwsProtectionSource to be an explicit nil

### UnsetAwsProtectionSource
`func (o *ProtectionSource) UnsetAwsProtectionSource()`

UnsetAwsProtectionSource ensures that no value is present for AwsProtectionSource, not even an explicit nil
### GetAzureProtectionSource

`func (o *ProtectionSource) GetAzureProtectionSource() AzureProtectionSource`

GetAzureProtectionSource returns the AzureProtectionSource field if non-nil, zero value otherwise.

### GetAzureProtectionSourceOk

`func (o *ProtectionSource) GetAzureProtectionSourceOk() (*AzureProtectionSource, bool)`

GetAzureProtectionSourceOk returns a tuple with the AzureProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureProtectionSource

`func (o *ProtectionSource) SetAzureProtectionSource(v AzureProtectionSource)`

SetAzureProtectionSource sets AzureProtectionSource field to given value.

### HasAzureProtectionSource

`func (o *ProtectionSource) HasAzureProtectionSource() bool`

HasAzureProtectionSource returns a boolean if a field has been set.

### SetAzureProtectionSourceNil

`func (o *ProtectionSource) SetAzureProtectionSourceNil(b bool)`

 SetAzureProtectionSourceNil sets the value for AzureProtectionSource to be an explicit nil

### UnsetAzureProtectionSource
`func (o *ProtectionSource) UnsetAzureProtectionSource()`

UnsetAzureProtectionSource ensures that no value is present for AzureProtectionSource, not even an explicit nil
### GetCassandraProtectionSource

`func (o *ProtectionSource) GetCassandraProtectionSource() CassandraProtectionSource`

GetCassandraProtectionSource returns the CassandraProtectionSource field if non-nil, zero value otherwise.

### GetCassandraProtectionSourceOk

`func (o *ProtectionSource) GetCassandraProtectionSourceOk() (*CassandraProtectionSource, bool)`

GetCassandraProtectionSourceOk returns a tuple with the CassandraProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraProtectionSource

`func (o *ProtectionSource) SetCassandraProtectionSource(v CassandraProtectionSource)`

SetCassandraProtectionSource sets CassandraProtectionSource field to given value.

### HasCassandraProtectionSource

`func (o *ProtectionSource) HasCassandraProtectionSource() bool`

HasCassandraProtectionSource returns a boolean if a field has been set.

### SetCassandraProtectionSourceNil

`func (o *ProtectionSource) SetCassandraProtectionSourceNil(b bool)`

 SetCassandraProtectionSourceNil sets the value for CassandraProtectionSource to be an explicit nil

### UnsetCassandraProtectionSource
`func (o *ProtectionSource) UnsetCassandraProtectionSource()`

UnsetCassandraProtectionSource ensures that no value is present for CassandraProtectionSource, not even an explicit nil
### GetCouchbaseProtectionSource

`func (o *ProtectionSource) GetCouchbaseProtectionSource() CouchbaseProtectionSource`

GetCouchbaseProtectionSource returns the CouchbaseProtectionSource field if non-nil, zero value otherwise.

### GetCouchbaseProtectionSourceOk

`func (o *ProtectionSource) GetCouchbaseProtectionSourceOk() (*CouchbaseProtectionSource, bool)`

GetCouchbaseProtectionSourceOk returns a tuple with the CouchbaseProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouchbaseProtectionSource

`func (o *ProtectionSource) SetCouchbaseProtectionSource(v CouchbaseProtectionSource)`

SetCouchbaseProtectionSource sets CouchbaseProtectionSource field to given value.

### HasCouchbaseProtectionSource

`func (o *ProtectionSource) HasCouchbaseProtectionSource() bool`

HasCouchbaseProtectionSource returns a boolean if a field has been set.

### SetCouchbaseProtectionSourceNil

`func (o *ProtectionSource) SetCouchbaseProtectionSourceNil(b bool)`

 SetCouchbaseProtectionSourceNil sets the value for CouchbaseProtectionSource to be an explicit nil

### UnsetCouchbaseProtectionSource
`func (o *ProtectionSource) UnsetCouchbaseProtectionSource()`

UnsetCouchbaseProtectionSource ensures that no value is present for CouchbaseProtectionSource, not even an explicit nil
### GetElastifileProtectionSource

`func (o *ProtectionSource) GetElastifileProtectionSource() ElastifileProtectionSource`

GetElastifileProtectionSource returns the ElastifileProtectionSource field if non-nil, zero value otherwise.

### GetElastifileProtectionSourceOk

`func (o *ProtectionSource) GetElastifileProtectionSourceOk() (*ElastifileProtectionSource, bool)`

GetElastifileProtectionSourceOk returns a tuple with the ElastifileProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastifileProtectionSource

`func (o *ProtectionSource) SetElastifileProtectionSource(v ElastifileProtectionSource)`

SetElastifileProtectionSource sets ElastifileProtectionSource field to given value.

### HasElastifileProtectionSource

`func (o *ProtectionSource) HasElastifileProtectionSource() bool`

HasElastifileProtectionSource returns a boolean if a field has been set.

### SetElastifileProtectionSourceNil

`func (o *ProtectionSource) SetElastifileProtectionSourceNil(b bool)`

 SetElastifileProtectionSourceNil sets the value for ElastifileProtectionSource to be an explicit nil

### UnsetElastifileProtectionSource
`func (o *ProtectionSource) UnsetElastifileProtectionSource()`

UnsetElastifileProtectionSource ensures that no value is present for ElastifileProtectionSource, not even an explicit nil
### GetEnvironment

`func (o *ProtectionSource) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProtectionSource) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProtectionSource) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProtectionSource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *ProtectionSource) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ProtectionSource) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExchangeProtectionSource

`func (o *ProtectionSource) GetExchangeProtectionSource() ExchangeProtectionSource`

GetExchangeProtectionSource returns the ExchangeProtectionSource field if non-nil, zero value otherwise.

### GetExchangeProtectionSourceOk

`func (o *ProtectionSource) GetExchangeProtectionSourceOk() (*ExchangeProtectionSource, bool)`

GetExchangeProtectionSourceOk returns a tuple with the ExchangeProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeProtectionSource

`func (o *ProtectionSource) SetExchangeProtectionSource(v ExchangeProtectionSource)`

SetExchangeProtectionSource sets ExchangeProtectionSource field to given value.

### HasExchangeProtectionSource

`func (o *ProtectionSource) HasExchangeProtectionSource() bool`

HasExchangeProtectionSource returns a boolean if a field has been set.

### SetExchangeProtectionSourceNil

`func (o *ProtectionSource) SetExchangeProtectionSourceNil(b bool)`

 SetExchangeProtectionSourceNil sets the value for ExchangeProtectionSource to be an explicit nil

### UnsetExchangeProtectionSource
`func (o *ProtectionSource) UnsetExchangeProtectionSource()`

UnsetExchangeProtectionSource ensures that no value is present for ExchangeProtectionSource, not even an explicit nil
### GetFlashBladeProtectionSource

`func (o *ProtectionSource) GetFlashBladeProtectionSource() FlashBladeProtectionSource`

GetFlashBladeProtectionSource returns the FlashBladeProtectionSource field if non-nil, zero value otherwise.

### GetFlashBladeProtectionSourceOk

`func (o *ProtectionSource) GetFlashBladeProtectionSourceOk() (*FlashBladeProtectionSource, bool)`

GetFlashBladeProtectionSourceOk returns a tuple with the FlashBladeProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlashBladeProtectionSource

`func (o *ProtectionSource) SetFlashBladeProtectionSource(v FlashBladeProtectionSource)`

SetFlashBladeProtectionSource sets FlashBladeProtectionSource field to given value.

### HasFlashBladeProtectionSource

`func (o *ProtectionSource) HasFlashBladeProtectionSource() bool`

HasFlashBladeProtectionSource returns a boolean if a field has been set.

### SetFlashBladeProtectionSourceNil

`func (o *ProtectionSource) SetFlashBladeProtectionSourceNil(b bool)`

 SetFlashBladeProtectionSourceNil sets the value for FlashBladeProtectionSource to be an explicit nil

### UnsetFlashBladeProtectionSource
`func (o *ProtectionSource) UnsetFlashBladeProtectionSource()`

UnsetFlashBladeProtectionSource ensures that no value is present for FlashBladeProtectionSource, not even an explicit nil
### GetGcpProtectionSource

`func (o *ProtectionSource) GetGcpProtectionSource() GcpProtectionSource`

GetGcpProtectionSource returns the GcpProtectionSource field if non-nil, zero value otherwise.

### GetGcpProtectionSourceOk

`func (o *ProtectionSource) GetGcpProtectionSourceOk() (*GcpProtectionSource, bool)`

GetGcpProtectionSourceOk returns a tuple with the GcpProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProtectionSource

`func (o *ProtectionSource) SetGcpProtectionSource(v GcpProtectionSource)`

SetGcpProtectionSource sets GcpProtectionSource field to given value.

### HasGcpProtectionSource

`func (o *ProtectionSource) HasGcpProtectionSource() bool`

HasGcpProtectionSource returns a boolean if a field has been set.

### SetGcpProtectionSourceNil

`func (o *ProtectionSource) SetGcpProtectionSourceNil(b bool)`

 SetGcpProtectionSourceNil sets the value for GcpProtectionSource to be an explicit nil

### UnsetGcpProtectionSource
`func (o *ProtectionSource) UnsetGcpProtectionSource()`

UnsetGcpProtectionSource ensures that no value is present for GcpProtectionSource, not even an explicit nil
### GetGpfsProtectionSource

`func (o *ProtectionSource) GetGpfsProtectionSource() GpfsProtectionSource`

GetGpfsProtectionSource returns the GpfsProtectionSource field if non-nil, zero value otherwise.

### GetGpfsProtectionSourceOk

`func (o *ProtectionSource) GetGpfsProtectionSourceOk() (*GpfsProtectionSource, bool)`

GetGpfsProtectionSourceOk returns a tuple with the GpfsProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpfsProtectionSource

`func (o *ProtectionSource) SetGpfsProtectionSource(v GpfsProtectionSource)`

SetGpfsProtectionSource sets GpfsProtectionSource field to given value.

### HasGpfsProtectionSource

`func (o *ProtectionSource) HasGpfsProtectionSource() bool`

HasGpfsProtectionSource returns a boolean if a field has been set.

### SetGpfsProtectionSourceNil

`func (o *ProtectionSource) SetGpfsProtectionSourceNil(b bool)`

 SetGpfsProtectionSourceNil sets the value for GpfsProtectionSource to be an explicit nil

### UnsetGpfsProtectionSource
`func (o *ProtectionSource) UnsetGpfsProtectionSource()`

UnsetGpfsProtectionSource ensures that no value is present for GpfsProtectionSource, not even an explicit nil
### GetHbaseProtectionSource

`func (o *ProtectionSource) GetHbaseProtectionSource() HBaseProtectionSource`

GetHbaseProtectionSource returns the HbaseProtectionSource field if non-nil, zero value otherwise.

### GetHbaseProtectionSourceOk

`func (o *ProtectionSource) GetHbaseProtectionSourceOk() (*HBaseProtectionSource, bool)`

GetHbaseProtectionSourceOk returns a tuple with the HbaseProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHbaseProtectionSource

`func (o *ProtectionSource) SetHbaseProtectionSource(v HBaseProtectionSource)`

SetHbaseProtectionSource sets HbaseProtectionSource field to given value.

### HasHbaseProtectionSource

`func (o *ProtectionSource) HasHbaseProtectionSource() bool`

HasHbaseProtectionSource returns a boolean if a field has been set.

### SetHbaseProtectionSourceNil

`func (o *ProtectionSource) SetHbaseProtectionSourceNil(b bool)`

 SetHbaseProtectionSourceNil sets the value for HbaseProtectionSource to be an explicit nil

### UnsetHbaseProtectionSource
`func (o *ProtectionSource) UnsetHbaseProtectionSource()`

UnsetHbaseProtectionSource ensures that no value is present for HbaseProtectionSource, not even an explicit nil
### GetHdfsProtectionSource

`func (o *ProtectionSource) GetHdfsProtectionSource() HdfsProtectionSource`

GetHdfsProtectionSource returns the HdfsProtectionSource field if non-nil, zero value otherwise.

### GetHdfsProtectionSourceOk

`func (o *ProtectionSource) GetHdfsProtectionSourceOk() (*HdfsProtectionSource, bool)`

GetHdfsProtectionSourceOk returns a tuple with the HdfsProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsProtectionSource

`func (o *ProtectionSource) SetHdfsProtectionSource(v HdfsProtectionSource)`

SetHdfsProtectionSource sets HdfsProtectionSource field to given value.

### HasHdfsProtectionSource

`func (o *ProtectionSource) HasHdfsProtectionSource() bool`

HasHdfsProtectionSource returns a boolean if a field has been set.

### SetHdfsProtectionSourceNil

`func (o *ProtectionSource) SetHdfsProtectionSourceNil(b bool)`

 SetHdfsProtectionSourceNil sets the value for HdfsProtectionSource to be an explicit nil

### UnsetHdfsProtectionSource
`func (o *ProtectionSource) UnsetHdfsProtectionSource()`

UnsetHdfsProtectionSource ensures that no value is present for HdfsProtectionSource, not even an explicit nil
### GetHiveProtectionSource

`func (o *ProtectionSource) GetHiveProtectionSource() HiveProtectionSource`

GetHiveProtectionSource returns the HiveProtectionSource field if non-nil, zero value otherwise.

### GetHiveProtectionSourceOk

`func (o *ProtectionSource) GetHiveProtectionSourceOk() (*HiveProtectionSource, bool)`

GetHiveProtectionSourceOk returns a tuple with the HiveProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiveProtectionSource

`func (o *ProtectionSource) SetHiveProtectionSource(v HiveProtectionSource)`

SetHiveProtectionSource sets HiveProtectionSource field to given value.

### HasHiveProtectionSource

`func (o *ProtectionSource) HasHiveProtectionSource() bool`

HasHiveProtectionSource returns a boolean if a field has been set.

### SetHiveProtectionSourceNil

`func (o *ProtectionSource) SetHiveProtectionSourceNil(b bool)`

 SetHiveProtectionSourceNil sets the value for HiveProtectionSource to be an explicit nil

### UnsetHiveProtectionSource
`func (o *ProtectionSource) UnsetHiveProtectionSource()`

UnsetHiveProtectionSource ensures that no value is present for HiveProtectionSource, not even an explicit nil
### GetHyperFlexProtectionSource

`func (o *ProtectionSource) GetHyperFlexProtectionSource() HyperFlexProtectionSource`

GetHyperFlexProtectionSource returns the HyperFlexProtectionSource field if non-nil, zero value otherwise.

### GetHyperFlexProtectionSourceOk

`func (o *ProtectionSource) GetHyperFlexProtectionSourceOk() (*HyperFlexProtectionSource, bool)`

GetHyperFlexProtectionSourceOk returns a tuple with the HyperFlexProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperFlexProtectionSource

`func (o *ProtectionSource) SetHyperFlexProtectionSource(v HyperFlexProtectionSource)`

SetHyperFlexProtectionSource sets HyperFlexProtectionSource field to given value.

### HasHyperFlexProtectionSource

`func (o *ProtectionSource) HasHyperFlexProtectionSource() bool`

HasHyperFlexProtectionSource returns a boolean if a field has been set.

### SetHyperFlexProtectionSourceNil

`func (o *ProtectionSource) SetHyperFlexProtectionSourceNil(b bool)`

 SetHyperFlexProtectionSourceNil sets the value for HyperFlexProtectionSource to be an explicit nil

### UnsetHyperFlexProtectionSource
`func (o *ProtectionSource) UnsetHyperFlexProtectionSource()`

UnsetHyperFlexProtectionSource ensures that no value is present for HyperFlexProtectionSource, not even an explicit nil
### GetHypervProtectionSource

`func (o *ProtectionSource) GetHypervProtectionSource() HypervProtectionSource`

GetHypervProtectionSource returns the HypervProtectionSource field if non-nil, zero value otherwise.

### GetHypervProtectionSourceOk

`func (o *ProtectionSource) GetHypervProtectionSourceOk() (*HypervProtectionSource, bool)`

GetHypervProtectionSourceOk returns a tuple with the HypervProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervProtectionSource

`func (o *ProtectionSource) SetHypervProtectionSource(v HypervProtectionSource)`

SetHypervProtectionSource sets HypervProtectionSource field to given value.

### HasHypervProtectionSource

`func (o *ProtectionSource) HasHypervProtectionSource() bool`

HasHypervProtectionSource returns a boolean if a field has been set.

### SetHypervProtectionSourceNil

`func (o *ProtectionSource) SetHypervProtectionSourceNil(b bool)`

 SetHypervProtectionSourceNil sets the value for HypervProtectionSource to be an explicit nil

### UnsetHypervProtectionSource
`func (o *ProtectionSource) UnsetHypervProtectionSource()`

UnsetHypervProtectionSource ensures that no value is present for HypervProtectionSource, not even an explicit nil
### GetId

`func (o *ProtectionSource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtectionSource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtectionSource) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtectionSource) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProtectionSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProtectionSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsilonProtectionSource

`func (o *ProtectionSource) GetIsilonProtectionSource() IsilonProtectionSource`

GetIsilonProtectionSource returns the IsilonProtectionSource field if non-nil, zero value otherwise.

### GetIsilonProtectionSourceOk

`func (o *ProtectionSource) GetIsilonProtectionSourceOk() (*IsilonProtectionSource, bool)`

GetIsilonProtectionSourceOk returns a tuple with the IsilonProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsilonProtectionSource

`func (o *ProtectionSource) SetIsilonProtectionSource(v IsilonProtectionSource)`

SetIsilonProtectionSource sets IsilonProtectionSource field to given value.

### HasIsilonProtectionSource

`func (o *ProtectionSource) HasIsilonProtectionSource() bool`

HasIsilonProtectionSource returns a boolean if a field has been set.

### SetIsilonProtectionSourceNil

`func (o *ProtectionSource) SetIsilonProtectionSourceNil(b bool)`

 SetIsilonProtectionSourceNil sets the value for IsilonProtectionSource to be an explicit nil

### UnsetIsilonProtectionSource
`func (o *ProtectionSource) UnsetIsilonProtectionSource()`

UnsetIsilonProtectionSource ensures that no value is present for IsilonProtectionSource, not even an explicit nil
### GetKubernetesProtectionSource

`func (o *ProtectionSource) GetKubernetesProtectionSource() KubernetesProtectionSource`

GetKubernetesProtectionSource returns the KubernetesProtectionSource field if non-nil, zero value otherwise.

### GetKubernetesProtectionSourceOk

`func (o *ProtectionSource) GetKubernetesProtectionSourceOk() (*KubernetesProtectionSource, bool)`

GetKubernetesProtectionSourceOk returns a tuple with the KubernetesProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProtectionSource

`func (o *ProtectionSource) SetKubernetesProtectionSource(v KubernetesProtectionSource)`

SetKubernetesProtectionSource sets KubernetesProtectionSource field to given value.

### HasKubernetesProtectionSource

`func (o *ProtectionSource) HasKubernetesProtectionSource() bool`

HasKubernetesProtectionSource returns a boolean if a field has been set.

### SetKubernetesProtectionSourceNil

`func (o *ProtectionSource) SetKubernetesProtectionSourceNil(b bool)`

 SetKubernetesProtectionSourceNil sets the value for KubernetesProtectionSource to be an explicit nil

### UnsetKubernetesProtectionSource
`func (o *ProtectionSource) UnsetKubernetesProtectionSource()`

UnsetKubernetesProtectionSource ensures that no value is present for KubernetesProtectionSource, not even an explicit nil
### GetKvmProtectionSource

`func (o *ProtectionSource) GetKvmProtectionSource() KvmProtectionSource`

GetKvmProtectionSource returns the KvmProtectionSource field if non-nil, zero value otherwise.

### GetKvmProtectionSourceOk

`func (o *ProtectionSource) GetKvmProtectionSourceOk() (*KvmProtectionSource, bool)`

GetKvmProtectionSourceOk returns a tuple with the KvmProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmProtectionSource

`func (o *ProtectionSource) SetKvmProtectionSource(v KvmProtectionSource)`

SetKvmProtectionSource sets KvmProtectionSource field to given value.

### HasKvmProtectionSource

`func (o *ProtectionSource) HasKvmProtectionSource() bool`

HasKvmProtectionSource returns a boolean if a field has been set.

### SetKvmProtectionSourceNil

`func (o *ProtectionSource) SetKvmProtectionSourceNil(b bool)`

 SetKvmProtectionSourceNil sets the value for KvmProtectionSource to be an explicit nil

### UnsetKvmProtectionSource
`func (o *ProtectionSource) UnsetKvmProtectionSource()`

UnsetKvmProtectionSource ensures that no value is present for KvmProtectionSource, not even an explicit nil
### GetMongodbProtectionSource

`func (o *ProtectionSource) GetMongodbProtectionSource() MongoDBProtectionSource`

GetMongodbProtectionSource returns the MongodbProtectionSource field if non-nil, zero value otherwise.

### GetMongodbProtectionSourceOk

`func (o *ProtectionSource) GetMongodbProtectionSourceOk() (*MongoDBProtectionSource, bool)`

GetMongodbProtectionSourceOk returns a tuple with the MongodbProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodbProtectionSource

`func (o *ProtectionSource) SetMongodbProtectionSource(v MongoDBProtectionSource)`

SetMongodbProtectionSource sets MongodbProtectionSource field to given value.

### HasMongodbProtectionSource

`func (o *ProtectionSource) HasMongodbProtectionSource() bool`

HasMongodbProtectionSource returns a boolean if a field has been set.

### SetMongodbProtectionSourceNil

`func (o *ProtectionSource) SetMongodbProtectionSourceNil(b bool)`

 SetMongodbProtectionSourceNil sets the value for MongodbProtectionSource to be an explicit nil

### UnsetMongodbProtectionSource
`func (o *ProtectionSource) UnsetMongodbProtectionSource()`

UnsetMongodbProtectionSource ensures that no value is present for MongodbProtectionSource, not even an explicit nil
### GetName

`func (o *ProtectionSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectionSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectionSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectionSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProtectionSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProtectionSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNasProtectionSource

`func (o *ProtectionSource) GetNasProtectionSource() NasProtectionSource`

GetNasProtectionSource returns the NasProtectionSource field if non-nil, zero value otherwise.

### GetNasProtectionSourceOk

`func (o *ProtectionSource) GetNasProtectionSourceOk() (*NasProtectionSource, bool)`

GetNasProtectionSourceOk returns a tuple with the NasProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasProtectionSource

`func (o *ProtectionSource) SetNasProtectionSource(v NasProtectionSource)`

SetNasProtectionSource sets NasProtectionSource field to given value.

### HasNasProtectionSource

`func (o *ProtectionSource) HasNasProtectionSource() bool`

HasNasProtectionSource returns a boolean if a field has been set.

### SetNasProtectionSourceNil

`func (o *ProtectionSource) SetNasProtectionSourceNil(b bool)`

 SetNasProtectionSourceNil sets the value for NasProtectionSource to be an explicit nil

### UnsetNasProtectionSource
`func (o *ProtectionSource) UnsetNasProtectionSource()`

UnsetNasProtectionSource ensures that no value is present for NasProtectionSource, not even an explicit nil
### GetNetappProtectionSource

`func (o *ProtectionSource) GetNetappProtectionSource() NetappProtectionSource`

GetNetappProtectionSource returns the NetappProtectionSource field if non-nil, zero value otherwise.

### GetNetappProtectionSourceOk

`func (o *ProtectionSource) GetNetappProtectionSourceOk() (*NetappProtectionSource, bool)`

GetNetappProtectionSourceOk returns a tuple with the NetappProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappProtectionSource

`func (o *ProtectionSource) SetNetappProtectionSource(v NetappProtectionSource)`

SetNetappProtectionSource sets NetappProtectionSource field to given value.

### HasNetappProtectionSource

`func (o *ProtectionSource) HasNetappProtectionSource() bool`

HasNetappProtectionSource returns a boolean if a field has been set.

### SetNetappProtectionSourceNil

`func (o *ProtectionSource) SetNetappProtectionSourceNil(b bool)`

 SetNetappProtectionSourceNil sets the value for NetappProtectionSource to be an explicit nil

### UnsetNetappProtectionSource
`func (o *ProtectionSource) UnsetNetappProtectionSource()`

UnsetNetappProtectionSource ensures that no value is present for NetappProtectionSource, not even an explicit nil
### GetNimbleProtectionSource

`func (o *ProtectionSource) GetNimbleProtectionSource() NimbleProtectionSource`

GetNimbleProtectionSource returns the NimbleProtectionSource field if non-nil, zero value otherwise.

### GetNimbleProtectionSourceOk

`func (o *ProtectionSource) GetNimbleProtectionSourceOk() (*NimbleProtectionSource, bool)`

GetNimbleProtectionSourceOk returns a tuple with the NimbleProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNimbleProtectionSource

`func (o *ProtectionSource) SetNimbleProtectionSource(v NimbleProtectionSource)`

SetNimbleProtectionSource sets NimbleProtectionSource field to given value.

### HasNimbleProtectionSource

`func (o *ProtectionSource) HasNimbleProtectionSource() bool`

HasNimbleProtectionSource returns a boolean if a field has been set.

### SetNimbleProtectionSourceNil

`func (o *ProtectionSource) SetNimbleProtectionSourceNil(b bool)`

 SetNimbleProtectionSourceNil sets the value for NimbleProtectionSource to be an explicit nil

### UnsetNimbleProtectionSource
`func (o *ProtectionSource) UnsetNimbleProtectionSource()`

UnsetNimbleProtectionSource ensures that no value is present for NimbleProtectionSource, not even an explicit nil
### GetOffice365ProtectionSource

`func (o *ProtectionSource) GetOffice365ProtectionSource() Office365ProtectionSource`

GetOffice365ProtectionSource returns the Office365ProtectionSource field if non-nil, zero value otherwise.

### GetOffice365ProtectionSourceOk

`func (o *ProtectionSource) GetOffice365ProtectionSourceOk() (*Office365ProtectionSource, bool)`

GetOffice365ProtectionSourceOk returns a tuple with the Office365ProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365ProtectionSource

`func (o *ProtectionSource) SetOffice365ProtectionSource(v Office365ProtectionSource)`

SetOffice365ProtectionSource sets Office365ProtectionSource field to given value.

### HasOffice365ProtectionSource

`func (o *ProtectionSource) HasOffice365ProtectionSource() bool`

HasOffice365ProtectionSource returns a boolean if a field has been set.

### SetOffice365ProtectionSourceNil

`func (o *ProtectionSource) SetOffice365ProtectionSourceNil(b bool)`

 SetOffice365ProtectionSourceNil sets the value for Office365ProtectionSource to be an explicit nil

### UnsetOffice365ProtectionSource
`func (o *ProtectionSource) UnsetOffice365ProtectionSource()`

UnsetOffice365ProtectionSource ensures that no value is present for Office365ProtectionSource, not even an explicit nil
### GetOracleProtectionSource

`func (o *ProtectionSource) GetOracleProtectionSource() OracleProtectionSource`

GetOracleProtectionSource returns the OracleProtectionSource field if non-nil, zero value otherwise.

### GetOracleProtectionSourceOk

`func (o *ProtectionSource) GetOracleProtectionSourceOk() (*OracleProtectionSource, bool)`

GetOracleProtectionSourceOk returns a tuple with the OracleProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleProtectionSource

`func (o *ProtectionSource) SetOracleProtectionSource(v OracleProtectionSource)`

SetOracleProtectionSource sets OracleProtectionSource field to given value.

### HasOracleProtectionSource

`func (o *ProtectionSource) HasOracleProtectionSource() bool`

HasOracleProtectionSource returns a boolean if a field has been set.

### SetOracleProtectionSourceNil

`func (o *ProtectionSource) SetOracleProtectionSourceNil(b bool)`

 SetOracleProtectionSourceNil sets the value for OracleProtectionSource to be an explicit nil

### UnsetOracleProtectionSource
`func (o *ProtectionSource) UnsetOracleProtectionSource()`

UnsetOracleProtectionSource ensures that no value is present for OracleProtectionSource, not even an explicit nil
### GetParentId

`func (o *ProtectionSource) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProtectionSource) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProtectionSource) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProtectionSource) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ProtectionSource) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ProtectionSource) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPhysicalProtectionSource

`func (o *ProtectionSource) GetPhysicalProtectionSource() PhysicalProtectionSource`

GetPhysicalProtectionSource returns the PhysicalProtectionSource field if non-nil, zero value otherwise.

### GetPhysicalProtectionSourceOk

`func (o *ProtectionSource) GetPhysicalProtectionSourceOk() (*PhysicalProtectionSource, bool)`

GetPhysicalProtectionSourceOk returns a tuple with the PhysicalProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalProtectionSource

`func (o *ProtectionSource) SetPhysicalProtectionSource(v PhysicalProtectionSource)`

SetPhysicalProtectionSource sets PhysicalProtectionSource field to given value.

### HasPhysicalProtectionSource

`func (o *ProtectionSource) HasPhysicalProtectionSource() bool`

HasPhysicalProtectionSource returns a boolean if a field has been set.

### SetPhysicalProtectionSourceNil

`func (o *ProtectionSource) SetPhysicalProtectionSourceNil(b bool)`

 SetPhysicalProtectionSourceNil sets the value for PhysicalProtectionSource to be an explicit nil

### UnsetPhysicalProtectionSource
`func (o *ProtectionSource) UnsetPhysicalProtectionSource()`

UnsetPhysicalProtectionSource ensures that no value is present for PhysicalProtectionSource, not even an explicit nil
### GetPureProtectionSource

`func (o *ProtectionSource) GetPureProtectionSource() PureProtectionSource`

GetPureProtectionSource returns the PureProtectionSource field if non-nil, zero value otherwise.

### GetPureProtectionSourceOk

`func (o *ProtectionSource) GetPureProtectionSourceOk() (*PureProtectionSource, bool)`

GetPureProtectionSourceOk returns a tuple with the PureProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureProtectionSource

`func (o *ProtectionSource) SetPureProtectionSource(v PureProtectionSource)`

SetPureProtectionSource sets PureProtectionSource field to given value.

### HasPureProtectionSource

`func (o *ProtectionSource) HasPureProtectionSource() bool`

HasPureProtectionSource returns a boolean if a field has been set.

### SetPureProtectionSourceNil

`func (o *ProtectionSource) SetPureProtectionSourceNil(b bool)`

 SetPureProtectionSourceNil sets the value for PureProtectionSource to be an explicit nil

### UnsetPureProtectionSource
`func (o *ProtectionSource) UnsetPureProtectionSource()`

UnsetPureProtectionSource ensures that no value is present for PureProtectionSource, not even an explicit nil
### GetSqlProtectionSource

`func (o *ProtectionSource) GetSqlProtectionSource() SqlProtectionSource`

GetSqlProtectionSource returns the SqlProtectionSource field if non-nil, zero value otherwise.

### GetSqlProtectionSourceOk

`func (o *ProtectionSource) GetSqlProtectionSourceOk() (*SqlProtectionSource, bool)`

GetSqlProtectionSourceOk returns a tuple with the SqlProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlProtectionSource

`func (o *ProtectionSource) SetSqlProtectionSource(v SqlProtectionSource)`

SetSqlProtectionSource sets SqlProtectionSource field to given value.

### HasSqlProtectionSource

`func (o *ProtectionSource) HasSqlProtectionSource() bool`

HasSqlProtectionSource returns a boolean if a field has been set.

### SetSqlProtectionSourceNil

`func (o *ProtectionSource) SetSqlProtectionSourceNil(b bool)`

 SetSqlProtectionSourceNil sets the value for SqlProtectionSource to be an explicit nil

### UnsetSqlProtectionSource
`func (o *ProtectionSource) UnsetSqlProtectionSource()`

UnsetSqlProtectionSource ensures that no value is present for SqlProtectionSource, not even an explicit nil
### GetViewProtectionSource

`func (o *ProtectionSource) GetViewProtectionSource() ViewProtectionSource`

GetViewProtectionSource returns the ViewProtectionSource field if non-nil, zero value otherwise.

### GetViewProtectionSourceOk

`func (o *ProtectionSource) GetViewProtectionSourceOk() (*ViewProtectionSource, bool)`

GetViewProtectionSourceOk returns a tuple with the ViewProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewProtectionSource

`func (o *ProtectionSource) SetViewProtectionSource(v ViewProtectionSource)`

SetViewProtectionSource sets ViewProtectionSource field to given value.

### HasViewProtectionSource

`func (o *ProtectionSource) HasViewProtectionSource() bool`

HasViewProtectionSource returns a boolean if a field has been set.

### SetViewProtectionSourceNil

`func (o *ProtectionSource) SetViewProtectionSourceNil(b bool)`

 SetViewProtectionSourceNil sets the value for ViewProtectionSource to be an explicit nil

### UnsetViewProtectionSource
`func (o *ProtectionSource) UnsetViewProtectionSource()`

UnsetViewProtectionSource ensures that no value is present for ViewProtectionSource, not even an explicit nil
### GetVmWareProtectionSource

`func (o *ProtectionSource) GetVmWareProtectionSource() VMwareProtectionSource`

GetVmWareProtectionSource returns the VmWareProtectionSource field if non-nil, zero value otherwise.

### GetVmWareProtectionSourceOk

`func (o *ProtectionSource) GetVmWareProtectionSourceOk() (*VMwareProtectionSource, bool)`

GetVmWareProtectionSourceOk returns a tuple with the VmWareProtectionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmWareProtectionSource

`func (o *ProtectionSource) SetVmWareProtectionSource(v VMwareProtectionSource)`

SetVmWareProtectionSource sets VmWareProtectionSource field to given value.

### HasVmWareProtectionSource

`func (o *ProtectionSource) HasVmWareProtectionSource() bool`

HasVmWareProtectionSource returns a boolean if a field has been set.

### SetVmWareProtectionSourceNil

`func (o *ProtectionSource) SetVmWareProtectionSourceNil(b bool)`

 SetVmWareProtectionSourceNil sets the value for VmWareProtectionSource to be an explicit nil

### UnsetVmWareProtectionSource
`func (o *ProtectionSource) UnsetVmWareProtectionSource()`

UnsetVmWareProtectionSource ensures that no value is present for VmWareProtectionSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


