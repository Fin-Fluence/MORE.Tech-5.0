FROM python:3.9 AS builder
COPY requirements.txt .
RUN pip install --user -r requirements.txt

FROM python:3.9-slim
WORKDIR /app

COPY --from=builder /root/.local /root/.local
COPY ./ .

ENV PATH=/root/.local:$PATH

CMD [ "python", "main.py" ]