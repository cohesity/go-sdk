# ADAttributeRestoreParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedPropertyVec** | Pointer to **[]string** | Array of LDAP property names to excluded from &#39;property_vec&#39;. Excluded property name cannot contain wildcard character &#39;*&#39;.  Property names are case insensitive. | [optional] 
**GuidpairVec** | Pointer to [**[]ADGuidPair**](ADGuidPair.md) | Array of source and destination object guid pairs to restore attributes. Pair source guid refers to guid in AD snapshot in source_server endpoint. Destination guid refers to guid in production AD. If destination guid is empty, then source guid in AD snapshot should exist in production AD. | [optional] 
**OptionFlags** | Pointer to **NullableInt32** | Attribute restore option flags of type ADAttributeOptionFlags. | [optional] 
**PropertyVec** | Pointer to **[]string** | Array of LDAP property(attribute) names. The name can be standard or custom property defined in AD schema partition. The property can contain wildcard character &#39;*&#39;. If this array is empty, then &#39;*&#39; is assigned, means restore all properties except default system excluded properties. Wildcards will be expanded. If &#39;memberOf&#39; property is included, group membership of the objects specified in &#39;guid_vec&#39; will be restored. Property that does not exist for an object is ignored and no error info is returned for that property. Property names are case insensitive. Caller may check the ADAttributeFlags.kSystem obtained during object compare to exclude system properties. | [optional] 
**SrcSysvolFolder** | Pointer to **NullableString** | When restoring a GPO, need to know the absolute path for SYSVOL folder. | [optional] 

## Methods

### NewADAttributeRestoreParam

`func NewADAttributeRestoreParam() *ADAttributeRestoreParam`

NewADAttributeRestoreParam instantiates a new ADAttributeRestoreParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADAttributeRestoreParamWithDefaults

`func NewADAttributeRestoreParamWithDefaults() *ADAttributeRestoreParam`

NewADAttributeRestoreParamWithDefaults instantiates a new ADAttributeRestoreParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedPropertyVec

`func (o *ADAttributeRestoreParam) GetExcludedPropertyVec() []string`

GetExcludedPropertyVec returns the ExcludedPropertyVec field if non-nil, zero value otherwise.

### GetExcludedPropertyVecOk

`func (o *ADAttributeRestoreParam) GetExcludedPropertyVecOk() (*[]string, bool)`

GetExcludedPropertyVecOk returns a tuple with the ExcludedPropertyVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPropertyVec

`func (o *ADAttributeRestoreParam) SetExcludedPropertyVec(v []string)`

SetExcludedPropertyVec sets ExcludedPropertyVec field to given value.

### HasExcludedPropertyVec

`func (o *ADAttributeRestoreParam) HasExcludedPropertyVec() bool`

HasExcludedPropertyVec returns a boolean if a field has been set.

### SetExcludedPropertyVecNil

`func (o *ADAttributeRestoreParam) SetExcludedPropertyVecNil(b bool)`

 SetExcludedPropertyVecNil sets the value for ExcludedPropertyVec to be an explicit nil

### UnsetExcludedPropertyVec
`func (o *ADAttributeRestoreParam) UnsetExcludedPropertyVec()`

UnsetExcludedPropertyVec ensures that no value is present for ExcludedPropertyVec, not even an explicit nil
### GetGuidpairVec

`func (o *ADAttributeRestoreParam) GetGuidpairVec() []ADGuidPair`

GetGuidpairVec returns the GuidpairVec field if non-nil, zero value otherwise.

### GetGuidpairVecOk

`func (o *ADAttributeRestoreParam) GetGuidpairVecOk() (*[]ADGuidPair, bool)`

GetGuidpairVecOk returns a tuple with the GuidpairVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidpairVec

`func (o *ADAttributeRestoreParam) SetGuidpairVec(v []ADGuidPair)`

SetGuidpairVec sets GuidpairVec field to given value.

### HasGuidpairVec

