# RegisterProtectionSourceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcropolisType** | Pointer to **NullableString** | Specifies the entity type if the environment is kAcropolis. overrideDescription: true | [optional] 
**AgentEndpoint** | Pointer to **NullableString** | Specifies the agent endpoint if it is different from the source endpoint. | [optional] 
**AwsCredentials** | Pointer to [**AwsCredentials**](AwsCredentials.md) |  | [optional] 
**AwsFleetParams** | Pointer to [**AwsFleetParams**](AwsFleetParams.md) |  | [optional] 
**AzureCredentials** | Pointer to [**AzureCredentials**](AzureCredentials.md) |  | [optional] 
**BlacklistedIpAddresses** | Pointer to **[]string** | Specifies the list of IP Addresses on the registered source to be blacklisted for doing any type of IO operations. | [optional] 
**ClusterNetworkInfo** | Pointer to [**FleetNetworkParams**](FleetNetworkParams.md) |  | [optional] 
**Endpoint** | Pointer to **NullableString** | Specifies the network endpoint of the Protection Source where it is reachable. It could be an URL or hostname or an IP address of the Protection Source. | [optional] 
**Environment** | Pointer to **NullableString** | Specifies the environment such as &#39;kPhysical&#39; or &#39;kVMware&#39; of the Protection Source. overrideDescription: true Supported environment types such as &#39;kView&#39;, &#39;kSQL&#39;, &#39;kVMware&#39;, etc. NOTE: &#39;kPuppeteer&#39; refers to Cohesity&#39;s Remote Adapter. &#39;kVMware&#39; indicates the VMware Protection Source environment. &#39;kHyperV&#39; indicates the HyperV Protection Source environment. &#39;kSQL&#39; indicates the SQL Protection Source environment. &#39;kView&#39; indicates the View Protection Source environment. &#39;kPuppeteer&#39; indicates the Cohesity&#39;s Remote Adapter. &#39;kPhysical&#39; indicates the physical Protection Source environment. &#39;kPure&#39; indicates the Pure Storage Protection Source environment. &#39;Nimble&#39; indicates the Nimble Storage Protection Source environment. &#39;kAzure&#39; indicates the Microsoft&#39;s Azure Protection Source environment. &#39;kNetapp&#39; indicates the Netapp Protection Source environment. &#39;kAgent&#39; indicates the Agent Protection Source environment. &#39;kGenericNas&#39; indicates the Generic Network Attached Storage Protection Source environment. &#39;kAcropolis&#39; indicates the Acropolis Protection Source environment. &#39;kPhsicalFiles&#39; indicates the Physical Files Protection Source environment. &#39;kIsilon&#39; indicates the Dell EMC&#39;s Isilon Protection Source environment. &#39;kGPFS&#39; indicates IBM&#39;s GPFS Protection Source environment. &#39;kKVM&#39; indicates the KVM Protection Source environment. &#39;kAWS&#39; indicates the AWS Protection Source environment. &#39;kExchange&#39; indicates the Exchange Protection Source environment. &#39;kHyperVVSS&#39; indicates the HyperV VSS Protection Source environment. &#39;kOracle&#39; indicates the Oracle Protection Source environment. &#39;kGCP&#39; indicates the Google Cloud Platform Protection Source environment. &#39;kFlashBlade&#39; indicates the Flash Blade Protection Source environment. &#39;kAWSNative&#39; indicates the AWS Native Protection Source environment. &#39;kO365&#39; indicates the Office 365 Protection Source environment. &#39;kO365Outlook&#39; indicates Office 365 outlook Protection Source environment. &#39;kHyperFlex&#39; indicates the Hyper Flex Protection Source environment. &#39;kGCPNative&#39; indicates the GCP Native Protection Source environment. &#39;kAzureNative&#39; indicates the Azure Native Protection Source environment. &#39;kKubernetes&#39; indicates a Kubernetes Protection Source environment. &#39;kElastifile&#39; indicates Elastifile Protection Source environment. &#39;kAD&#39; indicates Active Directory Protection Source environment. &#39;kRDSSnapshotManager&#39; indicates AWS RDS Protection Source environment. &#39;kCassandra&#39; indicates Cassandra Protection Source environment. &#39;kMongoDB&#39; indicates MongoDB Protection Source environment. &#39;kCouchbase&#39; indicates Couchbase Protection Source environment. &#39;kHdfs&#39; indicates Hdfs Protection Source environment. &#39;kHive&#39; indicates Hive Protection Source environment. &#39;kHBase&#39; indicates HBase Protection Source environment. | [optional] 
**ExchangeDagProtectionPreference** | Pointer to [**ExchangeDAGProtectionPreference**](ExchangeDAGProtectionPreference.md) |  | [optional] 
**ForceRegister** | Pointer to **NullableBool** | ForceRegister is applicable to Physical Environment. By default, the agent running on a physical host will fail the registration, if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster. This is a hidden parameter and should not be documented externally. | [optional] 
**GcpCredentials** | Pointer to [**GcpCredentials**](GcpCredentials.md) |  | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the optional OS type of the Protection Source (such as kWindows or kLinux). overrideDescription: true &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**HyperVType** | Pointer to **NullableString** | Specifies the entity type if the environment is kHyperV. overrideDescription: true | [optional] 
**KubernetesCredentials** | Pointer to [**KubernetesCredentials**](KubernetesCredentials.md) |  | [optional] 
**KubernetesType** | Pointer to **NullableString** | Specifies the entity type if the environment is kKubernetes. overrideDescription: true | [optional] 
**KvmType** | Pointer to **NullableString** | Specifies the entity type if the environment is kKVM. overrideDescription: true | [optional] 
**NasMountCredentials** | Pointer to [**NullableNasMountCredentialParams**](NasMountCredentialParams.md) | Specifies the server credentials to connect to a NetApp server. This field is required for mounting SMB volumes on NetApp servers. | [optional] 
**NetappType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kCluster,&#39; if the environment is kNetapp. | [optional] 
**NimbleType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kStorageArray&#39; if the environment is kNimble. overrideDescription: true | [optional] 
**Office365CredentialsList** | Pointer to [**[]Office365Credentials**](Office365Credentials.md) | Office365 Source Credentials.  Specifies credentials needed to authenticate &amp; authorize user for Office365 using MS Graph APIs. | [optional] 
**Office365Region** | Pointer to **NullableString** | Specifies the region for Office365. | [optional] 
**Office365Type** | Pointer to **NullableString** | Specifies the entity type such as &#39;kDomain&#39;, &#39;kOutlook&#39;, &#39;kMailbox&#39;, if the environment is kO365. | [optional] 
**Password** | Pointer to **NullableString** | Specifies password of the username to access the target source. | [optional] 
**PhysicalType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kPhysicalHost&#39; if the environment is kPhysical. overrideDescription: true | [optional] 
**PureType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kStorageArray&#39; if the environment is kPure. overrideDescription: true | [optional] 
**ReRegister** | Pointer to **NullableBool** | ReRegister is applicable to Physical Environment. By default, the agent running on a physical host will fail the registration, if it is already registered with the cluster. By setting this option to true, agent can be re-registered with the current cluster. | [optional] 
**SourceSideDedupEnabled** | Pointer to **NullableBool** | This controls whether to use source side dedup on the source or not. This is only applicable to sources which support source side dedup (e.g., Linux physical servers). | [optional] 
**SslVerification** | Pointer to [**SslVerification**](SslVerification.md) |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | Specifies the list of subnet IP addresses and CIDR prefix for enabeling network data transfer. Currently, only Subnet IP and NetbaskBits are valid input fields. All other fields provided as input will be ignored. | [optional] 
**ThrottlingPolicy** | Pointer to [**NullableThrottlingPolicyParameters**](ThrottlingPolicyParameters.md) | Specifies the throttling policy that should be applied to this Source. | [optional] 
**ThrottlingPolicyOverrides** | Pointer to [**[]ThrottlingPolicyOverride**](ThrottlingPolicyOverride.md) | Array of Throttling Policy Overrides for Datastores.  Specifies a list of Throttling Policy for datastores that override the common throttling policy specified for the registered Protection Source. For datastores not in this list, common policy will still apply. | [optional] 
**UseOAuthForExchangeOnline** | Pointer to **NullableBool** | Specifies whether OAuth should be used for authentication in case of Exchange Online. | [optional] 
**Username** | Pointer to **NullableString** | Specifies username to access the target source. | [optional] 
**VlanParams** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 
**VmwareType** | Pointer to **NullableString** | Specifies the entity type such as &#39;kVCenter&#39; if the environment is kKMware. overrideDescription: true | [optional] 

