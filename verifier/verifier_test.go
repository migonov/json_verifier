package verifier

import "testing"

func TestVerifier(t *testing.T) {
	cases := []struct {
		input   []byte
		isValid bool
	}{
		{
			input: []byte(`{
				"PolicyName": "root",
				"PolicyDocument": {
					"Version": "2012-10-17",
					"Statement": [
						{
							"Sid": "IamListAccess",
							"Effect": "Allow",
							"Action": [
								"iam:ListRoles",
								"iam:ListUsers"
							],
							"Resource": "*"
						}
					]
				}
			}`),
			isValid: false,
		},
		{
			input: []byte(`{
				"PolicyName": "root",
				"PolicyDocument": {
					"Version": "2012-10-17",
					"Statement": [
					  {
						"Sid": "FirstStatement",
						"Effect": "Allow",
						"Action": ["iam:ChangePassword"],
						"Resource": [
							"arn:aws:s3:::confidential-data",
							"arn:aws:s3:::confidential-data/*"
						  ]
					  },
					  {
						"Sid": "SecondStatement",
						"Effect": "Allow",
						"Action": "s3:ListAllMyBuckets",
						"Resource": "*"
					  },
					  {
						"Sid": "ThirdStatement",
						"Effect": "Allow",
						"Action": [
						  "s3:List*",
						  "s3:Get*"
						],
						"Resource": [
						  "arn:aws:s3:::confidential-data",
						  "arn:aws:s3:::confidential-data/*"
						],
						"Condition": {"Bool": {"aws:MultiFactorAuthPresent": "true"}}
					  }
					]
				  }
			}`),
			isValid: false,
		},
		{
			input: []byte(`{
				"PolicyName": "root",
				"PolicyDocument": {
					"Version": "2012-10-17",
					"Statement": [
					  {
						"Sid": "FirstStatement",
						"Effect": "Allow",
						"Action": ["iam:ChangePassword"],
						"Resource": [
							"arn:aws:s3:::confidential-data",
							"*"
						  ]
					  },
					  {
						"Sid": "ThirdStatement",
						"Effect": "Allow",
						"Action": [
						  "s3:List*",
						  "s3:Get*"
						],
						"Resource": [
						  "arn:aws:s3:::confidential-data",
						  "arn:aws:s3:::confidential-data/*"
						],
						"Condition": {"Bool": {"aws:MultiFactorAuthPresent": "true"}}
					  }
					]
				  }
			}`),
			isValid: false,
		},
		{
			input: []byte(`{
				"PolicyName": "root",
				"PolicyDocument": {
					"Version": "2012-10-17",
					"Statement": [
						{
							"Sid": "IamListAccess",
							"Effect": "Allow",
							"Action": [
								"iam:ListRoles",
								"iam:ListUsers"
							],
							"Resource": "arn:aws:s3:::confidential-data/*"
						}
					]
				}
			}`),
			isValid: true,
		},
		{
			input: []byte(`{
				"PolicyName": "root",
				"PolicyDocument": {
					"Version": "2012-10-17",
					"Statement": [
					  {
						"Sid": "FirstStatement",
						"Effect": "Allow",
						"Action": ["iam:ChangePassword"],
						"Resource": [
							"arn:aws:s3:::confidential-data",
							"arn:aws:s3:::confidential-data/*"
						  ]
					  },
					  {
						"Sid": "ThirdStatement",
						"Effect": "Allow",
						"Action": [
						  "s3:List*",
						  "s3:Get*"
						],
						"Resource": [
						  "arn:aws:s3:::confidential-data",
						  "arn:aws:s3:::confidential-data/*"
						],
						"Condition": {"Bool": {"aws:MultiFactorAuthPresent": "true"}}
					  }
					]
				  }
			}`),
			isValid: true,
		},
		{
			input: []byte(`{
				"PolicyName": "root",
				"PolicyDocument": {
					"Version": "2012-10-17",
					"Statement": [
						{
							"Sid": "IamListAccess",
							"Effect": "Allow",
							"Action": [
								"iam:ListRoles",
								"iam:ListUsers"
							]
						}
					]
				}
			}`),
			isValid: true,
		},
	}

	for _, testCase := range cases {
		isValid := Verify(testCase.input)
		if testCase.isValid != isValid {
			t.Fatalf("input: %s", testCase.input)
		}
	}
}

func TestVerifierBadJSON(t *testing.T) {
	inputJSON := []byte(`{
		"PolicyName": "root",
		"PolicyDocument": {
			"Statement": 1
		}
	}`)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("should panic for input %s", inputJSON)
		}
	}()

	Verify(inputJSON)
}
