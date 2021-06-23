/*
Copyright SUSE LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"

	hypperChart "github.com/rancher-sandbox/hypper/pkg/chart"
	helmChart "helm.sh/helm/v3/pkg/chart"
)

type tristate int

const (
	Unknown tristate = iota
	Present
	Absent
)

// Pkg is the minimum object the solver reasons about. It is comprised of a
// chart, its version, and its characteristics when installed (release name,
// namespace, etc).
// Note that each package is unique. The same chart, with the same default
// release name & ns, but different version, is a different package. E.g:
// prometheus-1.2.0 and prometheus-1.3.0 are different packages.
type Pkg struct {
	Name               string // Release name, or default chart release-name
	Version            string // sem ver (without a range)
	ChartHash          uint64 // hash of the chart contents
	Namespace          string // Installed ns, or default chart namespace
	DependsRel         []string // list of dependencies' fingerprints
	DependsOptionalRel []string // list of optional dependencies' fingerprints
	CurrentState       tristate // current state of the package
	DesiredState       tristate // desired state of the package
	Chart              *helmChart.Chart
}

func NewPkg(name, version, namespace string,
	currentState, desiredState tristate, chart *helmChart.Chart) *Pkg {

	p :=  &Pkg{
		Name:               name,
		Version:            version,
		ChartHash:          hypperChart.Hash(chart),
		Namespace:          namespace,
		DependsRel:         []string{},
		DependsOptionalRel: []string{},
		CurrentState:       currentState,
		DesiredState:       desiredState,
		Chart:              chart,
	}

	// codify dependency relations:


	return p

}

// NewPkgMock creates a new package, with a digest based in the package name,
// and a nil chart pointer.
// Useful for testing.
func NewPkgMock(name, version, namespace string,
	depends, dependsOptional []string,
	currentState, desiredState tristate) *Pkg {

	p := NewPkg(name, version, namespace, currentState, desiredState, nil)

	p.DependsRel = depends
	p.DependsOptionalRel = dependsOptional
	p.ChartHash = 0

	return p
}

// JSON serializes package p into JSON, returning a []byte
func (p *Pkg) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(p)
	return buffer.Bytes(), err
}

// GetFingerPrint returns a unique id of the package.
// Digest is present to help with packages (mariadb, postgres) that satisfy a
// metapackage, and get installed with the metapackage releaseName (e.g: rdbms)
func (p *Pkg) GetFingerPrint() string {
	return fmt.Sprintf("%s-%s-%d-%s", p.Name, p.Version, p.ChartHash, p.Namespace)
}

// GetBaseFingerPrint returns a unique id of the package minus version.
func (p *Pkg) GetBaseFingerPrint() string {
	return fmt.Sprintf("%s-%s", p.Name, p.Namespace)
}

// GetFingerPrintMock returns a fingerprint of a mock package, setting the
// p.ChartHash to 0.
func GetFingerPrintMock(name string, version string, ns string) string {
	// chart.Hash is always 0, as it is a mock
	return fmt.Sprintf("%s-%s-0-%s", name, version, ns)
}

// Encode encodes the package to string.
// It returns an ID which can be used to retrieve the package later on.
func (p *Pkg) Encode() (string, error) {

	encodedPackage, err := p.JSON()
	if err != nil {
		return "", err
	}

	return string(encodedPackage), nil
}
