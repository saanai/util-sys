package mysql

import (
	"time"

	"github.com/cenkalti/backoff"
	"github.com/saanai/util-sys/entity"
	"github.com/sirupsen/logrus"
)

// RetryableClient Mysql client wrapper for retry
type RetryableClient struct {
	c     *client
	retry uint64
}

var Rc *RetryableClient
var b backoff.BackOff

func init() {
	Rc = &RetryableClient{
		c:     c,
		retry: 5,
	}
	b = backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(Rc.retry))
}

func (rc *RetryableClient) CreateUser(user *entity.User) (err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to create user. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		err = rc.c.createUser(user)
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return err
	}

	return err
}

func (rc *RetryableClient) GetUserByEmail(email string) (user *entity.User, err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to get user by email. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		user, err = rc.c.getUserByEmail(email)
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (rc *RetryableClient) CreateSession(session *entity.Session) (err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to create session. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		err = rc.c.createSession(session)
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return err
	}

	return err
}

func (rc *RetryableClient) CheckSession(uuid string) (valid bool, err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to check session. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		_, err = rc.c.checkSession(uuid)
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return false, err
	}

	return true, err
}

func (rc *RetryableClient) GetAllThreads() (threads []entity.Thread, err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to get all threads. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		threads, err = rc.c.getAllThreads()
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return threads, err
	}

	return threads, err
}

func (rc *RetryableClient) GetNumRepliesOfAnyThread(threadId int64) (count int64, err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to get all threads. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		count, err = rc.c.getNumRepliesOfAnyThread(threadId)
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return count, err
	}

	return count, err
}

func (rc *RetryableClient) GetUserByUserId(userId int64) (user *entity.User, err error) {

	notify := func(err error, d time.Duration) {
		logrus.Warnf("failed to get all threads. retrying error:%v, backoff %v sec", err.Error(), d.Seconds())
	}
	operation := func() (err error) {
		user, err = rc.c.getUserByUserId(userId)
		if err != nil {
			return err
		}
		return err
	}

	err = backoff.RetryNotify(operation, b, notify)
	if err != nil {
		return user, err
	}

	return user, err
}
