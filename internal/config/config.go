package config

import (
	"flag"
	"tronglv_upload_svc/helper/s3"
	"tronglv_upload_svc/helper/server"

	"github.com/zeromicro/go-zero/core/conf"
)

func Load(file *string) Config {
	flag.Parse()
	var c Config
	conf.MustLoad(*file, &c, conf.UseEnv())
	return c
}

type Config struct {
	Server server.Config     `json:"server,optional"`
	AwsS3  s3.BucketS3Config `json:"aws-s3,optional"`
}

func (c Config) ServiceName() string {
	return c.Server.Http.Name
}
