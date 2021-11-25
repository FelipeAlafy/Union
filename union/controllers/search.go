package controllers

import (
	"strings"

	"com.github/FelipeAlafy/union/manager"
	"com.github/FelipeAlafy/union/osmanager"
)


func reload() {
	Instances := manager.GetClients(osmanager.GetInstances())
	var nInst []manager.Client
	if !FilterArchived {
		//in this case get all instances when it has Archived property equals false
		for _, ins := range Instances {
			filter := ins.FilterByArchived(FilterArchived)
			if filter {
				nInst = append(nInst, ins)
			}
		}
		instances = nInst
		clientLimit = len(nInst)
	} else {
		//In this case get all instances
		instances = Instances
		clientLimit = len(instances)
	}
}

func SearchForClients(name string, archived bool) (int) {
	FilterArchived = archived
	reload()
	//look for a client
	for i, ins := range instances {
		if strings.EqualFold(name, ins.Nome) {
			return i
		}
	}
	return 0
}