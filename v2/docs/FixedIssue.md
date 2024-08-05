# FixedIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Specifies a unique number of the bug. | [optional] 
**ReleaseNote** | Pointer to **string** | Specifies the description of fix made for the issue. | [optional] 

## Methods

### NewFixedIssue

`func NewFixedIssue() *FixedIssue`

NewFixedIssue instantiates a new FixedIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedIssueWithDefaults

`func NewFixedIssueWithDefaults() *FixedIssue`

NewFixedIssueWithDefaults instantiates a new FixedIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FixedIssue) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedIssue) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedIssue) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FixedIssue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReleaseNote

`func (o *FixedIssue) GetReleaseNote() string`

GetReleaseNote returns the ReleaseNote field if non-nil, zero value otherwise.

### GetReleaseNoteOk

`func (o *FixedIssue) GetReleaseNoteOk() (*string, bool)`

GetReleaseNoteOk returns a tuple with the ReleaseNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNote

`func (o *FixedIssue) SetReleaseNote(v string)`

SetReleaseNote sets ReleaseNote field to given value.

### HasReleaseNote

`func (o *FixedIssue) HasReleaseNote() bool`

HasReleaseNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


