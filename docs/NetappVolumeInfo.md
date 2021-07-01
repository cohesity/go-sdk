# NetappVolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateName** | Pointer to **NullableString** | Specifies the containing aggregate name of this volume. | [optional] 
**CapacityBytes** | Pointer to **NullableInt64** | Specifies the total capacity in bytes of this volume. | [optional] 
**CifsShares** | Pointer to [**[]CifsShareInfo**](CifsShareInfo.md) | Array of CIFS Shares.  Specifies the set of CIFS Shares exported for this volume. | [optional] 
**CreationTimeUsecs** | Pointer to **NullableInt64** | Specifies the creation time of the volume specified in Unix epoch time (in microseconds). | [optional] 
**DataProtocols** | Pointer to **[]string** | Array of Data Protocols.  Specifies the set of data protocols supported by this volume. &#39;kNfs&#39; indicates NFS connections. &#39;kCifs&#39; indicates SMB (CIFS) connections. &#39;kIscsi&#39; indicates iSCSI connections. &#39;kFc&#39; indicates Fiber Channel connections. &#39;kFcache&#39; indicates Flex Cache connections. &#39;kHttp&#39; indicates HTTP connections. &#39;kNdmp&#39; indicates NDMP connections. &#39;kManagement&#39; indicates non-data connections used for management purposes. &#39;kNvme&#39; indicates NVMe connections. | [optional] 
**ExportPolicyName** | Pointer to **NullableString** | Specifies the name of the export policy (which defines the access permissions for the mount client) that has been assigned to this volume. | [optional] 
**JunctionPath** | Pointer to **NullableString** | Specifies the junction path of this volume. This path can be used to mount this volume via protocols such as NFS. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the NetApp Vserver that this volume belongs to. | [optional] 
**SecurityInfo** | Pointer to [**VolumeSecurityInfo**](VolumeSecurityInfo.md) |  | [optional] 
**State** | Pointer to **NullableString** | Specifies the state of this volume. Specifies the state of a NetApp Volume. &#39;kOnline&#39; indicates the volume is online. Read and write access to this volume is allowed. &#39;kRestricted&#39; indicates the volume is restricted. Some operations, such as parity reconstruction, are allowed, but data access is not allowed. &#39;kOffline&#39; indicates the volume is offline. No access to the volume is allowed. &#39;kMixed&#39; indicates the volume is in mixed state, which means its aggregates are not all in the same state. | [optional] 
**Type** | Pointer to **NullableString** | Specifies the NetApp type of this volume. Specifies the type of a NetApp Volume. &#39;kReadWrite&#39; indicates read-write volume. &#39;kLoadSharing&#39; indicates load-sharing volume. &#39;kDataProtection&#39; indicates data-protection volume. &#39;kDataCache&#39; indicates data-cache volume. &#39;kTmp&#39; indicates temporary purpose. &#39;kUnknownType&#39; indicates unknown type. | [optional] 
**UsedBytes** | Pointer to **NullableInt64** | Specifies the total space (in bytes) used in this volume. | [optional] 

## Methods

### NewNetappVolumeInfo

`func NewNetappVolumeInfo() *NetappVolumeInfo`

NewNetappVolumeInfo instantiates a new NetappVolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetappVolumeInfoWithDefaults

`func NewNetappVolumeInfoWithDefaults() *NetappVolumeInfo`

NewNetappVolumeInfoWithDefaults instantiates a new NetappVolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregateName

`func (o *NetappVolumeInfo) GetAggregateName() string`

GetAggregateName returns the AggregateName field if non-nil, zero value otherwise.

### GetAggregateNameOk

`func (o *NetappVolumeInfo) GetAggregateNameOk() (*string, bool)`

GetAggregateNameOk returns a tuple with the AggregateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateName

`func (o *NetappVolumeInfo) SetAggregateName(v string)`

SetAggregateName sets AggregateName field to given value.

### HasAggregateName

`func (o *NetappVolumeInfo) HasAggregateName() bool`

HasAggregateName returns a boolean if a field has been set.

### SetAggregateNameNil

`func (o *NetappVolumeInfo) SetAggregateNameNil(b bool)`

 SetAggregateNameNil sets the value for AggregateName to be an explicit nil

