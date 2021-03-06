package command

import (
	"errors"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/alibaba/kt-connect/pkg/kt/cluster"
	"github.com/alibaba/kt-connect/pkg/kt/connect"
	"github.com/alibaba/kt-connect/pkg/kt/options"
	"github.com/alibaba/kt-connect/pkg/mockd/mock"
	"github.com/golang/mock/gomock"
	"github.com/urfave/cli"
)

func Test_newConnectCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	mockAction := mock.NewMockActionInterface(ctl)

	mockAction.EXPECT().Connect(gomock.Any()).Return(nil).AnyTimes()

	cases := []struct {
		testArgs               []string
		skipFlagParsing        bool
		useShortOptionHandling bool
		expectedErr            error
	}{
		{testArgs: []string{"connect", "--method", "socks5"}, skipFlagParsing: false, useShortOptionHandling: false, expectedErr: nil},
		{testArgs: []string{"connect"}, skipFlagParsing: false, useShortOptionHandling: false, expectedErr: nil},
	}

	for _, c := range cases {

		app := &cli.App{Writer: ioutil.Discard}
		set := flag.NewFlagSet("test", 0)
		_ = set.Parse(c.testArgs)

		context := cli.NewContext(app, set, nil)

		opts := options.NewDaemonOptions()
		opts.Debug = true
		command := newConnectCommand(opts, mockAction)
		err := command.Run(context)

		if c.expectedErr != nil {
			if err.Error() != c.expectedErr.Error() {
				t.Errorf("expected %v but is %v", c.expectedErr, err)
			}
		} else if err != c.expectedErr {
			t.Errorf("expected %v but is %v", c.expectedErr, err)
		}

	}
}

func Test_shouldConnectToCluster(t *testing.T) {

	ctl := gomock.NewController(t)
	kubernetes := mock.NewMockKubernetesInterface(ctl)
	shadow := mock.NewMockShadowInterface(ctl)

	kubernetes.EXPECT().CreateShadow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("172.168.0.2", "shadowName", nil).AnyTimes()
	kubernetes.EXPECT().ClusterCrids(gomock.Any()).Return([]string{"10.10.10.0/24"}, nil)

	shadow.EXPECT().Outbound("shadowName", "172.168.0.2", []string{"10.10.10.0/24"}).Return(nil)

	type args struct {
		shadow     connect.ShadowInterface
		kubernetes cluster.KubernetesInterface
		options    *options.DaemonOptions
	}

	opts := options.NewDaemonOptions()
	opts.Labels = "a:b"

	arg := args{
		shadow:     shadow,
		kubernetes: kubernetes,
		options:    opts,
	}

	if err := connectToCluster(arg.shadow, arg.kubernetes, arg.options); (err != nil) != false {
		t.Errorf("connectToCluster() error = %v, wantErr %v", err, false)
	}

}

func Test_shouldConnectClusterFailWhenFailCreateShadow(t *testing.T) {

	ctl := gomock.NewController(t)
	kubernetesInterface := mock.NewMockKubernetesInterface(ctl)
	shadowInterface := mock.NewMockShadowInterface(ctl)

	kubernetesInterface.EXPECT().CreateShadow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("", "", errors.New("")).AnyTimes()

	type args struct {
		shadow     connect.ShadowInterface
		kubernetes cluster.KubernetesInterface
		options    *options.DaemonOptions
	}

	arg := args{
		shadow:     shadowInterface,
		kubernetes: kubernetesInterface,
		options:    options.NewDaemonOptions(),
	}

	if err := connectToCluster(arg.shadow, arg.kubernetes, arg.options); (err != nil) != true {
		t.Errorf("connectToCluster() error = %v, wantErr %v", err, true)
	}

}

func Test_shouldConnectClusterFailWhenFailGetCrids(t *testing.T) {

	ctl := gomock.NewController(t)
	kubernetes := mock.NewMockKubernetesInterface(ctl)
	shadow := mock.NewMockShadowInterface(ctl)

	kubernetes.EXPECT().CreateShadow(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("172.168.0.2", "shadowName", nil).AnyTimes()
	kubernetes.EXPECT().ClusterCrids(gomock.Any()).Return([]string{}, errors.New("fail to get crid"))

	type args struct {
		shadow     connect.ShadowInterface
		kubernetes cluster.KubernetesInterface
		options    *options.DaemonOptions
	}

	opts := options.NewDaemonOptions()
	opts.Labels = "a:b"

	arg := args{
		shadow:     shadow,
		kubernetes: kubernetes,
		options:    opts,
	}

	if err := connectToCluster(arg.shadow, arg.kubernetes, arg.options); (err != nil) != true {
		t.Errorf("connectToCluster() error = %v, wantErr %v", err, true)
	}

}
