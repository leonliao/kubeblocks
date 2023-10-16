// Copyright (C) 2022-2023 ApeCloud Co., Ltd
//
// This file is part of KubeBlocks project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

parameters: {
	input_configs: *"`config.LogsCollector`" | string
	container_id: *"`container_id`" | string
	storage: *"file_storage/oteld" | string
}


output: {
	container_filelog: {
		rule: "type == \"container\" && config != nil && config.EnabledLogs"
    config: {
    	input_configs: parameters.input_configs
      container_id: parameters.container_id
      pod_id: "`endpoint`"
      storage: parameters.storage
      cluster_name: "`config.ClusterName`"
			component_name: "`config.ComponentName`"
			character_type: "`config.CharacterType`"
    }
	}

}
