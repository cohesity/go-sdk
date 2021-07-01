# TenantCreateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BifrostEnabled** | Pointer to **NullableBool** | Specifies whether bifrost (Ambassador proxy) is enabled for tenant. | [optional] 
**ClusterHostname** | Pointer to **NullableString** | The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant&#39;s network. Tenant&#39;s VLAN&#39;s hostname, if available can be used instead but it is mandatory to provide this value if there&#39;s no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname. | [optional] 
**ClusterIps** | Pointer to **[]string** | Set of IPs as seen from the tenant&#39;s network for the Cohesity cluster. Only one from &#39;ClusterHostname&#39; and &#39;ClusterIps&#39; is needed. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of this tenant. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the tenant. | [optional] 
**OrgSuffix** | Pointer to **NullableString** | Specifies the organization suffix needed to construct tenant id. Tenant id is not completely auto generated rather chosen by tenant/SP admin as we needed same tenant id on multiple clusters to support replication across clusters, etc. | [optional] 
**ParentTenantId** | Pointer to **NullableString** | Specifies the parent tenant of this tenant if available. | [optional] 
**SubscribeToAlertEmails** | Pointer to **NullableBool** | Service provider can optionally unsubscribe from the Tenant Alert Emails. | [optional] 

## Methods

### NewTenantCreateParameters

`func NewTenantCreateParameters() *TenantCreateParameters`

NewTenantCreateParameters instantiates a new TenantCreateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreateParametersWithDefaults

`func NewTenantCreateParametersWithDefaults() *TenantCreateParameters`

NewTenantCreateParametersWithDefaults instantiates a new TenantCreateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBifrostEnabled

`func (o *TenantCreateParameters) GetBifrostEnabled() bool`

GetBifrostEnabled returns the BifrostEnabled field if non-nil, zero value otherwise.

### GetBifrostEnabledOk

`func (o *TenantCreateParameters) GetBifrostEnabledOk() (*bool, bool)`

GetBifrostEnabledOk returns a tuple with the BifrostEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBifrostEnabled

`func (o *TenantCreateParameters) SetBifrostEnabled(v bool)`

SetBifrostEnabled sets BifrostEnabled field to given value.

### HasBifrostEnabled

`func (o *TenantCreateParameters) HasBifrostEnabled() bool`

HasBifrostEnabled returns a boolean if a field has been set.

### SetBifrostEnabledNil

`func (o *TenantCreateParameters) SetBifrostEnabledNil(b bool)`

 SetBifrostEnabledNil sets the value for BifrostEnabled to be an explicit nil

### UnsetBifrostEnabled
`func (o *TenantCreateParameters) UnsetBifrostEnabled()`

UnsetBifrostEnabled ensures that no value is present for BifrostEnabled, not even an explicit nil
### GetClusterHostname

`func (o *TenantCreateParameters) GetClusterHostname() string`

GetClusterHostname returns the ClusterHostname field if non-nil, zero value otherwise.

### GetClusterHostnameOk

`func (o *TenantCreateParameters) GetClusterHostnameOk() (*string, bool)`

GetClusterHostnameOk returns a tuple with the ClusterHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHostname

`func (o *TenantCreateParameters) SetClusterHostname(v string)`

SetClusterHostname sets ClusterHostname field to given value.

### HasClusterHostname

`func (o *TenantCreateParameters) HasClusterHostname() bool`

HasClusterHostname returns a boolean if a field has been set.

### SetClusterHostnameNil

`func (o *TenantCreateParameters) SetClusterHostnameNil(b bool)`

 SetClusterHostnameNil sets the value for ClusterHostname to be an explicit nil

### UnsetClusterHostname
`func (o *TenantCreateParameters) UnsetClusterHostname()`

UnsetClusterHostname ensures that no value is present for ClusterHostname, not even an explicit nil
### GetClusterIps

`func (o *TenantCreateParameters) GetClusterIps() []string`

GetClusterIps returns the ClusterIps field if non-nil, zero value otherwise.

### GetClusterIpsOk

`func (o *TenantCreateParameters) GetClusterIpsOk() (*[]string, bool)`

GetClusterIpsOk returns a tuple with the ClusterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIps

`func (o *TenantCreateParameters) SetClusterIps(v []string)`

