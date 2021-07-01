# PatternRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationDataType** | **NullableInt32** | Specifies the data type for which supported patterns can be fetched. | 
**ApplicationId** | **NullableInt64** | Specifies AWB Application ID. | 
**UserPattern** | [**SupportedPattern**](SupportedPattern.md) |  | 

## Methods

### NewPatternRequestBody

`func NewPatternRequestBody(applicationDataType NullableInt32, applicationId NullableInt64, userPattern SupportedPattern, ) *PatternRequestBody`

NewPatternRequestBody instantiates a new PatternRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatternRequestBodyWithDefaults

`func NewPatternRequestBodyWithDefaults() *PatternRequestBody`

NewPatternRequestBodyWithDefaults instantiates a new PatternRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationDataType

`func (o *PatternRequestBody) GetApplicationDataType() int32`

GetApplicationDataType returns the ApplicationDataType field if non-nil, zero value otherwise.

### GetApplicationDataTypeOk

`func (o *PatternRequestBody) GetApplicationDataTypeOk() (*int32, bool)`

GetApplicationDataTypeOk returns a tuple with the ApplicationDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDataType

`func (o *PatternRequestBody) SetApplicationDataType(v int32)`

SetApplicationDataType sets ApplicationDataType field to given value.


### SetApplicationDataTypeNil

`func (o *PatternRequestBody) SetApplicationDataTypeNil(b bool)`

 SetApplicationDataTypeNil sets the value for ApplicationDataType to be an explicit nil

### UnsetApplicationDataType
`func (o *PatternRequestBody) UnsetApplicationDataType()`

UnsetApplicationDataType ensures that no value is present for ApplicationDataType, not even an explicit nil
### GetApplicationId

`func (o *PatternRequestBody) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PatternRequestBody) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PatternRequestBody) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### SetApplicationIdNil

`func (o *PatternRequestBody) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *PatternRequestBody) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetUserPattern

`func (o *PatternRequestBody) GetUserPattern() SupportedPattern`

GetUserPattern returns the UserPattern field if non-nil, zero value otherwise.

### GetUserPatternOk

`func (o *PatternRequestBody) GetUserPatternOk() (*SupportedPattern, bool)`

GetUserPatternOk returns a tuple with the UserPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPattern

`func (o *PatternRequestBody) SetUserPattern(v SupportedPattern)`

SetUserPattern sets UserPattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