### UnsetAggregateName
`func (o *NetappVolumeInfo) UnsetAggregateName()`

UnsetAggregateName ensures that no value is present for AggregateName, not even an explicit nil
### GetCapacityBytes

`func (o *NetappVolumeInfo) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *NetappVolumeInfo) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *NetappVolumeInfo) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.

### HasCapacityBytes

`func (o *NetappVolumeInfo) HasCapacityBytes() bool`

HasCapacityBytes returns a boolean if a field has been set.

### SetCapacityBytesNil

`func (o *NetappVolumeInfo) SetCapacityBytesNil(b bool)`

 SetCapacityBytesNil sets the value for CapacityBytes to be an explicit nil

### UnsetCapacityBytes
`func (o *NetappVolumeInfo) UnsetCapacityBytes()`

UnsetCapacityBytes ensures that no value is present for CapacityBytes, not even an explicit nil
### GetCifsShares

`func (o *NetappVolumeInfo) GetCifsShares() []CifsShareInfo`

GetCifsShares returns the CifsShares field if non-nil, zero value otherwise.

### GetCifsSharesOk

`func (o *NetappVolumeInfo) GetCifsSharesOk() (*[]CifsShareInfo, bool)`

GetCifsSharesOk returns a tuple with the CifsShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsShares

`func (o *NetappVolumeInfo) SetCifsShares(v []CifsShareInfo)`

SetCifsShares sets CifsShares field to given value.

### HasCifsShares

`func (o *NetappVolumeInfo) HasCifsShares() bool`

HasCifsShares returns a boolean if a field has been set.

### SetCifsSharesNil

`func (o *NetappVolumeInfo) SetCifsSharesNil(b bool)`

 SetCifsSharesNil sets the value for CifsShares to be an explicit nil

### UnsetCifsShares
`func (o *NetappVolumeInfo) UnsetCifsShares()`

UnsetCifsShares ensures that no value is present for CifsShares, not even an explicit nil
### GetCreationTimeUsecs

`func (o *NetappVolumeInfo) GetCreationTimeUsecs() int64`

GetCreationTimeUsecs returns the CreationTimeUsecs field if non-nil, zero value otherwise.

### GetCreationTimeUsecsOk

`func (o *NetappVolumeInfo) GetCreationTimeUsecsOk() (*int64, bool)`

GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimeUsecs

`func (o *NetappVolumeInfo) SetCreationTimeUsecs(v int64)`

SetCreationTimeUsecs sets CreationTimeUsecs field to given value.

### HasCreationTimeUsecs

`func (o *NetappVolumeInfo) HasCreationTimeUsecs() bool`

HasCreationTimeUsecs returns a boolean if a field has been set.

### SetCreationTimeUsecsNil

`func (o *NetappVolumeInfo) SetCreationTimeUsecsNil(b bool)`

 SetCreationTimeUsecsNil sets the value for CreationTimeUsecs to be an explicit nil

### UnsetCreationTimeUsecs
`func (o *NetappVolumeInfo) UnsetCreationTimeUsecs()`

UnsetCreationTimeUsecs ensures that no value is present for CreationTimeUsecs, not even an explicit nil
### GetDataProtocols

`func (o *NetappVolumeInfo) GetDataProtocols() []string`

GetDataProtocols returns the DataProtocols field if non-nil, zero value otherwise.

### GetDataProtocolsOk

`func (o *NetappVolumeInfo) GetDataProtocolsOk() (*[]string, bool)`

GetDataProtocolsOk returns a tuple with the DataProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtocols

`func (o *NetappVolumeInfo) SetDataProtocols(v []string)`

SetDataProtocols sets DataProtocols field to given value.

### HasDataProtocols

`func (o *NetappVolumeInfo) HasDataProtocols() bool`

HasDataProtocols returns a boolean if a field has been set.

### SetDataProtocolsNil

`func (o *NetappVolumeInfo) SetDataProtocolsNil(b bool)`

 SetDataProtocolsNil sets the value for DataProtocols to be an explicit nil

### UnsetDataProtocols
`func (o *NetappVolumeInfo) UnsetDataProtocols()`

UnsetDataProtocols ensures that no value is present for DataProtocols, not even an explicit nil
### GetExportPolicyName

`func (o *NetappVolumeInfo) GetExportPolicyName() string`

GetExportPolicyName returns the ExportPolicyName field if non-nil, zero value otherwise.

### GetExportPolicyNameOk

`func (o *NetappVolumeInfo) GetExportPolicyNameOk() (*string, bool)`

GetExportPolicyNameOk returns a tuple with the ExportPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicyName

`func (o *NetappVolumeInfo) SetExportPolicyName(v string)`

SetExportPolicyName sets ExportPolicyName field to given value.

### HasExportPolicyName

`func (o *NetappVolumeInfo) HasExportPolicyName() bool`

HasExportPolicyName returns a boolean if a field has been set.

### SetExportPolicyNameNil

`func (o *NetappVolumeInfo) SetExportPolicyNameNil(b bool)`

 SetExportPolicyNameNil sets the value for ExportPolicyName to be an explicit nil

### UnsetExportPolicyName
`func (o *NetappVolumeInfo) UnsetExportPolicyName()`

UnsetExportPolicyName ensures that no value is present for ExportPolicyName, not even an explicit nil
### GetJunctionPath

`func (o *NetappVolumeInfo) GetJunctionPath() string`

GetJunctionPath returns the JunctionPath field if non-nil, zero value otherwise.

### GetJunctionPathOk

`func (o *NetappVolumeInfo) GetJunctionPathOk() (*string, bool)`

GetJunctionPathOk returns a tuple with the JunctionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJunctionPath

`func (o *NetappVolumeInfo) SetJunctionPath(v string)`

SetJunctionPath sets JunctionPath field to given value.

### HasJunctionPath

`func (o *NetappVolumeInfo) HasJunctionPath() bool`

HasJunctionPath returns a boolean if a field has been set.

### SetJunctionPathNil

`func (o *NetappVolumeInfo) SetJunctionPathNil(b bool)`

 SetJunctionPathNil sets the value for JunctionPath to be an explicit nil

### UnsetJunctionPath
`func (o *NetappVolumeInfo) UnsetJunctionPath()`

UnsetJunctionPath ensures that no value is present for JunctionPath, not even an explicit nil
### GetName

`func (o *NetappVolumeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetappVolumeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetappVolumeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetappVolumeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NetappVolumeInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NetappVolumeInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSecurityInfo

`func (o *NetappVolumeInfo) GetSecurityInfo() VolumeSecurityInfo`

GetSecurityInfo returns the SecurityInfo field if non-nil, zero value otherwise.

### GetSecurityInfoOk

`func (o *NetappVolumeInfo) GetSecurityInfoOk() (*VolumeSecurityInfo, bool)`

GetSecurityInfoOk returns a tuple with the SecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityInfo

`func (o *NetappVolumeInfo) SetSecurityInfo(v VolumeSecurityInfo)`

SetSecurityInfo sets SecurityInfo field to given value.

### HasSecurityInfo

`func (o *NetappVolumeInfo) HasSecurityInfo() bool`

HasSecurityInfo returns a boolean if a field has been set.

### GetState

`func (o *NetappVolumeInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetappVolumeInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetappVolumeInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NetappVolumeInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *NetappVolumeInfo) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *NetappVolumeInfo) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetType

`func (o *NetappVolumeInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetappVolumeInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetappVolumeInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetappVolumeInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NetappVolumeInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NetappVolumeInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsedBytes

`func (o *NetappVolumeInfo) GetUsedBytes() int64`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *NetappVolumeInfo) GetUsedBytesOk() (*int64, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *NetappVolumeInfo) SetUsedBytes(v int64)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *NetappVolumeInfo) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### SetUsedBytesNil

`func (o *NetappVolumeInfo) SetUsedBytesNil(b bool)`

 SetUsedBytesNil sets the value for UsedBytes to be an explicit nil

### UnsetUsedBytes
`func (o *NetappVolumeInfo) UnsetUsedBytes()`

UnsetUsedBytes ensures that no value is present for UsedBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


