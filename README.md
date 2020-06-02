# aws-s3-upload-pre-signed-demo

This a demo project for uploading files to an s3 bucket using a pre-signed url
issued by a lambda function.

The frontend is a static HTML + javascript page deployed also on the same bucket 


## install

Just clone the repo and make sure you have serverless framework installed

For deploying the lambda function first run make to compile the golang lambda:


```bash
make
```

and then:

```bash
sls deploy 

````

This will deploy the lambda and generate an endpoint like: https://xxxxxx.execute-api.us-east-2.amazonaws.com/dev/purl 

When called the lambda return a json with a random file name and a pre-signed-url

```bash
curl https://xxxxxxxxxxxx.execute-api.us-east-2.amazonaws.com/dev/purl
```

```json
{  
   "file":"tmp-52fdfc072182654f163f5f0f9a621d72",
   "url":"https://aws-upload-pre-signed-bucket.s3.us-east-2.amazonaws.com/tmp-52fd........"
}
```


Now to deploy the frontend, just copy the endpoint url and add it to the `client/dist/index.html`

and now deploy: 


```bash
sls client deploy

```








