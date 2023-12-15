FROM python:3.11

WORKDIR /code


COPY ./storyteller/requirements.txt /code/requirements.txt
RUN pip install --no-cache-dir --upgrade -r /code/requirements.txt
COPY ./storyteller/app /code/app

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", ${STORYTELLER_PORT}]
