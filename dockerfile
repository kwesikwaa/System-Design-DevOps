FROM python:latest

WORKDIR /appser

RUN pip install fastapi uvicorn

COPY appserver /appserver/

CMD ["uvicorn", "appserver.server:app", "--host", "0.0.0.0", "--port", "8000"]