`func (o *ADAttributeRestoreParam) HasGuidpairVec() bool`

HasGuidpairVec returns a boolean if a field has been set.

### SetGuidpairVecNil

`func (o *ADAttributeRestoreParam) SetGuidpairVecNil(b bool)`

 SetGuidpairVecNil sets the value for GuidpairVec to be an explicit nil

### UnsetGuidpairVec
`func (o *ADAttributeRestoreParam) UnsetGuidpairVec()`

UnsetGuidpairVec ensures that no value is present for GuidpairVec, not even an explicit nil
### GetOptionFlags

`func (o *ADAttributeRestoreParam) GetOptionFlags() int32`

GetOptionFlags returns the OptionFlags field if non-nil, zero value otherwise.

### GetOptionFlagsOk

`func (o *ADAttributeRestoreParam) GetOptionFlagsOk() (*int32, bool)`

GetOptionFlagsOk returns a tuple with the OptionFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFlags

`func (o *ADAttributeRestoreParam) SetOptionFlags(v int32)`

SetOptionFlags sets OptionFlags field to given value.

### HasOptionFlags

`func (o *ADAttributeRestoreParam) HasOptionFlags() bool`

HasOptionFlags returns a boolean if a field has been set.

### SetOptionFlagsNil

`func (o *ADAttributeRestoreParam) SetOptionFlagsNil(b bool)`

 SetOptionFlagsNil sets the value for OptionFlags to be an explicit nil

### UnsetOptionFlags
`func (o *ADAttributeRestoreParam) UnsetOptionFlags()`

UnsetOptionFlags ensures that no value is present for OptionFlags, not even an explicit nil
### GetPropertyVec

`func (o *ADAttributeRestoreParam) GetPropertyVec() []string`

GetPropertyVec returns the PropertyVec field if non-nil, zero value otherwise.

### GetPropertyVecOk

`func (o *ADAttributeRestoreParam) GetPropertyVecOk() (*[]string, bool)`

GetPropertyVecOk returns a tuple with the PropertyVec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyVec

`func (o *ADAttributeRestoreParam) SetPropertyVec(v []string)`

SetPropertyVec sets PropertyVec field to given value.

### HasPropertyVec

`func (o *ADAttributeRestoreParam) HasPropertyVec() bool`

HasPropertyVec returns a boolean if a field has been set.

### SetPropertyVecNil

`func (o *ADAttributeRestoreParam) SetPropertyVecNil(b bool)`

 SetPropertyVecNil sets the value for PropertyVec to be an explicit nil

### UnsetPropertyVec
`func (o *ADAttributeRestoreParam) UnsetPropertyVec()`

UnsetPropertyVec ensures that no value is present for PropertyVec, not even an explicit nil
### GetSrcSysvolFolder

`func (o *ADAttributeRestoreParam) GetSrcSysvolFolder() string`

GetSrcSysvolFolder returns the SrcSysvolFolder field if non-nil, zero value otherwise.

### GetSrcSysvolFolderOk

`func (o *ADAttributeRestoreParam) GetSrcSysvolFolderOk() (*string, bool)`

GetSrcSysvolFolderOk returns a tuple with the SrcSysvolFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcSysvolFolder

`func (o *ADAttributeRestoreParam) SetSrcSysvolFolder(v string)`

SetSrcSysvolFolder sets SrcSysvolFolder field to given value.

### HasSrcSysvolFolder

`func (o *ADAttributeRestoreParam) HasSrcSysvolFolder() bool`

HasSrcSysvolFolder returns a boolean if a field has been set.

### SetSrcSysvolFolderNil

`func (o *ADAttributeRestoreParam) SetSrcSysvolFolderNil(b bool)`

 SetSrcSysvolFolderNil sets the value for SrcSysvolFolder to be an explicit nil

### UnsetSrcSysvolFolder
`func (o *ADAttributeRestoreParam) UnsetSrcSysvolFolder()`

UnsetSrcSysvolFolder ensures that no value is present for SrcSysvolFolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


