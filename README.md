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

This will deploy the lambda and generate an endpoint like: http://xxxxxxxxxxxxxx.amazon.com/dev/purl 

When called the lambda return a json with a random file name and a pre-signed-url


Now to deploy the frontend, just copy the endpoint url and add it to the `client/dist/index.html`

and now deploy: 


```bash
sls client deploy

```








