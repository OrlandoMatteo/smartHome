package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GetConfig(jsonFile string) *Configs {
	jsonByte, err := ioutil.ReadFile(jsonFile)
	conf := &Configs{}
	if err == nil {
		json.Unmarshal(jsonByte, &conf)
	} else {
		fmt.Println(err)
	}
	return conf
}

func CreateSensor(conf Configs) map[string]DeviceDocker {
	devices := make(map[string]DeviceDocker)
	var dependsOn []string
	var links []string
	for _, v := range conf.Services {
		serviceName := v.Container_Name
		dependsOn = append(dependsOn, serviceName)
		links = append(links, serviceName)
	}
	for _, owner := range conf.Owners {
		for _, house := range owner.Houses {
			for k := 0; k < house.Quantity; k++ {
				deviceName := owner.Name + "_" + fmt.Sprint(k)
				envinronmentName := map[string]string{"FILENAME": "cfg/" + deviceName}
				volume := [1]string{"./cfg:/app/cfg/"}
				device := DeviceDocker{"orterra/gosensor", volume, deviceName, envinronmentName, dependsOn, links}
				devices[deviceName] = device
				CreateSensorConfig(owner.Name, deviceName, k, conf.CatalogAddress)

			}
		}
	}
	return devices
}

func CreateServices(conf Configs) map[string]Service {
	services := make(map[string]Service)
	for _, x := range conf.Services {
		service := Service{x.Container_Name, "orterra/" + x.Container_Name, x.Expose, x.Ports}
		services[x.Container_Name] = service
	}
	return services
}

func CreateCompose(conf Configs) {
	devices := CreateSensor(conf)
	services := CreateServices(conf)
	output := make(map[string]interface{})
	for k, v := range devices {
		output[k] = v
	}
	for k, v := range services {
		output[k] = v
	}
	out := yaml.MapSlice{{"version", "3.5"}, {"services", output}}

	data, err := yaml.Marshal(&out)
	if err != nil {
		fmt.Println("error:", err)
	}
	err2 := ioutil.WriteFile("docker-compose.yml", data, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func CreateSensorConfig(ownerName, deviceName string, ID int, catalogAddress string) {
	sensor := Sensor{"Temperature", "C", 10, deviceName, 1000, ownerName, ID, catalogAddress}
	data, err := yaml.Marshal(&sensor)
	if err != nil {
		fmt.Println("error:", err)
	}
	fileName := "cfg/" + deviceName + ".yml"
	err2 := ioutil.WriteFile(fileName, data, 0644)
	if err2 != nil {
		fmt.Println(err2)
	}

}

// TODO write q cfg file for every sensor
