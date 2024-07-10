//go:build !ignore_autogenerated

/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoTrigger) DeepCopyInto(out *AutoTrigger) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoTrigger.
func (in *AutoTrigger) DeepCopy() *AutoTrigger {
	if in == nil {
		return nil
	}
	out := new(AutoTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConstraint) DeepCopyInto(out *ConfigConstraint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConstraint.
func (in *ConfigConstraint) DeepCopy() *ConfigConstraint {
	if in == nil {
		return nil
	}
	out := new(ConfigConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigConstraint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConstraintList) DeepCopyInto(out *ConfigConstraintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConstraintList.
func (in *ConfigConstraintList) DeepCopy() *ConfigConstraintList {
	if in == nil {
		return nil
	}
	out := new(ConfigConstraintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigConstraintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConstraintSpec) DeepCopyInto(out *ConfigConstraintSpec) {
	*out = *in
	if in.ReloadAction != nil {
		in, out := &in.ReloadAction, &out.ReloadAction
		*out = new(ReloadAction)
		(*in).DeepCopyInto(*out)
	}
	if in.MergeReloadAndRestart != nil {
		in, out := &in.MergeReloadAndRestart, &out.MergeReloadAndRestart
		*out = new(bool)
		**out = **in
	}
	if in.ReloadStaticParamsBeforeRestart != nil {
		in, out := &in.ReloadStaticParamsBeforeRestart, &out.ReloadStaticParamsBeforeRestart
		*out = new(bool)
		**out = **in
	}
	if in.DownwardAPIChangeTriggeredActions != nil {
		in, out := &in.DownwardAPIChangeTriggeredActions, &out.DownwardAPIChangeTriggeredActions
		*out = make([]DownwardAPIChangeTriggeredAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ParametersSchema != nil {
		in, out := &in.ParametersSchema, &out.ParametersSchema
		*out = new(ParametersSchema)
		(*in).DeepCopyInto(*out)
	}
	if in.StaticParameters != nil {
		in, out := &in.StaticParameters, &out.StaticParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DynamicParameters != nil {
		in, out := &in.DynamicParameters, &out.DynamicParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImmutableParameters != nil {
		in, out := &in.ImmutableParameters, &out.ImmutableParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FileFormatConfig != nil {
		in, out := &in.FileFormatConfig, &out.FileFormatConfig
		*out = new(FileFormatConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConstraintSpec.
func (in *ConfigConstraintSpec) DeepCopy() *ConfigConstraintSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigConstraintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigConstraintStatus) DeepCopyInto(out *ConfigConstraintStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigConstraintStatus.
func (in *ConfigConstraintStatus) DeepCopy() *ConfigConstraintStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigConstraintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownwardAPIChangeTriggeredAction) DeepCopyInto(out *DownwardAPIChangeTriggeredAction) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.DownwardAPIVolumeFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ScriptConfig != nil {
		in, out := &in.ScriptConfig, &out.ScriptConfig
		*out = new(ScriptConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownwardAPIChangeTriggeredAction.
func (in *DownwardAPIChangeTriggeredAction) DeepCopy() *DownwardAPIChangeTriggeredAction {
	if in == nil {
		return nil
	}
	out := new(DownwardAPIChangeTriggeredAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileFormatConfig) DeepCopyInto(out *FileFormatConfig) {
	*out = *in
	in.FormatterAction.DeepCopyInto(&out.FormatterAction)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileFormatConfig.
func (in *FileFormatConfig) DeepCopy() *FileFormatConfig {
	if in == nil {
		return nil
	}
	out := new(FileFormatConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FormatterAction) DeepCopyInto(out *FormatterAction) {
	*out = *in
	if in.IniConfig != nil {
		in, out := &in.IniConfig, &out.IniConfig
		*out = new(IniConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FormatterAction.
func (in *FormatterAction) DeepCopy() *FormatterAction {
	if in == nil {
		return nil
	}
	out := new(FormatterAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IniConfig) DeepCopyInto(out *IniConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IniConfig.
func (in *IniConfig) DeepCopy() *IniConfig {
	if in == nil {
		return nil
	}
	out := new(IniConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterDeletedPolicy) DeepCopyInto(out *ParameterDeletedPolicy) {
	*out = *in
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterDeletedPolicy.
func (in *ParameterDeletedPolicy) DeepCopy() *ParameterDeletedPolicy {
	if in == nil {
		return nil
	}
	out := new(ParameterDeletedPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersDefinition) DeepCopyInto(out *ParametersDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersDefinition.
func (in *ParametersDefinition) DeepCopy() *ParametersDefinition {
	if in == nil {
		return nil
	}
	out := new(ParametersDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParametersDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersDefinitionList) DeepCopyInto(out *ParametersDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ParametersDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersDefinitionList.
func (in *ParametersDefinitionList) DeepCopy() *ParametersDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ParametersDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParametersDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersDefinitionSpec) DeepCopyInto(out *ParametersDefinitionSpec) {
	*out = *in
	if in.FileFormatConfig != nil {
		in, out := &in.FileFormatConfig, &out.FileFormatConfig
		*out = new(FileFormatConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ParametersSchema != nil {
		in, out := &in.ParametersSchema, &out.ParametersSchema
		*out = new(ParametersSchema)
		(*in).DeepCopyInto(*out)
	}
	if in.ReloadAction != nil {
		in, out := &in.ReloadAction, &out.ReloadAction
		*out = new(ReloadAction)
		(*in).DeepCopyInto(*out)
	}
	if in.AsEnvFrom != nil {
		in, out := &in.AsEnvFrom, &out.AsEnvFrom
		*out = new(bool)
		**out = **in
	}
	if in.MergeReloadAndRestart != nil {
		in, out := &in.MergeReloadAndRestart, &out.MergeReloadAndRestart
		*out = new(bool)
		**out = **in
	}
	if in.ReloadStaticParamsBeforeRestart != nil {
		in, out := &in.ReloadStaticParamsBeforeRestart, &out.ReloadStaticParamsBeforeRestart
		*out = new(bool)
		**out = **in
	}
	if in.StaticParameters != nil {
		in, out := &in.StaticParameters, &out.StaticParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DynamicParameters != nil {
		in, out := &in.DynamicParameters, &out.DynamicParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImmutableParameters != nil {
		in, out := &in.ImmutableParameters, &out.ImmutableParameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersDefinitionSpec.
func (in *ParametersDefinitionSpec) DeepCopy() *ParametersDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ParametersDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersDefinitionStatus) DeepCopyInto(out *ParametersDefinitionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersDefinitionStatus.
func (in *ParametersDefinitionStatus) DeepCopy() *ParametersDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ParametersDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersSchema) DeepCopyInto(out *ParametersSchema) {
	*out = *in
	if in.SchemaInJSON != nil {
		in, out := &in.SchemaInJSON, &out.SchemaInJSON
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersSchema.
func (in *ParametersSchema) DeepCopy() *ParametersSchema {
	if in == nil {
		return nil
	}
	out := new(ParametersSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReloadAction) DeepCopyInto(out *ReloadAction) {
	*out = *in
	if in.UnixSignalTrigger != nil {
		in, out := &in.UnixSignalTrigger, &out.UnixSignalTrigger
		*out = new(UnixSignalTrigger)
		**out = **in
	}
	if in.ShellTrigger != nil {
		in, out := &in.ShellTrigger, &out.ShellTrigger
		*out = new(ShellTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.TPLScriptTrigger != nil {
		in, out := &in.TPLScriptTrigger, &out.TPLScriptTrigger
		*out = new(TPLScriptTrigger)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoTrigger != nil {
		in, out := &in.AutoTrigger, &out.AutoTrigger
		*out = new(AutoTrigger)
		**out = **in
	}
	if in.TargetPodSelector != nil {
		in, out := &in.TargetPodSelector, &out.TargetPodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReloadAction.
func (in *ReloadAction) DeepCopy() *ReloadAction {
	if in == nil {
		return nil
	}
	out := new(ReloadAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptConfig) DeepCopyInto(out *ScriptConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptConfig.
func (in *ScriptConfig) DeepCopy() *ScriptConfig {
	if in == nil {
		return nil
	}
	out := new(ScriptConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShellTrigger) DeepCopyInto(out *ShellTrigger) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(bool)
		**out = **in
	}
	if in.BatchReload != nil {
		in, out := &in.BatchReload, &out.BatchReload
		*out = new(bool)
		**out = **in
	}
	if in.ToolsSetup != nil {
		in, out := &in.ToolsSetup, &out.ToolsSetup
		*out = new(ToolsSetup)
		(*in).DeepCopyInto(*out)
	}
	if in.ScriptConfig != nil {
		in, out := &in.ScriptConfig, &out.ScriptConfig
		*out = new(ScriptConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShellTrigger.
func (in *ShellTrigger) DeepCopy() *ShellTrigger {
	if in == nil {
		return nil
	}
	out := new(ShellTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TPLScriptTrigger) DeepCopyInto(out *TPLScriptTrigger) {
	*out = *in
	out.ScriptConfig = in.ScriptConfig
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TPLScriptTrigger.
func (in *TPLScriptTrigger) DeepCopy() *TPLScriptTrigger {
	if in == nil {
		return nil
	}
	out := new(TPLScriptTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolConfig) DeepCopyInto(out *ToolConfig) {
	*out = *in
	if in.AsContainerImage != nil {
		in, out := &in.AsContainerImage, &out.AsContainerImage
		*out = new(bool)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolConfig.
func (in *ToolConfig) DeepCopy() *ToolConfig {
	if in == nil {
		return nil
	}
	out := new(ToolConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToolsSetup) DeepCopyInto(out *ToolsSetup) {
	*out = *in
	if in.ToolConfigs != nil {
		in, out := &in.ToolConfigs, &out.ToolConfigs
		*out = make([]ToolConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToolsSetup.
func (in *ToolsSetup) DeepCopy() *ToolsSetup {
	if in == nil {
		return nil
	}
	out := new(ToolsSetup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnixSignalTrigger) DeepCopyInto(out *UnixSignalTrigger) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnixSignalTrigger.
func (in *UnixSignalTrigger) DeepCopy() *UnixSignalTrigger {
	if in == nil {
		return nil
	}
	out := new(UnixSignalTrigger)
	in.DeepCopyInto(out)
	return out
}
