package bm

import (
	"testing"
)

func TestDescribeDevice(t *testing.T) {

	client, _ := NewClientFromEnv()

	lanIps := []string{"10.0.0.4"}
	req := DescribeDeviceArgs{
		LanIps: &lanIps,
	}

	if devInfo, err := client.DescribeDevice(&req); err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("DescribeDevice Pass devInfo=%v", devInfo)
	}

}

func TestSubnetIp(t *testing.T) {
	client, _ := NewClientFromEnv()
	taskId, err := client.RegisterContainerSubnetIp("vpc-muinpf9p", "subnet-o4xwhqa8")
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = client.WaitUntiTaskDone(taskId, 60)
	if err != nil {
		t.Error(err.Error())
		return
	}

	taskId, err = client.ReleaseContainerSubnetIp("vpc-muinpf9p", "subnet-o4xwhqa8")
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = client.WaitUntiTaskDone(taskId, 60)
	if err != nil {
		t.Error(err.Error())
		return
	}

}
