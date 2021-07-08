# UpdateProtectionSourceParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentEndpoint** | Pointer to **NullableString** | Specifies the agent endpoint if it is different from the source endpoint. | [optional] 
**AwsCredentials** | Pointer to [**AwsCredentials**](AwsCredentials.md) |  | [optional] 
**AwsFleetParams** | Pointer to [**AwsFleetParams**](AwsFleetParams.md) |  | [optional] 
**AzureCredentials** | Pointer to [**AzureCredentials**](AzureCredentials.md) |  | [optional] 
**BlacklistedIpAddresses** | Pointer to **[]string** | Specifies the list of IP Addresses on the registered source to be blacklisted for doing any type of IO operations. | [optional] 
**ClusterNetworkInfo** | Pointer to [**FleetNetworkParams**](FleetNetworkParams.md) |  | [optional] 
**Endpoint** | Pointer to **NullableString** | Specifies the network endpoint of the Protection Source where it is reachable. It could be an URL or hostname or an IP address of the Protection Source. | [optional] 
**ExchangeDagProtectionPreference** | Pointer to [**ExchangeDAGProtectionPreference**](ExchangeDAGProtectionPreference.md) |  | [optional] 
**ForceRegister** | Pointer to **NullableBool** | ForceRegister is applicable to Physical Environment. By default, the agent running on a physical host will fail the registration, if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster. This is a hidden parameter and should not be documented externally. | [optional] 
**GcpCredentials** | Pointer to [**GcpCredentials**](GcpCredentials.md) |  | [optional] 
**HostType** | Pointer to **NullableString** | Specifies the optional OS type of the Protection Source (such as kWindows or kLinux). overrideDescription: true &#39;kLinux&#39; indicates the Linux operating system. &#39;kWindows&#39; indicates the Microsoft Windows operating system. &#39;kAix&#39; indicates the IBM AIX operating system. &#39;kSolaris&#39; indicates the Oracle Solaris operating system. &#39;kSapHana&#39; indicates the Sap Hana database system developed by SAP SE. &#39;kOther&#39; indicates the other types of operating system. | [optional] 
**KubernetesCredentials** | Pointer to [**KubernetesCredentials**](KubernetesCredentials.md) |  | [optional] 
**MinimumFreeSpaceGB** | Pointer to **NullableInt64** | Specifies the minimum space in GB after which backup jobs will be canceled due to low space. | [optional] 
**NasMountCredentials** | Pointer to [**NullableNasMountCredentialParams**](NasMountCredentialParams.md) | Specifies the server credentials to connect to a NetApp server. This field is required for mounting SMB volumes on NetApp servers. | [optional] 
**Office365CredentialsList** | Pointer to [**[]Office365Credentials**](Office365Credentials.md) | Office365 Source Credentials.  Specifies credentials needed to authenticate &amp; authorize user for Office365 using MS Graph APIs. | [optional] 
**Office365Region** | Pointer to **NullableString** | Specifies the region for Office365. | [optional] 
**Password** | Pointer to **NullableString** | Specifies password of the username to access the target source. | [optional] 
**ReRegister** | Pointer to **NullableBool** | ReRegister is applicable to Physical Environment. By default, the agent running on a physical host will fail the registration, if it is already registered with the cluster. By setting this option to true, agent can be re-registered with the current cluster. | [optional] 
**SourceSideDedupEnabled** | Pointer to **NullableBool** | This controls whether to use source side dedup on the source or not. This is only applicable to sources which support source side dedup (e.g., Linux physical servers). | [optional] 
**SslVerification** | Pointer to [**SslVerification**](SslVerification.md) |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | Specifies the list of subnet IP addresses and CIDR prefix for enabeling network data transfer. Currently, only Subnet IP and NetbaskBits are valid input fields. All other fields provided as input will be ignored. | [optional] 
**ThrottlingPolicy** | Pointer to [**NullableThrottlingPolicyParameters**](ThrottlingPolicyParameters.md) | Specifies the throttling policy that should be applied to this Source. | [optional] 
**ThrottlingPolicyOverrides** | Pointer to [**[]ThrottlingPolicyOverride**](ThrottlingPolicyOverride.md) | Array of Throttling Policy Overrides for Datastores.  Specifies a list of Throttling Policy for datastores that override the common throttling policy specified for the registered Protection Source. For datastores not in this list, common policy will still apply. | [optional] 
**UseOAuthForExchangeOnline** | Pointer to **NullableBool** | Specifies whether OAuth should be used for authentication in case of Exchange Online. | [optional] 
**Username** | Pointer to **NullableString** | Specifies username to access the target source. | [optional] 
**VlanParams** | Pointer to [**VlanParameters**](VlanParameters.md) |  | [optional] 

