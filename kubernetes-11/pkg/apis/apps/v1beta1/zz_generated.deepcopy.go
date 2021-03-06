// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	api_v1 "github.com/sourcegraph/monorepo-test-1/kubernetes-11/pkg/api/v1"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Deployment, InType: reflect.TypeOf(&Deployment{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentCondition, InType: reflect.TypeOf(&DeploymentCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentList, InType: reflect.TypeOf(&DeploymentList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentRollback, InType: reflect.TypeOf(&DeploymentRollback{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentSpec, InType: reflect.TypeOf(&DeploymentSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentStatus, InType: reflect.TypeOf(&DeploymentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_DeploymentStrategy, InType: reflect.TypeOf(&DeploymentStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollbackConfig, InType: reflect.TypeOf(&RollbackConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_RollingUpdateDeployment, InType: reflect.TypeOf(&RollingUpdateDeployment{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_Scale, InType: reflect.TypeOf(&Scale{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ScaleSpec, InType: reflect.TypeOf(&ScaleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_ScaleStatus, InType: reflect.TypeOf(&ScaleStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSet, InType: reflect.TypeOf(&StatefulSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSetList, InType: reflect.TypeOf(&StatefulSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSetSpec, InType: reflect.TypeOf(&StatefulSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1beta1_StatefulSetStatus, InType: reflect.TypeOf(&StatefulSetStatus{})},
	)
}

func DeepCopy_v1beta1_Deployment(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Deployment)
		out := out.(*Deployment)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_DeploymentSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentCondition)
		out := out.(*DeploymentCondition)
		*out = *in
		out.LastUpdateTime = in.LastUpdateTime.DeepCopy()
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentList)
		out := out.(*DeploymentList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Deployment, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_Deployment(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentRollback(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentRollback)
		out := out.(*DeploymentRollback)
		*out = *in
		if in.UpdatedAnnotations != nil {
			in, out := &in.UpdatedAnnotations, &out.UpdatedAnnotations
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentSpec)
		out := out.(*DeploymentSpec)
		*out = *in
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_DeploymentStrategy(&in.Strategy, &out.Strategy, c); err != nil {
			return err
		}
		if in.RevisionHistoryLimit != nil {
			in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
			*out = new(int32)
			**out = **in
		}
		if in.RollbackTo != nil {
			in, out := &in.RollbackTo, &out.RollbackTo
			*out = new(RollbackConfig)
			**out = **in
		}
		if in.ProgressDeadlineSeconds != nil {
			in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStatus)
		out := out.(*DeploymentStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]DeploymentCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_DeploymentCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_DeploymentStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DeploymentStrategy)
		out := out.(*DeploymentStrategy)
		*out = *in
		if in.RollingUpdate != nil {
			in, out := &in.RollingUpdate, &out.RollingUpdate
			*out = new(RollingUpdateDeployment)
			if err := DeepCopy_v1beta1_RollingUpdateDeployment(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_RollbackConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollbackConfig)
		out := out.(*RollbackConfig)
		*out = *in
		return nil
	}
}

func DeepCopy_v1beta1_RollingUpdateDeployment(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RollingUpdateDeployment)
		out := out.(*RollingUpdateDeployment)
		*out = *in
		if in.MaxUnavailable != nil {
			in, out := &in.MaxUnavailable, &out.MaxUnavailable
			*out = new(intstr.IntOrString)
			**out = **in
		}
		if in.MaxSurge != nil {
			in, out := &in.MaxSurge, &out.MaxSurge
			*out = new(intstr.IntOrString)
			**out = **in
		}
		return nil
	}
}

func DeepCopy_v1beta1_Scale(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Scale)
		out := out.(*Scale)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_ScaleStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_ScaleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleSpec)
		out := out.(*ScaleSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_v1beta1_ScaleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ScaleStatus)
		out := out.(*ScaleStatus)
		*out = *in
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSet)
		out := out.(*StatefulSet)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1beta1_StatefulSetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1beta1_StatefulSetStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetList)
		out := out.(*StatefulSetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]StatefulSet, len(*in))
			for i := range *in {
				if err := DeepCopy_v1beta1_StatefulSet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetSpec)
		out := out.(*StatefulSetSpec)
		*out = *in
		if in.Replicas != nil {
			in, out := &in.Replicas, &out.Replicas
			*out = new(int32)
			**out = **in
		}
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			if newVal, err := c.DeepCopy(*in); err != nil {
				return err
			} else {
				*out = newVal.(*v1.LabelSelector)
			}
		}
		if err := api_v1.DeepCopy_v1_PodTemplateSpec(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.VolumeClaimTemplates != nil {
			in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
			*out = make([]api_v1.PersistentVolumeClaim, len(*in))
			for i := range *in {
				if err := api_v1.DeepCopy_v1_PersistentVolumeClaim(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1beta1_StatefulSetStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StatefulSetStatus)
		out := out.(*StatefulSetStatus)
		*out = *in
		if in.ObservedGeneration != nil {
			in, out := &in.ObservedGeneration, &out.ObservedGeneration
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}
