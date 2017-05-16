// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package integration_new_tests

import (
	"fmt"
	"math/rand"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

const (
	postDeployPlanName            = "post-deploy-plan-name"
	postDeployErrandName          = "post-deploy-errand-name"
	withPostDeployErrandPlanID    = "post-deploy-errand-id"
	withoutPostDeployErrandPlanID = defaultPlanID
	secondDefaultPlanID           = "second-default-plan-id"
)

var _ = Describe("updating a service instance", func() {
	It("returns tracking data for an update operation", func() {
		updateTaskID := rand.Int()
		boshDeploysUpdatedManifest := func(env *BrokerEnvironment) {
			deploymentName := env.DeploymentName()
			env.Bosh.HasNoTasksFor(deploymentName)
			env.Bosh.HasManifestFor(deploymentName)

			env.Bosh.DeploysWithoutContextID(deploymentName, updateTaskID)
		}

		When(updatingPlan(defaultPlanID, secondDefaultPlanID)).
			With(secondDefaultPlanConfigured, NoCredhub, serviceAdapterGeneratesManifest, boshDeploysUpdatedManifest).
			theBroker(
				RespondsWith(http.StatusAccepted, OperationData(withoutErrand(broker.OperationTypeUpdate, updateTaskID))),
				LogsWithServiceId("updating instance %s"),
				LogsWithDeploymentName(fmt.Sprintf("Deployed. Operation: update, BoshTaskID: %d, DeploymentName: %%s, PlanID: %s", updateTaskID, secondDefaultPlanID)),
			)
	})

	It("runs the post-deployment errand if new plan has one", func() {
		updateTaskID := rand.Int()
		boshDeploysUpdatedManifest := func(env *BrokerEnvironment) {
			deploymentName := env.DeploymentName()
			env.Bosh.HasNoTasksFor(deploymentName)
			env.Bosh.HasManifestFor(deploymentName)

			env.Bosh.DeploysWithAContextID(deploymentName, updateTaskID)
		}

		When(updatingPlan(withoutPostDeployErrandPlanID, withPostDeployErrandPlanID)).
			With(postDeployErrandConfigured, NoCredhub, serviceAdapterGeneratesManifest, boshDeploysUpdatedManifest).
			theBroker(
				RespondsWith(http.StatusAccepted, OperationData(withErrand(broker.OperationTypeUpdate, updateTaskID, postDeployErrandName))),
				LogsWithServiceId("updating instance %s"),
				LogsWithDeploymentName(fmt.Sprintf("Deployed. Operation: update, BoshTaskID: %d, DeploymentName: %%s, PlanID: %s", updateTaskID, withPostDeployErrandPlanID)),
			)
	})

	It("does not run a post-deployment errand if new plan does not have one", func() {
		updateTaskID := rand.Int()
		boshDeploysUpdatedManifest := func(env *BrokerEnvironment) {
			deploymentName := env.DeploymentName()
			env.Bosh.HasNoTasksFor(deploymentName)
			env.Bosh.HasManifestFor(deploymentName)

			env.Bosh.DeploysWithoutContextID(deploymentName, updateTaskID)
		}

		When(updatingPlan(withPostDeployErrandPlanID, withoutPostDeployErrandPlanID)).
			With(postDeployErrandConfigured, NoCredhub, serviceAdapterGeneratesManifest, boshDeploysUpdatedManifest).
			theBroker(
				RespondsWith(http.StatusAccepted, OperationData(withoutErrand(broker.OperationTypeUpdate, updateTaskID))),
				LogsWithServiceId("updating instance %s"),
				LogsWithDeploymentName(fmt.Sprintf("Deployed. Operation: update, BoshTaskID: %d, DeploymentName: %%s, PlanID: %s", updateTaskID, withoutPostDeployErrandPlanID)),
			)
	})

})

// TODO Should we verify the parameters to GenerateManifest?

func updatingPlan(oldPlanID string, newPlanID string) func(env *BrokerEnvironment) *http.Request {
	return func(env *BrokerEnvironment) *http.Request {
		return env.Broker.UpdateServiceInstanceRequest(env.serviceInstanceID, oldPlanID, newPlanID)
	}
}

var serviceAdapterGeneratesManifest = func(sa *ServiceAdapter, id ServiceInstanceID) {
	sa.adapter.GenerateManifest().ToReturnManifest(rawManifestWithDeploymentName(id))
}

func rawManifestWithDeploymentName(id ServiceInstanceID) string {
	return "name: " + broker.DeploymentNameFrom(string(id))
}

func secondDefaultPlanConfigured(source *config.Config) *config.Config {
	source.ServiceCatalog.Plans = append(source.ServiceCatalog.Plans, config.Plan{
		ID: secondDefaultPlanID,
		InstanceGroups: []serviceadapter.InstanceGroup{
			{
				VMType:    "the-vm-type",
				Name:      "the-instance-group",
				Instances: 1,
				AZs:       []string{"the-az"},
			},
		},
	})

	return source
}

func postDeployErrandConfigured(source *config.Config) *config.Config {
	source.ServiceCatalog.Plans = append(source.ServiceCatalog.Plans, config.Plan{
		Name: postDeployPlanName,
		ID:   withPostDeployErrandPlanID,
		InstanceGroups: []serviceadapter.InstanceGroup{
			{
				Name:      "instance-group-name",
				VMType:    "post-deploy-errand-vm-type",
				Instances: 1,
				Networks:  []string{"net1"},
				AZs:       []string{"az1"},
			},
		},
		LifecycleErrands: &config.LifecycleErrands{
			PostDeploy: postDeployErrandName,
		},
	})

	return source
}

func withoutErrand(opType broker.OperationType, taskId int) types.GomegaMatcher {
	return MatchAllFields(
		Fields{
			"OperationType":        Equal(opType),
			"BoshTaskID":           Equal(taskId),
			"PlanID":               BeEmpty(),
			"PostDeployErrandName": BeEmpty(),
			"BoshContextID":        BeEmpty(),
		})
}

func withErrand(opType broker.OperationType, taskId int, errandName string) types.GomegaMatcher {
	return MatchAllFields(
		Fields{
			"OperationType":        Equal(opType),
			"BoshTaskID":           Equal(taskId),
			"PlanID":               BeEmpty(),
			"PostDeployErrandName": Equal(errandName),
			"BoshContextID":        Not(BeEmpty()),
		})
}