## Methods

### NewUpdateProtectionSourceParameters

`func NewUpdateProtectionSourceParameters() *UpdateProtectionSourceParameters`

NewUpdateProtectionSourceParameters instantiates a new UpdateProtectionSourceParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProtectionSourceParametersWithDefaults

`func NewUpdateProtectionSourceParametersWithDefaults() *UpdateProtectionSourceParameters`

NewUpdateProtectionSourceParametersWithDefaults instantiates a new UpdateProtectionSourceParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentEndpoint

`func (o *UpdateProtectionSourceParameters) GetAgentEndpoint() string`

GetAgentEndpoint returns the AgentEndpoint field if non-nil, zero value otherwise.

### GetAgentEndpointOk

`func (o *UpdateProtectionSourceParameters) GetAgentEndpointOk() (*string, bool)`

GetAgentEndpointOk returns a tuple with the AgentEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEndpoint

`func (o *UpdateProtectionSourceParameters) SetAgentEndpoint(v string)`

SetAgentEndpoint sets AgentEndpoint field to given value.

### HasAgentEndpoint

`func (o *UpdateProtectionSourceParameters) HasAgentEndpoint() bool`

HasAgentEndpoint returns a boolean if a field has been set.

### SetAgentEndpointNil

`func (o *UpdateProtectionSourceParameters) SetAgentEndpointNil(b bool)`

 SetAgentEndpointNil sets the value for AgentEndpoint to be an explicit nil

### UnsetAgentEndpoint
`func (o *UpdateProtectionSourceParameters) UnsetAgentEndpoint()`

UnsetAgentEndpoint ensures that no value is present for AgentEndpoint, not even an explicit nil
### GetAwsCredentials

`func (o *UpdateProtectionSourceParameters) GetAwsCredentials() AwsCredentials`

GetAwsCredentials returns the AwsCredentials field if non-nil, zero value otherwise.

### GetAwsCredentialsOk

`func (o *UpdateProtectionSourceParameters) GetAwsCredentialsOk() (*AwsCredentials, bool)`

GetAwsCredentialsOk returns a tuple with the AwsCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCredentials

`func (o *UpdateProtectionSourceParameters) SetAwsCredentials(v AwsCredentials)`

SetAwsCredentials sets AwsCredentials field to given value.

### HasAwsCredentials

`func (o *UpdateProtectionSourceParameters) HasAwsCredentials() bool`

HasAwsCredentials returns a boolean if a field has been set.

### GetAwsFleetParams

`func (o *UpdateProtectionSourceParameters) GetAwsFleetParams() AwsFleetParams`

GetAwsFleetParams returns the AwsFleetParams field if non-nil, zero value otherwise.

### GetAwsFleetParamsOk

`func (o *UpdateProtectionSourceParameters) GetAwsFleetParamsOk() (*AwsFleetParams, bool)`

GetAwsFleetParamsOk returns a tuple with the AwsFleetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsFleetParams

`func (o *UpdateProtectionSourceParameters) SetAwsFleetParams(v AwsFleetParams)`

SetAwsFleetParams sets AwsFleetParams field to given value.

### HasAwsFleetParams

`func (o *UpdateProtectionSourceParameters) HasAwsFleetParams() bool`

HasAwsFleetParams returns a boolean if a field has been set.

### GetAzureCredentials

`func (o *UpdateProtectionSourceParameters) GetAzureCredentials() AzureCredentials`

GetAzureCredentials returns the AzureCredentials field if non-nil, zero value otherwise.

### GetAzureCredentialsOk

`func (o *UpdateProtectionSourceParameters) GetAzureCredentialsOk() (*AzureCredentials, bool)`

GetAzureCredentialsOk returns a tuple with the AzureCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCredentials

`func (o *UpdateProtectionSourceParameters) SetAzureCredentials(v AzureCredentials)`

SetAzureCredentials sets AzureCredentials field to given value.

### HasAzureCredentials

`func (o *UpdateProtectionSourceParameters) HasAzureCredentials() bool`

HasAzureCredentials returns a boolean if a field has been set.

### GetBlacklistedIpAddresses

`func (o *UpdateProtectionSourceParameters) GetBlacklistedIpAddresses() []string`

GetBlacklistedIpAddresses returns the BlacklistedIpAddresses field if non-nil, zero value otherwise.

### GetBlacklistedIpAddressesOk

`func (o *UpdateProtectionSourceParameters) GetBlacklistedIpAddressesOk() (*[]string, bool)`

GetBlacklistedIpAddressesOk returns a tuple with the BlacklistedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedIpAddresses

`func (o *UpdateProtectionSourceParameters) SetBlacklistedIpAddresses(v []string)`

SetBlacklistedIpAddresses sets BlacklistedIpAddresses field to given value.

### HasBlacklistedIpAddresses

`func (o *UpdateProtectionSourceParameters) HasBlacklistedIpAddresses() bool`

HasBlacklistedIpAddresses returns a boolean if a field has been set.

### SetBlacklistedIpAddressesNil

`func (o *UpdateProtectionSourceParameters) SetBlacklistedIpAddressesNil(b bool)`

 SetBlacklistedIpAddressesNil sets the value for BlacklistedIpAddresses to be an explicit nil

### UnsetBlacklistedIpAddresses
`func (o *UpdateProtectionSourceParameters) UnsetBlacklistedIpAddresses()`

UnsetBlacklistedIpAddresses ensures that no value is present for BlacklistedIpAddresses, not even an explicit nil
### GetClusterNetworkInfo

`func (o *UpdateProtectionSourceParameters) GetClusterNetworkInfo() FleetNetworkParams`

GetClusterNetworkInfo returns the ClusterNetworkInfo field if non-nil, zero value otherwise.

### GetClusterNetworkInfoOk

`func (o *UpdateProtectionSourceParameters) GetClusterNetworkInfoOk() (*FleetNetworkParams, bool)`

GetClusterNetworkInfoOk returns a tuple with the ClusterNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetworkInfo

`func (o *UpdateProtectionSourceParameters) SetClusterNetworkInfo(v FleetNetworkParams)`

SetClusterNetworkInfo sets ClusterNetworkInfo field to given value.

### HasClusterNetworkInfo

`func (o *UpdateProtectionSourceParameters) HasClusterNetworkInfo() bool`

HasClusterNetworkInfo returns a boolean if a field has been set.

### GetEndpoint