## Methods

### NewRegisterProtectionSourceParameters

`func NewRegisterProtectionSourceParameters() *RegisterProtectionSourceParameters`

NewRegisterProtectionSourceParameters instantiates a new RegisterProtectionSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterProtectionSourceParametersWithDefaults

`func NewRegisterProtectionSourceParametersWithDefaults() *RegisterProtectionSourceParameters`

NewRegisterProtectionSourceParametersWithDefaults instantiates a new RegisterProtectionSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcropolisType

`func (o *RegisterProtectionSourceParameters) GetAcropolisType() string`

GetAcropolisType returns the AcropolisType field if non-nil, zero value otherwise.

### GetAcropolisTypeOk

`func (o *RegisterProtectionSourceParameters) GetAcropolisTypeOk() (*string, bool)`

GetAcropolisTypeOk returns a tuple with the AcropolisType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcropolisType

`func (o *RegisterProtectionSourceParameters) SetAcropolisType(v string)`

SetAcropolisType sets AcropolisType field to given value.

### HasAcropolisType

`func (o *RegisterProtectionSourceParameters) HasAcropolisType() bool`

HasAcropolisType returns a boolean if a field has been set.

### SetAcropolisTypeNil

`func (o *RegisterProtectionSourceParameters) SetAcropolisTypeNil(b bool)`

 SetAcropolisTypeNil sets the value for AcropolisType to be an explicit nil

### UnsetAcropolisType
`func (o *RegisterProtectionSourceParameters) UnsetAcropolisType()`

UnsetAcropolisType ensures that no value is present for AcropolisType, not even an explicit nil
### GetAgentEndpoint

`func (o *RegisterProtectionSourceParameters) GetAgentEndpoint() string`

GetAgentEndpoint returns the AgentEndpoint field if non-nil, zero value otherwise.

### GetAgentEndpointOk

`func (o *RegisterProtectionSourceParameters) GetAgentEndpointOk() (*string, bool)`

