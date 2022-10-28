package v1alpha1

import (
	"encoding/json"
)

// This file contains the UOR Collection core attributes for UOR descriptors.

// Core schema define the required attributes for a collection manifests.
type Core struct {
	// RegistryHint will allow registries to be
	// add to a user search domain in a discovered zone
	RegistryHint string      `json:"registryHint"`
	Components   []Component `json:"components"`
}

// Component schema defines information to create a component list.
// Based on Anchore Syft Package spec
type Component struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	Type      string   `json:"type"`
	FoundBy   string   `json:"foundBy"`
	Locations []string `json:"locations"`
	Licenses  []string `json:"licenses"`
	Language  string   `json:"language"`
	// Common Platform Enumeration
	CPEs []string `json:"cpes"`
	// Package URL
	PURL               string          `json:"purl"`
	AdditionalMetadata json.RawMessage `json:",inline"`
}

// Platform describes the platform which the artifact can be used on.
type Platform struct {
	// Architecture field specifies the CPU architecture, for example
	// `amd64` or `ppc64`.
	Architecture string `json:"architecture"`

	// OS specifies the operating system, for example `linux` or `windows`.
	OS string `json:"os"`

	// OSVersion is an optional field specifying the operating system
	// version, for example on Windows `10.0.14393.1066`.
	OSVersion string `json:"os.version,omitempty"`

	// OSFeatures is an optional field specifying an array of strings,
	// each listing a required OS feature (for example on Windows `win32k`).
	OSFeatures []string `json:"os.features,omitempty"`

	// Variant is an optional field specifying a variant of the CPU, for
	// example `v7` to specify ARMv7 when architecture is `arm`.
	Variant string `json:"variant,omitempty"`
}

// SchemaAttributes are the required attributes for a schema descriptor.
type SchemaAttributes struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}