`func (o *UpdateProtectionSourceParameters) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateProtectionSourceParameters) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateProtectionSourceParameters) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UpdateProtectionSourceParameters) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *UpdateProtectionSourceParameters) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *UpdateProtectionSourceParameters) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetExchangeDagProtectionPreference

`func (o *UpdateProtectionSourceParameters) GetExchangeDagProtectionPreference() ExchangeDAGProtectionPreference`

GetExchangeDagProtectionPreference returns the ExchangeDagProtectionPreference field if non-nil, zero value otherwise.

### GetExchangeDagProtectionPreferenceOk

`func (o *UpdateProtectionSourceParameters) GetExchangeDagProtectionPreferenceOk() (*ExchangeDAGProtectionPreference, bool)`

GetExchangeDagProtectionPreferenceOk returns a tuple with the ExchangeDagProtectionPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDagProtectionPreference

`func (o *UpdateProtectionSourceParameters) SetExchangeDagProtectionPreference(v ExchangeDAGProtectionPreference)`

SetExchangeDagProtectionPreference sets ExchangeDagProtectionPreference field to given value.

### HasExchangeDagProtectionPreference

`func (o *UpdateProtectionSourceParameters) HasExchangeDagProtectionPreference() bool`

HasExchangeDagProtectionPreference returns a boolean if a field has been set.

### GetForceRegister

`func (o *UpdateProtectionSourceParameters) GetForceRegister() bool`

GetForceRegister returns the ForceRegister field if non-nil, zero value otherwise.

### GetForceRegisterOk

`func (o *UpdateProtectionSourceParameters) GetForceRegisterOk() (*bool, bool)`

GetForceRegisterOk returns a tuple with the ForceRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRegister

`func (o *UpdateProtectionSourceParameters) SetForceRegister(v bool)`

SetForceRegister sets ForceRegister field to given value.

### HasForceRegister

`func (o *UpdateProtectionSourceParameters) HasForceRegister() bool`

HasForceRegister returns a boolean if a field has been set.

### SetForceRegisterNil

`func (o *UpdateProtectionSourceParameters) SetForceRegisterNil(b bool)`

 SetForceRegisterNil sets the value for ForceRegister to be an explicit nil

### UnsetForceRegister
`func (o *UpdateProtectionSourceParameters) UnsetForceRegister()`

UnsetForceRegister ensures that no value is present for ForceRegister, not even an explicit nil
### GetGcpCredentials

`func (o *UpdateProtectionSourceParameters) GetGcpCredentials() GcpCredentials`

GetGcpCredentials returns the GcpCredentials field if non-nil, zero value otherwise.

### GetGcpCredentialsOk

`func (o *UpdateProtectionSourceParameters) GetGcpCredentialsOk() (*GcpCredentials, bool)`

GetGcpCredentialsOk returns a tuple with the GcpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredentials

`func (o *UpdateProtectionSourceParameters) SetGcpCredentials(v GcpCredentials)`

SetGcpCredentials sets GcpCredentials field to given value.

### HasGcpCredentials

`func (o *UpdateProtectionSourceParameters) HasGcpCredentials() bool`

HasGcpCredentials returns a boolean if a field has been set.

### GetHostType

`func (o *UpdateProtectionSourceParameters) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *UpdateProtectionSourceParameters) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *UpdateProtectionSourceParameters) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *UpdateProtectionSourceParameters) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *UpdateProtectionSourceParameters) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *UpdateProtectionSourceParameters) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetKubernetesCredentials

`func (o *UpdateProtectionSourceParameters) GetKubernetesCredentials() KubernetesCredentials`

GetKubernetesCredentials returns the KubernetesCredentials field if non-nil, zero value otherwise.

### GetKubernetesCredentialsOk

`func (o *UpdateProtectionSourceParameters) GetKubernetesCredentialsOk() (*KubernetesCredentials, bool)`

GetKubernetesCredentialsOk returns a tuple with the KubernetesCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCredentials

`func (o *UpdateProtectionSourceParameters) SetKubernetesCredentials(v KubernetesCredentials)`

SetKubernetesCredentials sets KubernetesCredentials field to given value.

### HasKubernetesCredentials

`func (o *UpdateProtectionSourceParameters) HasKubernetesCredentials() bool`

HasKubernetesCredentials returns a boolean if a field has been set.

### GetMinimumFreeSpaceGB

`func (o *UpdateProtectionSourceParameters) GetMinimumFreeSpaceGB() int64`

GetMinimumFreeSpaceGB returns the MinimumFreeSpaceGB field if non-nil, zero value otherwise.

### GetMinimumFreeSpaceGBOk

`func (o *UpdateProtectionSourceParameters) GetMinimumFreeSpaceGBOk() (*int64, bool)`

GetMinimumFreeSpaceGBOk returns a tuple with the MinimumFreeSpaceGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumFreeSpaceGB

`func (o *UpdateProtectionSourceParameters) SetMinimumFreeSpaceGB(v int64)`

SetMinimumFreeSpaceGB sets MinimumFreeSpaceGB field to given value.

### HasMinimumFreeSpaceGB

`func (o *UpdateProtectionSourceParameters) HasMinimumFreeSpaceGB() bool`

HasMinimumFreeSpaceGB returns a boolean if a field has been set.

### SetMinimumFreeSpaceGBNil

`func (o *UpdateProtectionSourceParameters) SetMinimumFreeSpaceGBNil(b bool)`

 SetMinimumFreeSpaceGBNil sets the value for MinimumFreeSpaceGB to be an explicit nil

### UnsetMinimumFreeSpaceGB
`func (o *UpdateProtectionSourceParameters) UnsetMinimumFreeSpaceGB()`

