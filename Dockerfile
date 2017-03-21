#FROM scratch
FROM ubuntu
COPY ./bin/Mole_Docker /opt/mole/Mole
COPY ./conf /opt/mole/conf
COPY ./template /opt/mole/template
EXPOSE 8080
ENTRYPOINT ["/opt/mole/Mole"]
