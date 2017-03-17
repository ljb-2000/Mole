FROM scratch
COPY ./Mole /Mole
COPY ./conf /conf
EXPOSE 8080
ENTRYPOINT ["/Mole"]

