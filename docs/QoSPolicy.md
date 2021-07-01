# QoSPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysUseSsd** | Pointer to **NullableBool** | Specifies whether to always write to SSD even if SeqWriteSsdPct is 0. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies Id of the QoS Policy. | [optional] 
**MinRequests** | Pointer to **NullableInt32** | Specifies minimum number of requests,  corresponding to this Policy, executed in the QoS queue. | [optional] 
**Name** | Pointer to **NullableString** | Specifies Name of the Qos Policy. | [optional] 
**Priority** | Pointer to **NullableString** | Specifies Priority of the Qos Policy. Priority of QoS Policy as defined in cluster config proto. | [optional] 
**RandomWriteHydraPct** | Pointer to **NullableInt32** | Specifies percentage of a random write request belonging to this Policy that hits hydra. | [optional] 
**RandomWriteSsdPct** | Pointer to **NullableInt32** | Specifies percentage of a random write request belonging to this Policy that hits SSD. | [optional] 
**SeqWriteHydraPct** | Pointer to **NullableInt32** | Specifies percentage of a sequential write request belonging to this Policy that hits hydra. | [optional] 
**SeqWriteSsdPct** | Pointer to **NullableInt32** | Specifies percentage of a sequential write request belonging to this Policy that hits SSD. | [optional] 
**Weight** | Pointer to **NullableInt32** | Specifies Weight of the QoS Policy used in QoS queue. | [optional] 
**WorkLoadType** | Pointer to **NullableString** | Specifies Workload type attribute associated with this Policy. | [optional] 

## Methods

### NewQoSPolicy

`func NewQoSPolicy() *QoSPolicy`

NewQoSPolicy instantiates a new QoSPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoSPolicyWithDefaults

`func NewQoSPolicyWithDefaults() *QoSPolicy`

NewQoSPolicyWithDefaults instantiates a new QoSPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysUseSsd

`func (o *QoSPolicy) GetAlwaysUseSsd() bool`

GetAlwaysUseSsd returns the AlwaysUseSsd field if non-nil, zero value otherwise.

### GetAlwaysUseSsdOk

`func (o *QoSPolicy) GetAlwaysUseSsdOk() (*bool, bool)`

GetAlwaysUseSsdOk returns a tuple with the AlwaysUseSsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUseSsd

`func (o *QoSPolicy) SetAlwaysUseSsd(v bool)`

SetAlwaysUseSsd sets AlwaysUseSsd field to given value.

### HasAlwaysUseSsd

`func (o *QoSPolicy) HasAlwaysUseSsd() bool`

HasAlwaysUseSsd returns a boolean if a field has been set.

### SetAlwaysUseSsdNil

`func (o *QoSPolicy) SetAlwaysUseSsdNil(b bool)`

 SetAlwaysUseSsdNil sets the value for AlwaysUseSsd to be an explicit nil

### UnsetAlwaysUseSsd
`func (o *QoSPolicy) UnsetAlwaysUseSsd()`

UnsetAlwaysUseSsd ensures that no value is present for AlwaysUseSsd, not even an explicit nil
### GetId

`func (o *QoSPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QoSPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QoSPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QoSPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *QoSPolicy) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *QoSPolicy) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMinRequests

`func (o *QoSPolicy) GetMinRequests() int32`

GetMinRequests returns the MinRequests field if non-nil, zero value otherwise.

### GetMinRequestsOk

`func (o *QoSPolicy) GetMinRequestsOk() (*int32, bool)`

GetMinRequestsOk returns a tuple with the MinRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRequests

`func (o *QoSPolicy) SetMinRequests(v int32)`

SetMinRequests sets MinRequests field to given value.

### HasMinRequests

`func (o *QoSPolicy) HasMinRequests() bool`

HasMinRequests returns a boolean if a field has been set.

### SetMinRequestsNil

`func (o *QoSPolicy) SetMinRequestsNil(b bool)`

 SetMinRequestsNil sets the value for MinRequests to be an explicit nil

### UnsetMinRequests
`func (o *QoSPolicy) UnsetMinRequests()`

UnsetMinRequests ensures that no value is present for MinRequests, not even an explicit nil
### GetName

`func (o *QoSPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QoSPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QoSPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QoSPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *QoSPolicy) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *QoSPolicy) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPriority

