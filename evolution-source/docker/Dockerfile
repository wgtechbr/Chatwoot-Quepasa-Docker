FROM golang:latest

WORKDIR /opt/evolution-source/

COPY ./evolution-source/ /opt/evolution-source/

RUN chmod 755 /opt/evolution-source/docker/docker-entrypoint.sh
RUN sed -i 's/\r//' /opt/evolution-source/docker/docker-entrypoint.sh

RUN ls -la /opt/evolution-source/
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 31000

ENTRYPOINT ["/opt/evolution-source/docker/docker-entrypoint.sh"]
CMD ["sh"]