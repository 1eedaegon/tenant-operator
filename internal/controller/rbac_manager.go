package controller

import (
	"context"
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *TenantReconciler) EnsureRoleBinding(ctx context.Context, namespaceName string, groups []string, clusterRoleName string) error {
	log := log.FromContext(ctx)

	roleBindingName := fmt.Sprintf("%s-%s-rb", namespaceName, clusterRoleName)

	clusterRole := &rbacv1.ClusterRole{}

	err := r.Get(ctx, client.ObjectKey{Name: clusterRoleName}, clusterRole)

	if err != nil {
		log.Error(err, "Failed to get ClusterRole", "clusterRole", clusterRoleName)
	}

	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleBindingName,
			Namespace: namespaceName,
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			Name:     clusterRoleName,
			APIGroup: rbacv1.GroupName,
		},
		Subjects: make([]rbacv1.Subject, len(groups)),
	}

	for i, group := range groups {
		roleBinding.Subjects[i] = rbacv1.Subject{
			Kind:     "Group",
			Name:     group,
			APIGroup: rbacv1.GroupName,
		}
	}

	return nil
}
