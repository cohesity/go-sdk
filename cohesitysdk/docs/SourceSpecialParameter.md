# SourceSpecialParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdSpecialParameters** | Pointer to [**ApplicationSpecialParameters**](ApplicationSpecialParameters.md) |  | [optional] 
**ExchangeSpecialParameters** | Pointer to [**ApplicationSpecialParameters**](ApplicationSpecialParameters.md) |  | [optional] 
**OracleSpecialParameters** | Pointer to [**OracleSpecialParameters**](OracleSpecialParameters.md) |  | [optional] 
**PhysicalSpecialParameters** | Pointer to [**PhysicalSpecialParameters**](PhysicalSpecialParameters.md) |  | [optional] 
**SkipIndexing** | Pointer to **NullableBool** | Specifies not to index the objects in the Protection Source when backing up. | [optional] 
**SourceId** | Pointer to **NullableInt64** | Specifies the object id of the Protection Source that these special settings apply. This field must refer to a leaf node such a VM or a Physical Server. | [optional] 
**SqlSpecialParameters** | Pointer to [**ApplicationSpecialParameters**](ApplicationSpecialParameters.md) |  | [optional] 
**TruncateExchangeLog** | Pointer to **NullableBool** | If true, after the Cohesity Cluster successfully captures a Snapshot during a Job Run, the Cluster truncates the Exchange transaction logs on a Microsoft Exchange Server. The default value is false. This field is deprecated. Use the field in ApplicationParameters inside source specific parameter. deprecated: true | [optional] 
**VmCredentials** | Pointer to [**NullableCredentials**](Credentials.md) | Specifies the administrator credentials to log in to the guest Windows system of a VM that hosts the Microsoft Exchange Server. If truncateExchangeLog is set to true and the specified source is a VM, administrator credentials to log in to the guest Windows system of the VM must be provided to truncate the logs. This field is only applicable to Sources in the kVMware environment. This field is deprecated. Use the field in VmCredentials inside source specific parameter. deprecated: true | [optional] 
**VmwareSpecialParameters** | Pointer to [**VmwareSpecialParameters**](VmwareSpecialParameters.md) |  | [optional] 

## Methods

### NewSourceSpecialParameter

`func NewSourceSpecialParameter() *SourceSpecialParameter`

NewSourceSpecialParameter instantiates a new SourceSpecialParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSpecialParameterWithDefaults

`func NewSourceSpecialParameterWithDefaults() *SourceSpecialParameter`

NewSourceSpecialParameterWithDefaults instantiates a new SourceSpecialParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdSpecialParameters

`func (o *SourceSpecialParameter) GetAdSpecialParameters() ApplicationSpecialParameters`

GetAdSpecialParameters returns the AdSpecialParameters field if non-nil, zero value otherwise.

### GetAdSpecialParametersOk

`func (o *SourceSpecialParameter) GetAdSpecialParametersOk() (*ApplicationSpecialParameters, bool)`

GetAdSpecialParametersOk returns a tuple with the AdSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSpecialParameters

`func (o *SourceSpecialParameter) SetAdSpecialParameters(v ApplicationSpecialParameters)`

SetAdSpecialParameters sets AdSpecialParameters field to given value.

### HasAdSpecialParameters

`func (o *SourceSpecialParameter) HasAdSpecialParameters() bool`

HasAdSpecialParameters returns a boolean if a field has been set.

### GetExchangeSpecialParameters

`func (o *SourceSpecialParameter) GetExchangeSpecialParameters() ApplicationSpecialParameters`

GetExchangeSpecialParameters returns the ExchangeSpecialParameters field if non-nil, zero value otherwise.

### GetExchangeSpecialParametersOk

`func (o *SourceSpecialParameter) GetExchangeSpecialParametersOk() (*ApplicationSpecialParameters, bool)`

GetExchangeSpecialParametersOk returns a tuple with the ExchangeSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeSpecialParameters

`func (o *SourceSpecialParameter) SetExchangeSpecialParameters(v ApplicationSpecialParameters)`

SetExchangeSpecialParameters sets ExchangeSpecialParameters field to given value.

### HasExchangeSpecialParameters

`func (o *SourceSpecialParameter) HasExchangeSpecialParameters() bool`

HasExchangeSpecialParameters returns a boolean if a field has been set.

### GetOracleSpecialParameters

`func (o *SourceSpecialParameter) GetOracleSpecialParameters() OracleSpecialParameters`

GetOracleSpecialParameters returns the OracleSpecialParameters field if non-nil, zero value otherwise.

### GetOracleSpecialParametersOk

`func (o *SourceSpecialParameter) GetOracleSpecialParametersOk() (*OracleSpecialParameters, bool)`

GetOracleSpecialParametersOk returns a tuple with the OracleSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleSpecialParameters

`func (o *SourceSpecialParameter) SetOracleSpecialParameters(v OracleSpecialParameters)`

SetOracleSpecialParameters sets OracleSpecialParameters field to given value.

### HasOracleSpecialParameters

`func (o *SourceSpecialParameter) HasOracleSpecialParameters() bool`

HasOracleSpecialParameters returns a boolean if a field has been set.

### GetPhysicalSpecialParameters

`func (o *SourceSpecialParameter) GetPhysicalSpecialParameters() PhysicalSpecialParameters`

GetPhysicalSpecialParameters returns the PhysicalSpecialParameters field if non-nil, zero value otherwise.

### GetPhysicalSpecialParametersOk

`func (o *SourceSpecialParameter) GetPhysicalSpecialParametersOk() (*PhysicalSpecialParameters, bool)`

GetPhysicalSpecialParametersOk returns a tuple with the PhysicalSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalSpecialParameters

`func (o *SourceSpecialParameter) SetPhysicalSpecialParameters(v PhysicalSpecialParameters)`

