// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"alidns/log"
	"encoding/json"
	"flag"
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
	"strings"
	"time"
)

type Client struct {
	alidns20150109.Client
}

// CreateClient
/*
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-shanghai.aliyuncs.com")
	client := new(Client)
	err := client.Init(config)
	return client, err
}

// AddDomainRecord 添加记录
func (c *Client) AddDomainRecord(DomainName, Type, RR, Value string) (_result *alidns20150109.AddDomainRecordResponse, _err error) {
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: tea.String(DomainName),
		Type:       tea.String(Type),
		RR:         tea.String(RR),
		Value:      tea.String(Value),
	}
	runtime := &util.RuntimeOptions{}
	return c.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
}

// DeleteSubDomainRecords 删除记录
func (c *Client) DeleteSubDomainRecords(DomainName, Type, RR string) (_result *alidns20150109.DeleteSubDomainRecordsResponse, _err error) {
	deleteSubDomainRecordsRequest := &alidns20150109.DeleteSubDomainRecordsRequest{
		DomainName: tea.String(DomainName),
		Type:       tea.String(Type),
		RR:         tea.String(RR),
	}
	runtime := &util.RuntimeOptions{}
	return c.DeleteSubDomainRecordsWithOptions(deleteSubDomainRecordsRequest, runtime)
}

// DeleteDomainRecord 删除记录
func (c *Client) DeleteDomainRecord(RecordId *string) (_result *alidns20150109.DeleteDomainRecordResponse, _err error) {
	deleteDomainRecordRequest := &alidns20150109.DeleteDomainRecordRequest{
		RecordId: RecordId,
	}
	runtime := &util.RuntimeOptions{}
	return c.DeleteDomainRecordWithOptions(deleteDomainRecordRequest, runtime)
}

// DescribeDomainRecords 查询记录
func (c *Client) DescribeDomainRecords(DomainName, Type, RRKeyWord string) (_result *alidns20150109.DescribeDomainRecordsResponse, _err error) {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(DomainName),
		Type:       tea.String(Type),
		RRKeyWord:  tea.String(RRKeyWord),
	}
	runtime := &util.RuntimeOptions{}
	return c.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
}

var (
	AccessKey string
)

type Config struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

func init() {
	flag.StringVar(&AccessKey, "AK", "", "Set AccessKey")
	flag.Parse()
}

func main() {

	Home := os.Getenv("HOME")
	ConfigPath := Home + "/.alidns"
	ConfigFile := ConfigPath + "/ali.json"

	cfg := &Config{}

	if AccessKey != "" {
		ak := strings.Split(AccessKey, "=")
		if AccessKey == "" || len(ak) != 2 {
			log.Error("AccessKeyId and AccessKeySecret must be specified. (e.g -AK id=secret)")
			os.Exit(1)
		}

		cfg.AccessKeyId = ak[0]
		cfg.AccessKeySecret = ak[1]

		data, err := json.Marshal(cfg)
		if err != nil {
			log.Error("set AccessKey failed:", err.Error())
			os.Exit(1)
		}
		_ = os.MkdirAll(ConfigPath, 0755)
		_ = os.WriteFile(ConfigFile, data, 0755)
		log.Info("set AccessKey success")
		os.Exit(0)
	}

	data, err := os.ReadFile(ConfigFile)
	if err == nil {
		err = json.Unmarshal(data, cfg)
	}
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	AccessKeyId := tea.String(cfg.AccessKeyId)
	AccessKeySecret := tea.String(cfg.AccessKeySecret)

	if *util.Empty(AccessKeyId) || *util.Empty(AccessKeySecret) {
		log.Error("need set AccessKey")
		os.Exit(1)
	}

	client, err := CreateClient(AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}

	CertbotDomain := os.Getenv("CERTBOT_DOMAIN")
	CertbotValidation := os.Getenv("CERTBOT_VALIDATION")
	//CertbotToken := os.Getenv("CERTBOT_TOKEN")
	CertbotAuthOutput := os.Getenv("CERTBOT_AUTH_OUTPUT")

	if CertbotDomain == "" || CertbotValidation == "" {
		flag.Usage()
		os.Exit(0)
	}

	comps := strings.Split(CertbotDomain, ".")
	if len(comps) < 2 {
		log.Error("domain name error:", CertbotDomain)
		os.Exit(1)
	}

	n := len(comps) - 2
	Type := "TXT"
	Domain := strings.Join(comps[n:], ".")
	RR := "_acme-challenge"

	if n > 0 {
		RR += fmt.Sprintf(".%s", strings.Join(comps[:n], "."))
	}

	add := strings.Count(CertbotAuthOutput, CertbotValidation) == 0

	if add { // 添加
		log.Info("ADD -", CertbotDomain, CertbotValidation)
		_, err1 := client.AddDomainRecord(Domain, Type, RR, CertbotValidation)
		if err1 == nil {
			log.Info("success")
		} else {
			log.Error("failed:", err1)
		}

		time.Sleep(3 * time.Second)

	} else {
		log.Info("DEL -", CertbotDomain, CertbotValidation)
		resp, err1 := client.DescribeDomainRecords(Domain, Type, RR)
		if err1 != nil {
			log.Error("query failed:", err1)
			os.Exit(1)
		}

		for _, record := range resp.Body.DomainRecords.Record {
			if CertbotValidation == tea.StringValue(record.Value) && RR == tea.StringValue(record.RR) {
				_, err2 := client.DeleteDomainRecord(record.RecordId)
				if err2 == nil {
					log.Info("success")
				} else {
					log.Error("failed:", err2)
				}
				break
			}
		}
	}
}
