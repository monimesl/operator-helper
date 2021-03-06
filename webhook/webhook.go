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

package webhook

import (
	"github.com/monimesl/operator-helper/config"
	"github.com/monimesl/operator-helper/reconciler"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
)

func Context() reconciler.Context {
	return reconciler.GetContext()
}

// Configure configures the webhook for the added CR types
func Configure(manager ctrl.Manager, apiTypes ...runtime.Object) error {
	if config.WebHooksEnabled() {
		for _, apiType := range apiTypes {
			log.Printf("configuring the webhook: %T\n", apiType)
			if err := ctrl.NewWebhookManagedBy(manager).For(apiType).Complete(); err != nil {
				return err
			}
		}
	} else {
		log.Printf("Cannot configure webhooks as it's disabled")
	}
	return nil
}
