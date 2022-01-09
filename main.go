package main

import (
	sap_api_caller "sap-api-integrations-maintenance-task-list-reads/SAP_API_Caller"
	"sap-api-integrations-maintenance-task-list-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Maintenance_Task_List_Operation_Text_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "HeaderEquipmentPlant", "StrategyPackage", "StrategyPackageText",
			"Operation", "OperationText", "OperationMaterial",
		}
	}

	caller.AsyncGetMaintenanceTaskList(
		inoutSDC.MaintenanceTaskList.TaskListType,
		inoutSDC.MaintenanceTaskList.TaskListGroup,
		inoutSDC.MaintenanceTaskList.TaskListGroupCounter,
		inoutSDC.MaintenanceTaskList.TaskListVersionCounter,
		inoutSDC.MaintenanceTaskList.Equipment,
		inoutSDC.MaintenanceTaskList.Plant,
		inoutSDC.MaintenanceTaskList.StrategyPackage.TaskListSequence,
		inoutSDC.MaintenanceTaskList.StrategyPackage.MaintenancePackageText,
		inoutSDC.MaintenanceTaskList.StrategyPackage.Operation.TechnicalObject,
		inoutSDC.MaintenanceTaskList.StrategyPackage.Operation.OperationText,
		accepter,
	)
}