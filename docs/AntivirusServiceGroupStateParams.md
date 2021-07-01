# AntivirusServiceGroupStateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **NullableBool** | Specifies the enable flag to enable the Antivirus service group. | 
**Id** | **NullableInt64** | Specifies the Id of the Antivirus service group. | 

## Methods

### NewAntivirusServiceGroupStateParams

`func NewAntivirusServiceGroupStateParams(enable NullableBool, id NullableInt64, ) *AntivirusServiceGroupStateParams`

NewAntivirusServiceGroupStateParams instantiates a new AntivirusServiceGroupStateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAntivirusServiceGroupStateParamsWithDefaults

`func NewAntivirusServiceGroupStateParamsWithDefaults() *AntivirusServiceGroupStateParams`

NewAntivirusServiceGroupStateParamsWithDefaults instantiates a new AntivirusServiceGroupStateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *AntivirusServiceGroupStateParams) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *AntivirusServiceGroupStateParams) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *AntivirusServiceGroupStateParams) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### SetEnableNil

`func (o *AntivirusServiceGroupStateParams) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *AntivirusServiceGroupStateParams) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil
### GetId

`func (o *AntivirusServiceGroupStateParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AntivirusServiceGroupStateParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AntivirusServiceGroupStateParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AntivirusServiceGroupStateParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AntivirusServiceGroupStateParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


