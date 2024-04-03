package controller

import (
	"context"

	multitenancyv1 "github.com/1eedaegon/tenant-oprator/api/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	tenantOperatorAnnotation = "tenant-operator"
)

func (r *TenantReconciler) ensureNamespace(ctx context.Context, tenant *multitenancyv1.Tenant, namespaceName string) error {
	log := log.FromContext(ctx)
	namespace := &corev1.Namespace{}

	err := r.Get(ctx, client.ObjectKey{Name: namespaceName}, namespace)

	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Creating namespace", "namespace", namespaceName)
			namespace := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: namespaceName,
					Annotations: map[string]string{
						"adminEmail": tenant.Spec.AdminEmail,
						"managed-by": tenantOperatorAnnotation,
					},
				},
			}
			if err = r.Create(ctx, namespace); err != nil {
				return err
			}
		}
	} else {
		log.Info("Namespace alreay exists", "namespace", namespaceName)
	}
	return nil
}
