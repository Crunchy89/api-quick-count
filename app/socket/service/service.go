package service

import (
	"github.com/centrifugal/centrifuge"
	"github.com/sirupsen/logrus"
)

func Connect(socket *centrifuge.Node, logger *logrus.Entry) {
	socket.OnConnect(func(client *centrifuge.Client) {
		client.OnSubscribe(func(e centrifuge.SubscribeEvent, cb centrifuge.SubscribeCallback) {
			logger.Info("client subscribe on ", e.Channel)
			cb(centrifuge.SubscribeReply{}, nil)
		})
		client.OnPublish(func(e centrifuge.PublishEvent, cb centrifuge.PublishCallback) {
			cb(centrifuge.PublishReply{}, nil)
		})
	})
}
