# HeliosCloudSpinTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsParams** | Pointer to [**HeliosAwsCloudSpinParams**](HeliosAwsCloudSpinParams.md) |  | [optional] 
**AzureParams** | Pointer to [**HeliosAzureCloudSpinParams**](HeliosAzureCloudSpinParams.md) |  | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the unique id of the cloud spin entity. | [optional] 
**Name** | Pointer to **NullableString** | Specifies the name of the already added cloud spin target. | [optional] [readonly] 

## Methods

### NewHeliosCloudSpinTarget

`func NewHeliosCloudSpinTarget() *HeliosCloudSpinTarget`

NewHeliosCloudSpinTarget instantiates a new HeliosCloudSpinTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosCloudSpinTargetWithDefaults

`func NewHeliosCloudSpinTargetWithDefaults() *HeliosCloudSpinTarget`

NewHeliosCloudSpinTargetWithDefaults instantiates a new HeliosCloudSpinTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsParams

`func (o *HeliosCloudSpinTarget) GetAwsParams() HeliosAwsCloudSpinParams`

GetAwsParams returns the AwsParams field if non-nil, zero value otherwise.

### GetAwsParamsOk

`func (o *HeliosCloudSpinTarget) GetAwsParamsOk() (*HeliosAwsCloudSpinParams, bool)`

GetAwsParamsOk returns a tuple with the AwsParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsParams

`func (o *HeliosCloudSpinTarget) SetAwsParams(v HeliosAwsCloudSpinParams)`

SetAwsParams sets AwsParams field to given value.

### HasAwsParams

`func (o *HeliosCloudSpinTarget) HasAwsParams() bool`

HasAwsParams returns a boolean if a field has been set.

### GetAzureParams

`func (o *HeliosCloudSpinTarget) GetAzureParams() HeliosAzureCloudSpinParams`

GetAzureParams returns the AzureParams field if non-nil, zero value otherwise.

### GetAzureParamsOk

`func (o *HeliosCloudSpinTarget) GetAzureParamsOk() (*HeliosAzureCloudSpinParams, bool)`

GetAzureParamsOk returns a tuple with the AzureParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureParams

`func (o *HeliosCloudSpinTarget) SetAzureParams(v HeliosAzureCloudSpinParams)`

SetAzureParams sets AzureParams field to given value.

### HasAzureParams

`func (o *HeliosCloudSpinTarget) HasAzureParams() bool`

HasAzureParams returns a boolean if a field has been set.

### GetId

`func (o *HeliosCloudSpinTarget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeliosCloudSpinTarget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeliosCloudSpinTarget) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HeliosCloudSpinTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HeliosCloudSpinTarget) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HeliosCloudSpinTarget) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *HeliosCloudSpinTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeliosCloudSpinTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeliosCloudSpinTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeliosCloudSpinTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HeliosCloudSpinTarget) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HeliosCloudSpinTarget) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


