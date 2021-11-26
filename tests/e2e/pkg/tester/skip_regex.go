/*
Copyright 2021 The Kubernetes Authors.

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

package tester

import (
	"regexp"
	"strings"

	"k8s.io/kops/upup/pkg/fi/utils"
)

const (
	skipRegexBase = "\\[Slow\\]|\\[Serial\\]|\\[Disruptive\\]|\\[Flaky\\]|\\[Feature:.+\\]|\\[HPA\\]|\\[Driver:.nfs\\]|Dashboard|Gluster|RuntimeClass|RuntimeHandler"
)

func (t *Tester) setSkipRegexFlag() error {
	if t.SkipRegex != "" {
		return nil
	}

	cluster, err := t.getKopsCluster()
	if err != nil {
		return err
	}

	skipRegex := skipRegexBase

	networking := cluster.Spec.Networking
	switch {
	case networking.Kubenet != nil, networking.Canal != nil, networking.Weave != nil, networking.Cilium != nil:
		skipRegex += "|Services.*rejected.*endpoints"
	}
	if networking.Cilium != nil {
		// https://github.com/cilium/cilium/issues/10002
		skipRegex += "|TCP.CLOSE_WAIT"
		// https://github.com/cilium/cilium/issues/15361
		skipRegex += "|external.IP.is.not.assigned.to.a.node"
		// https://github.com/cilium/cilium/issues/14287
		skipRegex += "|same.port.number.but.different.protocols|same.hostPort.but.different.hostIP.and.protocol"
		if strings.Contains(cluster.Spec.KubernetesVersion, "v1.23.0") || strings.Contains(cluster.Spec.KubernetesVersion, "v1.24.0") {
			// Reassess after https://github.com/kubernetes/kubernetes/pull/102643 is merged
			// ref:
			// https://github.com/kubernetes/kubernetes/issues/96717
			// https://github.com/cilium/cilium/issues/5719
			skipRegex += "|should.create.a.Pod.with.SCTP.HostPort"
		}
	} else if networking.Calico != nil {
		skipRegex += "|Services.*functioning.*NodePort"
	} else if networking.Kuberouter != nil {
		skipRegex += "|load-balancer|hairpin|affinity\\stimeout|service\\.kubernetes\\.io|CLOSE_WAIT"
	} else if networking.Kubenet != nil {
		skipRegex += "|Services.*affinity"
	}

	if cluster.Spec.CloudProvider == "gce" {
		// Firewall tests expect a specific format for cluster and control plane host names
		// which kOps does not match
		// ref: https://github.com/kubernetes/kubernetes/blob/1bd00776b5d78828a065b5c21e7003accc308a06/test/e2e/framework/providers/gce/firewall.go#L92-L100
		skipRegex += "|Firewall"
		// kube-dns tests are not skipped automatically if a cluster uses CoreDNS instead
		skipRegex += "|kube-dns"
		// this test assumes the cluster runs COS but kOps uses Ubuntu by default
		// ref: https://github.com/kubernetes/test-infra/pull/22190
		skipRegex += "|should.be.mountable.when.non-attachable"
	}

	if cluster.Spec.CloudProvider == "aws" && utils.IsIPv6CIDR(cluster.Spec.NonMasqueradeCIDR) {
		// AWS VPC Classic ELBs are IPv4 only
		// ref: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-internet-facing-load-balancers.html#internet-facing-ip-addresses
		skipRegex += "|should.not.disrupt.a.cloud.load-balancer.s.connectivity.during.rollout"
	}

	// Ensure it is valid regex
	if _, err := regexp.Compile(skipRegex); err != nil {
		return err
	}
	t.SkipRegex = skipRegex
	return nil
}
