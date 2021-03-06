// +build integration

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tests

import (
	"os"
	"testing"

	"github.com/apache/incubator-openwhisk-wskdeploy/tests/src/integration/common"
	"github.com/stretchr/testify/assert"
)

func TestManagedDeployment(t *testing.T) {
	path := "/src/github.com/apache/incubator-openwhisk-wskdeploy/tests/src/integration/managed-deployment/"
	manifestPath := os.Getenv("GOPATH") + path + "manifest.yaml"
	deploymentPath := ""
	wskdeploy := common.NewWskdeploy()
	_, err := wskdeploy.ManagedDeployment(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")

	manifestPath = os.Getenv("GOPATH") + path + "00-manifest-minus-second-package.yaml"
	wskdeploy = common.NewWskdeploy()
	_, err = wskdeploy.ManagedDeployment(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")

	manifestPath = os.Getenv("GOPATH") + path + "01-manifest-minus-sequence-2.yaml"
	wskdeploy = common.NewWskdeploy()
	_, err = wskdeploy.ManagedDeployment(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")

	manifestPath = os.Getenv("GOPATH") + path + "02-manifest-minus-action-3.yaml"
	wskdeploy = common.NewWskdeploy()
	_, err = wskdeploy.ManagedDeployment(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")

	manifestPath = os.Getenv("GOPATH") + path + "03-manifest-minus-trigger.yaml"
	wskdeploy = common.NewWskdeploy()
	_, err = wskdeploy.ManagedDeployment(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")

	manifestPath = os.Getenv("GOPATH") + path + "04-manifest-minus-package.yaml"
	wskdeploy = common.NewWskdeploy()
	_, err = wskdeploy.ManagedDeployment(manifestPath, deploymentPath)
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")

	manifestPath = os.Getenv("GOPATH") + path + "05-manifest-headless.yaml"
	wskdeploy = common.NewWskdeploy()
	_, err = wskdeploy.HeadlessManagedDeployment(manifestPath, deploymentPath, "Headless Managed")
	assert.Equal(t, nil, err, "Failed to deploy based on the manifest and deployment files.")
}
