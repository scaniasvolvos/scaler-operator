/*
Copyright 2026 Pranav Deshpande.

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

package controller

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apiv1alpha1 "github.com/scaniasvolvos/scaler-operator/api/v1alpha1"
)

var _ = Describe("Scaler Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}

		deploymentNamespacedName := types.NamespacedName{
			Name:      "test-deployment",
			Namespace: "default",
		}

		scaler := &apiv1alpha1.Scaler{}

		BeforeEach(func() {
			By("creating the test Deployment")
			dep := &appsv1.Deployment{}
			err := k8sClient.Get(ctx, deploymentNamespacedName, dep)
			if err != nil && errors.IsNotFound(err) {
				dep = &appsv1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test-deployment",
						Namespace: "default",
					},
					Spec: appsv1.DeploymentSpec{
						Replicas: ptr.To(int32(1)),
						Selector: &metav1.LabelSelector{
							MatchLabels: map[string]string{"app": "test"},
						},
						Template: corev1.PodTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Labels: map[string]string{"app": "test"},
							},
							Spec: corev1.PodSpec{
								Containers: []corev1.Container{
									{
										Name:  "test",
										Image: "busybox",
									},
								},
							},
						},
					},
				}
				Expect(k8sClient.Create(ctx, dep)).To(Succeed())
			}

			By("creating the custom resource for the Kind Scaler")
			err = k8sClient.Get(ctx, typeNamespacedName, scaler)
			if err != nil && errors.IsNotFound(err) {
				resource := &apiv1alpha1.Scaler{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					Spec: apiv1alpha1.ScalerSpec{
						Start:    0,
						End:      23,
						Replicas: 1,
						Deployments: []apiv1alpha1.NamespacedName{
							{
								Name:      "test-deployment",
								Namespace: "default",
							},
						},
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func() {
			resource := &apiv1alpha1.Scaler{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance Scaler")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())

			dep := &appsv1.Deployment{}
			err = k8sClient.Get(ctx, deploymentNamespacedName, dep)
			if err == nil {
				By("Cleanup the test Deployment")
				Expect(k8sClient.Delete(ctx, dep)).To(Succeed())
			}
		})
		It("should successfully reconcile the resource", func() {
			By("Reconciling the created resource")
			controllerReconciler := &ScalerReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
