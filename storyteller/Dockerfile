FROM python:3.11

WORKDIR /code


COPY ./requirements.txt /code/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt
COPY ./app /code/app

# Only using env variables that are useful for the test
ENV EDENAI_TOKEN=${EDENAI_TOKEN}
ENV EDENAI_PROVIDER=${EDENAI_PROVIDER}
ENV EDNAI_URL=${EDENAI_URL}

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", ${STORYTELLER_PORT}]