GetAgentEndpointOk returns a tuple with the AgentEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEndpoint

`func (o *RegisterProtectionSourceParameters) SetAgentEndpoint(v string)`

SetAgentEndpoint sets AgentEndpoint field to given value.

### HasAgentEndpoint

`func (o *RegisterProtectionSourceParameters) HasAgentEndpoint() bool`

HasAgentEndpoint returns a boolean if a field has been set.

### SetAgentEndpointNil

`func (o *RegisterProtectionSourceParameters) SetAgentEndpointNil(b bool)`

 SetAgentEndpointNil sets the value for AgentEndpoint to be an explicit nil

### UnsetAgentEndpoint
`func (o *RegisterProtectionSourceParameters) UnsetAgentEndpoint()`

UnsetAgentEndpoint ensures that no value is present for AgentEndpoint, not even an explicit nil
### GetAwsCredentials

`func (o *RegisterProtectionSourceParameters) GetAwsCredentials() AwsCredentials`

GetAwsCredentials returns the AwsCredentials field if non-nil, zero value otherwise.

### GetAwsCredentialsOk

`func (o *RegisterProtectionSourceParameters) GetAwsCredentialsOk() (*AwsCredentials, bool)`

GetAwsCredentialsOk returns a tuple with the AwsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentials

`func (o *RegisterProtectionSourceParameters) SetAwsCredentials(v AwsCredentials)`

SetAwsCredentials sets AwsCredentials field to given value.

### HasAwsCredentials

`func (o *RegisterProtectionSourceParameters) HasAwsCredentials() bool`

HasAwsCredentials returns a boolean if a field has been set.

### GetAwsFleetParams

`func (o *RegisterProtectionSourceParameters) GetAwsFleetParams() AwsFleetParams`

GetAwsFleetParams returns the AwsFleetParams field if non-nil, zero value otherwise.

### GetAwsFleetParamsOk

`func (o *RegisterProtectionSourceParameters) GetAwsFleetParamsOk() (*AwsFleetParams, bool)`

GetAwsFleetParamsOk returns a tuple with the AwsFleetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsFleetParams

`func (o *RegisterProtectionSourceParameters) SetAwsFleetParams(v AwsFleetParams)`

SetAwsFleetParams sets AwsFleetParams field to given value.

### HasAwsFleetParams

`func (o *RegisterProtectionSourceParameters) HasAwsFleetParams() bool`

HasAwsFleetParams returns a boolean if a field has been set.

### GetAzureCredentials

`func (o *RegisterProtectionSourceParameters) GetAzureCredentials() AzureCredentials`

GetAzureCredentials returns the AzureCredentials field if non-nil, zero value otherwise.

### GetAzureCredentialsOk

`func (o *RegisterProtectionSourceParameters) GetAzureCredentialsOk() (*AzureCredentials, bool)`

GetAzureCredentialsOk returns a tuple with the AzureCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentials

`func (o *RegisterProtectionSourceParameters) SetAzureCredentials(v AzureCredentials)`

SetAzureCredentials sets AzureCredentials field to given value.

### HasAzureCredentials

`func (o *RegisterProtectionSourceParameters) HasAzureCredentials() bool`

HasAzureCredentials returns a boolean if a field has been set.

### GetBlacklistedIpAddresses

`func (o *RegisterProtectionSourceParameters) GetBlacklistedIpAddresses() []string`

GetBlacklistedIpAddresses returns the BlacklistedIpAddresses field if non-nil, zero value otherwise.

### GetBlacklistedIpAddressesOk

`func (o *RegisterProtectionSourceParameters) GetBlacklistedIpAddressesOk() (*[]string, bool)`

GetBlacklistedIpAddressesOk returns a tuple with the BlacklistedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIpAddresses

`func (o *RegisterProtectionSourceParameters) SetBlacklistedIpAddresses(v []string)`

SetBlacklistedIpAddresses sets BlacklistedIpAddresses field to given value.

### HasBlacklistedIpAddresses

`func (o *RegisterProtectionSourceParameters) HasBlacklistedIpAddresses() bool`

HasBlacklistedIpAddresses returns a boolean if a field has been set.

### SetBlacklistedIpAddressesNil

`func (o *RegisterProtectionSourceParameters) SetBlacklistedIpAddressesNil(b bool)`

 SetBlacklistedIpAddressesNil sets the value for BlacklistedIpAddresses to be an explicit nil

### UnsetBlacklistedIpAddresses
`func (o *RegisterProtectionSourceParameters) UnsetBlacklistedIpAddresses()`

UnsetBlacklistedIpAddresses ensures that no value is present for BlacklistedIpAddresses, not even an explicit nil
### GetClusterNetworkInfo

`func (o *RegisterProtectionSourceParameters) GetClusterNetworkInfo() FleetNetworkParams`

GetClusterNetworkInfo returns the ClusterNetworkInfo field if non-nil, zero value otherwise.

### GetClusterNetworkInfoOk

`func (o *RegisterProtectionSourceParameters) GetClusterNetworkInfoOk() (*FleetNetworkParams, bool)`

GetClusterNetworkInfoOk returns a tuple with the ClusterNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetworkInfo

`func (o *RegisterProtectionSourceParameters) SetClusterNetworkInfo(v FleetNetworkParams)`

