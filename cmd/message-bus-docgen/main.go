package main

import (
	"github.com/wissemmansouri/OpenIT.one-AppManagement/codegen/message_bus"
	"github.com/wissemmansouri/OpenIT.one-AppManagement/common"
	"github.com/wissemmansouri/OpenIT.one-Common/external"
	"github.com/samber/lo"
)

func main() {
	eventTypes := lo.Map(common.EventTypes, func(item message_bus.EventType, index int) external.EventType {
		return external.EventType{
			Name:     item.Name,
			SourceID: item.SourceID,
			PropertyTypeList: lo.Map(
				item.PropertyTypeList, func(item message_bus.PropertyType, index int) external.PropertyType {
					return external.PropertyType{
						Name:        item.Name,
						Description: item.Description,
						Example:     item.Example,
					}
				},
			),
		}
	})

	external.PrintEventTypesAsMarkdown(common.AppManagementServiceName, common.AppManagementVersion, eventTypes)
}
