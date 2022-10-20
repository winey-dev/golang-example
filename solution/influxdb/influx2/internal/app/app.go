package app

import (
	"context"
	"fmt"
	"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"

	"influx2/config"
	"influx2/internal/pkg/argument"
)

var ctx = context.Background()

type App struct {
	agent AgentService
}

type AppService interface {
	Start(cfg *config.Config, arg argument.Arguments) error
}

type AgentService interface {
	Run(client influxdb2.Client) error
}

func NewApp(agent AgentService) AppService {
	return &App{agent: agent}
}

func (a *App) Start(cfg *config.Config, arg argument.Arguments) error {
	dbCfg := cfg.InfluxDB

	// New Client
	client := influxdb2.NewClient(dbCfg.EndPoint, dbCfg.Token)
	defer client.Close()

	/*
		// InfluxDB 2.0 SignIn/SignOut
		err := client.UsersAPI().SignIn(ctx, dbCfg.UserName, dbCfg.Password)
		if err != nil {
			fmt.Printf("influxdb2 sign in failed. err=%v\n", err)
			return err
		}
		defer client.UsersAPI().SignOut(ctx)
	*/

	// InfluxDB 2.0  Exists bucket check,
	// If not exists, Create bucket
	/*
		var (
			bucket *domain.Bucket
		)
	*/

	//exsits bucket check
	//bucket, err = client.BucketsAPI().FindBucketByName(ctx, dbCfg.Bucket)
	orgObj, err := client.OrganizationsAPI().FindOrganizationByName(ctx, dbCfg.OrgName)
	if err != nil {
		fmt.Printf("get org name(%s) failed. err=%v\n", dbCfg.OrgName, err)
		return err
	}

	dbCfg.OrgID = *orgObj.Id
	_, err = client.BucketsAPI().FindBucketByName(ctx, dbCfg.Bucket)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// create database
			fmt.Printf("create bucket name (%s)\n", dbCfg.Bucket)

			// 604800 = > 1week
			// create buckets
			//bucket, err = client.BucketsAPI().CreateBucketWithNameWithID(ctx, dbCfg.OrgID, dbCfg.Bucket, domain.RetentionRule{EverySeconds: 604800})
			_, err = client.BucketsAPI().CreateBucketWithNameWithID(ctx, dbCfg.OrgID, dbCfg.Bucket, domain.RetentionRule{EverySeconds: 604800})
			if err != nil {
				fmt.Printf("create bucket name (%s) failed. err=%v\n", dbCfg.Bucket, err)
				return err
			}
		} else {
			fmt.Printf("find by bucket name (%s) failed. err=%v\n", dbCfg.Bucket, err)
			return err
		}
	}

	err = a.agent.Run(client)
	if err != nil {
		fmt.Printf("agent run failed. err=%v\n", err)
		return err
	}

	return nil
}
