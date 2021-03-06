/*
 * Copyright 2021 - now, the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pvc

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// New creates a new pvc
func New(namespace, name string, labels map[string]string, spec v1.PersistentVolumeClaimSpec) v1.PersistentVolumeClaim {
	return v1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
			Labels:    labels,
		},
		Spec:   spec,
		Status: v1.PersistentVolumeClaimStatus{},
	}
}

// ListAllWithMatchingLabels list the pvcs matching the labels
func ListAllWithMatchingLabels(cl client.Client, namespace string, labels map[string]string) (*v1.PersistentVolumeClaimList, error) {
	selector, err := metav1.LabelSelectorAsSelector(&metav1.LabelSelector{
		MatchLabels: labels,
	})
	if err != nil {
		return nil, fmt.Errorf("error on creating selector from label selector: %w", err)
	}
	list := &v1.PersistentVolumeClaimList{}
	listOpts := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: selector,
	}
	err = cl.List(context.TODO(), list, listOpts)
	if err != nil {
		return nil, err
	}
	return list, nil
}