UnsetMinimumFreeSpaceGB ensures that no value is present for MinimumFreeSpaceGB, not even an explicit nil
### GetNasMountCredentials

`func (o *UpdateProtectionSourceParameters) GetNasMountCredentials() NasMountCredentialParams`

GetNasMountCredentials returns the NasMountCredentials field if non-nil, zero value otherwise.

### GetNasMountCredentialsOk

`func (o *UpdateProtectionSourceParameters) GetNasMountCredentialsOk() (*NasMountCredentialParams, bool)`

GetNasMountCredentialsOk returns a tuple with the NasMountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasMountCredentials

`func (o *UpdateProtectionSourceParameters) SetNasMountCredentials(v NasMountCredentialParams)`

SetNasMountCredentials sets NasMountCredentials field to given value.

### HasNasMountCredentials

`func (o *UpdateProtectionSourceParameters) HasNasMountCredentials() bool`

HasNasMountCredentials returns a boolean if a field has been set.

### SetNasMountCredentialsNil

`func (o *UpdateProtectionSourceParameters) SetNasMountCredentialsNil(b bool)`

 SetNasMountCredentialsNil sets the value for NasMountCredentials to be an explicit nil

### UnsetNasMountCredentials
`func (o *UpdateProtectionSourceParameters) UnsetNasMountCredentials()`

UnsetNasMountCredentials ensures that no value is present for NasMountCredentials, not even an explicit nil
### GetOffice365CredentialsList

`func (o *UpdateProtectionSourceParameters) GetOffice365CredentialsList() []Office365Credentials`

GetOffice365CredentialsList returns the Office365CredentialsList field if non-nil, zero value otherwise.

### GetOffice365CredentialsListOk

`func (o *UpdateProtectionSourceParameters) GetOffice365CredentialsListOk() (*[]Office365Credentials, bool)`

GetOffice365CredentialsListOk returns a tuple with the Office365CredentialsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365CredentialsList

`func (o *UpdateProtectionSourceParameters) SetOffice365CredentialsList(v []Office365Credentials)`

SetOffice365CredentialsList sets Office365CredentialsList field to given value.

### HasOffice365CredentialsList

`func (o *UpdateProtectionSourceParameters) HasOffice365CredentialsList() bool`

HasOffice365CredentialsList returns a boolean if a field has been set.

### SetOffice365CredentialsListNil

`func (o *UpdateProtectionSourceParameters) SetOffice365CredentialsListNil(b bool)`

 SetOffice365CredentialsListNil sets the value for Office365CredentialsList to be an explicit nil

### UnsetOffice365CredentialsList
`func (o *UpdateProtectionSourceParameters) UnsetOffice365CredentialsList()`

UnsetOffice365CredentialsList ensures that no value is present for Office365CredentialsList, not even an explicit nil
### GetOffice365Region

`func (o *UpdateProtectionSourceParameters) GetOffice365Region() string`

GetOffice365Region returns the Office365Region field if non-nil, zero value otherwise.

### GetOffice365RegionOk

`func (o *UpdateProtectionSourceParameters) GetOffice365RegionOk() (*string, bool)`

GetOffice365RegionOk returns a tuple with the Office365Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice365Region

`func (o *UpdateProtectionSourceParameters) SetOffice365Region(v string)`

SetOffice365Region sets Office365Region field to given value.

### HasOffice365Region

`func (o *UpdateProtectionSourceParameters) HasOffice365Region() bool`

HasOffice365Region returns a boolean if a field has been set.

### SetOffice365RegionNil

`func (o *UpdateProtectionSourceParameters) SetOffice365RegionNil(b bool)`

 SetOffice365RegionNil sets the value for Office365Region to be an explicit nil

### UnsetOffice365Region
`func (o *UpdateProtectionSourceParameters) UnsetOffice365Region()`

UnsetOffice365Region ensures that no value is present for Office365Region, not even an explicit nil
### GetPassword

