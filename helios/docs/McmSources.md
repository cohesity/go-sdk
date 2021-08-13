# McmSources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]McmSource**](McmSource.md) | Specifies the list of Protection Sources. | [optional] 

## Methods

### NewMcmSources

`func NewMcmSources() *McmSources`

NewMcmSources instantiates a new McmSources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmSourcesWithDefaults

`func NewMcmSourcesWithDefaults() *McmSources`

NewMcmSourcesWithDefaults instantiates a new McmSources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *McmSources) GetSources() []McmSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *McmSources) GetSourcesOk() (*[]McmSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *McmSources) SetSources(v []McmSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *McmSources) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *McmSources) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *McmSources) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


