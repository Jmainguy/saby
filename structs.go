package main

type Volume struct {
	VolumeMounts []VolumeMount `yaml:"volumeMounts"`
}
type VolumeMount struct {
	Mount      interface{} `yaml:"mount"`
	MountPath  string      `yaml:"mountPath"`
	SecretName string      `yaml:"secretName"`
}
