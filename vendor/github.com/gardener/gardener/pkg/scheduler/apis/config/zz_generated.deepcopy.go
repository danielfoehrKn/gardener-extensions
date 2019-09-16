// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketSchedulerConfiguration) DeepCopyInto(out *BackupBucketSchedulerConfiguration) {
	*out = *in
	out.RetrySyncPeriod = in.RetrySyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketSchedulerConfiguration.
func (in *BackupBucketSchedulerConfiguration) DeepCopy() *BackupBucketSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupBucketSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntrySchedulerConfiguration) DeepCopyInto(out *BackupEntrySchedulerConfiguration) {
	*out = *in
	out.RetrySyncPeriod = in.RetrySyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntrySchedulerConfiguration.
func (in *BackupEntrySchedulerConfiguration) DeepCopy() *BackupEntrySchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupEntrySchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConfiguration) DeepCopyInto(out *DiscoveryConfiguration) {
	*out = *in
	if in.DiscoveryCacheDir != nil {
		in, out := &in.DiscoveryCacheDir, &out.DiscoveryCacheDir
		*out = new(string)
		**out = **in
	}
	if in.HTTPCacheDir != nil {
		in, out := &in.HTTPCacheDir, &out.HTTPCacheDir
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConfiguration.
func (in *DiscoveryConfiguration) DeepCopy() *DiscoveryConfiguration {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaderElectionConfiguration = in.LeaderElectionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerConfiguration) DeepCopyInto(out *SchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ClientConnection = in.ClientConnection
	out.LeaderElection = in.LeaderElection
	in.Discovery.DeepCopyInto(&out.Discovery)
	out.Server = in.Server
	in.Schedulers.DeepCopyInto(&out.Schedulers)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerConfiguration.
func (in *SchedulerConfiguration) DeepCopy() *SchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerControllerConfiguration) DeepCopyInto(out *SchedulerControllerConfiguration) {
	*out = *in
	if in.BackupBucket != nil {
		in, out := &in.BackupBucket, &out.BackupBucket
		*out = new(BackupBucketSchedulerConfiguration)
		**out = **in
	}
	if in.Shoot != nil {
		in, out := &in.Shoot, &out.Shoot
		*out = new(ShootSchedulerConfiguration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerControllerConfiguration.
func (in *SchedulerControllerConfiguration) DeepCopy() *SchedulerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SchedulerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	out.HTTP = in.HTTP
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootSchedulerConfiguration) DeepCopyInto(out *ShootSchedulerConfiguration) {
	*out = *in
	out.RetrySyncPeriod = in.RetrySyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootSchedulerConfiguration.
func (in *ShootSchedulerConfiguration) DeepCopy() *ShootSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}
