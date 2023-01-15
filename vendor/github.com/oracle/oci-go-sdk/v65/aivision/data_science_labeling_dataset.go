// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vision API
//
// Using Vision, you can upload images to detect and classify objects in them. If you have lots of images, you can process them in batch using asynchronous API endpoints. Vision's features are thematically split between Document AI for document-centric images, and Image Analysis for object and scene-based images. Pretrained models and custom models are supported.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataScienceLabelingDataset The dataset created by the Data Labeling Service.
type DataScienceLabelingDataset struct {

	// OCID of the Data Labeling dataset.
	DatasetId *string `mandatory:"false" json:"datasetId"`
}

func (m DataScienceLabelingDataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataScienceLabelingDataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataScienceLabelingDataset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataScienceLabelingDataset DataScienceLabelingDataset
	s := struct {
		DiscriminatorParam string `json:"datasetType"`
		MarshalTypeDataScienceLabelingDataset
	}{
		"DATA_SCIENCE_LABELING",
		(MarshalTypeDataScienceLabelingDataset)(m),
	}

	return json.Marshal(&s)
}
