FROM public.ecr.aws/lambda/go:1

COPY . .

COPY api ${LAMBDA_TASK_ROOT}

CMD [ "api" ]
