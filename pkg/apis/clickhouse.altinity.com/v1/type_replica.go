// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

func (replica *ChiReplica) InheritTemplates(shard *ChiShard) {
	(&replica.Templates).MergeFrom(&shard.Templates)
}

func (replica *ChiReplica) GetPodTemplate() (*ChiPodTemplate, bool) {
	name := replica.Templates.PodTemplate
	template, ok := replica.Chi.GetPodTemplate(name)
	return template, ok
}

func (replica *ChiReplica) GetVolumeClaimTemplate() (*ChiVolumeClaimTemplate, bool) {
	name := replica.Templates.VolumeClaimTemplate
	template, ok := replica.Chi.GetVolumeClaimTemplate(name)
	return template, ok
}

func (replica *ChiReplica) GetServiceTemplate() (*ChiServiceTemplate, bool) {
	name := replica.Templates.ReplicaServiceTemplate
	template, ok := replica.Chi.GetServiceTemplate(name)
	return template, ok
}
