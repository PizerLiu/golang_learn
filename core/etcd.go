package core

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"pizer_project/conf"
	"pizer_project/globle"
	"pizer_project/utils"
	"strconv"
	"time"
)

var (
	etcdConf *utils.CfgFileConfig
)

//es操作方法结构体
type EtcdExecute struct {
	etcd *clientv3.Client
}

func (execute *EtcdExecute) RegisterServer(serverName string, endpoint string) {
	/*
		服务注册到etcd
	*/
	var endpoints []string
	if getResp, err := execute.etcd.Get(context.Background(), "/grpc_server_register/"+serverName, clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		// ttl设置为60
		resp, err := execute.etcd.Grant(context.TODO(), int64(60))
		if err != nil {
			log.Fatal("create clientv3 lease failed: %v", err)
		}
		onRegisterServerNum := getResp.Count
		if onRegisterServerNum == 0 {
			if _, err := execute.etcd.Put(
				context.TODO(),
				"/grpc_server_register/"+serverName+"/"+strconv.Itoa(int(resp.ID)),
				endpoint,
				clientv3.WithLease(resp.ID)); err != nil {
				log.Fatal("register server '%s' with ttl to clientv3 failed: %s", endpoint, err.Error())
			}
		} else {
			for _, ev := range getResp.Kvs {
				endpoints = append(endpoints, string(ev.Value))
				if endpoint == string(ev.Value) {
					//已经记录该endpoint,刷新下
					if _, err := execute.etcd.Put(
						context.TODO(),
						string(ev.Key),
						endpoint,
						clientv3.WithLease(resp.ID)); err != nil {
						log.Fatal("register server '%s' with ttl to clientv3 failed: %s", endpoint, err.Error())
					}
					return
				}
			}
			if _, err := execute.etcd.Put(context.TODO(),
				"/grpc_server_register/"+serverName+"/"+strconv.Itoa(int(resp.ID)),
				endpoint,
				clientv3.WithLease(resp.ID)); err != nil {
				log.Fatal("register server '%s' with ttl to clientv3 failed: %s", endpoint, err.Error())
			}
		}
		if _, err := execute.etcd.KeepAlive(context.TODO(), resp.ID); err != nil {
			log.Fatal("refresh service '%s' with ttl to clientv3 failed: %s", serverName, err.Error())
		}
	}
}

func (execute *EtcdExecute) GetServerRegisterByName(serverName string) string {
	var endpoints []string
	if getResp, err := execute.etcd.Get(context.Background(), "/grpc_server_register/"+serverName, clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		// 获取成功, 我们遍历所有的kvs
		for _, ev := range getResp.Kvs {
			endpoints = append(endpoints, string(ev.Value))
		}
	}
	//负载均衡选择，从endpoints中选择一个endpoint
	return endpoints[0]
}

//etcd操作方法 [add key-value]
func (execute *EtcdExecute) AddKeyValue(key string, value string) bool {
	_, err := execute.etcd.Put(context.Background(), key, value)
	if err != nil {
		return false
	}
	return true
}

// todo
//  0. Watcher机制，client端更新提供服务的server list; 监测记录在etcd的配置文件，并更新
//  1. etcd进行低耦合的心跳保持，TTL key，续约方式
//  2. 客户端通过etcd拿服务注册server时，如果拿到后请求失败，重新从etcd中拿一个新的server
//  3. etcd初始化这里，分为：server端和client端
//  4. leader 选举的期间， 集群不能处理任何写
//  5. 服务发现etcd如果挂掉了，需要在server本地缓存一份服务list

func etcdInit() *EtcdExecute {
	etcdConf = conf.GetBaseConf()
	addresses := etcdConf.Get(globle.CONST_CONFIG_SECTION_ETCD, "addresses")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addresses},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("create clientv3 client failed: %v", err)
	}

	etcdExe := EtcdExecute{cli}

	return &etcdExe
}
