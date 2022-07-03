README
======

```bash
chmod +x build.sh

./build.sh
aws lambda update-function-code --function-name Hello \
    --zip-file fileb://./deployment.zip \
    --region us-west-2
```