SetClusterNetworkInfo sets ClusterNetworkInfo field to given value.

### HasClusterNetworkInfo

`func (o *RegisterProtectionSourceParameters) HasClusterNetworkInfo() bool`

HasClusterNetworkInfo returns a boolean if a field has been set.

### GetEndpoint

`func (o *RegisterProtectionSourceParameters) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *RegisterProtectionSourceParameters) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *RegisterProtectionSourceParameters) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *RegisterProtectionSourceParameters) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *RegisterProtectionSourceParameters) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *RegisterProtectionSourceParameters) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetEnvironment

`func (o *RegisterProtectionSourceParameters) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RegisterProtectionSourceParameters) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RegisterProtectionSourceParameters) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RegisterProtectionSourceParameters) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *RegisterProtectionSourceParameters) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *RegisterProtectionSourceParameters) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetExchangeDagProtectionPreference

`func (o *RegisterProtectionSourceParameters) GetExchangeDagProtectionPreference() ExchangeDAGProtectionPreference`

GetExchangeDagProtectionPreference returns the ExchangeDagProtectionPreference field if non-nil, zero value otherwise.

### GetExchangeDagProtectionPreferenceOk

`func (o *RegisterProtectionSourceParameters) GetExchangeDagProtectionPreferenceOk() (*ExchangeDAGProtectionPreference, bool)`

GetExchangeDagProtectionPreferenceOk returns a tuple with the ExchangeDagProtectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDagProtectionPreference

`func (o *RegisterProtectionSourceParameters) SetExchangeDagProtectionPreference(v ExchangeDAGProtectionPreference)`

SetExchangeDagProtectionPreference sets ExchangeDagProtectionPreference field to given value.

### HasExchangeDagProtectionPreference

`func (o *RegisterProtectionSourceParameters) HasExchangeDagProtectionPreference() bool`

HasExchangeDagProtectionPreference returns a boolean if a field has been set.

### GetForceRegister

`func (o *RegisterProtectionSourceParameters) GetForceRegister() bool`

GetForceRegister returns the ForceRegister field if non-nil, zero value otherwise.

### GetForceRegisterOk

`func (o *RegisterProtectionSourceParameters) GetForceRegisterOk() (*bool, bool)`

GetForceRegisterOk returns a tuple with the ForceRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRegister

`func (o *RegisterProtectionSourceParameters) SetForceRegister(v bool)`

SetForceRegister sets ForceRegister field to given value.

### HasForceRegister

`func (o *RegisterProtectionSourceParameters) HasForceRegister() bool`

HasForceRegister returns a boolean if a field has been set.

### SetForceRegisterNil

`func (o *RegisterProtectionSourceParameters) SetForceRegisterNil(b bool)`

 SetForceRegisterNil sets the value for ForceRegister to be an explicit nil

### UnsetForceRegister
`func (o *RegisterProtectionSourceParameters) UnsetForceRegister()`

UnsetForceRegister ensures that no value is present for ForceRegister, not even an explicit nil
### GetGcpCredentials

`func (o *RegisterProtectionSourceParameters) GetGcpCredentials() GcpCredentials`

GetGcpCredentials returns the GcpCredentials field if non-nil, zero value otherwise.

### GetGcpCredentialsOk

`func (o *RegisterProtectionSourceParameters) GetGcpCredentialsOk() (*GcpCredentials, bool)`

GetGcpCredentialsOk returns a tuple with the GcpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentials

`func (o *RegisterProtectionSourceParameters) SetGcpCredentials(v GcpCredentials)`

SetGcpCredentials sets GcpCredentials field to given value.

### HasGcpCredentials

`func (o *RegisterProtectionSourceParameters) HasGcpCredentials() bool`

HasGcpCredentials returns a boolean if a field has been set.

### GetHostType

