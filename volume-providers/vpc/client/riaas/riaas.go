/*******************************************************************************
 * IBM Confidential
 * OCO Source Materials
 * IBM Cloud Container Service, 5737-D43
 * (C) Copyright IBM Corp. 2018 All Rights Reserved.
 * The source code for this program is not  published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with
 * the U.S. Copyright Office.
 ******************************************************************************/

package riaas

import (
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/client"
	"github.com/IBM/ibmcloud-storage-volume-lib/volume-providers/vpc/client/volume"
)

// RegionalAPI is the main interface for the RIAAS API client. From here, service
// objects for the individual parts of the API can be obtained
//go:generate counterfeiter -o fakes/regional_api.go --fake-name RegionalAPI . RegionalAPI
type RegionalAPI interface {
	Login(token string) error

	Volume() volume.VolumeManager
}

var _ RegionalAPI = &Session{}

// Session is a base implementation of the RegionalAPI interface
type Session struct {
	client client.Client
	config Config
}

// New creates a new Session volume, using the supplied config
func New(config Config) (*Session, error) {
	riaasClient := client.New(config.baseURL(), config.httpClient())

	if config.DebugWriter != nil {
		riaasClient.WithDebug(config.DebugWriter)
	}
	return &Session{
		client: riaasClient,
		config: config,
	}, nil
}

// Login configures the session with the supplied Authentication token
// which is used for all requests to the API
func (s *Session) Login(token string) error {

	s.client.WithAuthToken(token)
	return nil
}

// Volume returns the Volume service for managing volumes
func (s *Session) Volume() volume.VolumeManager {
	return volume.New(s.client)
}

// RegionalAPIClientProvider declares an interface for a provider that can supply a new
// RegionalAPI client session
//go:generate counterfeiter -o fakes/client_provider.go --fake-name RegionalAPIClientProvider . RegionalAPIClientProvider
type RegionalAPIClientProvider interface {
	New(config Config) (RegionalAPI, error)
}

// DefaultRegionalAPIClientProvider declares a basic client provider that delegates to
// New(). Can be used for dependency injection.
type DefaultRegionalAPIClientProvider struct {
}

var _ RegionalAPIClientProvider = DefaultRegionalAPIClientProvider{}

// New creates a new Session volume, using the supplied config
func (d DefaultRegionalAPIClientProvider) New(config Config) (RegionalAPI, error) {
	return New(config)
}
