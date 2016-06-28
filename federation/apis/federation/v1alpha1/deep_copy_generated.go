// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	api "k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1alpha1_Cluster,
		DeepCopy_v1alpha1_ClusterCondition,
		DeepCopy_v1alpha1_ClusterList,
		DeepCopy_v1alpha1_ClusterSpec,
		DeepCopy_v1alpha1_ClusterStatus,
		DeepCopy_v1alpha1_ServerAddressByClientCIDR,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1alpha1_Cluster(in Cluster, out *Cluster, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1alpha1_ClusterSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1alpha1_ClusterStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1alpha1_ClusterCondition(in ClusterCondition, out *ClusterCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	if err := unversioned.DeepCopy_unversioned_Time(in.LastProbeTime, &out.LastProbeTime, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Time(in.LastTransitionTime, &out.LastTransitionTime, c); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func DeepCopy_v1alpha1_ClusterList(in ClusterList, out *ClusterList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Cluster, len(in))
		for i := range in {
			if err := DeepCopy_v1alpha1_Cluster(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1alpha1_ClusterSpec(in ClusterSpec, out *ClusterSpec, c *conversion.Cloner) error {
	if in.ServerAddressByClientCIDRs != nil {
		in, out := in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
		*out = make([]ServerAddressByClientCIDR, len(in))
		for i := range in {
			if err := DeepCopy_v1alpha1_ServerAddressByClientCIDR(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.ServerAddressByClientCIDRs = nil
	}
	if in.SecretRef != nil {
		in, out := in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		if err := v1.DeepCopy_v1_LocalObjectReference(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.SecretRef = nil
	}
	return nil
}

func DeepCopy_v1alpha1_ClusterStatus(in ClusterStatus, out *ClusterStatus, c *conversion.Cloner) error {
	if in.Conditions != nil {
		in, out := in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(in))
		for i := range in {
			if err := DeepCopy_v1alpha1_ClusterCondition(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.Zones != nil {
		in, out := in.Zones, &out.Zones
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Zones = nil
	}
	out.Region = in.Region
	return nil
}

func DeepCopy_v1alpha1_ServerAddressByClientCIDR(in ServerAddressByClientCIDR, out *ServerAddressByClientCIDR, c *conversion.Cloner) error {
	out.ClientCIDR = in.ClientCIDR
	out.ServerAddress = in.ServerAddress
	return nil
}