`func (o *QoSPolicy) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *QoSPolicy) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *QoSPolicy) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *QoSPolicy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *QoSPolicy) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *QoSPolicy) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetRandomWriteHydraPct

`func (o *QoSPolicy) GetRandomWriteHydraPct() int32`

GetRandomWriteHydraPct returns the RandomWriteHydraPct field if non-nil, zero value otherwise.

### GetRandomWriteHydraPctOk

`func (o *QoSPolicy) GetRandomWriteHydraPctOk() (*int32, bool)`

GetRandomWriteHydraPctOk returns a tuple with the RandomWriteHydraPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomWriteHydraPct

`func (o *QoSPolicy) SetRandomWriteHydraPct(v int32)`

SetRandomWriteHydraPct sets RandomWriteHydraPct field to given value.

### HasRandomWriteHydraPct

`func (o *QoSPolicy) HasRandomWriteHydraPct() bool`

HasRandomWriteHydraPct returns a boolean if a field has been set.

### SetRandomWriteHydraPctNil

`func (o *QoSPolicy) SetRandomWriteHydraPctNil(b bool)`

 SetRandomWriteHydraPctNil sets the value for RandomWriteHydraPct to be an explicit nil

### UnsetRandomWriteHydraPct
`func (o *QoSPolicy) UnsetRandomWriteHydraPct()`

UnsetRandomWriteHydraPct ensures that no value is present for RandomWriteHydraPct, not even an explicit nil
### GetRandomWriteSsdPct

`func (o *QoSPolicy) GetRandomWriteSsdPct() int32`

GetRandomWriteSsdPct returns the RandomWriteSsdPct field if non-nil, zero value otherwise.

### GetRandomWriteSsdPctOk

`func (o *QoSPolicy) GetRandomWriteSsdPctOk() (*int32, bool)`

GetRandomWriteSsdPctOk returns a tuple with the RandomWriteSsdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomWriteSsdPct

`func (o *QoSPolicy) SetRandomWriteSsdPct(v int32)`

SetRandomWriteSsdPct sets RandomWriteSsdPct field to given value.

### HasRandomWriteSsdPct

`func (o *QoSPolicy) HasRandomWriteSsdPct() bool`

HasRandomWriteSsdPct returns a boolean if a field has been set.

### SetRandomWriteSsdPctNil

`func (o *QoSPolicy) SetRandomWriteSsdPctNil(b bool)`

 SetRandomWriteSsdPctNil sets the value for RandomWriteSsdPct to be an explicit nil

### UnsetRandomWriteSsdPct
`func (o *QoSPolicy) UnsetRandomWriteSsdPct()`

UnsetRandomWriteSsdPct ensures that no value is present for RandomWriteSsdPct, not even an explicit nil
### GetSeqWriteHydraPct

`func (o *QoSPolicy) GetSeqWriteHydraPct() int32`

GetSeqWriteHydraPct returns the SeqWriteHydraPct field if non-nil, zero value otherwise.

### GetSeqWriteHydraPctOk

`func (o *QoSPolicy) GetSeqWriteHydraPctOk() (*int32, bool)`

GetSeqWriteHydraPctOk returns a tuple with the SeqWriteHydraPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqWriteHydraPct

`func (o *QoSPolicy) SetSeqWriteHydraPct(v int32)`

SetSeqWriteHydraPct sets SeqWriteHydraPct field to given value.

### HasSeqWriteHydraPct

`func (o *QoSPolicy) HasSeqWriteHydraPct() bool`

HasSeqWriteHydraPct returns a boolean if a field has been set.

### SetSeqWriteHydraPctNil

`func (o *QoSPolicy) SetSeqWriteHydraPctNil(b bool)`

 SetSeqWriteHydraPctNil sets the value for SeqWriteHydraPct to be an explicit nil

### UnsetSeqWriteHydraPct
`func (o *QoSPolicy) UnsetSeqWriteHydraPct()`

UnsetSeqWriteHydraPct ensures that no value is present for SeqWriteHydraPct, not even an explicit nil
### GetSeqWriteSsdPct

`func (o *QoSPolicy) GetSeqWriteSsdPct() int32`

GetSeqWriteSsdPct returns the SeqWriteSsdPct field if non-nil, zero value otherwise.

### GetSeqWriteSsdPctOk

`func (o *QoSPolicy) GetSeqWriteSsdPctOk() (*int32, bool)`

GetSeqWriteSsdPctOk returns a tuple with the SeqWriteSsdPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqWriteSsdPct

`func (o *QoSPolicy) SetSeqWriteSsdPct(v int32)`

SetSeqWriteSsdPct sets SeqWriteSsdPct field to given value.

### HasSeqWriteSsdPct

`func (o *QoSPolicy) HasSeqWriteSsdPct() bool`

HasSeqWriteSsdPct returns a boolean if a field has been set.

### SetSeqWriteSsdPctNil

`func (o *QoSPolicy) SetSeqWriteSsdPctNil(b bool)`

 SetSeqWriteSsdPctNil sets the value for SeqWriteSsdPct to be an explicit nil

### UnsetSeqWriteSsdPct
`func (o *QoSPolicy) UnsetSeqWriteSsdPct()`

UnsetSeqWriteSsdPct ensures that no value is present for SeqWriteSsdPct, not even an explicit nil
### GetWeight

`func (o *QoSPolicy) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *QoSPolicy) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *QoSPolicy) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *QoSPolicy) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *QoSPolicy) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *QoSPolicy) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWorkLoadType

`func (o *QoSPolicy) GetWorkLoadType() string`

GetWorkLoadType returns the WorkLoadType field if non-nil, zero value otherwise.

### GetWorkLoadTypeOk

`func (o *QoSPolicy) GetWorkLoadTypeOk() (*string, bool)`

GetWorkLoadTypeOk returns a tuple with the WorkLoadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLoadType

`func (o *QoSPolicy) SetWorkLoadType(v string)`

SetWorkLoadType sets WorkLoadType field to given value.

### HasWorkLoadType

`func (o *QoSPolicy) HasWorkLoadType() bool`

HasWorkLoadType returns a boolean if a field has been set.

### SetWorkLoadTypeNil

`func (o *QoSPolicy) SetWorkLoadTypeNil(b bool)`

 SetWorkLoadTypeNil sets the value for WorkLoadType to be an explicit nil

### UnsetWorkLoadType
`func (o *QoSPolicy) UnsetWorkLoadType()`

UnsetWorkLoadType ensures that no value is present for WorkLoadType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


