package cvm

type DescribeSecurityGroupExResponse struct {
	Code     int                                 `json:"code"`
	CodeDesc string                              `json:"codeDesc"`
	Message  string                              `json:"message"`
	Data     DescribeSecurityGroupExResponseData `json:"data"`
}

type DescribeSecurityGroupExResponseData struct {
	TotalCount       int                 `json:"totalNum"`
	SecurityGroupSet []SecurityGroupInfo `json:"detail"`
}

type SecurityGroupInfo struct {
	SgId           string `json:"sgId"`
	SgName         string `json:"sgName"`
	SgRemark       string `json:"sgRemark"`
	CreateTime     string `json:"createTime"`
	ProjectId      string `json:"projectId"`
	AssociateCount int    `json:"beAssociateCount"`
}

type DescribeSecurityGroupExArgs struct {
	ProjectId  int    `qcloud_arg:"projectId,omitempty"`
	SgId       string `qcloud_arg:"sgId,omitempty"`
	SgName     string `qcloud_arg:"sgName,omitempty"`
	Offset     int    `qcloud_arg:"offset,omitempty"`
	OffsetLine int    `qcloud_arg:"offsetLine,omitempty"`
	Limit      int    `qcloud_arg:"limit,omitempty"`
}

func (client *Client) DescribeSecurityGroupEx(args *DescribeSecurityGroupExArgs) (*DescribeSecurityGroupExResponse, error) {
	resp := &DescribeSecurityGroupExResponse{}
	err := client.Invoke("DescribeSecurityGroupEx", args, resp)
	if err != nil {
		return &DescribeSecurityGroupExResponse{}, err
	}
	return resp, nil
}

type CreateSecurityGroupPolicyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateSecurityGroupPolicyArgs struct {
	SgId      string                 `qcloud_arg:"sgId,required"`
	Version   int                    `qcloud_arg:"version,omitempty"`
	Direction string                 `qcloud_arg:"direction,required"`
	Index     int                    `qcloud_arg:"index,required"`
	Polices   *[]SecurityGroupPolicy `qcloud_arg:"policys,required"`
}

type SecurityGroupPolicy struct {
	IpProtocol    string `qcloud_arg:"ipProtocol" json:"ipProtocol"`
	CidrIp        string `qcloud_arg:"cidrIp,omitempty" json:"cidrIp"`
	SgId          string `qcloud_arg:"sgId,omitempty" json:"sgId"`
	AddressModule string `qcloud_arg:"addressModule,omitempty" json:"addressModule"`
	PortRange     string `qcloud_arg:"portRange,omitempty" json:"portRange"`
	ServiceModule string `qcloud_arg:"serviceModule,omitempty" json:"serviceModule"`
	Description   string `qcloud_arg:"desc,omitempty" json:"desc"`
	Action        string `qcloud_arg:"action,required" json:"action"`
}

type SecurityGroupPolicyWithIndex struct {
	SecurityGroupPolicy
	Index int `qcloud_arg:"index" json:"index"`
	//Version int `qcloud_arg:"version,omitempty" json:"version"`
}

func (client *Client) CreateSecurityGroupPolicy(args *CreateSecurityGroupPolicyArgs) (*CreateSecurityGroupPolicyResponse, error) {
	resp := &CreateSecurityGroupPolicyResponse{}
	err := client.Invoke("CreateSecurityGroupPolicy", args, resp)
	if err != nil {
		return &CreateSecurityGroupPolicyResponse{}, err
	}
	return resp, nil
}

type DescribeSecurityGroupPolicesResponse struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data struct {
		Ingress []SecurityGroupPolicyWithIndex `json:"ingress"`
		Egress  []SecurityGroupPolicyWithIndex `json:"egress"`
		Version int                            `json:"version"`
	} `json:"data"`
}

type DescribeSecurityGroupPolicesArgs struct {
	SgId string `qcloud_arg:"sgId,required"`
}

func (client *Client) DescribeSecurityGroupPolices(args *DescribeSecurityGroupPolicesArgs) (*DescribeSecurityGroupPolicesResponse, error) {
	resp := &DescribeSecurityGroupPolicesResponse{}
	err := client.Invoke("DescribeSecurityGroupPolicys", args, resp)
	if err != nil {
		return &DescribeSecurityGroupPolicesResponse{}, err
	}
	return resp, nil
}

type DeleteSecurityGroupPolicyResponse struct {
	Code     int    `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

type DeleteSecurityGroupPolicyArgs struct {
	SgId      string `qcloud_arg:"sgId,required"`
	Direction string `qcloud_arg:"direction,required"`
	Indexes   *[]int `qcloud_arg:"indexes,required"`
	Version   int    `qcloud_arg:"version,omitempty"`
}

func (client *Client) DeleteSecurityGroupPolicy(args *DeleteSecurityGroupPolicyArgs) (*DeleteSecurityGroupPolicyResponse, error) {
	resp := &DeleteSecurityGroupPolicyResponse{}
	err := client.Invoke("DeleteSecurityGroupPolicy", args, resp)
	if err != nil {
		return &DeleteSecurityGroupPolicyResponse{}, err
	}
	return resp, nil
}
