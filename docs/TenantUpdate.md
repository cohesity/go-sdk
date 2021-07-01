# TenantUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BifrostEnabled** | Pointer to **NullableBool** | Specifies whether bifrost (Ambassador proxy) is enabled for tenant. | [optional] 
**ClusterHostname** | Pointer to **NullableString** | The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant&#39;s network. Tenant&#39;s VLAN&#39;s hostname, if available can be used instead but it is mandatory to provide this value if there&#39;s no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname. | [optional] 
**ClusterIps** | Pointer to **[]string** | Set of IPs as seen from the tenant&#39;s network for the Cohesity cluster. Only one from &#39;ClusterHostname&#39; and &#39;ClusterIps&#39; is needed. | [optional] 
**Description** | Pointer to **NullableString** | Specifies the description of this tenant. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the tenant. | [optional] 
**SubscribeToAlertEmails** | Pointer to **NullableBool** | Service provider can optionally unsubscribe from the Tenant Alert Emails. | [optional] 
**TenantId** | Pointer to **NullableString** | Specifies the unique id of the tenant. | [optional] 

## Methods

### NewTenantUpdate

`func NewTenantUpdate() *TenantUpdate`

NewTenantUpdate instantiates a new TenantUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUpdateWithDefaults

`func NewTenantUpdateWithDefaults() *TenantUpdate`

NewTenantUpdateWithDefaults instantiates a new TenantUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBifrostEnabled

`func (o *TenantUpdate) GetBifrostEnabled() bool`

GetBifrostEnabled returns the BifrostEnabled field if non-nil, zero value otherwise.

### GetBifrostEnabledOk

`func (o *TenantUpdate) GetBifrostEnabledOk() (*bool, bool)`

GetBifrostEnabledOk returns a tuple with the BifrostEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBifrostEnabled

`func (o *TenantUpdate) SetBifrostEnabled(v bool)`

SetBifrostEnabled sets BifrostEnabled field to given value.

### HasBifrostEnabled

`func (o *TenantUpdate) HasBifrostEnabled() bool`

HasBifrostEnabled returns a boolean if a field has been set.

### SetBifrostEnabledNil

`func (o *TenantUpdate) SetBifrostEnabledNil(b bool)`

 SetBifrostEnabledNil sets the value for BifrostEnabled to be an explicit nil

### UnsetBifrostEnabled
`func (o *TenantUpdate) UnsetBifrostEnabled()`

UnsetBifrostEnabled ensures that no value is present for BifrostEnabled, not even an explicit nil
### GetClusterHostname

`func (o *TenantUpdate) GetClusterHostname() string`

GetClusterHostname returns the ClusterHostname field if non-nil, zero value otherwise.

### GetClusterHostnameOk

`func (o *TenantUpdate) GetClusterHostnameOk() (*string, bool)`

GetClusterHostnameOk returns a tuple with the ClusterHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHostname

`func (o *TenantUpdate) SetClusterHostname(v string)`

SetClusterHostname sets ClusterHostname field to given value.

### HasClusterHostname

`func (o *TenantUpdate) HasClusterHostname() bool`

HasClusterHostname returns a boolean if a field has been set.

### SetClusterHostnameNil

`func (o *TenantUpdate) SetClusterHostnameNil(b bool)`

 SetClusterHostnameNil sets the value for ClusterHostname to be an explicit nil

### UnsetClusterHostname
`func (o *TenantUpdate) UnsetClusterHostname()`

UnsetClusterHostname ensures that no value is present for ClusterHostname, not even an explicit nil
### GetClusterIps

`func (o *TenantUpdate) GetClusterIps() []string`

GetClusterIps returns the ClusterIps field if non-nil, zero value otherwise.

### GetClusterIpsOk

`func (o *TenantUpdate) GetClusterIpsOk() (*[]string, bool)`

GetClusterIpsOk returns a tuple with the ClusterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIps

`func (o *TenantUpdate) SetClusterIps(v []string)`

SetClusterIps sets ClusterIps field to given value.

### HasClusterIps

`func (o *TenantUpdate) HasClusterIps() bool`

HasClusterIps returns a boolean if a field has been set.

### SetClusterIpsNil

`func (o *TenantUpdate) SetClusterIpsNil(b bool)`

 SetClusterIpsNil sets the value for ClusterIps to be an explicit nil

### UnsetClusterIps
`func (o *TenantUpdate) UnsetClusterIps()`

UnsetClusterIps ensures that no value is present for ClusterIps, not even an explicit nil
### GetDescription

`func (o *TenantUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *TenantUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscribeToAlertEmails

`func (o *TenantUpdate) GetSubscribeToAlertEmails() bool`

GetSubscribeToAlertEmails returns the SubscribeToAlertEmails field if non-nil, zero value otherwise.

### GetSubscribeToAlertEmailsOk

`func (o *TenantUpdate) GetSubscribeToAlertEmailsOk() (*bool, bool)`

GetSubscribeToAlertEmailsOk returns a tuple with the SubscribeToAlertEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeToAlertEmails

`func (o *TenantUpdate) SetSubscribeToAlertEmails(v bool)`

SetSubscribeToAlertEmails sets SubscribeToAlertEmails field to given value.

### HasSubscribeToAlertEmails

`func (o *TenantUpdate) HasSubscribeToAlertEmails() bool`

HasSubscribeToAlertEmails returns a boolean if a field has been set.

### SetSubscribeToAlertEmailsNil

`func (o *TenantUpdate) SetSubscribeToAlertEmailsNil(b bool)`

 SetSubscribeToAlertEmailsNil sets the value for SubscribeToAlertEmails to be an explicit nil

### UnsetSubscribeToAlertEmails
`func (o *TenantUpdate) UnsetSubscribeToAlertEmails()`

UnsetSubscribeToAlertEmails ensures that no value is present for SubscribeToAlertEmails, not even an explicit nil
### GetTenantId

`func (o *TenantUpdate) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantUpdate) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantUpdate) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantUpdate) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantUpdate) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantUpdate) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


