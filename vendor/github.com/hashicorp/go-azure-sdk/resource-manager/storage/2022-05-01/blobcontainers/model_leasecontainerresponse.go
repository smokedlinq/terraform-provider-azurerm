package blobcontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LeaseContainerResponse struct {
	LeaseId          *string `json:"leaseId,omitempty"`
	LeaseTimeSeconds *string `json:"leaseTimeSeconds,omitempty"`
}
