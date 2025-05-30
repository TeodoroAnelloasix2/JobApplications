FROM ubuntu
RUN apt-get update && apt-get install -y python3 git vim iputils-ping \
&& echo "1.0.0" >> ./version

CMD ["echo", "-e","Welcome to this container"]