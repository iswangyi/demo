package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

const message = `
mesg = {"apiKey":"823f3995baecc0aa77c073571294bd97","secureToken":"","providerUrl":"http://keystone.openstack.svc.cluster.local/v3","tenantUserName":"admin","tenantUserPass":"Admin@ES20!8","tenantName":"","tenantId":"4afe31713d49470a8b35d48242a85485","domain":"default","region":"RegionOne","asynScan":False,"managementUser":"","managementPassword":"","serverIp":"","serverPort":"","subscriptionId":"","msgId":"13ce8b02-24b7-4310-b53e-7d58522f8da7","name":"","availabilityZone":"","siteUri":"","resourceGroupName":"","joinMaas":"","resourceGroup":"","ascmAK":"","ascmSK":"","asoAK":"","asoSK":"","projectId":"","cloudVersion":"V6","providerType":"ESCloud","virtEnvType":"","virtEnvUuid":"","cloudEnvId":70,"opUser":"","orgSid":"","proxyEnabled":"0","httpProxyHost":"","httpProxyPort":"","httpProxyUsername":"","httpProxyPassword":"","httpsProxyHost":"","httpsProxyPort":"","httpsProxyUsername":"","httpsProxyPassword":"","operationalUrl":"","operationalUser":"","operationalPassword":"","operationalPort":"","options":"","rcLinkId":"","orgFlag":False,"appId":"","uin":"","companyId":"","taskId":"","attrData":"","tenantUserId":"","manageOneUrl":"","manageOneUser":"","manageOnePassword":"","manageOnePort":"","monitorUrl":"","monitorUser":"","monitorPassword":"","monitorPort":"","monitorLimitCount":-1,"monitorLimitCountExpireTime":-1,"fiveMinuteDataHour":120,"endpointConvertAddress":"","rclinkType":False,"proxyEnable":False}
        mesg = {"apiKey":"","secureToken":"","providerUrl":"10.69.3.253","tenantUserName":"","tenantUserPass":"","tenantName":"","tenantId":"","domain":"","region":"","asynScan":False,"managementUser":"gesysman","managementPassword":"12#$qwER","serverIp":"","serverPort":"7443","subscriptionId":"","msgId":"20fc4db9-b49c-441e-93d6-525ea77267cc","name":"","availabilityZone":"","siteUri":"","resourceGroupName":"","joinMaas":"","resourceGroup":"","ascmAK":"","ascmSK":"","asoAK":"","asoSK":"","projectId":"","cloudVersion":"8.0","providerType":"FusionCompute","virtEnvType":"","virtEnvUuid":"","cloudEnvId":64,"opUser":"","orgSid":"","proxyEnabled":"0","httpProxyHost":"","httpProxyPort":"","httpProxyUsername":"","httpProxyPassword":"","httpsProxyHost":"","httpsProxyPort":"","httpsProxyUsername":"","httpsProxyPassword":"","operationalUrl":"","operationalUser":"","operationalPassword":"","operationalPort":"","options":"","rcLinkId":"","orgFlag":False,"appId":"","uin":"","companyId":"","taskId":"62ceb42894e4b83a2e5929bc","attrData":"","tenantUserId":"","manageOneUrl":"","manageOneUser":"","manageOnePassword":"","manageOnePort":"","monitorUrl":"","monitorUser":"","monitorPassword":"","monitorPort":"","monitorLimitCount":-1,"monitorLimitCountExpireTime":-1,"fiveMinuteDataHour":120,"endpointConvertAddress":"","previousFireTime":1657713404261,"fireTime":1657713704286,"rclinkType":False,"proxyEnable":False}
`

func main() {
	conn, err := amqp.Dial("amqp://admin:jpcrgLURgM00K_YjN@10.96.1.65:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"go_demo_rabbitmq", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message + time.Now().String()

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < 1000; i++ {
				time.Sleep(1 * time.Second)
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(body),
					})
				failOnError(err, "Failed to publish a message")
				log.Printf(" [x] Sent %s\n", body)
			}
			wg.Done()

		}(&wg)
	}
	wg.Wait()
}
