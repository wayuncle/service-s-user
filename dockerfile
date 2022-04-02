FROM centos
WORKDIR /
ADD service-b-demo /
ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh" ]
CMD [ "/service-b-demo"]

