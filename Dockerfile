FROM scratch
COPY ./bin/Mole_Docker /Mole
COPY ./conf /conf
EXPOSE 8080
ENTRYPOINT ["/Mole"]

