# ApplicationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TruncateExchangeLog** | Pointer to **NullableBool** | If true, after the Cohesity Cluster successfully captures a Snapshot during a Job Run, the Cluster truncates the Exchange transaction logs on a Microsoft Exchange Server. The default value is false. | [optional] 

## Methods

### NewApplicationParameters

`func NewApplicationParameters() *ApplicationParameters`

NewApplicationParameters instantiates a new ApplicationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationParametersWithDefaults

`func NewApplicationParametersWithDefaults() *ApplicationParameters`

NewApplicationParametersWithDefaults instantiates a new ApplicationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTruncateExchangeLog

`func (o *ApplicationParameters) GetTruncateExchangeLog() bool`

GetTruncateExchangeLog returns the TruncateExchangeLog field if non-nil, zero value otherwise.

### GetTruncateExchangeLogOk

`func (o *ApplicationParameters) GetTruncateExchangeLogOk() (*bool, bool)`

GetTruncateExchangeLogOk returns a tuple with the TruncateExchangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLog

`func (o *ApplicationParameters) SetTruncateExchangeLog(v bool)`

SetTruncateExchangeLog sets TruncateExchangeLog field to given value.

### HasTruncateExchangeLog

`func (o *ApplicationParameters) HasTruncateExchangeLog() bool`

HasTruncateExchangeLog returns a boolean if a field has been set.

### SetTruncateExchangeLogNil

`func (o *ApplicationParameters) SetTruncateExchangeLogNil(b bool)`

 SetTruncateExchangeLogNil sets the value for TruncateExchangeLog to be an explicit nil

### UnsetTruncateExchangeLog
`func (o *ApplicationParameters) UnsetTruncateExchangeLog()`

UnsetTruncateExchangeLog ensures that no value is present for TruncateExchangeLog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