SetPhysicalSpecialParameters sets PhysicalSpecialParameters field to given value.

### HasPhysicalSpecialParameters

`func (o *SourceSpecialParameter) HasPhysicalSpecialParameters() bool`

HasPhysicalSpecialParameters returns a boolean if a field has been set.

### GetSkipIndexing

`func (o *SourceSpecialParameter) GetSkipIndexing() bool`

GetSkipIndexing returns the SkipIndexing field if non-nil, zero value otherwise.

### GetSkipIndexingOk

`func (o *SourceSpecialParameter) GetSkipIndexingOk() (*bool, bool)`

GetSkipIndexingOk returns a tuple with the SkipIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIndexing

`func (o *SourceSpecialParameter) SetSkipIndexing(v bool)`

SetSkipIndexing sets SkipIndexing field to given value.

### HasSkipIndexing

`func (o *SourceSpecialParameter) HasSkipIndexing() bool`

HasSkipIndexing returns a boolean if a field has been set.

### SetSkipIndexingNil

`func (o *SourceSpecialParameter) SetSkipIndexingNil(b bool)`

 SetSkipIndexingNil sets the value for SkipIndexing to be an explicit nil

### UnsetSkipIndexing
`func (o *SourceSpecialParameter) UnsetSkipIndexing()`

UnsetSkipIndexing ensures that no value is present for SkipIndexing, not even an explicit nil
### GetSourceId

`func (o *SourceSpecialParameter) GetSourceId() int64`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceSpecialParameter) GetSourceIdOk() (*int64, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceSpecialParameter) SetSourceId(v int64)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SourceSpecialParameter) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *SourceSpecialParameter) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *SourceSpecialParameter) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetSqlSpecialParameters

`func (o *SourceSpecialParameter) GetSqlSpecialParameters() ApplicationSpecialParameters`

GetSqlSpecialParameters returns the SqlSpecialParameters field if non-nil, zero value otherwise.

### GetSqlSpecialParametersOk

`func (o *SourceSpecialParameter) GetSqlSpecialParametersOk() (*ApplicationSpecialParameters, bool)`

GetSqlSpecialParametersOk returns a tuple with the SqlSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlSpecialParameters

`func (o *SourceSpecialParameter) SetSqlSpecialParameters(v ApplicationSpecialParameters)`

SetSqlSpecialParameters sets SqlSpecialParameters field to given value.

### HasSqlSpecialParameters

`func (o *SourceSpecialParameter) HasSqlSpecialParameters() bool`

HasSqlSpecialParameters returns a boolean if a field has been set.

### GetTruncateExchangeLog

`func (o *SourceSpecialParameter) GetTruncateExchangeLog() bool`

GetTruncateExchangeLog returns the TruncateExchangeLog field if non-nil, zero value otherwise.

### GetTruncateExchangeLogOk

`func (o *SourceSpecialParameter) GetTruncateExchangeLogOk() (*bool, bool)`

GetTruncateExchangeLogOk returns a tuple with the TruncateExchangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLog

`func (o *SourceSpecialParameter) SetTruncateExchangeLog(v bool)`

SetTruncateExchangeLog sets TruncateExchangeLog field to given value.

### HasTruncateExchangeLog

`func (o *SourceSpecialParameter) HasTruncateExchangeLog() bool`

HasTruncateExchangeLog returns a boolean if a field has been set.

### SetTruncateExchangeLogNil

`func (o *SourceSpecialParameter) SetTruncateExchangeLogNil(b bool)`

 SetTruncateExchangeLogNil sets the value for TruncateExchangeLog to be an explicit nil

### UnsetTruncateExchangeLog
`func (o *SourceSpecialParameter) UnsetTruncateExchangeLog()`

UnsetTruncateExchangeLog ensures that no value is present for TruncateExchangeLog, not even an explicit nil
### GetVmCredentials

`func (o *SourceSpecialParameter) GetVmCredentials() Credentials`

GetVmCredentials returns the VmCredentials field if non-nil, zero value otherwise.

### GetVmCredentialsOk

`func (o *SourceSpecialParameter) GetVmCredentialsOk() (*Credentials, bool)`

GetVmCredentialsOk returns a tuple with the VmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCredentials

`func (o *SourceSpecialParameter) SetVmCredentials(v Credentials)`

SetVmCredentials sets VmCredentials field to given value.

### HasVmCredentials

`func (o *SourceSpecialParameter) HasVmCredentials() bool`

HasVmCredentials returns a boolean if a field has been set.

### SetVmCredentialsNil

`func (o *SourceSpecialParameter) SetVmCredentialsNil(b bool)`

 SetVmCredentialsNil sets the value for VmCredentials to be an explicit nil

### UnsetVmCredentials
`func (o *SourceSpecialParameter) UnsetVmCredentials()`

UnsetVmCredentials ensures that no value is present for VmCredentials, not even an explicit nil
### GetVmwareSpecialParameters

`func (o *SourceSpecialParameter) GetVmwareSpecialParameters() VmwareSpecialParameters`

GetVmwareSpecialParameters returns the VmwareSpecialParameters field if non-nil, zero value otherwise.

### GetVmwareSpecialParametersOk

`func (o *SourceSpecialParameter) GetVmwareSpecialParametersOk() (*VmwareSpecialParameters, bool)`

GetVmwareSpecialParametersOk returns a tuple with the VmwareSpecialParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareSpecialParameters

`func (o *SourceSpecialParameter) SetVmwareSpecialParameters(v VmwareSpecialParameters)`

SetVmwareSpecialParameters sets VmwareSpecialParameters field to given value.

### HasVmwareSpecialParameters

`func (o *SourceSpecialParameter) HasVmwareSpecialParameters() bool`

HasVmwareSpecialParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