`func (o *RegisterProtectionSourceParameters) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *RegisterProtectionSourceParameters) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *RegisterProtectionSourceParameters) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *RegisterProtectionSourceParameters) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *RegisterProtectionSourceParameters) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *RegisterProtectionSourceParameters) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetHyperVType

`func (o *RegisterProtectionSourceParameters) GetHyperVType() string`

GetHyperVType returns the HyperVType field if non-nil, zero value otherwise.

### GetHyperVTypeOk

`func (o *RegisterProtectionSourceParameters) GetHyperVTypeOk() (*string, bool)`

GetHyperVTypeOk returns a tuple with the HyperVType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperVType

`func (o *RegisterProtectionSourceParameters) SetHyperVType(v string)`

SetHyperVType sets HyperVType field to given value.

### HasHyperVType

`func (o *RegisterProtectionSourceParameters) HasHyperVType() bool`

HasHyperVType returns a boolean if a field has been set.

### SetHyperVTypeNil

`func (o *RegisterProtectionSourceParameters) SetHyperVTypeNil(b bool)`

 SetHyperVTypeNil sets the value for HyperVType to be an explicit nil

### UnsetHyperVType
`func (o *RegisterProtectionSourceParameters) UnsetHyperVType()`

UnsetHyperVType ensures that no value is present for HyperVType, not even an explicit nil
### GetKubernetesCredentials

`func (o *RegisterProtectionSourceParameters) GetKubernetesCredentials() KubernetesCredentials`

GetKubernetesCredentials returns the KubernetesCredentials field if non-nil, zero value otherwise.

### GetKubernetesCredentialsOk

`func (o *RegisterProtectionSourceParameters) GetKubernetesCredentialsOk() (*KubernetesCredentials, bool)`

GetKubernetesCredentialsOk returns a tuple with the KubernetesCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCredentials

`func (o *RegisterProtectionSourceParameters) SetKubernetesCredentials(v KubernetesCredentials)`

SetKubernetesCredentials sets KubernetesCredentials field to given value.

### HasKubernetesCredentials

`func (o *RegisterProtectionSourceParameters) HasKubernetesCredentials() bool`

HasKubernetesCredentials returns a boolean if a field has been set.

### GetKubernetesType

`func (o *RegisterProtectionSourceParameters) GetKubernetesType() string`

GetKubernetesType returns the KubernetesType field if non-nil, zero value otherwise.

### GetKubernetesTypeOk

`func (o *RegisterProtectionSourceParameters) GetKubernetesTypeOk() (*string, bool)`

GetKubernetesTypeOk returns a tuple with the KubernetesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesType

`func (o *RegisterProtectionSourceParameters) SetKubernetesType(v string)`

SetKubernetesType sets KubernetesType field to given value.

### HasKubernetesType

`func (o *RegisterProtectionSourceParameters) HasKubernetesType() bool`

HasKubernetesType returns a boolean if a field has been set.

### SetKubernetesTypeNil

`func (o *RegisterProtectionSourceParameters) SetKubernetesTypeNil(b bool)`

 SetKubernetesTypeNil sets the value for KubernetesType to be an explicit nil

### UnsetKubernetesType
`func (o *RegisterProtectionSourceParameters) UnsetKubernetesType()`

UnsetKubernetesType ensures that no value is present for KubernetesType, not even an explicit nil
### GetKvmType

`func (o *RegisterProtectionSourceParameters) GetKvmType() string`

GetKvmType returns the KvmType field if non-nil, zero value otherwise.

### GetKvmTypeOk

`func (o *RegisterProtectionSourceParameters) GetKvmTypeOk() (*string, bool)`

GetKvmTypeOk returns a tuple with the KvmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmType

`func (o *RegisterProtectionSourceParameters) SetKvmType(v string)`

SetKvmType sets KvmType field to given value.

### HasKvmType

`func (o *RegisterProtectionSourceParameters) HasKvmType() bool`

HasKvmType returns a boolean if a field has been set.

### SetKvmTypeNil

`func (o *RegisterProtectionSourceParameters) SetKvmTypeNil(b bool)`

 SetKvmTypeNil sets the value for KvmType to be an explicit nil

### UnsetKvmType
`func (o *RegisterProtectionSourceParameters) UnsetKvmType()`

UnsetKvmType ensures that no value is present for KvmType, not even an explicit nil
### GetNasMountCredentials

`func (o *RegisterProtectionSourceParameters) GetNasMountCredentials() NasMountCredentialParams`

GetNasMountCredentials returns the NasMountCredentials field if non-nil, zero value otherwise.

### GetNasMountCredentialsOk

`func (o *RegisterProtectionSourceParameters) GetNasMountCredentialsOk() (*NasMountCredentialParams, bool)`

GetNasMountCredentialsOk returns a tuple with the NasMountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasMountCredentials

`func (o *RegisterProtectionSourceParameters) SetNasMountCredentials(v NasMountCredentialParams)`

SetNasMountCredentials sets NasMountCredentials field to given value.

### HasNasMountCredentials

`func (o *RegisterProtectionSourceParameters) HasNasMountCredentials() bool`

HasNasMountCredentials returns a boolean if a field has been set.

### SetNasMountCredentialsNil

`func (o *RegisterProtectionSourceParameters) SetNasMountCredentialsNil(b bool)`

 SetNasMountCredentialsNil sets the value for NasMountCredentials to be an explicit nil

### UnsetNasMountCredentials
`func (o *RegisterProtectionSourceParameters) UnsetNasMountCredentials()`

UnsetNasMountCredentials ensures that no value is present for NasMountCredentials, not even an explicit nil
### GetNetappType

`func (o *RegisterProtectionSourceParameters) GetNetappType() string`

GetNetappType returns the NetappType field if non-nil, zero value otherwise.

### GetNetappTypeOk

`func (o *RegisterProtectionSourceParameters) GetNetappTypeOk() (*string, bool)`

GetNetappTypeOk returns a tuple with the NetappType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetappType

`func (o *RegisterProtectionSourceParameters) SetNetappType(v string)`

SetNetappType sets NetappType field to given value.

### HasNetappType

`func (o *RegisterProtectionSourceParameters) HasNetappType() bool`

HasNetappType returns a boolean if a field has been set.

### SetNetappTypeNil

`func (o *RegisterProtectionSourceParameters) SetNetappTypeNil(b bool)`

 SetNetappTypeNil sets the value for NetappType to be an explicit nil

### UnsetNetappType
`func (o *RegisterProtectionSourceParameters) UnsetNetappType()`

UnsetNetappType ensures that no value is present for NetappType, not even an explicit nil
### GetNimbleType

`func (o *RegisterProtectionSourceParameters) GetNimbleType() string`

GetNimbleType returns the NimbleType field if non-nil, zero value otherwise.

### GetNimbleTypeOk

`func (o *RegisterProtectionSourceParameters) GetNimbleTypeOk() (*string, bool)`

GetNimbleTypeOk returns a tuple with the NimbleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNimbleType

`func (o *RegisterProtectionSourceParameters) SetNimbleType(v string)`

SetNimbleType sets NimbleType field to given value.

### HasNimbleType

`func (o *RegisterProtectionSourceParameters) HasNimbleType() bool`

HasNimbleType returns a boolean if a field has been set.

### SetNimbleTypeNil

`func (o *RegisterProtectionSourceParameters) SetNimbleTypeNil(b bool)`

 SetNimbleTypeNil sets the value for NimbleType to be an explicit nil

### UnsetNimbleType
`func (o *RegisterProtectionSourceParameters) UnsetNimbleType()`

UnsetNimbleType ensures that no value is present for NimbleType, not even an explicit nil
### GetOffice365CredentialsList

`func (o *RegisterProtectionSourceParameters) GetOffice365CredentialsList() []Office365Credentials`

GetOffice365CredentialsList returns the Office365CredentialsList field if non-nil, zero value otherwise.

### GetOffice365CredentialsListOk

`func (o *RegisterProtectionSourceParameters) GetOffice365CredentialsListOk() (*[]Office365Credentials, bool)`

GetOffice365CredentialsListOk returns a tuple with the Office365CredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365CredentialsList

`func (o *RegisterProtectionSourceParameters) SetOffice365CredentialsList(v []Office365Credentials)`

SetOffice365CredentialsList sets Office365CredentialsList field to given value.

### HasOffice365CredentialsList

`func (o *RegisterProtectionSourceParameters) HasOffice365CredentialsList() bool`

HasOffice365CredentialsList returns a boolean if a field has been set.

### SetOffice365CredentialsListNil

`func (o *RegisterProtectionSourceParameters) SetOffice365CredentialsListNil(b bool)`

 SetOffice365CredentialsListNil sets the value for Office365CredentialsList to be an explicit nil

### UnsetOffice365CredentialsList
`func (o *RegisterProtectionSourceParameters) UnsetOffice365CredentialsList()`

UnsetOffice365CredentialsList ensures that no value is present for Office365CredentialsList, not even an explicit nil
### GetOffice365Region

`func (o *RegisterProtectionSourceParameters) GetOffice365Region() string`

GetOffice365Region returns the Office365Region field if non-nil, zero value otherwise.

### GetOffice365RegionOk

`func (o *RegisterProtectionSourceParameters) GetOffice365RegionOk() (*string, bool)`

GetOffice365RegionOk returns a tuple with the Office365Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Region

`func (o *RegisterProtectionSourceParameters) SetOffice365Region(v string)`

SetOffice365Region sets Office365Region field to given value.

### HasOffice365Region

`func (o *RegisterProtectionSourceParameters) HasOffice365Region() bool`

HasOffice365Region returns a boolean if a field has been set.

### SetOffice365RegionNil

`func (o *RegisterProtectionSourceParameters) SetOffice365RegionNil(b bool)`

 SetOffice365RegionNil sets the value for Office365Region to be an explicit nil

### UnsetOffice365Region
`func (o *RegisterProtectionSourceParameters) UnsetOffice365Region()`

UnsetOffice365Region ensures that no value is present for Office365Region, not even an explicit nil
### GetOffice365Type

`func (o *RegisterProtectionSourceParameters) GetOffice365Type() string`

GetOffice365Type returns the Office365Type field if non-nil, zero value otherwise.

### GetOffice365TypeOk

`func (o *RegisterProtectionSourceParameters) GetOffice365TypeOk() (*string, bool)`

GetOffice365TypeOk returns a tuple with the Office365Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Type

`func (o *RegisterProtectionSourceParameters) SetOffice365Type(v string)`

SetOffice365Type sets Office365Type field to given value.

### HasOffice365Type

`func (o *RegisterProtectionSourceParameters) HasOffice365Type() bool`

HasOffice365Type returns a boolean if a field has been set.

### SetOffice365TypeNil

`func (o *RegisterProtectionSourceParameters) SetOffice365TypeNil(b bool)`

 SetOffice365TypeNil sets the value for Office365Type to be an explicit nil

### UnsetOffice365Type
`func (o *RegisterProtectionSourceParameters) UnsetOffice365Type()`

UnsetOffice365Type ensures that no value is present for Office365Type, not even an explicit nil
### GetPassword

`func (o *RegisterProtectionSourceParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterProtectionSourceParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterProtectionSourceParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterProtectionSourceParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RegisterProtectionSourceParameters) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisterProtectionSourceParameters) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPhysicalType

