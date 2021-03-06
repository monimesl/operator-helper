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

package annotation

import (
	"strconv"
)

const (
	annPrometheusScrape = "prometheus.io/scrape"
	annPrometheusPort   = "prometheus.io/port"
)

// DecorateForPrometheus adds prometheus scraping annotations
func DecorateForPrometheus(ann map[string]string, scrap bool, port int) map[string]string {
	if ann == nil {
		ann = map[string]string{}
	}
	ann[annPrometheusScrape] = strconv.FormatBool(scrap)
	if scrap {
		ann[annPrometheusPort] = strconv.FormatInt(int64(port), 10)
	}
	return ann
}
