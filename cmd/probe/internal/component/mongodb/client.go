/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

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

package mongodb

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func NewMongodbClient(ctx context.Context, config *Config) (*mongo.Client, error) {
	if len(config.hosts) == 0 {
		return nil, errors.New("Get replset client whitout hosts")
	}

	opts := options.Client().
		SetHosts(config.hosts).
		SetReplicaSet(config.replSetName).
		SetAuth(options.Credential{
			Password: config.password,
			Username: config.username,
		}).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority(), writeconcern.J(true))).
		SetReadPreference(readpref.Primary()).
		SetDirect(config.direct)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "connect to mongodb")
	}
	return client, err
}

func NewReplSetClient(ctx context.Context, hosts []string) (*mongo.Client, error) {
	config := GetConfig().DeepCopy()
	config.hosts = hosts
	config.direct = false
	return NewMongodbClient(ctx, config)

}

func MongosClient(ctx context.Context, hosts []string) (*mongo.Client, error) {
	config := GetConfig().DeepCopy()
	config.hosts = hosts
	config.direct = false
	config.replSetName = ""

	return NewMongodbClient(ctx, config)
}

func StandaloneClient(ctx context.Context, host string) (*mongo.Client, error) {
	config := GetConfig().DeepCopy()
	config.hosts = []string{host}
	config.direct = true
	config.replSetName = ""

	return NewMongodbClient(ctx, config)
}