`func (o *RegisterProtectionSourceParameters) GetPhysicalType() string`

GetPhysicalType returns the PhysicalType field if non-nil, zero value otherwise.

### GetPhysicalTypeOk

`func (o *RegisterProtectionSourceParameters) GetPhysicalTypeOk() (*string, bool)`

GetPhysicalTypeOk returns a tuple with the PhysicalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalType

`func (o *RegisterProtectionSourceParameters) SetPhysicalType(v string)`

SetPhysicalType sets PhysicalType field to given value.

### HasPhysicalType

`func (o *RegisterProtectionSourceParameters) HasPhysicalType() bool`

HasPhysicalType returns a boolean if a field has been set.

### SetPhysicalTypeNil

`func (o *RegisterProtectionSourceParameters) SetPhysicalTypeNil(b bool)`

 SetPhysicalTypeNil sets the value for PhysicalType to be an explicit nil

### UnsetPhysicalType
`func (o *RegisterProtectionSourceParameters) UnsetPhysicalType()`

UnsetPhysicalType ensures that no value is present for PhysicalType, not even an explicit nil
### GetPureType

`func (o *RegisterProtectionSourceParameters) GetPureType() string`

GetPureType returns the PureType field if non-nil, zero value otherwise.

### GetPureTypeOk

`func (o *RegisterProtectionSourceParameters) GetPureTypeOk() (*string, bool)`

