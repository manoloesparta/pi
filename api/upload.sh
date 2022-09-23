GOOS=linux go build

ACCOUNT=$1
TAG=$2

sudo docker build -t piandme:$TAG .

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/piandme

docker tag piandme:$TAG $ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/piandme:$TAG

docker push $ACCOUNT.dkr.ecr.us-east-1.amazonaws.com/piandme:$TAG
