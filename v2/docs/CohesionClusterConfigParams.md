# CohesionClusterConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceId** | Pointer to **NullableString** | Specifies the global unique appliance ID issued by AWS. | [optional] [readonly] 
**AwsControlPlaneUrl** | Pointer to **NullableString** | Specifies the AWS backup control plane URL. | [optional] [readonly] 
**AwsRegionId** | Pointer to **NullableString** | Specifies the AWS region to which this appliance is connected to. | [optional] [readonly] 
**AwsSQSUrl** | Pointer to **NullableString** | Specifies the AWS Simple Queue Service URL from which the appliance receives the necessary command messages. | [optional] [readonly] 
**AwsStorageEndpointUrl** | Pointer to **NullableString** | Specifies the AWS storage endpoint to which data will be archived. | [optional] [readonly] 

## Methods

### NewCohesionClusterConfigParams

`func NewCohesionClusterConfigParams() *CohesionClusterConfigParams`

NewCohesionClusterConfigParams instantiates a new CohesionClusterConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCohesionClusterConfigParamsWithDefaults

`func NewCohesionClusterConfigParamsWithDefaults() *CohesionClusterConfigParams`

NewCohesionClusterConfigParamsWithDefaults instantiates a new CohesionClusterConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceId

`func (o *CohesionClusterConfigParams) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *CohesionClusterConfigParams) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *CohesionClusterConfigParams) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.

### HasApplianceId

`func (o *CohesionClusterConfigParams) HasApplianceId() bool`

HasApplianceId returns a boolean if a field has been set.

### SetApplianceIdNil

`func (o *CohesionClusterConfigParams) SetApplianceIdNil(b bool)`

 SetApplianceIdNil sets the value for ApplianceId to be an explicit nil

### UnsetApplianceId
`func (o *CohesionClusterConfigParams) UnsetApplianceId()`

UnsetApplianceId ensures that no value is present for ApplianceId, not even an explicit nil
### GetAwsControlPlaneUrl

`func (o *CohesionClusterConfigParams) GetAwsControlPlaneUrl() string`

GetAwsControlPlaneUrl returns the AwsControlPlaneUrl field if non-nil, zero value otherwise.

### GetAwsControlPlaneUrlOk

`func (o *CohesionClusterConfigParams) GetAwsControlPlaneUrlOk() (*string, bool)`

GetAwsControlPlaneUrlOk returns a tuple with the AwsControlPlaneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsControlPlaneUrl

`func (o *CohesionClusterConfigParams) SetAwsControlPlaneUrl(v string)`

SetAwsControlPlaneUrl sets AwsControlPlaneUrl field to given value.

### HasAwsControlPlaneUrl

`func (o *CohesionClusterConfigParams) HasAwsControlPlaneUrl() bool`

HasAwsControlPlaneUrl returns a boolean if a field has been set.

### SetAwsControlPlaneUrlNil

`func (o *CohesionClusterConfigParams) SetAwsControlPlaneUrlNil(b bool)`

 SetAwsControlPlaneUrlNil sets the value for AwsControlPlaneUrl to be an explicit nil

### UnsetAwsControlPlaneUrl
`func (o *CohesionClusterConfigParams) UnsetAwsControlPlaneUrl()`

UnsetAwsControlPlaneUrl ensures that no value is present for AwsControlPlaneUrl, not even an explicit nil
### GetAwsRegionId

`func (o *CohesionClusterConfigParams) GetAwsRegionId() string`

GetAwsRegionId returns the AwsRegionId field if non-nil, zero value otherwise.

### GetAwsRegionIdOk

`func (o *CohesionClusterConfigParams) GetAwsRegionIdOk() (*string, bool)`

GetAwsRegionIdOk returns a tuple with the AwsRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegionId

`func (o *CohesionClusterConfigParams) SetAwsRegionId(v string)`

SetAwsRegionId sets AwsRegionId field to given value.

### HasAwsRegionId

`func (o *CohesionClusterConfigParams) HasAwsRegionId() bool`

HasAwsRegionId returns a boolean if a field has been set.

### SetAwsRegionIdNil

`func (o *CohesionClusterConfigParams) SetAwsRegionIdNil(b bool)`

 SetAwsRegionIdNil sets the value for AwsRegionId to be an explicit nil

### UnsetAwsRegionId
`func (o *CohesionClusterConfigParams) UnsetAwsRegionId()`

UnsetAwsRegionId ensures that no value is present for AwsRegionId, not even an explicit nil
### GetAwsSQSUrl

`func (o *CohesionClusterConfigParams) GetAwsSQSUrl() string`

GetAwsSQSUrl returns the AwsSQSUrl field if non-nil, zero value otherwise.

### GetAwsSQSUrlOk

`func (o *CohesionClusterConfigParams) GetAwsSQSUrlOk() (*string, bool)`

GetAwsSQSUrlOk returns a tuple with the AwsSQSUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSQSUrl

`func (o *CohesionClusterConfigParams) SetAwsSQSUrl(v string)`

SetAwsSQSUrl sets AwsSQSUrl field to given value.

### HasAwsSQSUrl

`func (o *CohesionClusterConfigParams) HasAwsSQSUrl() bool`

HasAwsSQSUrl returns a boolean if a field has been set.

### SetAwsSQSUrlNil

`func (o *CohesionClusterConfigParams) SetAwsSQSUrlNil(b bool)`

 SetAwsSQSUrlNil sets the value for AwsSQSUrl to be an explicit nil

### UnsetAwsSQSUrl
`func (o *CohesionClusterConfigParams) UnsetAwsSQSUrl()`

UnsetAwsSQSUrl ensures that no value is present for AwsSQSUrl, not even an explicit nil
### GetAwsStorageEndpointUrl

`func (o *CohesionClusterConfigParams) GetAwsStorageEndpointUrl() string`

GetAwsStorageEndpointUrl returns the AwsStorageEndpointUrl field if non-nil, zero value otherwise.

### GetAwsStorageEndpointUrlOk

`func (o *CohesionClusterConfigParams) GetAwsStorageEndpointUrlOk() (*string, bool)`

GetAwsStorageEndpointUrlOk returns a tuple with the AwsStorageEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsStorageEndpointUrl

`func (o *CohesionClusterConfigParams) SetAwsStorageEndpointUrl(v string)`

SetAwsStorageEndpointUrl sets AwsStorageEndpointUrl field to given value.

### HasAwsStorageEndpointUrl

`func (o *CohesionClusterConfigParams) HasAwsStorageEndpointUrl() bool`

HasAwsStorageEndpointUrl returns a boolean if a field has been set.

### SetAwsStorageEndpointUrlNil

`func (o *CohesionClusterConfigParams) SetAwsStorageEndpointUrlNil(b bool)`

 SetAwsStorageEndpointUrlNil sets the value for AwsStorageEndpointUrl to be an explicit nil

### UnsetAwsStorageEndpointUrl
`func (o *CohesionClusterConfigParams) UnsetAwsStorageEndpointUrl()`

UnsetAwsStorageEndpointUrl ensures that no value is present for AwsStorageEndpointUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