GetPureTypeOk returns a tuple with the PureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPureType

`func (o *RegisterProtectionSourceParameters) SetPureType(v string)`

SetPureType sets PureType field to given value.

### HasPureType

`func (o *RegisterProtectionSourceParameters) HasPureType() bool`

HasPureType returns a boolean if a field has been set.

### SetPureTypeNil

`func (o *RegisterProtectionSourceParameters) SetPureTypeNil(b bool)`

 SetPureTypeNil sets the value for PureType to be an explicit nil

### UnsetPureType
`func (o *RegisterProtectionSourceParameters) UnsetPureType()`

UnsetPureType ensures that no value is present for PureType, not even an explicit nil
### GetReRegister

`func (o *RegisterProtectionSourceParameters) GetReRegister() bool`

GetReRegister returns the ReRegister field if non-nil, zero value otherwise.

### GetReRegisterOk

`func (o *RegisterProtectionSourceParameters) GetReRegisterOk() (*bool, bool)`

GetReRegisterOk returns a tuple with the ReRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegister

`func (o *RegisterProtectionSourceParameters) SetReRegister(v bool)`

SetReRegister sets ReRegister field to given value.

### HasReRegister

`func (o *RegisterProtectionSourceParameters) HasReRegister() bool`

HasReRegister returns a boolean if a field has been set.

### SetReRegisterNil

`func (o *RegisterProtectionSourceParameters) SetReRegisterNil(b bool)`

 SetReRegisterNil sets the value for ReRegister to be an explicit nil

### UnsetReRegister
`func (o *RegisterProtectionSourceParameters) UnsetReRegister()`

UnsetReRegister ensures that no value is present for ReRegister, not even an explicit nil
### GetSourceSideDedupEnabled

`func (o *RegisterProtectionSourceParameters) GetSourceSideDedupEnabled() bool`

GetSourceSideDedupEnabled returns the SourceSideDedupEnabled field if non-nil, zero value otherwise.

### GetSourceSideDedupEnabledOk

`func (o *RegisterProtectionSourceParameters) GetSourceSideDedupEnabledOk() (*bool, bool)`

GetSourceSideDedupEnabledOk returns a tuple with the SourceSideDedupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDedupEnabled

`func (o *RegisterProtectionSourceParameters) SetSourceSideDedupEnabled(v bool)`

SetSourceSideDedupEnabled sets SourceSideDedupEnabled field to given value.

### HasSourceSideDedupEnabled

`func (o *RegisterProtectionSourceParameters) HasSourceSideDedupEnabled() bool`

HasSourceSideDedupEnabled returns a boolean if a field has been set.

### SetSourceSideDedupEnabledNil

`func (o *RegisterProtectionSourceParameters) SetSourceSideDedupEnabledNil(b bool)`

 SetSourceSideDedupEnabledNil sets the value for SourceSideDedupEnabled to be an explicit nil

### UnsetSourceSideDedupEnabled
`func (o *RegisterProtectionSourceParameters) UnsetSourceSideDedupEnabled()`

UnsetSourceSideDedupEnabled ensures that no value is present for SourceSideDedupEnabled, not even an explicit nil
### GetSslVerification

`func (o *RegisterProtectionSourceParameters) GetSslVerification() SslVerification`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *RegisterProtectionSourceParameters) GetSslVerificationOk() (*SslVerification, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *RegisterProtectionSourceParameters) SetSslVerification(v SslVerification)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *RegisterProtectionSourceParameters) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetSubnets

