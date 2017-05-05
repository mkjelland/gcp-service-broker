// Copyright the Service Broker Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////
//

package test_utils

import (
	"gcp-service-broker/brokerapi/brokers/models"
	"gcp-service-broker/db_service"
	"github.com/jinzhu/gorm"
)

func CreateTestDb() {
	testDb, _ := gorm.Open("sqlite3", "test.db")
	testDb.CreateTable(models.ServiceInstanceDetails{})
	testDb.CreateTable(models.ServiceBindingCredentials{})
	testDb.CreateTable(models.PlanDetails{})
	testDb.CreateTable(models.ProvisionRequestDetails{})
	testDb.CreateTable(models.CloudOperation{})
	testDb.CreateTable(models.ServiceDefaults{})

	db_service.DbConnection = testDb
}
