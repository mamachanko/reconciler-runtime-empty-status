/*
Copyright 2022.
*/

package controllers

import (
	"context"
	"github.com/mamachanko/reconciler-runtime-empty-status/api/v1alpha1"
	"github.com/vmware-labs/reconciler-runtime/reconcilers"
)

func ThingReconciler(c reconcilers.Config) *reconcilers.ParentReconciler {
	c.Log = c.Log.WithName("Thing")

	return &reconcilers.ParentReconciler{
		Type: &v1alpha1.Thing{},
		Reconciler: func(c reconcilers.Config) *reconcilers.SyncReconciler {
			c.Log = c.Log.WithName("ThingSync")

			return &reconcilers.SyncReconciler{
				Sync: func(ctx context.Context, parent *v1alpha1.Thing) error {
					c.Log.Info("syncing...")
					return nil
				},
				Config: c,
			}
		}(c),
		Config: c,
	}
}
