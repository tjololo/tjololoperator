/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appv1alpha1 "tjololoperator/api/v1alpha1"
)

// PlaygroundReconciler reconciles a Playground object
type PlaygroundReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=app.tjololo.net,resources=playgrounds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app.tjololo.net,resources=playgrounds/status,verbs=get;update;patch

func (r *PlaygroundReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("playground", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PlaygroundReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1alpha1.Playground{}).
		Complete(r)
}
