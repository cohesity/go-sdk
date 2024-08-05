# IpmiSel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelNonVerboseEntries** | Pointer to [**[]IpmiSelNonVerboseEntry**](IpmiSelNonVerboseEntry.md) | Specifies the list of sel entries when verbose&#x3D;false. | [optional] 
**SelVerboseEntries** | Pointer to [**[]IpmiSelVerboseEntry**](IpmiSelVerboseEntry.md) | Specifies the list of sel entries when verbose&#x3D;true. | [optional] 

## Methods

### NewIpmiSel

`func NewIpmiSel() *IpmiSel`

NewIpmiSel instantiates a new IpmiSel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpmiSelWithDefaults

`func NewIpmiSelWithDefaults() *IpmiSel`

NewIpmiSelWithDefaults instantiates a new IpmiSel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelNonVerboseEntries

`func (o *IpmiSel) GetSelNonVerboseEntries() []IpmiSelNonVerboseEntry`

GetSelNonVerboseEntries returns the SelNonVerboseEntries field if non-nil, zero value otherwise.

### GetSelNonVerboseEntriesOk

`func (o *IpmiSel) GetSelNonVerboseEntriesOk() (*[]IpmiSelNonVerboseEntry, bool)`

GetSelNonVerboseEntriesOk returns a tuple with the SelNonVerboseEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelNonVerboseEntries

`func (o *IpmiSel) SetSelNonVerboseEntries(v []IpmiSelNonVerboseEntry)`

SetSelNonVerboseEntries sets SelNonVerboseEntries field to given value.

### HasSelNonVerboseEntries

`func (o *IpmiSel) HasSelNonVerboseEntries() bool`

HasSelNonVerboseEntries returns a boolean if a field has been set.

### SetSelNonVerboseEntriesNil

`func (o *IpmiSel) SetSelNonVerboseEntriesNil(b bool)`

 SetSelNonVerboseEntriesNil sets the value for SelNonVerboseEntries to be an explicit nil

### UnsetSelNonVerboseEntries
`func (o *IpmiSel) UnsetSelNonVerboseEntries()`

UnsetSelNonVerboseEntries ensures that no value is present for SelNonVerboseEntries, not even an explicit nil
### GetSelVerboseEntries

`func (o *IpmiSel) GetSelVerboseEntries() []IpmiSelVerboseEntry`

GetSelVerboseEntries returns the SelVerboseEntries field if non-nil, zero value otherwise.

### GetSelVerboseEntriesOk

`func (o *IpmiSel) GetSelVerboseEntriesOk() (*[]IpmiSelVerboseEntry, bool)`

GetSelVerboseEntriesOk returns a tuple with the SelVerboseEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelVerboseEntries

`func (o *IpmiSel) SetSelVerboseEntries(v []IpmiSelVerboseEntry)`

SetSelVerboseEntries sets SelVerboseEntries field to given value.

### HasSelVerboseEntries

`func (o *IpmiSel) HasSelVerboseEntries() bool`

HasSelVerboseEntries returns a boolean if a field has been set.

### SetSelVerboseEntriesNil

`func (o *IpmiSel) SetSelVerboseEntriesNil(b bool)`

 SetSelVerboseEntriesNil sets the value for SelVerboseEntries to be an explicit nil

### UnsetSelVerboseEntries
`func (o *IpmiSel) UnsetSelVerboseEntries()`

UnsetSelVerboseEntries ensures that no value is present for SelVerboseEntries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


