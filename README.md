# Lambda Invoke Payload

Invokes a Lambda function with a payload. Intended to be used with a Cloudwatch event schedule.

## Running

`INVOKE_FUNCTION=my-function PAYLOAD="{\"test\":\"invoke\"}" go run main.go`

## Options

- `INVOKE_FUNCTION`: Either the Lambda function name or ARN
- `PAYLOAD`: The payload to send to the lambda (default is empty)

## Sample Cloudformation Template

```
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: Lambda Invoker

Resources:
  LambdaRole:
    Type: AWS::IAM::Role
    Properties:
        Path: "/"
        ManagedPolicyArns:
          - arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
          - arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole
        AssumeRolePolicyDocument:
          Version: "2012-10-17"
          Statement:
            - Sid: "AllowLambdaServiceToAssumeRole"
              Effect: "Allow"
              Action:
                - "sts:AssumeRole"
              Principal:
                Service:
                  - "lambda.amazonaws.com"
        Policies:
        - PolicyName: AWSLambdaInvoke
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
            - Effect: Allow
              Action:
                - "lambda:InvokeFunction"
              Resource:
                - "*"
  LambdaInvoker:
    Type: AWS::Serverless::Function
    Properties:
      Runtime: go1.x
      Timeout: 15
      MemorySize: 128
      CodeUri: bin
      Handler: main
      FunctionName: lambda-invoker
      Description: Lambda Invoker
      Events:
        Timer:
          Type: Schedule
          Properties:
            Schedule: rate(5 minutes)
      Environment:
        Variables:
          INVOKE_ARN: "my-function-name"
          PAYLOAD: "{\"test\":\"invoke\"}"
```
