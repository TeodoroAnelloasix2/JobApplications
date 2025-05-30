FROM ubuntu:22.04

RUN apt-get update && apt install -y python3 && mkdir /personal

WORKDIR /personal

COPY ./program.py .

CMD ["python3","program.py"]