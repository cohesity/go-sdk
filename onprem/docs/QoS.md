# QoS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | Pointer to **NullableInt64** | Specifies the name of the QoS Policy used for the View. | [optional] 
**PrincipalName** | Pointer to **NullableString** | Specifies the name of the QoS Policy used for the View such as &#39;TestAndDev High&#39;, &#39;Backup Target SSD&#39;, &#39;Backup Target High&#39; &#39;TestAndDev Low&#39; and &#39;Backup Target Low&#39;. For a complete list and descriptions, see the &#39;Create or Edit Views&#39; topic in the documentation. If not specified, the default is &#39;Backup Target Low&#39;. | [optional] 

## Methods

### NewQoS

`func NewQoS() *QoS`

NewQoS instantiates a new QoS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQoSWithDefaults

`func NewQoSWithDefaults() *QoS`

NewQoSWithDefaults instantiates a new QoS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *QoS) GetPrincipalId() int64`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *QoS) GetPrincipalIdOk() (*int64, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *QoS) SetPrincipalId(v int64)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *QoS) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### SetPrincipalIdNil

`func (o *QoS) SetPrincipalIdNil(b bool)`

 SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil

### UnsetPrincipalId
`func (o *QoS) UnsetPrincipalId()`

UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
### GetPrincipalName

`func (o *QoS) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *QoS) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *QoS) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *QoS) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### SetPrincipalNameNil

`func (o *QoS) SetPrincipalNameNil(b bool)`

 SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil

### UnsetPrincipalName
`func (o *QoS) UnsetPrincipalName()`

UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


