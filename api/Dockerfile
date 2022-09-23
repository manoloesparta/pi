FROM public.ecr.aws/lambda/go:1

COPY . .

RUN go get

RUN GOOS=linux go build

COPY api ${LAMBDA_TASK_ROOT}

CMD [ "api" ]
