package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	var list []string
	var volume Volume
	data, err := ioutil.ReadFile("items.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &volume)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range volume.VolumeMounts {
		var msg string
		msg = fmt.Sprintf(`{"name":"%s","mountPath":"%s","readonly":true}`, v.SecretName, v.MountPath)
		list = append(list, msg)
	}
	fmt.Println(list)
}
