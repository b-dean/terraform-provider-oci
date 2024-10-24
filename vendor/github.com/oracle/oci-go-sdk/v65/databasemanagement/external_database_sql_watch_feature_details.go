// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDatabaseSqlWatchFeatureDetails The details required to enable the SQL Watch feature.
type ExternalDatabaseSqlWatchFeatureDetails struct {
	ConnectorDetails ConnectorDetails `mandatory:"true" json:"connectorDetails"`
}

// GetConnectorDetails returns ConnectorDetails
func (m ExternalDatabaseSqlWatchFeatureDetails) GetConnectorDetails() ConnectorDetails {
	return m.ConnectorDetails
}

func (m ExternalDatabaseSqlWatchFeatureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDatabaseSqlWatchFeatureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalDatabaseSqlWatchFeatureDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalDatabaseSqlWatchFeatureDetails ExternalDatabaseSqlWatchFeatureDetails
	s := struct {
		DiscriminatorParam string `json:"feature"`
		MarshalTypeExternalDatabaseSqlWatchFeatureDetails
	}{
		"SQLWATCH",
		(MarshalTypeExternalDatabaseSqlWatchFeatureDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDatabaseSqlWatchFeatureDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ConnectorDetails connectordetails `json:"connectorDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.ConnectorDetails.UnmarshalPolymorphicJSON(model.ConnectorDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectorDetails = nn.(ConnectorDetails)
	} else {
		m.ConnectorDetails = nil
	}

	return
}