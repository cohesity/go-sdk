// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourceRegistrationUpdateRequestParams Source Registration update parameters.
//
// Specifies the Source registration Update request parameters.
//
// swagger:model SourceRegistrationUpdateRequestParams
type SourceRegistrationUpdateRequestParams struct {
	CommonSourceRegistrationRequestParams

	// Specifies the parameters to register a VMware Protection Source.
	VmwareParams *VmwareSourceRegistrationParams `json:"vmwareParams,omitempty"`

	// Specifies the parameters to register a Physical Protection Source.
	PhysicalParams *PhysicalSourceRegistrationParams `json:"physicalParams,omitempty"`

	// Specifies the parameters to register a Generic Nas Source.
	GenericNasParams *GenericNasRegistrationParams `json:"genericNasParams,omitempty"`

	// Specifies the parameters to register an Isilon Source.
	IsilonParams *IsilonRegistrationParams `json:"isilonParams,omitempty"`

	// Specifies the parameters to register an Netapp Source.
	NetappParams *NetappRegistrationParams `json:"netappParams,omitempty"`

	// Specifies the parameters to register an Elastifile Source.
	ElastifileParams *ElastifileRegistrationParams `json:"elastifileParams,omitempty"`

	// Specifies the parameters to register an Flashblade Source.
	FlashbladeParams *FlashbladeRegistrationParams `json:"flashbladeParams,omitempty"`

	// Specifies the parameters to register an GPFS Source.
	GpfsParams *GpfsRegistrationParams `json:"gpfsParams,omitempty"`

	// Specifies the parameters to update the registration of a Cassandra Protection Source.
	CassandraParams *CassandraSourceRegistrationParams `json:"cassandraParams,omitempty"`

	// Specifies the parameters to register a MongoDB Protection Source.
	MongodbParams *MongoDBSourceRegistrationParams `json:"mongodbParams,omitempty"`

	// Specifies the parameters to register a Couchbase Protection Source.
	CouchbaseParams *CouchbaseSourceRegistrationParams `json:"couchbaseParams,omitempty"`

	// Specifies the parameters to update registration of an HDFS Protection Source. Update for HDFS supports two operations - Reconfigure and Update Kerberos. For the Reconfigure operation you have to specify host, config directory and either SSHCredentials or SSHPrivateKey. All other fields in the request will be ignored. For the Update Kerberos operation you should set only the Kerberos field and not set any other field.
	HdfsParams *HdfsSourceRegistrationParams `json:"hdfsParams,omitempty"`

	// Specifies the parameters to update registration of an Hbase Protection Source. Update Hbase supports two operations - Reconfigure and Update Kerberos. For the Reconfigure operation you have to specify host, config directory and either SSHCredentials or SSHPrivateKey. All other fields in the request will be ignored. For the Update Kerberos operation you should set only the Kerberos field and not set any other field.
	HbaseParams *HbaseSourceRegistrationParams `json:"hbaseParams,omitempty"`

	// Specifies the parameters to update registration of an Hive Protection Source. Update Hive supports two operations - Reconfigure and Update Kerberos. For the Reconfigure operation you have to specify host, config directory and either SSHCredentials or SSHPrivateKey. All other fields in the request will be ignored. For the Update Kerberos operation you should set only the Kerberos field and not set any other field.
	HiveParams *HiveSourceRegistrationParams `json:"hiveParams,omitempty"`

	// Specifies the parameters to update the registration of a Universal Data Adapter Protection Source.
	UdaParams *UdaSourceRegistrationParams `json:"udaParams,omitempty"`

	// Specifies the parameters to update the registration of an office 365 Source.
	Office365Params *Office365SourceRegistrationParams `json:"office365Params,omitempty"`

	// Specifies the parameters to register an AWS source.
	AwsParams *AwsSourceRegistrationParams `json:"awsParams,omitempty"`

	// Specifies the parameters to register a HyperV Protection Source.
	HypervParams *HyperVSourceRegistrationParams `json:"hypervParams,omitempty"`

	// Specifies the parameters to update the registration of a SFDC Adapter Protection Source.
	SfdcParams *SfdcSourceRegistrationParams `json:"sfdcParams,omitempty"`

	// Specifies the last time this protection source was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection source was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.
	LastModifiedTimestampUsecs *int64 `json:"lastModifiedTimestampUsecs,omitempty"`

	// Specifies the parameters to register an Azure source.
	AzureParams *AzureSourceRegistrationParams `json:"azureParams,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SourceRegistrationUpdateRequestParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CommonSourceRegistrationRequestParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CommonSourceRegistrationRequestParams = aO0

	// AO1
	var dataAO1 struct {
		VmwareParams *VmwareSourceRegistrationParams `json:"vmwareParams,omitempty"`

		PhysicalParams *PhysicalSourceRegistrationParams `json:"physicalParams,omitempty"`

		GenericNasParams *GenericNasRegistrationParams `json:"genericNasParams,omitempty"`

		IsilonParams *IsilonRegistrationParams `json:"isilonParams,omitempty"`

		NetappParams *NetappRegistrationParams `json:"netappParams,omitempty"`

		ElastifileParams *ElastifileRegistrationParams `json:"elastifileParams,omitempty"`

		FlashbladeParams *FlashbladeRegistrationParams `json:"flashbladeParams,omitempty"`

		GpfsParams *GpfsRegistrationParams `json:"gpfsParams,omitempty"`

		CassandraParams *CassandraSourceRegistrationParams `json:"cassandraParams,omitempty"`

		MongodbParams *MongoDBSourceRegistrationParams `json:"mongodbParams,omitempty"`

		CouchbaseParams *CouchbaseSourceRegistrationParams `json:"couchbaseParams,omitempty"`

		HdfsParams *HdfsSourceRegistrationParams `json:"hdfsParams,omitempty"`

		HbaseParams *HbaseSourceRegistrationParams `json:"hbaseParams,omitempty"`

		HiveParams *HiveSourceRegistrationParams `json:"hiveParams,omitempty"`

		UdaParams *UdaSourceRegistrationParams `json:"udaParams,omitempty"`

		Office365Params *Office365SourceRegistrationParams `json:"office365Params,omitempty"`

		AwsParams *AwsSourceRegistrationParams `json:"awsParams,omitempty"`

		HypervParams *HyperVSourceRegistrationParams `json:"hypervParams,omitempty"`

		SfdcParams *SfdcSourceRegistrationParams `json:"sfdcParams,omitempty"`

		LastModifiedTimestampUsecs *int64 `json:"lastModifiedTimestampUsecs,omitempty"`

		AzureParams *AzureSourceRegistrationParams `json:"azureParams,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.VmwareParams = dataAO1.VmwareParams

	m.PhysicalParams = dataAO1.PhysicalParams

	m.GenericNasParams = dataAO1.GenericNasParams

	m.IsilonParams = dataAO1.IsilonParams

	m.NetappParams = dataAO1.NetappParams

	m.ElastifileParams = dataAO1.ElastifileParams

	m.FlashbladeParams = dataAO1.FlashbladeParams

	m.GpfsParams = dataAO1.GpfsParams

	m.CassandraParams = dataAO1.CassandraParams

	m.MongodbParams = dataAO1.MongodbParams

	m.CouchbaseParams = dataAO1.CouchbaseParams

	m.HdfsParams = dataAO1.HdfsParams

	m.HbaseParams = dataAO1.HbaseParams

	m.HiveParams = dataAO1.HiveParams

	m.UdaParams = dataAO1.UdaParams

	m.Office365Params = dataAO1.Office365Params

	m.AwsParams = dataAO1.AwsParams

	m.HypervParams = dataAO1.HypervParams

	m.SfdcParams = dataAO1.SfdcParams

	m.LastModifiedTimestampUsecs = dataAO1.LastModifiedTimestampUsecs

	m.AzureParams = dataAO1.AzureParams

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SourceRegistrationUpdateRequestParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CommonSourceRegistrationRequestParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		VmwareParams *VmwareSourceRegistrationParams `json:"vmwareParams,omitempty"`

		PhysicalParams *PhysicalSourceRegistrationParams `json:"physicalParams,omitempty"`

		GenericNasParams *GenericNasRegistrationParams `json:"genericNasParams,omitempty"`

		IsilonParams *IsilonRegistrationParams `json:"isilonParams,omitempty"`

		NetappParams *NetappRegistrationParams `json:"netappParams,omitempty"`

		ElastifileParams *ElastifileRegistrationParams `json:"elastifileParams,omitempty"`

		FlashbladeParams *FlashbladeRegistrationParams `json:"flashbladeParams,omitempty"`

		GpfsParams *GpfsRegistrationParams `json:"gpfsParams,omitempty"`

		CassandraParams *CassandraSourceRegistrationParams `json:"cassandraParams,omitempty"`

		MongodbParams *MongoDBSourceRegistrationParams `json:"mongodbParams,omitempty"`

		CouchbaseParams *CouchbaseSourceRegistrationParams `json:"couchbaseParams,omitempty"`

		HdfsParams *HdfsSourceRegistrationParams `json:"hdfsParams,omitempty"`

		HbaseParams *HbaseSourceRegistrationParams `json:"hbaseParams,omitempty"`

		HiveParams *HiveSourceRegistrationParams `json:"hiveParams,omitempty"`

		UdaParams *UdaSourceRegistrationParams `json:"udaParams,omitempty"`

		Office365Params *Office365SourceRegistrationParams `json:"office365Params,omitempty"`

		AwsParams *AwsSourceRegistrationParams `json:"awsParams,omitempty"`

		HypervParams *HyperVSourceRegistrationParams `json:"hypervParams,omitempty"`

		SfdcParams *SfdcSourceRegistrationParams `json:"sfdcParams,omitempty"`

		LastModifiedTimestampUsecs *int64 `json:"lastModifiedTimestampUsecs,omitempty"`

		AzureParams *AzureSourceRegistrationParams `json:"azureParams,omitempty"`
	}

	dataAO1.VmwareParams = m.VmwareParams

	dataAO1.PhysicalParams = m.PhysicalParams

	dataAO1.GenericNasParams = m.GenericNasParams

	dataAO1.IsilonParams = m.IsilonParams

	dataAO1.NetappParams = m.NetappParams

	dataAO1.ElastifileParams = m.ElastifileParams

	dataAO1.FlashbladeParams = m.FlashbladeParams

	dataAO1.GpfsParams = m.GpfsParams

	dataAO1.CassandraParams = m.CassandraParams

	dataAO1.MongodbParams = m.MongodbParams

	dataAO1.CouchbaseParams = m.CouchbaseParams

	dataAO1.HdfsParams = m.HdfsParams

	dataAO1.HbaseParams = m.HbaseParams

	dataAO1.HiveParams = m.HiveParams

	dataAO1.UdaParams = m.UdaParams

	dataAO1.Office365Params = m.Office365Params

	dataAO1.AwsParams = m.AwsParams

	dataAO1.HypervParams = m.HypervParams

	dataAO1.SfdcParams = m.SfdcParams

	dataAO1.LastModifiedTimestampUsecs = m.LastModifiedTimestampUsecs

	dataAO1.AzureParams = m.AzureParams

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this source registration update request params
func (m *SourceRegistrationUpdateRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSourceRegistrationRequestParams
	if err := m.CommonSourceRegistrationRequestParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmwareParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenericNasParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsilonParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetappParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElastifileParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlashbladeParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpfsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCassandraParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongodbParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCouchbaseParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdfsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHbaseParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHiveParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUdaParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffice365Params(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHypervParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSfdcParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateVmwareParams(formats strfmt.Registry) error {

	if swag.IsZero(m.VmwareParams) { // not required
		return nil
	}

	if m.VmwareParams != nil {
		if err := m.VmwareParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validatePhysicalParams(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalParams) { // not required
		return nil
	}

	if m.PhysicalParams != nil {
		if err := m.PhysicalParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("physicalParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateGenericNasParams(formats strfmt.Registry) error {

	if swag.IsZero(m.GenericNasParams) { // not required
		return nil
	}

	if m.GenericNasParams != nil {
		if err := m.GenericNasParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("genericNasParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("genericNasParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateIsilonParams(formats strfmt.Registry) error {

	if swag.IsZero(m.IsilonParams) { // not required
		return nil
	}

	if m.IsilonParams != nil {
		if err := m.IsilonParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isilonParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isilonParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateNetappParams(formats strfmt.Registry) error {

	if swag.IsZero(m.NetappParams) { // not required
		return nil
	}

	if m.NetappParams != nil {
		if err := m.NetappParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netappParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netappParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateElastifileParams(formats strfmt.Registry) error {

	if swag.IsZero(m.ElastifileParams) { // not required
		return nil
	}

	if m.ElastifileParams != nil {
		if err := m.ElastifileParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elastifileParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elastifileParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateFlashbladeParams(formats strfmt.Registry) error {

	if swag.IsZero(m.FlashbladeParams) { // not required
		return nil
	}

	if m.FlashbladeParams != nil {
		if err := m.FlashbladeParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flashbladeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flashbladeParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateGpfsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.GpfsParams) { // not required
		return nil
	}

	if m.GpfsParams != nil {
		if err := m.GpfsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpfsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gpfsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateCassandraParams(formats strfmt.Registry) error {

	if swag.IsZero(m.CassandraParams) { // not required
		return nil
	}

	if m.CassandraParams != nil {
		if err := m.CassandraParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateMongodbParams(formats strfmt.Registry) error {

	if swag.IsZero(m.MongodbParams) { // not required
		return nil
	}

	if m.MongodbParams != nil {
		if err := m.MongodbParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongodbParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongodbParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateCouchbaseParams(formats strfmt.Registry) error {

	if swag.IsZero(m.CouchbaseParams) { // not required
		return nil
	}

	if m.CouchbaseParams != nil {
		if err := m.CouchbaseParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("couchbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("couchbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateHdfsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HdfsParams) { // not required
		return nil
	}

	if m.HdfsParams != nil {
		if err := m.HdfsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hdfsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hdfsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateHbaseParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HbaseParams) { // not required
		return nil
	}

	if m.HbaseParams != nil {
		if err := m.HbaseParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateHiveParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HiveParams) { // not required
		return nil
	}

	if m.HiveParams != nil {
		if err := m.HiveParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hiveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hiveParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateUdaParams(formats strfmt.Registry) error {

	if swag.IsZero(m.UdaParams) { // not required
		return nil
	}

	if m.UdaParams != nil {
		if err := m.UdaParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("udaParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("udaParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateOffice365Params(formats strfmt.Registry) error {

	if swag.IsZero(m.Office365Params) { // not required
		return nil
	}

	if m.Office365Params != nil {
		if err := m.Office365Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("office365Params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("office365Params")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateAwsParams(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsParams) { // not required
		return nil
	}

	if m.AwsParams != nil {
		if err := m.AwsParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateHypervParams(formats strfmt.Registry) error {

	if swag.IsZero(m.HypervParams) { // not required
		return nil
	}

	if m.HypervParams != nil {
		if err := m.HypervParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hypervParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hypervParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateSfdcParams(formats strfmt.Registry) error {

	if swag.IsZero(m.SfdcParams) { // not required
		return nil
	}

	if m.SfdcParams != nil {
		if err := m.SfdcParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfdcParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sfdcParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) validateAzureParams(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureParams) { // not required
		return nil
	}

	if m.AzureParams != nil {
		if err := m.AzureParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this source registration update request params based on the context it is used
func (m *SourceRegistrationUpdateRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonSourceRegistrationRequestParams
	if err := m.CommonSourceRegistrationRequestParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVmwareParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhysicalParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGenericNasParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsilonParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetappParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElastifileParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlashbladeParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpfsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCassandraParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongodbParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCouchbaseParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHdfsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHbaseParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHiveParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUdaParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOffice365Params(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHypervParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSfdcParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateVmwareParams(ctx context.Context, formats strfmt.Registry) error {

	if m.VmwareParams != nil {

		if swag.IsZero(m.VmwareParams) { // not required
			return nil
		}

		if err := m.VmwareParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vmwareParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vmwareParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidatePhysicalParams(ctx context.Context, formats strfmt.Registry) error {

	if m.PhysicalParams != nil {

		if swag.IsZero(m.PhysicalParams) { // not required
			return nil
		}

		if err := m.PhysicalParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("physicalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("physicalParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateGenericNasParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GenericNasParams != nil {

		if swag.IsZero(m.GenericNasParams) { // not required
			return nil
		}

		if err := m.GenericNasParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("genericNasParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("genericNasParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateIsilonParams(ctx context.Context, formats strfmt.Registry) error {

	if m.IsilonParams != nil {

		if swag.IsZero(m.IsilonParams) { // not required
			return nil
		}

		if err := m.IsilonParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("isilonParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("isilonParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateNetappParams(ctx context.Context, formats strfmt.Registry) error {

	if m.NetappParams != nil {

		if swag.IsZero(m.NetappParams) { // not required
			return nil
		}

		if err := m.NetappParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netappParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netappParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateElastifileParams(ctx context.Context, formats strfmt.Registry) error {

	if m.ElastifileParams != nil {

		if swag.IsZero(m.ElastifileParams) { // not required
			return nil
		}

		if err := m.ElastifileParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elastifileParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elastifileParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateFlashbladeParams(ctx context.Context, formats strfmt.Registry) error {

	if m.FlashbladeParams != nil {

		if swag.IsZero(m.FlashbladeParams) { // not required
			return nil
		}

		if err := m.FlashbladeParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flashbladeParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flashbladeParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateGpfsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.GpfsParams != nil {

		if swag.IsZero(m.GpfsParams) { // not required
			return nil
		}

		if err := m.GpfsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gpfsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gpfsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateCassandraParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CassandraParams != nil {

		if swag.IsZero(m.CassandraParams) { // not required
			return nil
		}

		if err := m.CassandraParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cassandraParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cassandraParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateMongodbParams(ctx context.Context, formats strfmt.Registry) error {

	if m.MongodbParams != nil {

		if swag.IsZero(m.MongodbParams) { // not required
			return nil
		}

		if err := m.MongodbParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongodbParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongodbParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateCouchbaseParams(ctx context.Context, formats strfmt.Registry) error {

	if m.CouchbaseParams != nil {

		if swag.IsZero(m.CouchbaseParams) { // not required
			return nil
		}

		if err := m.CouchbaseParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("couchbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("couchbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateHdfsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HdfsParams != nil {

		if swag.IsZero(m.HdfsParams) { // not required
			return nil
		}

		if err := m.HdfsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hdfsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hdfsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateHbaseParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HbaseParams != nil {

		if swag.IsZero(m.HbaseParams) { // not required
			return nil
		}

		if err := m.HbaseParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hbaseParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hbaseParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateHiveParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HiveParams != nil {

		if swag.IsZero(m.HiveParams) { // not required
			return nil
		}

		if err := m.HiveParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hiveParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hiveParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateUdaParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UdaParams != nil {

		if swag.IsZero(m.UdaParams) { // not required
			return nil
		}

		if err := m.UdaParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("udaParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("udaParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateOffice365Params(ctx context.Context, formats strfmt.Registry) error {

	if m.Office365Params != nil {

		if swag.IsZero(m.Office365Params) { // not required
			return nil
		}

		if err := m.Office365Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("office365Params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("office365Params")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateAwsParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsParams != nil {

		if swag.IsZero(m.AwsParams) { // not required
			return nil
		}

		if err := m.AwsParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateHypervParams(ctx context.Context, formats strfmt.Registry) error {

	if m.HypervParams != nil {

		if swag.IsZero(m.HypervParams) { // not required
			return nil
		}

		if err := m.HypervParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hypervParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hypervParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateSfdcParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SfdcParams != nil {

		if swag.IsZero(m.SfdcParams) { // not required
			return nil
		}

		if err := m.SfdcParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sfdcParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sfdcParams")
			}
			return err
		}
	}

	return nil
}

func (m *SourceRegistrationUpdateRequestParams) contextValidateAzureParams(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureParams != nil {

		if swag.IsZero(m.AzureParams) { // not required
			return nil
		}

		if err := m.AzureParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceRegistrationUpdateRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceRegistrationUpdateRequestParams) UnmarshalBinary(b []byte) error {
	var res SourceRegistrationUpdateRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