`func (o *UpdateProtectionSourceParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateProtectionSourceParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateProtectionSourceParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateProtectionSourceParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateProtectionSourceParameters) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateProtectionSourceParameters) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetReRegister

`func (o *UpdateProtectionSourceParameters) GetReRegister() bool`

GetReRegister returns the ReRegister field if non-nil, zero value otherwise.

### GetReRegisterOk

`func (o *UpdateProtectionSourceParameters) GetReRegisterOk() (*bool, bool)`

GetReRegisterOk returns a tuple with the ReRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReRegister

`func (o *UpdateProtectionSourceParameters) SetReRegister(v bool)`

SetReRegister sets ReRegister field to given value.

### HasReRegister

`func (o *UpdateProtectionSourceParameters) HasReRegister() bool`

HasReRegister returns a boolean if a field has been set.

### SetReRegisterNil

`func (o *UpdateProtectionSourceParameters) SetReRegisterNil(b bool)`

 SetReRegisterNil sets the value for ReRegister to be an explicit nil

### UnsetReRegister
`func (o *UpdateProtectionSourceParameters) UnsetReRegister()`

UnsetReRegister ensures that no value is present for ReRegister, not even an explicit nil
### GetSourceSideDedupEnabled

`func (o *UpdateProtectionSourceParameters) GetSourceSideDedupEnabled() bool`

GetSourceSideDedupEnabled returns the SourceSideDedupEnabled field if non-nil, zero value otherwise.

### GetSourceSideDedupEnabledOk

`func (o *UpdateProtectionSourceParameters) GetSourceSideDedupEnabledOk() (*bool, bool)`

GetSourceSideDedupEnabledOk returns a tuple with the SourceSideDedupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSideDedupEnabled

`func (o *UpdateProtectionSourceParameters) SetSourceSideDedupEnabled(v bool)`

SetSourceSideDedupEnabled sets SourceSideDedupEnabled field to given value.

### HasSourceSideDedupEnabled

`func (o *UpdateProtectionSourceParameters) HasSourceSideDedupEnabled() bool`

HasSourceSideDedupEnabled returns a boolean if a field has been set.

### SetSourceSideDedupEnabledNil

`func (o *UpdateProtectionSourceParameters) SetSourceSideDedupEnabledNil(b bool)`

 SetSourceSideDedupEnabledNil sets the value for SourceSideDedupEnabled to be an explicit nil

### UnsetSourceSideDedupEnabled
`func (o *UpdateProtectionSourceParameters) UnsetSourceSideDedupEnabled()`

UnsetSourceSideDedupEnabled ensures that no value is present for SourceSideDedupEnabled, not even an explicit nil
### GetSslVerification

`func (o *UpdateProtectionSourceParameters) GetSslVerification() SslVerification`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *UpdateProtectionSourceParameters) GetSslVerificationOk() (*SslVerification, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *UpdateProtectionSourceParameters) SetSslVerification(v SslVerification)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *UpdateProtectionSourceParameters) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetSubnets

`func (o *UpdateProtectionSourceParameters) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *UpdateProtectionSourceParameters) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *UpdateProtectionSourceParameters) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *UpdateProtectionSourceParameters) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *UpdateProtectionSourceParameters) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *UpdateProtectionSourceParameters) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetThrottlingPolicy

`func (o *UpdateProtectionSourceParameters) GetThrottlingPolicy() ThrottlingPolicyParameters`

GetThrottlingPolicy returns the ThrottlingPolicy field if non-nil, zero value otherwise.

### GetThrottlingPolicyOk

`func (o *UpdateProtectionSourceParameters) GetThrottlingPolicyOk() (*ThrottlingPolicyParameters, bool)`

GetThrottlingPolicyOk returns a tuple with the ThrottlingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicy

`func (o *UpdateProtectionSourceParameters) SetThrottlingPolicy(v ThrottlingPolicyParameters)`

SetThrottlingPolicy sets ThrottlingPolicy field to given value.

### HasThrottlingPolicy

`func (o *UpdateProtectionSourceParameters) HasThrottlingPolicy() bool`

HasThrottlingPolicy returns a boolean if a field has been set.

### SetThrottlingPolicyNil

`func (o *UpdateProtectionSourceParameters) SetThrottlingPolicyNil(b bool)`

 SetThrottlingPolicyNil sets the value for ThrottlingPolicy to be an explicit nil

### UnsetThrottlingPolicy
`func (o *UpdateProtectionSourceParameters) UnsetThrottlingPolicy()`

