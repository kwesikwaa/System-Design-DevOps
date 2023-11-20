FROM python:latest

WORKDIR /usr/src/appserver

# copy requirement.txt if you have it

RUN pip install fastapi uvicorn

# good to have the copy below installing dependencies 

COPY appserver .

EXPOSE 8000

CMD ["uvicorn", "appserver.server:app", "--host", "0.0.0.0", "--port", "8000"]