package common

import corev1 "k8s.io/api/core/v1"

// ContainerBuilder container builder
// contains: image, imagePullPolicy, resource, ports should be required
// optional: name, command, commandArgs, containerEnv, volumeMount, livenessProbe, readinessProbe should be optional,
// optional fields should be implemented by the struct that embeds ContainerBuilder
// optional fields name usually should not be set, because container name can generate by deployment, statefulSet, daemonSet..
type ContainerBuilder struct {
	Image           string
	ImagePullPolicy corev1.PullPolicy
	Resources       corev1.ResourceRequirements
}

func NewContainerBuilder(
	Image string,
	ImagePullPolicy corev1.PullPolicy,
	Resource corev1.ResourceRequirements,
) *ContainerBuilder {
	return &ContainerBuilder{
		Image:           Image,
		ImagePullPolicy: ImagePullPolicy,
		Resources:       Resource,
	}
}

func (b *ContainerBuilder) Build(handler interface{}) corev1.Container {
	container := corev1.Container{
		Image:           b.Image,
		Resources:       b.Resources,
		ImagePullPolicy: b.ImagePullPolicy,
	}
	if containerName, ok := handler.(ContainerName); ok {
		container.Name = containerName.ContainerName()
	}
	if command, ok := handler.(Command); ok {
		container.Command = command.Command()
	}
	if containerPorts, ok := handler.(ContainerPorts); ok {
		container.Ports = containerPorts.ContainerPorts()
	}
	if commandArgs, ok := handler.(CommandArgs); ok {
		container.Args = commandArgs.CommandArgs()
	}
	if containerEnv, ok := handler.(ContainerEnv); ok {
		container.Env = containerEnv.ContainerEnv()
	}
	if volumeMount, ok := handler.(VolumeMount); ok {
		container.VolumeMounts = volumeMount.VolumeMount()
	}
	if livenessProbe, ok := handler.(LivenessProbe); ok {
		container.LivenessProbe = livenessProbe.LivenessProbe()
	}
	if readinessProbe, ok := handler.(ReadinessProbe); ok {
		container.ReadinessProbe = readinessProbe.ReadinessProbe()
	}
	return container
}

type ContainerName interface {
	ContainerName() string
}

type Command interface {
	Command() []string
}

type CommandArgs interface {
	CommandArgs() []string
}

type ContainerPorts interface {
	ContainerPorts() []corev1.ContainerPort
}

type ContainerEnv interface {
	ContainerEnv() []corev1.EnvVar
}

type VolumeMount interface {
	VolumeMount() []corev1.VolumeMount
}

type LivenessProbe interface {
	LivenessProbe() *corev1.Probe
}

type ReadinessProbe interface {
	ReadinessProbe() *corev1.Probe
}

// ContainerComponent use for define container name
type ContainerComponent string
