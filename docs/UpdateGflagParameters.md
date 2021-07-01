# UpdateGflagParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveNow** | Pointer to **NullableBool** | Specifies whether to apply the change immediately. If set to true, the gflag change will work without restarting the service. | [optional] 
**Gflags** | Pointer to [**[]Gflag**](Gflag.md) | Specifies a list of gflags. These will be added to the existing flags for the service. The values will be overwritten if required. If no value for gflag is specified, this gflag will be reset to default value. If no gflag is specified, all gflags for this service will be reset to default value. | [optional] 
**Reason** | Pointer to **NullableString** | Specifies the reason for clearing gflags. | [optional] 
**ServiceName** | **NullableString** | Specifies the service name. &#39;kApollo&#39; is a service for reclaiming freed disk sectors on Nodes in the SnapFS distributed file system. &#39;kBridge&#39; is a service for managing the SnapFS distributed file system. &#39;kGenie&#39; is a service that is responsible for monitoring hardware health on the Cluster. &#39;kGenieGofer&#39; is a service that links the Genie service to other services on the Cluster. &#39;kMagneto&#39; is the data protection service of the Cohesity Data Platform. &#39;kIris&#39; is the service which serves REST API calls to the UI, CLI, and any scripts written by customers. &#39;kIrisProxy&#39; is a service that links the Iris service to other services on the Cluster. &#39;kScribe&#39; is the service responsible for storing filesystem metadata. &#39;kStats&#39; is the service that is responsible for retrieving and aggregating disk metrics across the Cluster. &#39;kYoda&#39; is an elastic search indexing service. &#39;kAlerts&#39; is a publisher and subscribing service for alerts. &#39;kKeychain&#39; is a service for managing disk encryption keys. &#39;kLogWatcher&#39; is a service that scans the log directory and reduces the number of logs if required. &#39;kStatsCollector&#39; is a service that periodically logs system stats. &#39;kGandalf&#39; is a distributed lock service and coordination manager. &#39;kNexus&#39; indicates the Nexus service. This is the service that is responsible for creation of Clusters and configuration of Nodes and networking. &#39;kNexusProxy&#39; is a service that links the Nexus service to other services on the Cluster. &#39;kStorageProxy&#39; is a service for accessing data on external entities. &#39;kTricorder&#39; is a diagnostic health testing service for Clusters. &#39;kRtClient&#39; is a reverse tunneling client service. &#39;kVaultProxy&#39; is a service for managing external targets that Clusters can be backed up to. &#39;kSmbProxy&#39; is an SMB protocol service. &#39;kBridgeProxy&#39; is the service that links the Bridge service to other services on the Cluster. &#39;kLibrarian&#39; is an elastic search indexing service. &#39;kGroot&#39; is a service for managing replication of SQL databases across multiple nodes in a Cluster. &#39;kEagleAgent&#39; is a service that is responsible for retrieving information on Cluster health. &#39;kAthena&#39; is a service for running distributed containerized applications on the Cohesity Data Platform. &#39;kBifrostBroker&#39; is a service for communicating with the Cohesity proxies for multitenancy. &#39;kSmb2Proxy&#39; is a new SMB protocol service. &#39;kOs&#39; can be specified in order to do a full reboot. &#39;kAtom&#39; is a service for receiving data for the Continuous Data Protection. | 

## Methods

### NewUpdateGflagParameters

`func NewUpdateGflagParameters(serviceName NullableString, ) *UpdateGflagParameters`

NewUpdateGflagParameters instantiates a new UpdateGflagParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGflagParametersWithDefaults

`func NewUpdateGflagParametersWithDefaults() *UpdateGflagParameters`

NewUpdateGflagParametersWithDefaults instantiates a new UpdateGflagParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveNow

`func (o *UpdateGflagParameters) GetEffectiveNow() bool`

GetEffectiveNow returns the EffectiveNow field if non-nil, zero value otherwise.

### GetEffectiveNowOk

`func (o *UpdateGflagParameters) GetEffectiveNowOk() (*bool, bool)`

GetEffectiveNowOk returns a tuple with the EffectiveNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveNow

`func (o *UpdateGflagParameters) SetEffectiveNow(v bool)`

SetEffectiveNow sets EffectiveNow field to given value.

### HasEffectiveNow

`func (o *UpdateGflagParameters) HasEffectiveNow() bool`

HasEffectiveNow returns a boolean if a field has been set.

### SetEffectiveNowNil

`func (o *UpdateGflagParameters) SetEffectiveNowNil(b bool)`

 SetEffectiveNowNil sets the value for EffectiveNow to be an explicit nil

### UnsetEffectiveNow
`func (o *UpdateGflagParameters) UnsetEffectiveNow()`

UnsetEffectiveNow ensures that no value is present for EffectiveNow, not even an explicit nil
### GetGflags

`func (o *UpdateGflagParameters) GetGflags() []Gflag`

GetGflags returns the Gflags field if non-nil, zero value otherwise.

### GetGflagsOk

`func (o *UpdateGflagParameters) GetGflagsOk() (*[]Gflag, bool)`

GetGflagsOk returns a tuple with the Gflags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGflags

`func (o *UpdateGflagParameters) SetGflags(v []Gflag)`

SetGflags sets Gflags field to given value.

### HasGflags

`func (o *UpdateGflagParameters) HasGflags() bool`

HasGflags returns a boolean if a field has been set.

### SetGflagsNil

`func (o *UpdateGflagParameters) SetGflagsNil(b bool)`

 SetGflagsNil sets the value for Gflags to be an explicit nil

### UnsetGflags
`func (o *UpdateGflagParameters) UnsetGflags()`

UnsetGflags ensures that no value is present for Gflags, not even an explicit nil
### GetReason

`func (o *UpdateGflagParameters) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateGflagParameters) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateGflagParameters) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateGflagParameters) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *UpdateGflagParameters) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UpdateGflagParameters) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetServiceName

`func (o *UpdateGflagParameters) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UpdateGflagParameters) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UpdateGflagParameters) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### SetServiceNameNil

`func (o *UpdateGflagParameters) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *UpdateGflagParameters) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


