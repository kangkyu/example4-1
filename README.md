README
======

Following tutorial of Hands-On Serverless Applications with Go
https://www.packtpub.com/product/hands-on-serverless-applications-with-go/9781789134612


```bash
chmod +x build.sh

./build.sh
aws lambda update-function-code --function-name Hello \
    --zip-file fileb://./deployment.zip \
    --region us-west-2
```