SetClusterIps sets ClusterIps field to given value.

### HasClusterIps

`func (o *TenantCreateParameters) HasClusterIps() bool`

HasClusterIps returns a boolean if a field has been set.

### SetClusterIpsNil

`func (o *TenantCreateParameters) SetClusterIpsNil(b bool)`

 SetClusterIpsNil sets the value for ClusterIps to be an explicit nil

### UnsetClusterIps
`func (o *TenantCreateParameters) UnsetClusterIps()`

UnsetClusterIps ensures that no value is present for ClusterIps, not even an explicit nil
### GetDescription

`func (o *TenantCreateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantCreateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantCreateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantCreateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantCreateParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantCreateParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *TenantCreateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantCreateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantCreateParameters) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantCreateParameters) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrgSuffix

`func (o *TenantCreateParameters) GetOrgSuffix() string`

GetOrgSuffix returns the OrgSuffix field if non-nil, zero value otherwise.

### GetOrgSuffixOk

`func (o *TenantCreateParameters) GetOrgSuffixOk() (*string, bool)`

GetOrgSuffixOk returns a tuple with the OrgSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSuffix

`func (o *TenantCreateParameters) SetOrgSuffix(v string)`

SetOrgSuffix sets OrgSuffix field to given value.

### HasOrgSuffix

`func (o *TenantCreateParameters) HasOrgSuffix() bool`

HasOrgSuffix returns a boolean if a field has been set.

### SetOrgSuffixNil

`func (o *TenantCreateParameters) SetOrgSuffixNil(b bool)`

 SetOrgSuffixNil sets the value for OrgSuffix to be an explicit nil

### UnsetOrgSuffix
`func (o *TenantCreateParameters) UnsetOrgSuffix()`

UnsetOrgSuffix ensures that no value is present for OrgSuffix, not even an explicit nil
### GetParentTenantId

`func (o *TenantCreateParameters) GetParentTenantId() string`

GetParentTenantId returns the ParentTenantId field if non-nil, zero value otherwise.

### GetParentTenantIdOk

`func (o *TenantCreateParameters) GetParentTenantIdOk() (*string, bool)`

GetParentTenantIdOk returns a tuple with the ParentTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTenantId

`func (o *TenantCreateParameters) SetParentTenantId(v string)`

SetParentTenantId sets ParentTenantId field to given value.

### HasParentTenantId

`func (o *TenantCreateParameters) HasParentTenantId() bool`

HasParentTenantId returns a boolean if a field has been set.

### SetParentTenantIdNil

`func (o *TenantCreateParameters) SetParentTenantIdNil(b bool)`

 SetParentTenantIdNil sets the value for ParentTenantId to be an explicit nil

### UnsetParentTenantId
`func (o *TenantCreateParameters) UnsetParentTenantId()`

UnsetParentTenantId ensures that no value is present for ParentTenantId, not even an explicit nil
### GetSubscribeToAlertEmails

`func (o *TenantCreateParameters) GetSubscribeToAlertEmails() bool`

GetSubscribeToAlertEmails returns the SubscribeToAlertEmails field if non-nil, zero value otherwise.

### GetSubscribeToAlertEmailsOk

`func (o *TenantCreateParameters) GetSubscribeToAlertEmailsOk() (*bool, bool)`

GetSubscribeToAlertEmailsOk returns a tuple with the SubscribeToAlertEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeToAlertEmails

`func (o *TenantCreateParameters) SetSubscribeToAlertEmails(v bool)`

SetSubscribeToAlertEmails sets SubscribeToAlertEmails field to given value.

### HasSubscribeToAlertEmails

`func (o *TenantCreateParameters) HasSubscribeToAlertEmails() bool`

HasSubscribeToAlertEmails returns a boolean if a field has been set.

### SetSubscribeToAlertEmailsNil

`func (o *TenantCreateParameters) SetSubscribeToAlertEmailsNil(b bool)`

 SetSubscribeToAlertEmailsNil sets the value for SubscribeToAlertEmails to be an explicit nil

### UnsetSubscribeToAlertEmails
`func (o *TenantCreateParameters) UnsetSubscribeToAlertEmails()`

UnsetSubscribeToAlertEmails ensures that no value is present for SubscribeToAlertEmails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


