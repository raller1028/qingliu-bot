FROM golang:1.19 as builder

WORKDIR /app

COPY . .
# CGO_ENABLED=0 一定不要忘记,会因为依赖库的问题导构建后无法运行
RUN CGO_ENABLED=0 go build -v -o /go/bin/lark-bot


FROM python:3.8-slim

# Keeps Python from generating .pyc files in the container
ENV PYTHONDONTWRITEBYTECODE=1

# Turns off buffering for easier container logging
ENV PYTHONUNBUFFERED=1

WORKDIR /app

COPY --from=builder /go/bin/lark-bot /go/bin/lark-bot
# Install pip requirements
COPY ./qingliuSession/requirements.txt .
RUN python -m pip install -r requirements.txt

COPY . /app

# Creates a non-root user with an explicit UID and adds permission to access the /app folder
# For more info, please refer to https://aka.ms/vscode-docker-python-configure-containers
RUN adduser -u 5678 --disabled-password --gecos "" appuser && chown -R appuser /app
USER appuser

LABEL VERSION="0.1"

EXPOSE 8000


ENTRYPOINT ["bash","/app/start.sh"]