`func (o *RegisterProtectionSourceParameters) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *RegisterProtectionSourceParameters) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *RegisterProtectionSourceParameters) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *RegisterProtectionSourceParameters) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *RegisterProtectionSourceParameters) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *RegisterProtectionSourceParameters) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetThrottlingPolicy

`func (o *RegisterProtectionSourceParameters) GetThrottlingPolicy() ThrottlingPolicyParameters`

GetThrottlingPolicy returns the ThrottlingPolicy field if non-nil, zero value otherwise.

### GetThrottlingPolicyOk

`func (o *RegisterProtectionSourceParameters) GetThrottlingPolicyOk() (*ThrottlingPolicyParameters, bool)`

GetThrottlingPolicyOk returns a tuple with the ThrottlingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicy

`func (o *RegisterProtectionSourceParameters) SetThrottlingPolicy(v ThrottlingPolicyParameters)`

SetThrottlingPolicy sets ThrottlingPolicy field to given value.

### HasThrottlingPolicy

`func (o *RegisterProtectionSourceParameters) HasThrottlingPolicy() bool`

HasThrottlingPolicy returns a boolean if a field has been set.

### SetThrottlingPolicyNil

`func (o *RegisterProtectionSourceParameters) SetThrottlingPolicyNil(b bool)`

 SetThrottlingPolicyNil sets the value for ThrottlingPolicy to be an explicit nil

### UnsetThrottlingPolicy
`func (o *RegisterProtectionSourceParameters) UnsetThrottlingPolicy()`

UnsetThrottlingPolicy ensures that no value is present for ThrottlingPolicy, not even an explicit nil
### GetThrottlingPolicyOverrides

`func (o *RegisterProtectionSourceParameters) GetThrottlingPolicyOverrides() []ThrottlingPolicyOverride`

GetThrottlingPolicyOverrides returns the ThrottlingPolicyOverrides field if non-nil, zero value otherwise.

### GetThrottlingPolicyOverridesOk

`func (o *RegisterProtectionSourceParameters) GetThrottlingPolicyOverridesOk() (*[]ThrottlingPolicyOverride, bool)`

GetThrottlingPolicyOverridesOk returns a tuple with the ThrottlingPolicyOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicyOverrides

`func (o *RegisterProtectionSourceParameters) SetThrottlingPolicyOverrides(v []ThrottlingPolicyOverride)`

SetThrottlingPolicyOverrides sets ThrottlingPolicyOverrides field to given value.

### HasThrottlingPolicyOverrides

`func (o *RegisterProtectionSourceParameters) HasThrottlingPolicyOverrides() bool`

HasThrottlingPolicyOverrides returns a boolean if a field has been set.

### SetThrottlingPolicyOverridesNil

`func (o *RegisterProtectionSourceParameters) SetThrottlingPolicyOverridesNil(b bool)`

 SetThrottlingPolicyOverridesNil sets the value for ThrottlingPolicyOverrides to be an explicit nil

### UnsetThrottlingPolicyOverrides
`func (o *RegisterProtectionSourceParameters) UnsetThrottlingPolicyOverrides()`

UnsetThrottlingPolicyOverrides ensures that no value is present for ThrottlingPolicyOverrides, not even an explicit nil
### GetUseOAuthForExchangeOnline

`func (o *RegisterProtectionSourceParameters) GetUseOAuthForExchangeOnline() bool`

GetUseOAuthForExchangeOnline returns the UseOAuthForExchangeOnline field if non-nil, zero value otherwise.

### GetUseOAuthForExchangeOnlineOk

`func (o *RegisterProtectionSourceParameters) GetUseOAuthForExchangeOnlineOk() (*bool, bool)`

GetUseOAuthForExchangeOnlineOk returns a tuple with the UseOAuthForExchangeOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOAuthForExchangeOnline

`func (o *RegisterProtectionSourceParameters) SetUseOAuthForExchangeOnline(v bool)`

SetUseOAuthForExchangeOnline sets UseOAuthForExchangeOnline field to given value.

### HasUseOAuthForExchangeOnline

`func (o *RegisterProtectionSourceParameters) HasUseOAuthForExchangeOnline() bool`

HasUseOAuthForExchangeOnline returns a boolean if a field has been set.

### SetUseOAuthForExchangeOnlineNil

`func (o *RegisterProtectionSourceParameters) SetUseOAuthForExchangeOnlineNil(b bool)`

 SetUseOAuthForExchangeOnlineNil sets the value for UseOAuthForExchangeOnline to be an explicit nil

### UnsetUseOAuthForExchangeOnline
`func (o *RegisterProtectionSourceParameters) UnsetUseOAuthForExchangeOnline()`

UnsetUseOAuthForExchangeOnline ensures that no value is present for UseOAuthForExchangeOnline, not even an explicit nil
### GetUsername

`func (o *RegisterProtectionSourceParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterProtectionSourceParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterProtectionSourceParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisterProtectionSourceParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RegisterProtectionSourceParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RegisterProtectionSourceParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetVlanParams

`func (o *RegisterProtectionSourceParameters) GetVlanParams() VlanParameters`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *RegisterProtectionSourceParameters) GetVlanParamsOk() (*VlanParameters, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *RegisterProtectionSourceParameters) SetVlanParams(v VlanParameters)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *RegisterProtectionSourceParameters) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.

### GetVmwareType

`func (o *RegisterProtectionSourceParameters) GetVmwareType() string`

GetVmwareType returns the VmwareType field if non-nil, zero value otherwise.

### GetVmwareTypeOk

`func (o *RegisterProtectionSourceParameters) GetVmwareTypeOk() (*string, bool)`

GetVmwareTypeOk returns a tuple with the VmwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareType

`func (o *RegisterProtectionSourceParameters) SetVmwareType(v string)`

SetVmwareType sets VmwareType field to given value.

### HasVmwareType

`func (o *RegisterProtectionSourceParameters) HasVmwareType() bool`

HasVmwareType returns a boolean if a field has been set.

### SetVmwareTypeNil

`func (o *RegisterProtectionSourceParameters) SetVmwareTypeNil(b bool)`

 SetVmwareTypeNil sets the value for VmwareType to be an explicit nil

### UnsetVmwareType
`func (o *RegisterProtectionSourceParameters) UnsetVmwareType()`

UnsetVmwareType ensures that no value is present for VmwareType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