UnsetThrottlingPolicy ensures that no value is present for ThrottlingPolicy, not even an explicit nil
### GetThrottlingPolicyOverrides

`func (o *UpdateProtectionSourceParameters) GetThrottlingPolicyOverrides() []ThrottlingPolicyOverride`

GetThrottlingPolicyOverrides returns the ThrottlingPolicyOverrides field if non-nil, zero value otherwise.

### GetThrottlingPolicyOverridesOk

`func (o *UpdateProtectionSourceParameters) GetThrottlingPolicyOverridesOk() (*[]ThrottlingPolicyOverride, bool)`

GetThrottlingPolicyOverridesOk returns a tuple with the ThrottlingPolicyOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingPolicyOverrides

`func (o *UpdateProtectionSourceParameters) SetThrottlingPolicyOverrides(v []ThrottlingPolicyOverride)`

SetThrottlingPolicyOverrides sets ThrottlingPolicyOverrides field to given value.

### HasThrottlingPolicyOverrides

`func (o *UpdateProtectionSourceParameters) HasThrottlingPolicyOverrides() bool`

HasThrottlingPolicyOverrides returns a boolean if a field has been set.

### SetThrottlingPolicyOverridesNil

`func (o *UpdateProtectionSourceParameters) SetThrottlingPolicyOverridesNil(b bool)`

 SetThrottlingPolicyOverridesNil sets the value for ThrottlingPolicyOverrides to be an explicit nil

### UnsetThrottlingPolicyOverrides
`func (o *UpdateProtectionSourceParameters) UnsetThrottlingPolicyOverrides()`

UnsetThrottlingPolicyOverrides ensures that no value is present for ThrottlingPolicyOverrides, not even an explicit nil
### GetUseOAuthForExchangeOnline

`func (o *UpdateProtectionSourceParameters) GetUseOAuthForExchangeOnline() bool`

GetUseOAuthForExchangeOnline returns the UseOAuthForExchangeOnline field if non-nil, zero value otherwise.

### GetUseOAuthForExchangeOnlineOk

`func (o *UpdateProtectionSourceParameters) GetUseOAuthForExchangeOnlineOk() (*bool, bool)`

GetUseOAuthForExchangeOnlineOk returns a tuple with the UseOAuthForExchangeOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOAuthForExchangeOnline

`func (o *UpdateProtectionSourceParameters) SetUseOAuthForExchangeOnline(v bool)`

SetUseOAuthForExchangeOnline sets UseOAuthForExchangeOnline field to given value.

### HasUseOAuthForExchangeOnline

`func (o *UpdateProtectionSourceParameters) HasUseOAuthForExchangeOnline() bool`

HasUseOAuthForExchangeOnline returns a boolean if a field has been set.

### SetUseOAuthForExchangeOnlineNil

`func (o *UpdateProtectionSourceParameters) SetUseOAuthForExchangeOnlineNil(b bool)`

 SetUseOAuthForExchangeOnlineNil sets the value for UseOAuthForExchangeOnline to be an explicit nil

### UnsetUseOAuthForExchangeOnline
`func (o *UpdateProtectionSourceParameters) UnsetUseOAuthForExchangeOnline()`

UnsetUseOAuthForExchangeOnline ensures that no value is present for UseOAuthForExchangeOnline, not even an explicit nil
### GetUsername

`func (o *UpdateProtectionSourceParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateProtectionSourceParameters) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateProtectionSourceParameters) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateProtectionSourceParameters) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateProtectionSourceParameters) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateProtectionSourceParameters) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetVlanParams

`func (o *UpdateProtectionSourceParameters) GetVlanParams() VlanParameters`

GetVlanParams returns the VlanParams field if non-nil, zero value otherwise.

### GetVlanParamsOk

`func (o *UpdateProtectionSourceParameters) GetVlanParamsOk() (*VlanParameters, bool)`

GetVlanParamsOk returns a tuple with the VlanParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParams

`func (o *UpdateProtectionSourceParameters) SetVlanParams(v VlanParameters)`

SetVlanParams sets VlanParams field to given value.

### HasVlanParams

`func (o *UpdateProtectionSourceParameters) HasVlanParams() bool`

HasVlanParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


